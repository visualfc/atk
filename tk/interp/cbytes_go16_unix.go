// Copyright 2018 visualfc. All rights reserved.

// +build !go1.7,!windows

package interp

import (
	"unsafe"
)

/*
#include <stdlib.h>
#include <string.h>
*/
import "C"

func toCBytes(data []byte) unsafe.Pointer {
	size := C.size_t(len(data))
	ptr := C.malloc(size)
	C.memcpy(ptr, unsafe.Pointer(&data[0]), size)
	return ptr
}
