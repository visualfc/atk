// Copyright 2017 visualfc. All rights reserved.

package tk

import (
	"log"
	"runtime"

	"github.com/visualfc/go-tk/tk/interp"
)

var (
	mainInterp    *interp.Interp
	mainWindow    *Window
	fnErrorHandle func(error) = func(err error) {
		log.Println(err)
	}
)

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
	//hide console for macOS bundle
	mainInterp.Eval("if {[info commands console] == \"console\"} {console hide}")

	for _, fn := range init_func_list {
		fn()
	}

	mainWindow = &Window{}
	mainWindow.SetInternalId(".")
	mainWindow.registerWindowInfo()
	RegisterWidget(mainWindow)
	mainWindow.Hide()
	return nil
}

var (
	init_func_list []func()
)

func registerInit(fn func()) {
	init_func_list = append(init_func_list, fn)
}

func SetErrorHandle(fn func(error)) {
	fnErrorHandle = fn
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

func TclLibary() (path string) {
	path, _ = evalAsString("set tcl_library")
	return
}

func TkLibrary() (path string) {
	path, _ = evalAsString("set tk_library")
	return
}

func MainLoop(fn func()) {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	interp.MainLoop(fn)
}

func Async(fn func()) {
	interp.Async(fn)
}

func Update() {
	eval("update")
}

func Quit() {
	DestroyWidget(mainWindow)
}

func eval(script string) error {
	err := mainInterp.Eval(script)
	if err != nil {
		dumpError(err)
	}
	return err
}

func evalAsString(script string) (string, error) {
	r, err := mainInterp.EvalAsString(script)
	if err != nil {
		dumpError(err)
	}
	return r, err
}

func evalAsInt(script string) (int, error) {
	r, err := mainInterp.EvalAsInt(script)
	if err != nil {
		dumpError(err)
	}
	return r, err
}

func evalAsFloat64(script string) (float64, error) {
	r, err := mainInterp.EvalAsFloat64(script)
	if err != nil {
		dumpError(err)
	}
	return r, err
}

func evalAsBool(script string) (bool, error) {
	r, err := mainInterp.EvalAsBool(script)
	if err != nil {
		dumpError(err)
	}
	return r, err
}

func dumpError(err error) {
	if fnErrorHandle != nil {
		fnErrorHandle(err)
	}
}
