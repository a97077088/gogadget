package gogadget

/*
#include "frida-gumjs.h"
*/
import "C"
import (
	"errors"
	"unsafe"
)

type GumScriptBackend struct {
	ptr uintptr
}

func GumScriptBackendFormPtr(_ptr uintptr) *GumScriptBackend {
	return &GumScriptBackend{_ptr}
}
func (this *GumScriptBackend) CPtr() uintptr {
	return this.ptr
}
func (this *GumScriptBackend) CTypePtr() *C.GumScriptBackend {
	return (*C.GumScriptBackend)(unsafe.Pointer(this.CPtr()))
}

func NewGumScriptBackend() *GumScriptBackend {
	C.gum_init_embedded()
	return GumScriptBackendFormPtr(uintptr(unsafe.Pointer(C.gum_script_backend_obtain_duk())))
}

func (this *GumScriptBackend) Free() {
	C.gum_deinit_embedded()
}

func (this *GumScriptBackend) Create_script(_name string, _script string) (*GumScript, error) {
	cname := C.CString(_name)
	defer C.free(unsafe.Pointer(cname))
	script := C.CString(_script)
	defer C.free(unsafe.Pointer(script))

	var gerr *GError
	pscript := C.gum_script_backend_create_sync(this.CTypePtr(), cname, script, (*C.GCancellable)(C.NULL), (**C.GError)(unsafe.Pointer(&gerr)))
	if gerr != nil {
		return nil, errors.New(gerr.Message())
	}
	return GumScriptFormPtr(uintptr(unsafe.Pointer(pscript))), nil
}
