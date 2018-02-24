// Copyright 2018 visualfc. All rights reserved.

package interp

//go:generate go run $GOROOT/src/syscall/mksyscall_windows.go -systemdll=false -output zinterp_windows.go interp_windows.go

import (
	"errors"
	"fmt"
	"image"
	"image/draw"
	"os"
	"syscall"
	"unsafe"
)

//NOTE: BytePtrToString replace cgo C.GoStringN

type Tcl_Interp struct{}
type Tcl_ThreadId struct{}
type Tcl_Obj struct{}
type Tcl_Command struct{}
type Tk_PhotoHandle struct{}

type Tcl_WideInt int64
type Tcl_Double float64

type Tcl_Event struct {
	Proc    uintptr
	NextPtr *Tcl_Event
}

type Tk_PhotoImageBlock struct {
	pixelPtr  *byte
	width     int32
	height    int32
	pitch     int32
	pixelSize int32
	offset    [4]int32
}

// windows api calls

//sys	Tcl_CreateInterp() (interp *Tcl_Interp) = tcl86t.Tcl_CreateInterp
//sys	Tcl_DeleteInterp(interp *Tcl_Interp) = tcl86t.Tcl_DeleteInterp

//sys	Tcl_Alloc(size uint) (r *Tcl_Event) = tcl86t.Tcl_Alloc
//sys	Tcl_Eval(interp *Tcl_Interp, script *byte) (r int32) = tcl86t.Tcl_Eval
//sys	Tcl_EvalEx(interp *Tcl_Interp, script *byte, length int32, flags int32) (r int32) = tcl86t.Tcl_EvalEx
//sys	Tcl_GetStringResult(interp *Tcl_Interp) (ret *byte) = tcl86t.Tcl_GetStringResult
//sys	Tcl_GetObjResult(interp *Tcl_Interp) (obj *Tcl_Obj) = tcl86t.Tcl_GetObjResult
//sys	Tcl_GetWideIntFromObj(interp *Tcl_Interp, obj *Tcl_Obj, out *Tcl_WideInt) (status int32) = tcl86t.Tcl_GetWideIntFromObj
//-sys	Tcl_GetLongFromObj(interp *Tcl_Interp, obj *Tcl_Obj, out *int) (status int32) = tcl86t.Tcl_GetLongFromObj
//sys	Tcl_GetDoubleFromObj(interp *Tcl_Interp, obj *Tcl_Obj, out *Tcl_Double) (status int32) = tcl86t.Tcl_GetDoubleFromObj
//sys	Tcl_GetBooleanFromObj(interp *Tcl_Interp, obj *Tcl_Obj, out *int32) (status int32) = tcl86t.Tcl_GetBooleanFromObj
//sys	Tcl_GetStringFromObj(obj *Tcl_Obj, length *int32) (ret *byte) = tcl86t.Tcl_GetStringFromObj
//sys	Tcl_NewWideIntObj(value Tcl_WideInt) (obj *Tcl_Obj) = tcl86t.Tcl_NewWideIntObj
//-sys	Tcl_NewLongObj(value int) (obj *Tcl_Obj) = tcl86t.Tcl_NewLongObj
//sys	Tcl_NewDoubleObj(value Tcl_Double) (obj *Tcl_Obj) = tcl86t.Tcl_NewDoubleObj
//sys	Tcl_NewBooleanObj(value int32) (obj *Tcl_Obj) = tcl86t.Tcl_NewBooleanObj
//sys	Tcl_NewStringObj(bytes *byte, length int32) (obj *Tcl_Obj) = tcl86t.Tcl_NewStringObj
//sys	Tcl_Init(interp *Tcl_Interp) (r int32) = tcl86t.Tcl_Init
//sys	Tcl_GetCurrentThread() (threadid *Tcl_ThreadId) = tcl86t.Tcl_GetCurrentThread
//sys	Tcl_ThreadQueueEvent(threadId *Tcl_ThreadId, evPtr *Tcl_Event, positon Tcl_QueuePosition) = tcl86t.Tcl_ThreadQueueEvent
//sys	Tcl_ThreadAlert(threadId *Tcl_ThreadId) = tcl86t.Tcl_ThreadAlert
//sys	Tcl_CreateObjCommand(interp *Tcl_Interp, cmdName *byte, proc uintptr, clientData uintptr, deleteProc uintptr) (cmd *Tcl_Command) = tcl86t.Tcl_CreateObjCommand
//sys	Tcl_CreateCommand(interp *Tcl_Interp, cmdName *byte, proc uintptr, clientData uintptr, deleteProc uintptr) (cmd *Tcl_Command) = tcl86t.Tcl_CreateCommand
//sys	Tcl_SetObjResult(interp *Tcl_Interp, resultObjPtr *Tcl_Obj) = tcl86t.Tcl_SetObjResult
//sys	Tcl_WrongNumArgs(interp *Tcl_Interp, objc int32, objv uintptr, message *byte) = tcl86t.Tcl_WrongNumArgs
//sys	Tcl_NewListObj(objc int, objv **Tcl_Obj)(obj *Tcl_Obj) = tcl86t.Tcl_NewListObj
//sys	Tcl_ListObjLength(interp *Tcl_Interp, listobj *Tcl_Obj, length *int32) (status int32) = tcl86t.Tcl_ListObjLength
//sys	Tcl_ListObjIndex(interp *Tcl_Interp, listobj *Tcl_Obj, index int32, out **Tcl_Obj) (status int32) = tcl86t.Tcl_ListObjIndex
//sys	Tcl_ListObjGetElements(interp *Tcl_Interp, listobj *Tcl_Obj, objc *int32, objv ***Tcl_Obj)(status int32) = tcl86t.Tcl_ListObjGetElements
//sys	Tcl_SetListObj(listobj *Tcl_Obj, objc int, objv **Tcl_Obj) = tcl86t.Tcl_SetListObj
//sys	Tcl_ListObjAppendElement(interp *Tcl_Interp, listobj *Tcl_Obj, obj *Tcl_Obj) (status int32) = tcl86t.Tcl_ListObjAppendElement
//sys	Tcl_ListObjReplace(interp *Tcl_Interp, listobj *Tcl_Obj, first int32, count int32, objc int32, objv **Tcl_Obj) (status int32) = tcl86t.Tcl_ListObjReplace
//sys	Tcl_GetVar2Ex(interp *Tcl_Interp,part1 *byte, part2 *byte, flags int32) (obj *Tcl_Obj) = tcl86t.Tcl_GetVar2Ex
//sys	Tcl_SetVar(interp *Tcl_Interp,name *byte, value *byte, flags int32) (r *byte) = tcl86t.Tcl_SetVar
//sys	Tcl_SetVar2Ex(interp *Tcl_Interp,part1 *byte, part2 *byte, value *Tcl_Obj, flags int32) (r *byte) = tcl86t.Tcl_SetVar2Ex
//sys	Tcl_UnsetVar(interp *Tcl_Interp,part1 *byte, flags int32) (status int32) = tcl86t.Tcl_UnsetVar

//sys	Tk_Init(interp *Tcl_Interp) (r int32) = tk86t.Tk_Init
//sys	Tk_MainLoop() = tk86t.Tk_MainLoop
//sys	Tk_FindPhoto(interp *Tcl_Interp, imageName *byte) (handle *Tk_PhotoHandle) = tk86t.Tk_FindPhoto
//sys	Tk_PhotoBlank(handle *Tk_PhotoHandle) = tk86t.Tk_PhotoBlank
//sys	Tk_PhotoSetSize(interp *Tcl_Interp,handle *Tk_PhotoHandle, width int32, height int32) (status int32) = tk86t.Tk_PhotoSetSize
//sys	Tk_PhotoGetSize(hanlde *Tk_PhotoHandle, widthPtr *int32, heightPtr *int32) = tk86t.Tk_PhotoGetSize
//sys	Tk_PhotoExpand(interp *Tcl_Interp,handle *Tk_PhotoHandle, width int32, height int32) (status int32) = tk86t.Tk_PhotoExpand
//sys	Tk_PhotoGetImage(handle *Tk_PhotoHandle, blockPtr *Tk_PhotoImageBlock) (status int32) = tk86t.Tk_PhotoGetImage
//sys	Tk_PhotoPutBlock(interp *Tcl_Interp, handle *Tk_PhotoHandle,blockPtr *Tk_PhotoImageBlock, x int32, y int32, width int32, height int32, compRule int32) (status int32) = tk86t.Tk_PhotoPutBlock
//sys	Tk_PhotoPutZoomedBlock(interp *Tcl_Interp, handle *Tk_PhotoHandle,blockPtr *Tk_PhotoImageBlock, x int32, y int32, width int32, height int32, zoomX int32, zoomY int32, subsampleX int32, subsampleY int32, compRule int32) (status int32) = tk86t.Tk_PhotoPutZoomedBlock

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
		err := errors.New("Tcl_Init failed")
		return err
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
		err := errors.New("Tk_Init failed")
		return err
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

func (p *Interp) GetObjResult() *Obj {
	obj := Tcl_GetObjResult(p.interp)
	return &Obj{obj, p.interp}
}

func (p *Interp) GetListObjResult() *ListObj {
	return &ListObj{Tcl_GetObjResult(p.interp), p.interp}
}

func (p *Interp) Eval(script string) error {
	s, err := syscall.BytePtrFromString(script)
	if err != nil {
		return err
	}
	if Tcl_EvalEx(p.interp, s, int32(len(script)), 0) != TCL_OK {
		err := errors.New(p.GetStringResult())
		return err
	}
	return nil
}

//typedef int (Tcl_ObjCmdProc) (ClientData clientData, *Tcl_Interp *interp, int objc, struct *Tcl_Obj *const *objv);
func _go_tcl_objcmd_proc(clientData uintptr, interp *Tcl_Interp, objc int, objv unsafe.Pointer) int {
	objs := (*(*[1 << 20]*Tcl_Obj)(objv))[1:objc]
	var args []string
	for _, obj := range objs {
		args = append(args, objToString(interp, obj))
	}
	result, err := globalCommandMap.Invoke(clientData, args)
	if err != nil {
		cs, _ := syscall.BytePtrFromString(err.Error())
		Tcl_WrongNumArgs(interp, 1, uintptr(objv), cs)
		return TCL_ERROR
	}
	if result != "" {
		Tcl_SetObjResult(interp, stringToObj(result))
	}
	return TCL_OK
}

//typedef void (Tcl_CmdDeleteProc) (ClientData clientData);
func _go_tcl_cmddelete_proc(clientData uintptr) int {
	globalCommandMap.UnRegister(clientData)
	return 0
}

func _go_tcl_action_proc(id uintptr, interp *Tcl_Interp, objc int, objv unsafe.Pointer) int {
	objs := (*(*[1 << 20]*Tcl_Obj)(objv))[1:objc]
	var args []string
	for _, obj := range objs {
		args = append(args, objToString(interp, obj))
	}
	err := globalActionMap.Invoke(id, args)
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
		err := fmt.Errorf("CreateCommand %v failed", name)
		return 0, err
	}
	return id, nil
}

func (p *Interp) InvokeCommand(id uintptr, args []string) (string, error) {
	return globalCommandMap.Invoke(id, args)
}

func (p *Interp) CreateAction(name string, action func([]string)) (uintptr, error) {
	s, err := syscall.BytePtrFromString(name)
	if err != nil {
		return 0, err
	}
	id := globalActionMap.Register(action)
	cmd := Tcl_CreateObjCommand(p.interp, s, syscall.NewCallbackCDecl(_go_tcl_action_proc), id, syscall.NewCallbackCDecl(_go_tcl_actiono_delete_proc))
	if cmd == nil {
		err := fmt.Errorf("CreateAction %v failed", name)
		return 0, err
	}
	return id, nil
}

func (p *Interp) InvokeAction(id uintptr, args []string) error {
	return globalActionMap.Invoke(id, args)
}

const (
	TCL_LEAVE_ERR_MSG  = 0x200
	TCL_GLOBAL_ONLY    = 1
	TCL_NAMESPACE_ONLY = 2
	TCL_APPEND_VALUE   = 4
	TCL_LIST_ELEMENT   = 8
)

func (p *Interp) GetVar(name string, global bool) *Obj {
	cname, err := syscall.BytePtrFromString(name)
	if err != nil {
		return nil
	}
	var flag int32 = TCL_LEAVE_ERR_MSG
	if global {
		flag |= TCL_GLOBAL_ONLY
	}
	obj := Tcl_GetVar2Ex(p.interp, cname, nil, flag)
	if obj == nil {
		return nil
	}
	return &Obj{obj, p.interp}
}

func (p *Interp) GetList(name string, global bool) *ListObj {
	return (*ListObj)(p.GetVar(name, global))
}

func (p *Interp) SetStringList(name string, list []string, global bool) error {
	obj := NewListObj(p)
	obj.AppendStringList(list)
	return p.SetVarObj(name, (*Obj)(obj), global)
}

func (p *Interp) AppendStringListList(name string, list []string, global bool) error {
	cname, err := syscall.BytePtrFromString(name)
	if err != nil {
		return err
	}
	var flag int32 = TCL_LEAVE_ERR_MSG | TCL_APPEND_VALUE | TCL_LIST_ELEMENT
	if global {
		flag |= TCL_GLOBAL_ONLY
	}
	for _, value := range list {
		cvalue, _ := syscall.BytePtrFromString(value)
		Tcl_SetVar(p.interp, cname, cvalue, flag)
	}
	return nil
}

func (p *Interp) AppendStringList(name string, value string, global bool) error {
	cname, err := syscall.BytePtrFromString(name)
	if err != nil {
		return err
	}
	cvalue, err := syscall.BytePtrFromString(value)
	if err != nil {
		return err
	}
	var flag int32 = TCL_LEAVE_ERR_MSG | TCL_APPEND_VALUE | TCL_LIST_ELEMENT
	if global {
		flag |= TCL_GLOBAL_ONLY
	}
	r := Tcl_SetVar(p.interp, cname, cvalue, flag)
	if r == nil {
		return p.GetErrorResult()
	}
	return nil
}

func (p *Interp) SetVarObj(name string, obj *Obj, global bool) error {
	if obj == nil {
		return os.ErrInvalid
	}
	cname, err := syscall.BytePtrFromString(name)
	if err != nil {
		return err
	}
	var flag int32 = TCL_LEAVE_ERR_MSG
	if global {
		flag |= TCL_GLOBAL_ONLY
	}
	r := Tcl_SetVar2Ex(p.interp, cname, nil, obj.obj, flag)
	if r == nil {
		return p.GetErrorResult()
	}
	return nil
}

func (p *Interp) SetVarListObj(name string, obj *ListObj, global bool) error {
	return p.SetVarObj(name, (*Obj)(obj), global)
}

func (p *Interp) SetStringVar(name string, value string, global bool) error {
	cname, err := syscall.BytePtrFromString(name)
	if err != nil {
		return err
	}
	cvalue, err := syscall.BytePtrFromString(value)
	if err != nil {
		return err
	}
	var flag int32 = TCL_LEAVE_ERR_MSG
	if global {
		flag |= TCL_GLOBAL_ONLY
	}
	r := Tcl_SetVar(p.interp, cname, cvalue, flag)
	if r == nil {
		return p.GetErrorResult()
	}
	return nil
}

func (p *Interp) AppendStringVar(name string, value string, global bool) error {
	cname, err := syscall.BytePtrFromString(name)
	if err != nil {
		return err
	}
	cvalue, err := syscall.BytePtrFromString(value)
	if err != nil {
		return err
	}
	var flag int32 = TCL_LEAVE_ERR_MSG | TCL_APPEND_VALUE
	if global {
		flag |= TCL_GLOBAL_ONLY
	}
	r := Tcl_SetVar(p.interp, cname, cvalue, flag)
	if r == nil {
		return p.GetErrorResult()
	}
	return nil
}

func (p *Interp) UnsetVar(name string, global bool) error {
	cname, err := syscall.BytePtrFromString(name)
	if err != nil {
		return err
	}
	var flag int32 = TCL_LEAVE_ERR_MSG
	if global {
		flag |= TCL_GLOBAL_ONLY
	}
	r := Tcl_UnsetVar(p.interp, cname, flag)
	if r != TCL_OK {
		return p.GetErrorResult()
	}
	return nil
}

type Obj struct {
	obj    *Tcl_Obj
	interp *Tcl_Interp
}

func NewRawObj(obj *Tcl_Obj, interp *Tcl_Interp) *Obj {
	return &Obj{obj, interp}
}

func (o *Obj) ToFloat64() float64 {
	var out Tcl_Double
	status := Tcl_GetDoubleFromObj(o.interp, o.obj, &out)
	if status == TCL_OK {
		return float64(out)
	}
	return 0
}

func (o *Obj) ToInt64() int64 {
	var out Tcl_WideInt
	status := Tcl_GetWideIntFromObj(o.interp, o.obj, &out)
	if status == TCL_OK {
		return int64(out)
	}
	return 0
}

func (o *Obj) ToInt() int {
	return int(o.ToInt64())
}

func (o *Obj) ToUint() uint {
	return uint(o.ToInt64())
}

func (o *Obj) ToBool() bool {
	var out int32
	status := Tcl_GetBooleanFromObj(o.interp, o.obj, &out)
	if status == TCL_OK {
		return out == 1
	}
	return false
}

func (o *Obj) ToString() string {
	var n int32
	out := Tcl_GetStringFromObj(o.obj, &n)
	return BytePtrToString(out, n)
}

func NewStringObj(value string, p *Interp) *Obj {
	s, err := syscall.BytePtrFromString(value)
	if err != nil {
		return nil
	}
	return &Obj{Tcl_NewStringObj(s, int32(len(value))), p.interp}
}

//NOTE: Tcl_NewDoubleObj test error on windows
func NewFloat64Obj(value float64, p *Interp) *Obj {
	//return &Obj{Tcl_NewDoubleObj(Tcl_Double(value)), p.interp}
	return NewStringObj(fmt.Sprintf("%v", value), p)
}

//NOTE: Tcl_NewWideIntObj test error on windows 32 bit
func NewInt64Obj(value int64, p *Interp) *Obj {
	return NewStringObj(fmt.Sprintf("%v", value), p)
	//return &Obj{Tcl_NewWideIntObj(Tcl_WideInt(value)), p.interp}
}

//NOTE: use int to string for amd64/i386
func NewIntObj(value int, p *Interp) *Obj {
	return NewStringObj(fmt.Sprintf("%v", value), p)
	//	return &Obj{Tcl_NewLongObj(value), p.interp}
}

func NewBoolObj(value bool, p *Interp) *Obj {
	if value {
		return &Obj{Tcl_NewBooleanObj(1), p.interp}
	} else {
		return &Obj{Tcl_NewBooleanObj(0), p.interp}
	}
}

func objToString(interp *Tcl_Interp, obj *Tcl_Obj) string {
	var n int32
	out := Tcl_GetStringFromObj(obj, &n)
	return BytePtrToString(out, n)
	//return C.GoStringN((*C.char)(unsafe.Pointer(out)), (C.int)(n))
}

func stringToObj(value string) *Tcl_Obj {
	s, err := syscall.BytePtrFromString(value)
	if err != nil {
		return nil
	}
	return Tcl_NewStringObj(s, int32(len(value)))
}

type ListObj Obj

func NewListObj(p *Interp) *ListObj {
	o := Tcl_NewListObj(0, nil)
	return &ListObj{o, p.interp}
}

func (o *ListObj) Length() int {
	var length int32
	Tcl_ListObjLength(o.interp, o.obj, &length)
	return int(length)
}

func (o *ListObj) IndexObj(index int) *Obj {
	var obj *Tcl_Obj
	r := Tcl_ListObjIndex(o.interp, o.obj, int32(index), &obj)
	if r != TCL_OK || obj == nil {
		return nil
	}
	return &Obj{obj, o.interp}
}

func (o *ListObj) IndexString(index int) string {
	var obj *Tcl_Obj
	r := Tcl_ListObjIndex(o.interp, o.obj, int32(index), &obj)
	if r != TCL_OK || obj == nil {
		return ""
	}
	return objToString(o.interp, obj)
}

func (o *ListObj) ToObjList() (list []*Obj) {
	var objs **Tcl_Obj
	var objnum int32
	Tcl_ListObjGetElements(o.interp, o.obj, &objnum, &objs)
	if objnum == 0 {
		return
	}
	lst := (*[1 << 20]*Tcl_Obj)(unsafe.Pointer(objs))[:int(objnum):int(objnum)]
	for _, v := range lst {
		list = append(list, &Obj{v, o.interp})
	}
	return
}

func (o *ListObj) ToStringList() (list []string) {
	var objs **Tcl_Obj
	var objnum int32
	Tcl_ListObjGetElements(o.interp, o.obj, &objnum, &objs)
	if objnum == 0 {
		return
	}
	lst := (*[1 << 20]*Tcl_Obj)(unsafe.Pointer(objs))[:int(objnum):int(objnum)]
	var n int32
	for _, obj := range lst {
		out := Tcl_GetStringFromObj(obj, &n)
		list = append(list, BytePtrToString(out, n))
	}
	return
}

func (o *ListObj) ToIntList() (list []int) {
	var objs **Tcl_Obj
	var objnum int32
	Tcl_ListObjGetElements(o.interp, o.obj, &objnum, &objs)
	if objnum == 0 {
		return
	}
	lst := (*[1 << 20]*Tcl_Obj)(unsafe.Pointer(objs))[:int(objnum):int(objnum)]
	var out Tcl_WideInt
	for _, obj := range lst {
		Tcl_GetWideIntFromObj(o.interp, obj, &out)
		list = append(list, int(out))
	}
	return
}

func (o *ListObj) SetStringList(list []string) {
	Tcl_SetListObj(o.obj, 0, nil)
	o.AppendStringList(list)
}

func (o *ListObj) AppendStringList(list []string) {
	for _, v := range list {
		cs, _ := syscall.BytePtrFromString(v)
		obj := Tcl_NewStringObj(cs, int32(len(v)))
		Tcl_ListObjAppendElement(o.interp, o.obj, obj)
	}
}

func (o *ListObj) AppendObj(obj *Obj) bool {
	if obj == nil {
		return false
	}
	Tcl_ListObjAppendElement(o.interp, o.obj, obj.obj)
	return true
}

func (o *ListObj) AppendString(s string) {
	Tcl_ListObjAppendElement(o.interp, o.obj, stringToObj(s))
}

func (o *ListObj) InsertObj(index int, obj *Obj) {
	Tcl_ListObjReplace(o.interp, o.obj, int32(index), 0, 1, &obj.obj)
}

func (o *ListObj) InsertString(index int, s string) {
	obj := stringToObj(s)
	Tcl_ListObjReplace(o.interp, o.obj, int32(index), 0, 1, &obj)
}

func (o *ListObj) SetIndexObj(index int, obj *Obj) bool {
	if obj == nil {
		return false
	}
	Tcl_ListObjReplace(o.interp, o.obj, int32(index), 1, 1, &obj.obj)
	return true
}

func (o *ListObj) SetIndexString(index int, s string) {
	obj := stringToObj(s)
	Tcl_ListObjReplace(o.interp, o.obj, int32(index), 1, 1, &obj)
}

func (o *ListObj) Remove(first int, count int) {
	Tcl_ListObjReplace(o.interp, o.obj, int32(first), int32(count), 0, nil)
}

type Photo struct {
	handle *Tk_PhotoHandle
	interp *Interp
}

func FindPhoto(interp *Interp, imageName string) *Photo {
	cs, err := syscall.BytePtrFromString(imageName)
	if err != nil {
		return nil
	}
	handle := Tk_FindPhoto(interp.interp, cs)
	if handle == nil {
		return nil
	}
	return &Photo{handle, interp}
}

func (p *Photo) Blank() {
	Tk_PhotoBlank(p.handle)
}

func (p *Photo) SetSize(width int, height int) error {
	status := Tk_PhotoSetSize(p.interp.interp, p.handle, int32(width), int32(height))
	if status != TCL_OK {
		return p.interp.GetErrorResult()
	}
	return nil
}

func (p *Photo) Size() (int, int) {
	var width, height int32
	Tk_PhotoGetSize(p.handle, &width, &height)
	return int(width), int(height)
}

func (p *Photo) Expand(width int, height int) error {
	status := Tk_PhotoExpand(p.interp.interp, p.handle, int32(width), int32(height))
	if status != TCL_OK {
		return p.interp.GetErrorResult()
	}
	return nil
}

func (p *Photo) ToImage() image.Image {
	var block Tk_PhotoImageBlock
	Tk_PhotoGetImage(p.handle, &block)
	if block.width == 0 || block.height == 0 {
		return nil
	}
	r := image.Rect(0, 0, int(block.width), int(block.height))
	//pix := GoBytes(unsafe.Pointer(block.pixelPtr), int(4*block.width*block.height))
	//pix := make([]uint8
	img := image.NewNRGBA(r)
	data := (*([1 << 20]byte))(unsafe.Pointer(block.pixelPtr))[:4*block.width*block.height]
	copy(img.Pix, data)
	return img
}

const (
	TK_PHOTO_COMPOSITE_OVERLAY = 0
	TK_PHOTO_COMPOSITE_SET     = 1
)

func (p *Photo) PutImage(img image.Image) error {
	dstImage, ok := img.(*image.NRGBA)
	if !ok {
		dstImage = image.NewNRGBA(img.Bounds())
		draw.Draw(dstImage, dstImage.Bounds(), img, img.Bounds().Min, draw.Src)
	}
	block := Tk_PhotoImageBlock{
		&dstImage.Pix[0],
		int32(dstImage.Rect.Max.X),
		int32(dstImage.Rect.Max.Y),
		int32(dstImage.Stride),
		4,
		[...]int32{0, 1, 2, 3},
	}
	status := Tk_PhotoPutBlock(p.interp.interp, p.handle, &block, 0, 0,
		int32(dstImage.Rect.Max.X), int32(dstImage.Rect.Max.Y),
		TK_PHOTO_COMPOSITE_SET)
	if status != TCL_OK {
		return p.interp.GetErrorResult()
	}
	return nil
}

func (p *Photo) PutZoomedImage(img image.Image, zoomX, zoomY, subsampleX, subsampleY int) error {
	dstImage, ok := img.(*image.NRGBA)
	if !ok {
		dstImage = image.NewNRGBA(img.Bounds())
		draw.Draw(dstImage, dstImage.Bounds(), img, img.Bounds().Min, draw.Src)
	}

	block := Tk_PhotoImageBlock{
		&dstImage.Pix[0],
		int32(dstImage.Rect.Max.X),
		int32(dstImage.Rect.Max.Y),
		int32(dstImage.Stride),
		4,
		[...]int32{0, 1, 2, 3},
	}
	status := Tk_PhotoPutZoomedBlock(p.interp.interp, p.handle, &block,
		0, 0, int32(dstImage.Rect.Max.X), int32(dstImage.Rect.Max.Y),
		int32(zoomX), int32(zoomY), int32(subsampleX), int32(subsampleY),
		TK_PHOTO_COMPOSITE_SET)
	if status != TCL_OK {
		return p.interp.GetErrorResult()
	}
	return nil
}
