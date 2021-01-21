// Copyright 2018 visualfc. All rights reserved.

package tk

import (
	"fmt"
	"syscall"
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
	procGetScaleFactorForDevice, err := dll.FindProc("GetScaleFactorForDevice")
	if err != nil {
		return
	}
	procSetProcessDpiAwareness.Call(1)
	r0, _, _ := procGetScaleFactorForDevice.Call()
	mainInterp.Eval(fmt.Sprintf("tk scaling %v", float64(r0)/75))
}
