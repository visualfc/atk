// Copyright 2017 visualfc. All rights reserved.

package tk

import (
	"github.com/visualfc/go-tk/tk/interp"
)

var (
	mainInterp     *interp.Interp
	mainWindow     *Window
	fnErrorHandle  func(error)
	defaultMaxSize Size
	defaultMinSize Size
)

func Init() error {
	return InitEx("", "")
}

func InitEx(tcl_library string, tk_library string) (err error) {
	mainInterp, err = interp.NewInterp()
	if err != nil {
		return err
	}
	mainInterp.SetErrorHandle(fnErrorHandle)
	err = mainInterp.InitTcl(tcl_library)
	if err != nil {
		return err
	}
	err = mainInterp.InitTk(tk_library)
	if err != nil {
		return err
	}
	//hide console for macOS bundle
	mainInterp.Eval("if {[info commands console] == \"console\"} {console hide}")

	mainWindow = &Window{"."}
	mainWindow.registerWindowInfo()
	RegisterWidget(mainWindow)

	var w, h int
	w, h = mainWindow.MaximumSize()
	defaultMaxSize = Size{w, h}
	w, h = mainWindow.MinimumSize()
	defaultMinSize = Size{w, h}
	mainWindow.Hide()
	return nil
}

func SetErrorHandle(fn func(error)) {
	fnErrorHandle = fn
	if mainInterp != nil {
		mainInterp.SetErrorHandle(fn)
	}
}

func MainInterp() *interp.Interp {
	return mainInterp
}

func TclVersion() (ver string) {
	return mainInterp.TclVersion()
}

func TkVersion() (ver string) {
	return mainInterp.TkVersion()
}

func MainLoop(fn func()) {
	interp.MainLoop(fn)
}

func Async(fn func()) {
	interp.Async(fn)
}

func Update() {
	eval("update")
}

func Quit() {
	eval("destroy .")
}

func eval(script string) error {
	return mainInterp.Eval(script)
}

func evalAsString(script string) (string, error) {
	return mainInterp.EvalAsString(script)
}

func evalAsInt(script string) (int, error) {
	return mainInterp.EvalAsInt(script)
}

func evalAsFloat64(script string) (float64, error) {
	return mainInterp.EvalAsFloat64(script)
}

func evalAsBool(script string) (bool, error) {
	return mainInterp.EvalAsBool(script)
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

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}
