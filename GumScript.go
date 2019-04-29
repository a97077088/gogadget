package gogadget

/*
#include "frida-gumjs.h"
extern void on_message(GumScript * script, gchar * message, GBytes * data, gpointer user_data);
*/
import "C"
import (
	"encoding/json"
	"fmt"
	"sync"
	"unsafe"
)

var calls = sync.Map{}

//export on_message
func on_message(script *C.GumScript, message *C.gchar, data *C.GBytes, user_data C.gpointer) {
	fmt.Println("script call")
	defer C.g_bytes_unref(data)
	key := fmt.Sprintf("%d", uintptr(unsafe.Pointer(script)))
	jsobj := make(map[string]interface{})
	err := json.Unmarshal([]byte(C.GoString(message)), &jsobj)
	if err != nil {
		return
	}
	fv, ok := calls.Load(key)
	if ok == false {
		return
	}
	var gobytes []byte
	if uintptr(unsafe.Pointer(data)) != uintptr(0) {
		var nsize int
		pbuf := C.g_bytes_get_data(data, (*C.ulong)(unsafe.Pointer(&nsize)))
		gobytes = C.GoBytes(unsafe.Pointer(pbuf), C.int(nsize))
	}

	f := fv.(func(_script *GumScript, _message map[string]interface{}, _data []byte, _userdata uintptr))
	f(GumScriptFormPtr(uintptr(unsafe.Pointer(script))), jsobj, gobytes, uintptr(unsafe.Pointer(user_data)))
}

type GumScript struct {
	ptr uintptr
}

func GumScriptFormPtr(_ptr uintptr) *GumScript {
	return &GumScript{_ptr}
}
func (this *GumScript) CPtr() uintptr {
	return this.ptr
}
func (this *GumScript) CTypePtr() *C.GumScript {
	return (*C.GumScript)(unsafe.Pointer(this.CPtr()))
}
func (this *GumScript) On(_func func(_script *GumScript, _message map[string]interface{}, _data []byte, _userdata uintptr)) {
	key := fmt.Sprintf("%d", this.ptr)
	calls.Store(key, _func)
	C.gum_script_set_message_handler(this.CTypePtr(), (C.GDestroyNotify)(C.on_message), (C.gpointer(uintptr(0))), (C.GDestroyNotify)(C.NULL))
}

func (this *GumScript) Load() {
	var gcl *C.GCancellable
	C.gum_script_load_sync(this.CTypePtr(), gcl)
}
func (this *GumScript) UnLoad() {
	var gcl *C.GCancellable
	C.gum_script_unload_sync(this.CTypePtr(), gcl)
	C.g_object_unref((C.gpointer)(unsafe.Pointer(this.CTypePtr())))
}
