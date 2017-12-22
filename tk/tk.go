// Copyright 2017 visualfc. All rights reserved.

package tk

import (
	"github.com/visualfc/go-tk/tk/interp"
)

type Size struct {
	Width  int
	Height int
}

var (
	mainInterp     *interp.Interp
	root           *Window
	defaultMaxSize Size
	defaultMinSize Size
)

type WidgetId string

type Widget interface {
	Id() WidgetId
}

func Init() error {
	return InitEx("", "")
}

func InitEx(tcl_library string, tk_library string) (err error) {
	mainInterp, err = interp.NewInterp()
	if err != nil {
		return err
	}
	err = mainInterp.InitTcl(tcl_library)
	if err != nil {
		return err
	}
	err = mainInterp.InitTk(tk_library)
	if err != nil {
		return err
	}
	root = RootWindow()
	var w, h int
	w, h = root.MaximumSize()
	defaultMaxSize = Size{w, h}
	w, h = root.MinimumSize()
	defaultMinSize = Size{w, h}
	root.Hide()
	return nil
}

func TclVersion() (ver string) {
	return mainInterp.TclVersion()
}

func TkVersion() (ver string) {
	return mainInterp.TkVersion()
}

func eval(script string) error {
	return mainInterp.Eval(script)
}

func evalAsString(script string) (string, error) {
	err := mainInterp.Eval(script)
	if err != nil {
		return "", err
	}
	return mainInterp.GetStringResult(), nil
}

func evalAsInt(script string) (int, error) {
	err := mainInterp.Eval(script)
	if err != nil {
		return 0, err
	}
	return mainInterp.GetIntResult(), nil
}

func parserTwoInt(s string) (n1 int, n2 int) {
	var p = &n1
	for _, r := range s {
		if r == ' ' {
			p = &n2
		} else {
			*p = *p*10 + int(r-'0')
		}
	}
	return
}

func evalAsFloat64(script string) (float64, error) {
	err := mainInterp.Eval(script)
	if err != nil {
		return 0, err
	}
	return mainInterp.GetFloat64Result(), nil
}

func evalAsBool(script string) (bool, error) {
	err := mainInterp.Eval(script)
	if err != nil {
		return false, err
	}
	return mainInterp.GetBoolResult(), nil
}

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

func MainLoop(fn func()) {
	interp.MainLoop(fn)
}

func Async(fn func()) {
	interp.Async(fn)
}
