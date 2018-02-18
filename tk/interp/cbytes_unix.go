// Copyright 2018 visualfc. All rights reserved.

// +build go1.7,!windows

package interp

import (
	"unsafe"
)

import "C"

func toCBytes(data []byte) unsafe.Pointer {
	return C.CBytes(data)
}
