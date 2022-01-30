// Copyright 2018 visualfc. All rights reserved.

// +build !windows,tclgo !windows,!cgo

package interp

import (
	"unsafe"

	"modernc.org/libc"
	"modernc.org/libc/sys/types"
)

func toCBytes(tls *libc.TLS, b []byte) uintptr {
	if len(b) == 0 {
		return 0
	}

	p := libc.Xcalloc(tls, types.Size_t(len(b)), 1)
	copy((*libc.RawMem)(unsafe.Pointer(p))[:len(b):len(b)], b)
	return p
}
