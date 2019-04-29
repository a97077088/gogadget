package gogadget

/*
#include "frida-gumjs.h"
*/
import "C"

type GumCodeSigningPolicy C.GumCodeSigningPolicy

const(
	GUM_CODE_SIGNING_OPTIONAL=iota
	GUM_CODE_SIGNING_REQUIRED
)
