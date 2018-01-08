// Copyright 2018 visualfc. All rights reserved.

// +build !go1.9

package interp

import (
	"golang.org/x/sync/syncmap"
)

var (
	globalAsyncEvent syncmap.Map
)
