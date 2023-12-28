// Copyright 2018 visualfc. All rights reserved.

package tk

import (
	"fmt"
	"log"
	"runtime"

	"github.com/visualfc/atk/tk/interp"
)

var (
	tkHasInit            bool
	tkWindowInitAutoHide bool
	mainInterp           *interp.Interp
	rootWindow           *Window
	fnErrorHandle        func(error) = func(err error) {
		log.Println(err)
	}
)

func Init() error {
	return InitEx(true, "", "")
}

func InitEx(tk_window_init_hide bool, tcl_library string, tk_library string) (err error) {
	mainInterp, err = interp.NewInterp()
	if err != nil {
		dumpError(err)
		return err
	}
	err = mainInterp.InitTcl(tcl_library)
	if err != nil {
		dumpError(err)
		return err
	}
	err = mainInterp.InitTk(tk_library)
	if err != nil {
		dumpError(err)
		return err
	}

	tkWindowInitAutoHide = tk_window_init_hide
	//hide console for macOS bundle
	mainInterp.Eval("if {[info commands console] == \"console\"} {console hide}")

	for _, fn := range init_func_list {
		fn()
	}
	rootWindow = &Window{}
	rootWindow.Attach(".")
	if tkWindowInitAutoHide {
		rootWindow.Hide()
	}
	//hide wish menu on macos
	rootWindow.SetMenu(NewMenu(nil))
	tkHasInit = true
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

func MainLoop(fn func()) error {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	if !tkHasInit {
		err := Init()
		if err != nil {
			dumpError(err)
			return err
		}
	}
	interp.MainLoop(fn)
	return nil
}

func Async(fn func()) {
	interp.Async(fn)
}

func Update() {
	eval("update")
}

func Quit() {
	Async(func() {
		DestroyWidget(rootWindow)
	})
}

func eval(script string) error {
	err := mainInterp.Eval(script)
	if err != nil {
		dumpError(fmt.Errorf("script: %q, error: %q", script, err))
	}
	return err
}

func evalEx(script string, dump bool) error {
	err := mainInterp.Eval(script)
	if dump && err != nil {
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

func evalAsStringEx(script string, dump bool) (string, error) {
	r, err := mainInterp.EvalAsString(script)
	if dump && err != nil {
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

func evalAsIntEx(script string, dump bool) (int, error) {
	r, err := mainInterp.EvalAsInt(script)
	if dump && err != nil {
		dumpError(err)
	}
	return r, err
}

func evalAsUint(script string) (uint, error) {
	r, err := mainInterp.EvalAsUint(script)
	if err != nil {
		dumpError(err)
	}
	return r, err
}

func evalAsUintEx(script string, dump bool) (uint, error) {
	r, err := mainInterp.EvalAsUint(script)
	if dump && err != nil {
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

func evalAsFloat64Ex(script string, dump bool) (float64, error) {
	r, err := mainInterp.EvalAsFloat64(script)
	if dump && err != nil {
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

func evalAsBoolEx(script string, dump bool) (bool, error) {
	r, err := mainInterp.EvalAsBool(script)
	if dump && err != nil {
		dumpError(err)
	}
	return r, err
}

func evalAsStringList(script string) ([]string, error) {
	r, err := mainInterp.EvalAsStringList(script)
	if err != nil {
		dumpError(err)
	}
	return r, err
}

func evalAsStringListEx(script string, dump bool) ([]string, error) {
	r, err := mainInterp.EvalAsStringList(script)
	if dump && err != nil {
		dumpError(err)
	}
	return r, err
}

func evalAsIntList(script string) ([]int, error) {
	r, err := mainInterp.EvalAsIntList(script)
	if err != nil {
		dumpError(err)
	}
	return r, err
}

func evalAsIntListEx(script string, dump bool) ([]int, error) {
	r, err := mainInterp.EvalAsIntList(script)
	if dump && err != nil {
		dumpError(err)
	}
	return r, err
}

func dumpError(err error) {
	if fnErrorHandle != nil {
		fnErrorHandle(fmt.Errorf("%v", err))
	}
}

func setObjText(obj string, text string) {
	mainInterp.SetStringVar(obj, text, false)
}

func setObjTextList(obj string, list []string) {
	mainInterp.SetStringList(obj, list, false)
}
