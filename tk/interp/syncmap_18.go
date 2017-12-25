// Copyright 2017 visualfc. All rights reserved.

// +build !go1.9,!go1.10

package interp

import (
	"golang.org/x/sync/syncmap"
)

var (
	globalAsyncEvent syncmap.Map
)
