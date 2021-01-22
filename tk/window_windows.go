// Copyright 2018 visualfc. All rights reserved.

package tk

import (
	"fmt"
	"syscall"

	"github.com/visualfc/atk/tk/user32"
)

func (w *Window) ShowMaximized() error {
	return eval(fmt.Sprintf("wm state %v zoomed", w.id))
}

func (w *Window) IsMaximized() bool {
	r, _ := evalAsString(fmt.Sprintf("wm state %v", w.id))
	return r == "zoomed"
}

func init() {
	registerInit(setHighDPI)
}

func setHighDPI() {
	dll, err := syscall.LoadDLL("Shcore.dll")
	if err != nil {
		return
	}
	procSetProcessDpiAwareness, err := dll.FindProc("SetProcessDpiAwareness")
	if err != nil {
		return
	}
	cxLogical, cyLogical, cxPhysical, cyPhysical := user32.GetWindowScreen()
	tkScreenWidth = cxPhysical
	tkScreenHeight = cyPhysical
	_ = cyLogical
	tkScale = float64(cxPhysical) / float64(cxLogical)
	procSetProcessDpiAwareness.Call(1)
	mainInterp.Eval(fmt.Sprintf("tk scaling %v", tkScale))
}
