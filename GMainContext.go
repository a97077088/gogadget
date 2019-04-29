package gogadget


/*
#include "frida-gumjs.h"
*/
import "C"
import "unsafe"

type GMainContext struct {
	ptr uintptr
}

func GMain_ContextFormPtr(_ptr uintptr) *GMainContext {
	return &GMainContext{_ptr}
}
func (this *GMainContext) CPtr() uintptr {
	return this.ptr
}
func (this *GMainContext) CTypePtr() *C.GMainContext {
	return (*C.GMainContext)(unsafe.Pointer(this.CPtr()))
}

func (this *GMainContext) Iteration() bool {
	b:=int(C.g_main_context_iteration(this.CTypePtr(),C.FALSE))
	if b==0{
		return false
	}else{
		return true
	}
}
func (this *GMainContext) Pending()bool {
	b:=int(C.g_main_context_pending(this.CTypePtr()))
	if b==0{
		return false
	}else{
		return true
	}
}

func NewGMainContext()*GMainContext{
	return GMain_ContextFormPtr(uintptr(unsafe.Pointer(C.g_main_context_new())))
}

func GMainContext_Get_thread_default()*GMainContext{
	return GMain_ContextFormPtr(uintptr(unsafe.Pointer(C.g_main_context_get_thread_default())))
}
