// Copyright 2017 visualfc. All rights reserved.

package interp

//go:generate go run $GOROOT/src/syscall/mksyscall_windows.go -systemdll=false -output zinterp_windows.go interp_windows.go

import (
	"errors"
	"fmt"
	"os"
	"syscall"
	"unsafe"
)

//NOTE: BytePtrToString replace cgo C.GoStringN
//import "C"

type Tcl_Interp struct{}
type Tcl_ThreadId struct{}
type Tcl_Obj struct{}
type Tcl_Command struct{}

type Tcl_WideInt int64
type Tcl_Double float64

type Tcl_Event struct {
	Proc    uintptr
	NextPtr *Tcl_Event
}

// windows api calls

//sys	Tcl_CreateInterp() (interp *Tcl_Interp) = tcl86t.Tcl_CreateInterp
//sys	Tcl_DeleteInterp(interp *Tcl_Interp) = tcl86t.Tcl_DeleteInterp

//sys	Tcl_Alloc(size uint) (r *Tcl_Event) = tcl86t.Tcl_Alloc
//sys	Tcl_Eval(interp *Tcl_Interp, script *byte) (r int32) = tcl86t.Tcl_Eval
//sys	Tcl_GetStringResult(interp *Tcl_Interp) (ret *byte) = tcl86t.Tcl_GetStringResult
//sys	Tcl_GetObjResult(interp *Tcl_Interp) (obj *Tcl_Obj) = tcl86t.Tcl_GetObjResult
//sys	Tcl_GetWideIntFromObj(interp *Tcl_Interp, obj *Tcl_Obj, out *Tcl_WideInt) (status int32) = tcl86t.Tcl_GetWideIntFromObj
//sys	Tcl_GetDoubleFromObj(interp *Tcl_Interp, obj *Tcl_Obj, out *Tcl_Double) (status int32) = tcl86t.Tcl_GetDoubleFromObj
//sys	Tcl_GetBooleanFromObj(interp *Tcl_Interp, obj *Tcl_Obj, out *int32) (status int32) = tcl86t.Tcl_GetBooleanFromObj
//sys	Tcl_GetStringFromObj(obj *Tcl_Obj, length *int32) (ret *byte) = tcl86t.Tcl_GetStringFromObj
//sys	Tcl_NewWideIntObj(value int64) (obj *Tcl_Obj) = tcl86t.Tcl_NewWideIntObj
//sys	Tcl_NewDoubleObj(value float64) (obj *Tcl_Obj) = tcl86t.Tcl_NewDoubleObj
//sys	Tcl_NewBooleanObj(value bool) (obj *Tcl_Obj) = tcl86t.Tcl_NewBooleanObj
//sys	Tcl_NewStringObj(bytes *byte, length int32) (obj *Tcl_Obj) = tcl86t.Tcl_NewStringObj
//sys	Tcl_Init(interp *Tcl_Interp) (r int32) = tcl86t.Tcl_Init
//sys	Tcl_GetCurrentThread() (threadid *Tcl_ThreadId) = tcl86t.Tcl_GetCurrentThread
//sys	Tcl_ThreadQueueEvent(threadId *Tcl_ThreadId, evPtr *Tcl_Event, positon Tcl_QueuePosition) = tcl86t.Tcl_ThreadQueueEvent
//sys	Tcl_ThreadAlert(threadId *Tcl_ThreadId) = tcl86t.Tcl_ThreadAlert
//sys	Tcl_CreateObjCommand(interp *Tcl_Interp, cmdName *byte, proc uintptr, clientData uintptr, deleteProc uintptr) (cmd *Tcl_Command) = tcl86t.Tcl_CreateObjCommand
//sys	Tcl_CreateCommand(interp *Tcl_Interp, cmdName *byte, proc uintptr, clientData uintptr, deleteProc uintptr) (cmd *Tcl_Command) = tcl86t.Tcl_CreateCommand
//sys	Tcl_SetObjResult(interp *Tcl_Interp, resultObjPtr *Tcl_Obj) = tcl86t.Tcl_SetObjResult
//sys	Tcl_WrongNumArgs(interp *Tcl_Interp, objc int32, objv uintptr, message *byte) = tcl86t.Tcl_WrongNumArgs

//sys	Tk_Init(interp *Tcl_Interp) (r int32) = tk86t.Tk_Init
//sys	Tk_MainLoop() = tk86t.Tk_MainLoop

var (
	mainLoopThreadId *Tcl_ThreadId
)

func _go_async_event_handler(ev *Tcl_Event, flags int32) int {
	if flags != TCL_ALL_EVENTS {
		return 0
	}
	if fn, ok := globalAsyncEvent.Load(unsafe.Pointer(ev)); ok {
		fn.(func())()
		globalAsyncEvent.Delete(unsafe.Pointer(ev))
	}
	return 1
}

func IsMainThread() bool {
	return Tcl_GetCurrentThread() == mainLoopThreadId
}

func async_send_event(tid *Tcl_ThreadId, fn func()) {
	var ev *Tcl_Event
	ev = Tcl_Alloc(uint(unsafe.Sizeof(*ev)))
	ev.Proc = syscall.NewCallbackCDecl(_go_async_event_handler)
	ev.NextPtr = nil
	globalAsyncEvent.Store(unsafe.Pointer(ev), fn)
	Tcl_ThreadQueueEvent(tid, ev, TCL_QUEUE_TAIL)
	Tcl_ThreadAlert(tid)
}

func Async(fn func()) {
	if fn == nil {
		return
	}
	async_send_event(mainLoopThreadId, fn)
}

func MainLoop(fn func()) {
	mainLoopThreadId = Tcl_GetCurrentThread()
	if fn != nil {
		fn()
	}
	Tk_MainLoop()
	mainLoopThreadId = nil
}

type Interp struct {
	interp *Tcl_Interp
}

func NewInterp() (*Interp, error) {
	err := modtcl86t.Load()
	if err != nil {
		return nil, err
	}
	interp := Tcl_CreateInterp()
	if interp == nil {
		return nil, errors.New("Tcl_CreateInterp failed")
	}
	return &Interp{interp}, nil
}

func (p *Interp) InitTcl(tcl_library string) error {
	if tcl_library != "" {
		p.Eval(fmt.Sprintf("set tcl_library {%s}", tcl_library))
	}
	if Tcl_Init(p.interp) != TCL_OK {
		return errors.New("Tcl_Init failed")
	}
	return nil
}

func (p *Interp) InitTk(tk_library string) error {
	err := modtk86t.Load()
	if err != nil {
		return err
	}
	if tk_library != "" {
		p.Eval(fmt.Sprintf("set tk_library {%s}", tk_library))
	}
	if Tk_Init(p.interp) != TCL_OK {
		return errors.New("Tk_Init failed")
	}
	return nil
}

func (p *Interp) Destroy() error {
	if p == nil || p.interp == nil {
		return os.ErrInvalid
	}
	Tcl_DeleteInterp(p.interp)
	p.interp = nil
	return nil
}

func BytePtrToString(data *byte, length int32) string {
	if length == 0 {
		return ""
	}
	p := (*[1 << 30]byte)(unsafe.Pointer(data))[:length]
	a := make([]byte, length)
	copy(a, p)
	return string(a)
}

func (p *Interp) GetStringResult() string {
	obj := Tcl_GetObjResult(p.interp)
	var n int32
	out := Tcl_GetStringFromObj(obj, &n)
	return BytePtrToString(out, n)
	//return C.GoStringN((*C.char)(unsafe.Pointer(r)), C.int(out))
}

func (p *Interp) GetInt64Result() int64 {
	obj := Tcl_GetObjResult(p.interp)
	var out Tcl_WideInt
	status := Tcl_GetWideIntFromObj(p.interp, obj, &out)
	if status == TCL_OK {
		return int64(out)
	}
	return 0
}

func (p *Interp) GetIntResult() int {
	obj := Tcl_GetObjResult(p.interp)
	var out Tcl_WideInt
	status := Tcl_GetWideIntFromObj(p.interp, obj, &out)
	if status == TCL_OK {
		return int(out)
	}
	return 0
}

func (p *Interp) GetFloat64Result() float64 {
	obj := Tcl_GetObjResult(p.interp)
	var out Tcl_Double
	status := Tcl_GetDoubleFromObj(p.interp, obj, &out)
	if status == TCL_OK {
		return float64(out)
	}
	return 0
}

func (p *Interp) GetBoolResult() bool {
	obj := Tcl_GetObjResult(p.interp)
	var out int32
	status := Tcl_GetBooleanFromObj(p.interp, obj, &out)
	if status == TCL_OK {
		return out == 1
	}
	return false
}

func (p *Interp) Eval(script string) error {
	s, err := syscall.BytePtrFromString(script)
	if err != nil {
		return err
	}
	if Tcl_Eval(p.interp, s) != TCL_OK {
		return errors.New(p.GetStringResult())
	}
	return nil
}

//typedef int (Tcl_ObjCmdProc) (ClientData clientData, *Tcl_Interp *interp, int objc, struct *Tcl_Obj *const *objv);
func _go_tcl_objcmd_proc(clientData uintptr, interp *Tcl_Interp, objc int, objv unsafe.Pointer) int {
	objs := (*(*[1 << 20]*Tcl_Obj)(objv))[1:objc]
	var args []string
	for _, obj := range objs {
		args = append(args, ObjToString(interp, obj))
	}
	result, err := globalCommandMap.Invoke(clientData, args)
	if err != nil {
		cs, _ := syscall.BytePtrFromString(err.Error())
		Tcl_WrongNumArgs(interp, 1, uintptr(objv), cs)
		return TCL_ERROR
	}
	if result != "" {
		Tcl_SetObjResult(interp, StringToObj(result))
	}
	return TCL_OK
}

//typedef void (Tcl_CmdDeleteProc) (ClientData clientData);
func _go_tcl_cmddelete_proc(clientData uintptr) int {
	globalCommandMap.UnRegister(clientData)
	return 0
}

func _go_tcl_action_proc(id uintptr, interp *Tcl_Interp, objc int, objv unsafe.Pointer) int {
	err := globalActionMap.Invoke(id)
	if err != nil {
		cs, _ := syscall.BytePtrFromString(err.Error())
		Tcl_WrongNumArgs(interp, 1, uintptr(objv), cs)
		return TCL_ERROR
	}
	return TCL_OK
}

func _go_tcl_actiono_delete_proc(id uintptr) int {
	globalActionMap.UnRegister(id)
	return 0
}

//Tcl_Command Tcl_CreateObjCommand(*Tcl_Interp *interp, const char *cmdName, Tcl_ObjCmdProc *proc, ClientData clientData, Tcl_CmdDeleteProc *deleteProc);
func (p *Interp) CreateCommand(name string, fn func([]string) (string, error)) (uintptr, error) {
	s, err := syscall.BytePtrFromString(name)
	if err != nil {
		return 0, err
	}
	id := globalCommandMap.Register(fn)
	cmd := Tcl_CreateObjCommand(p.interp, s, syscall.NewCallbackCDecl(_go_tcl_objcmd_proc), id, syscall.NewCallbackCDecl(_go_tcl_cmddelete_proc))
	if cmd == nil {
		return 0, fmt.Errorf("CreateObjCommand %s failed", name)
	}
	return id, nil
}

func (p *Interp) InvokeCommand(id uintptr, args []string) (string, error) {
	return globalCommandMap.Invoke(id, args)
}

func (p *Interp) CreateAction(name string, action func()) error {
	s, err := syscall.BytePtrFromString(name)
	if err != nil {
		return err
	}
	id := globalActionMap.Register(action)
	cmd := Tcl_CreateObjCommand(p.interp, s, syscall.NewCallbackCDecl(_go_tcl_action_proc), id, syscall.NewCallbackCDecl(_go_tcl_actiono_delete_proc))
	if cmd == nil {
		return fmt.Errorf("CreateObjCommand %s failed", name)
	}
	return nil
}

func (p *Interp) CreateActionByType(typ string, action func()) (name string, err error) {
	id := globalActionMap.Register(action)
	name = fmt.Sprintf("_go_%v_%v", typ, id)
	s, err := syscall.BytePtrFromString(name)
	if err != nil {
		return name, err
	}
	cmd := Tcl_CreateObjCommand(p.interp, s, syscall.NewCallbackCDecl(_go_tcl_action_proc), id, syscall.NewCallbackCDecl(_go_tcl_actiono_delete_proc))
	if cmd == nil {
		return name, fmt.Errorf("CreateObjCommand %s failed", name)
	}
	return name, nil
}

func ObjToInt64(interp *Tcl_Interp, obj *Tcl_Obj) int64 {
	var out Tcl_WideInt
	status := Tcl_GetWideIntFromObj(interp, obj, &out)
	if status == TCL_OK {
		return int64(out)
	}
	return 0
}

func ObjToString(interp *Tcl_Interp, obj *Tcl_Obj) string {
	var n int32
	out := Tcl_GetStringFromObj(obj, &n)
	return BytePtrToString(out, n)
	//return C.GoStringN((*C.char)(unsafe.Pointer(out)), (C.int)(n))
}

func StringToObj(value string) *Tcl_Obj {
	s, err := syscall.BytePtrFromString(value)
	if err != nil {
		return nil
	}
	return Tcl_NewStringObj(s, int32(len(value)))
}

func Int64ToObj(value int64) *Tcl_Obj {
	return Tcl_NewWideIntObj(value)
}
