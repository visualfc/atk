// Copyright 2018 visualfc. All rights reserved.

// +build go1.9

package interp

import (
	"sync"
)

var (
	globalAsyncEvent sync.Map
)
