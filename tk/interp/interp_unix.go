// Copyright 2017 visualfc. All rights reserved.

// +build !windows

package interp

import (
	"errors"
	"fmt"
	"os"
	"unsafe"
)

/*
#cgo darwin CFLAGS: -I/Library/Frameworks/Tcl.framework/Headers -I/Library/Frameworks/Tk.framework/Headers
#cgo darwin LDFLAGS: -lTcl -lTk
#cgo linux CFLAGS: -I/usr/include
#cgo linux LDFLAGS: -ltcl8.6 -ltk8.6
#include <tcl.h>
#include <tk.h>
#include <stdlib.h>

extern int _go_async_event_handler(Tcl_Event*, int);

static Tcl_Event* _c_create_async_event()
{
	Tcl_Event *ev = (Tcl_Event*)Tcl_Alloc(sizeof(Tcl_Event));
	ev->proc = &_go_async_event_handler;
	ev->nextPtr = 0;
	return ev;
}

static void _c_send_async_event(Tcl_ThreadId tid, Tcl_Event *ev)
{
	Tcl_ThreadQueueEvent(tid, ev, TCL_QUEUE_TAIL);
	Tcl_ThreadAlert(tid);
}

extern  int _go_tcl_objcmd_proc(void *clientData, Tcl_Interp *interp, int objc, void *objv);
extern void _go_tcl_deletecmd_proc(void *clientData);
static  int _c_tcl_objcmd_proc(void *clientData, Tcl_Interp *interp, int objc, Tcl_Obj *const *objv)
{
	return _go_tcl_objcmd_proc(clientData, interp, objc, (Tcl_Obj**)objv);
}
static Tcl_Command _c_create_obj_command(Tcl_Interp *interp, char *name, void* clientData)
{
	return Tcl_CreateObjCommand(interp,name,_c_tcl_objcmd_proc,clientData,&_go_tcl_deletecmd_proc);
}

extern  int _go_tcl_actioncmd_proc(void *clientData, Tcl_Interp *interp, int objc, void *objv);
extern void _go_tcl_deleteaction_proc(void *clientData);
static  int _c_tcl_actioncmd_proc(void *clientData, Tcl_Interp *interp, int objc, Tcl_Obj *const *objv)
{
	return _go_tcl_actioncmd_proc(clientData, interp, objc, (Tcl_Obj**)objv);
}
static Tcl_Command _c_create_action_command(Tcl_Interp *interp, char *name, void* clientData)
{
	return Tcl_CreateObjCommand(interp,name,_c_tcl_actioncmd_proc,clientData,&_go_tcl_deleteaction_proc);
}

static void _c_wrong_num_args(Tcl_Interp *interp, int objc, void *objv, char *message)
{
	Tcl_WrongNumArgs(interp, objc, (Tcl_Obj**)objv, message)	;
}

*/
import "C"

var (
	mainLoopThreadId C.Tcl_ThreadId
)

//export _go_tcl_objcmd_proc
func _go_tcl_objcmd_proc(clientData unsafe.Pointer, interp *C.Tcl_Interp, objc C.int, objv unsafe.Pointer) C.int {
	objs := (*(*[1 << 20]*C.Tcl_Obj)(objv))[1:objc]
	var args []string
	for _, obj := range objs {
		args = append(args, ObjToString(interp, obj))
	}
	result, err := globalCommandMap.Invoke(uintptr(clientData), args)
	if err != nil {
		cs := C.CString(err.Error())
		defer C.free(unsafe.Pointer(cs))
		C._c_wrong_num_args(interp, 1, objv, cs)
		return TCL_ERROR
	}
	if result != "" {
		C.Tcl_SetObjResult(interp, StringToObj(result))
	}
	return TCL_OK
}

//export _go_tcl_deletecmd_proc
func _go_tcl_deletecmd_proc(clientData unsafe.Pointer) {
	globalCommandMap.UnRegister(uintptr(clientData))
	return
}

//export _go_tcl_actioncmd_proc
func _go_tcl_actioncmd_proc(clientData unsafe.Pointer, interp *C.Tcl_Interp, objc C.int, objv unsafe.Pointer) C.int {
	err := globalActionMap.Invoke(uintptr(clientData))
	if err != nil {
		cs := C.CString(err.Error())
		defer C.free(unsafe.Pointer(cs))
		C._c_wrong_num_args(interp, 1, objv, cs)
		return TCL_ERROR
	}
	return TCL_OK
}

//export _go_tcl_deleteaction_proc
func _go_tcl_deleteaction_proc(clientData unsafe.Pointer) {
	globalActionMap.UnRegister(uintptr(clientData))
	return
}

//export _go_async_event_handler
func _go_async_event_handler(ev *C.Tcl_Event, flags C.int) C.int {
	if flags != C.TK_ALL_EVENTS {
		return 0
	}
	if fn, ok := globalAsyncEvent.Load(unsafe.Pointer(ev)); ok {
		fn.(func())()
		globalAsyncEvent.Delete(unsafe.Pointer(ev))
	}
	return 1
}

func IsMainThread() bool {
	return C.Tcl_GetCurrentThread() == mainLoopThreadId
}

func async_send_event(tid C.Tcl_ThreadId, fn func()) {
	ev := C._c_create_async_event()
	globalAsyncEvent.Store(unsafe.Pointer(ev), fn)
	C._c_send_async_event(tid, ev)
}

func Async(fn func()) {
	if fn == nil {
		return
	}
	async_send_event(mainLoopThreadId, fn)
}

func MainLoop(fn func()) {
	mainLoopThreadId = C.Tcl_GetCurrentThread()
	if fn != nil {
		fn()
	}
	C.Tk_MainLoop()
	mainLoopThreadId = nil
}

type Interp struct {
	interp        *C.Tcl_Interp
	fnErrorHandle func(error)
}

func NewInterp() (*Interp, error) {
	interp := C.Tcl_CreateInterp()
	if interp == nil {
		return nil, errors.New("Tcl_CreateInterp failed")
	}
	return &Interp{interp}, nil
}

func (p *Interp) InitTcl(tcl_library string) error {
	if tcl_library != "" {
		p.Eval(fmt.Sprintf("set tcl_library {%s}", tcl_library))
	}
	if C.Tcl_Init(p.interp) != TCL_OK {
		err := errors.New("Tcl_Init failed")
		if p.fnErrorHandle != nil {
			p.fnErrorHandle(err)
		}
		return err
	}
	return nil
}

func (p *Interp) InitTk(tk_library string) error {
	if tk_library != "" {
		p.Eval(fmt.Sprintf("set tk_library {%s}", tk_library))
	}
	if C.Tk_Init(p.interp) != TCL_OK {
		err := errors.New("Tk_Init failed")
		if p.fnErrorHandle != nil {
			p.fnErrorHandle(err)
		}
		return err
	}
	return nil
}

func (p *Interp) Destroy() error {
	if p == nil || p.interp == nil {
		return os.ErrInvalid
	}
	C.Tcl_DeleteInterp(p.interp)
	p.interp = nil
	return nil
}

func (p *Interp) GetStringResult() string {
	obj := C.Tcl_GetObjResult(p.interp)
	var out C.int
	r := C.Tcl_GetStringFromObj(obj, &out)
	return C.GoStringN(r, out)
}

func (p *Interp) GetIntResult() int {
	obj := C.Tcl_GetObjResult(p.interp)
	var out C.Tcl_WideInt
	status := C.Tcl_GetWideIntFromObj(p.interp, obj, &out)
	if status == C.TCL_OK {
		return int(out)
	}
	return 0
}

func (p *Interp) GetInt64Result() int64 {
	obj := C.Tcl_GetObjResult(p.interp)
	var out C.Tcl_WideInt
	status := C.Tcl_GetWideIntFromObj(p.interp, obj, &out)
	if status == C.TCL_OK {
		return int64(out)
	}
	return 0
}

func (p *Interp) GetFloat64Result() float64 {
	obj := C.Tcl_GetObjResult(p.interp)
	var out C.double
	status := C.Tcl_GetDoubleFromObj(p.interp, obj, &out)
	if status == C.TCL_OK {
		return float64(out)
	}
	return 0
}

func (p *Interp) GetBoolResult() bool {
	obj := C.Tcl_GetObjResult(p.interp)
	var out C.int
	status := C.Tcl_GetBooleanFromObj(p.interp, obj, &out)
	if status == C.TCL_OK {
		return out == 1
	}
	return false
}

func (p *Interp) Eval(script string) error {
	cs := C.CString(script)
	defer C.free(unsafe.Pointer(cs))
	if C.Tcl_Eval(p.interp, cs) != TCL_OK {
		err := errors.New(p.GetStringResult())
		if p.fnErrorHandle != nil {
			p.fnErrorHandle(err)
		}
		return err
	}
	return nil
}

func (p *Interp) CreateCommand(name string, fn func([]string) (string, error)) (uintptr, error) {
	cs := C.CString(name)
	defer C.free(unsafe.Pointer(cs))

	id := globalCommandMap.Register(fn)
	cmd := C._c_create_obj_command(p.interp, cs, unsafe.Pointer(id))
	if cmd == nil {
		err := fmt.Errorf("CreateCommand %v failed", name)
		if p.fnErrorHandle != nil {
			p.fnErrorHandle(err)
		}
		return 0, err
	}
	return id, nil
}

func (p *Interp) InvokeCommand(id uintptr, args []string) (string, error) {
	return globalCommandMap.Invoke(id, args)
}

func (p *Interp) CreateAction(name string, fn func()) (uintptr, error) {
	cs := C.CString(name)
	defer C.free(unsafe.Pointer(cs))

	id := globalActionMap.Register(fn)
	cmd := C._c_create_action_command(p.interp, cs, unsafe.Pointer(id))
	if cmd == nil {
		err := fmt.Errorf("CreateAction %v failed", name)
		if p.fnErrorHandle != nil {
			p.fnErrorHandle(err)
		}
		return 0, err
	}
	return id, nil
}

func (p *Interp) CreateActionByType(typ string, fn func()) (string, error) {
	id := globalActionMap.Register(fn)
	name := fmt.Sprintf("_go_%v_%v", typ, id)
	cs := C.CString(name)
	defer C.free(unsafe.Pointer(cs))
	cmd := C._c_create_action_command(p.interp, cs, unsafe.Pointer(id))
	if cmd == nil {
		err := fmt.Errorf("CreateAction %v failed", typ)
		if p.fnErrorHandle != nil {
			p.fnErrorHandle(err)
		}
		return "", err
	}
	return name, nil
}

func (p *Interp) InvokeAction(id uintptr) error {
	return globalActionMap.Invoke(id)
}

func ObjToInt(interp *C.Tcl_Interp, obj *C.Tcl_Obj) int {
	var out C.Tcl_WideInt
	status := C.Tcl_GetWideIntFromObj(interp, obj, &out)
	if status == TCL_OK {
		return int(out)
	}
	return 0
}

func ObjToString(interp *C.Tcl_Interp, obj *C.Tcl_Obj) string {
	var n C.int
	out := C.Tcl_GetStringFromObj(obj, &n)
	return C.GoStringN(out, n)
}

func StringToObj(value string) *C.Tcl_Obj {
	cs := C.CString(value)
	defer C.free(unsafe.Pointer(cs))
	return C.Tcl_NewStringObj(cs, C.int(len(value)))
}

func Int64ToObj(value int64) *C.Tcl_Obj {
	return C.Tcl_NewWideIntObj(C.Tcl_WideInt(value))
}
