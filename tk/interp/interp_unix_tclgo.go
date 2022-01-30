// Copyright 2018 visualfc. All rights reserved.

// +build !windows,tclgo !windows,!cgo

package interp

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"os"
	"unsafe"

	"modernc.org/libc"
	gotcl "modernc.org/tcl"
	tcl "modernc.org/tcl/lib"
	gotk "modernc.org/tk"
	tklib "modernc.org/tk/lib"
)

var (
	mainLoopThreadId tcl.Tcl_ThreadId
	mainLoopTLS      *libc.TLS
)

func create_async_event(tls *libc.TLS) uintptr {
	ev := tcl.XTcl_Alloc(tls, uint32(unsafe.Sizeof(tcl.Tcl_Event{})))
	(*tcl.Tcl_Event)(unsafe.Pointer(ev)).Fproc = *(*uintptr)(unsafe.Pointer(&struct {
		f func(tls *libc.TLS, ev uintptr, flags int32) int32
	}{async_event_handler}))
	(*tcl.Tcl_Event)(unsafe.Pointer(ev)).FnextPtr = 0
	return ev
}

func send_async_event(tls *libc.TLS, tid, ev uintptr) {
	tcl.XTcl_ThreadQueueEvent(tls, tid, ev, tcl.TCL_QUEUE_TAIL)
	tcl.XTcl_ThreadAlert(tls, tid)
}

func tcl_objcmd_proc(tls *libc.TLS, clientData, interp uintptr, objc int32, objv uintptr) int32 {
	var args []string
	p := objv
	for i := int32(0); i < objc; i++ {
		args = append(args, objToString(tls, interp, *(*uintptr)(unsafe.Pointer(p))))
		p += unsafe.Sizeof(uintptr(0))
	}
	result, err := globalCommandMap.Invoke(clientData, args[1:])
	if err != nil {
		cs := toCString(err.Error())
		defer libc.Xfree(tls, cs)
		tcl.XTcl_WrongNumArgs(tls, interp, objc, objv, cs)
		return tcl.TCL_ERROR
	}
	if result != "" {
		tcl.XTcl_SetObjResult(tls, interp, stringToObj(tls, result))
	}
	return tcl.TCL_OK
}

func toCString(s string) uintptr {
	p, err := libc.CString(s)
	if err != nil {
		panic("OOM")
	}

	return p
}

func tcl_deletecmd_proc(tls *libc.TLS, clientData uintptr) {
	globalCommandMap.UnRegister(clientData)
}

func tcl_actioncmd_proc(tls *libc.TLS, clientData, interp uintptr, objc int32, objv uintptr) int32 {
	var args []string
	p := objv
	for i := int32(0); i < objc; i++ {
		args = append(args, objToString(tls, interp, *(*uintptr)(unsafe.Pointer(p))))
		p += unsafe.Sizeof(uintptr(0))
	}
	err := globalActionMap.Invoke(clientData, args[1:])
	if err != nil {
		cs := toCString(err.Error())
		defer libc.Xfree(tls, cs)
		tcl.XTcl_WrongNumArgs(tls, interp, objc, objv, cs)
		return tcl.TCL_ERROR
	}
	return tcl.TCL_OK
}

func tcl_deleteaction_proc(tls *libc.TLS, clientData uintptr) {
	globalActionMap.UnRegister(clientData)
}

func async_event_handler(tls *libc.TLS, ev uintptr, flags int32) int32 {
	if flags != tklib.TK_ALL_EVENTS {
		return 0
	}

	if fn, ok := globalAsyncEvent.Load(ev); ok {
		fn.(func())()
		globalAsyncEvent.Delete(ev)
	}
	return 1
}

func IsMainThread() bool {
	return tcl.XTcl_GetCurrentThread(mainLoopTLS) == mainLoopThreadId
}

func async_send_event(tls *libc.TLS, tid uintptr, fn func()) {
	ev := create_async_event(tls)
	globalAsyncEvent.Store(ev, fn)
	send_async_event(tls, tid, ev)
}

func Async(fn func()) {
	if fn == nil {
		return
	}

	async_send_event(mainLoopTLS, mainLoopThreadId, fn)
}

func MainLoop(fn func()) {
	mainLoopThreadId = tcl.XTcl_GetCurrentThread(mainLoopTLS)
	if fn != nil {
		fn()
	}
	tklib.XTk_MainLoop(mainLoopTLS)
	mainLoopThreadId = 0
}

type Interp struct {
	interp uintptr
	tls    *libc.TLS
}

func NewInterp() (*Interp, error) {
	tls := libc.NewTLS()
	interp := tcl.XTcl_CreateInterp(tls)
	if interp == 0 {
		return nil, errors.New("Tcl_CreateInterp failed")
	}

	mainLoopTLS = tls

	return &Interp{interp: interp, tls: tls}, nil
}

func (p *Interp) SupportTk86() bool { return true }

func (p *Interp) InitTcl(tcl_library string) (err error) {
	if tcl_library == "" {
		tcl_library, err = gotcl.MountLibraryVFS()
		if err != nil {
			tcl_library = ""
		}
	}
	if tcl_library != "" {
		p.Eval(fmt.Sprintf("set tcl_library {%s}", tcl_library))
	}
	if tcl.XTcl_Init(p.tls, p.interp) != tcl.TCL_OK {
		err := errors.New("Tcl_Init failed:\n" + p.GetStringResult())
		return err
	}

	return nil
}

func (p *Interp) InitTk(tk_library string) (err error) {
	if tk_library == "" {
		tk_library, err = gotk.MountLibraryVFS()
		if err != nil {
			tk_library = ""
		}
	}
	if tk_library != "" {
		p.Eval(fmt.Sprintf("set tk_library {%s}", tk_library))
	}
	if tklib.XTk_Init(p.tls, p.interp) != tcl.TCL_OK {
		err := errors.New("Tk_Init failed:\n" + p.GetStringResult())
		return err
	}

	return nil
}

func (p *Interp) Destroy() error {
	if p == nil || p.interp == 0 {
		return os.ErrInvalid
	}

	tcl.XTcl_DeleteInterp(p.tls, p.interp)
	p.interp = 0
	p.tls.Close()
	p.tls = nil
	return nil
}

func (p *Interp) GetObjResult() *Obj {
	return &Obj{tcl.XTcl_GetObjResult(p.tls, p.interp), p}
}

func (p *Interp) GetListObjResult() *ListObj {
	return &ListObj{tcl.XTcl_GetObjResult(p.tls, p.interp), p}
}

func (p *Interp) Eval(script string) error {
	cs := toCString(script)
	defer libc.Xfree(p.tls, cs)
	if tcl.XTcl_EvalEx(p.tls, p.interp, cs, int32(len(script)), 0) != tcl.TCL_OK {
		err := errors.New(p.GetStringResult())
		return err
	}

	return nil
}

func (p *Interp) CreateCommand(name string, fn func([]string) (string, error)) (uintptr, error) {
	cs := toCString(name)
	defer libc.Xfree(p.tls, cs)
	id := globalCommandMap.Register(fn)
	cmd := tcl.XTcl_CreateObjCommand(
		p.tls,
		p.interp,
		cs,
		*(*uintptr)(unsafe.Pointer(&struct {
			f func(tls *libc.TLS, clientData, interp uintptr, objc int32, objv uintptr) int32
		}{tcl_objcmd_proc})),
		id,
		*(*uintptr)(unsafe.Pointer(&struct {
			f func(tls *libc.TLS, clientData uintptr)
		}{tcl_deletecmd_proc})),
	)
	if cmd == 0 {
		err := fmt.Errorf("CreateCommand %v failed", name)
		return 0, err
	}

	return id, nil
}

func (p *Interp) InvokeCommand(id uintptr, args []string) (string, error) {
	return globalCommandMap.Invoke(id, args)
}

func (p *Interp) CreateAction(name string, fn func([]string)) (uintptr, error) {
	cs := toCString(name)
	defer libc.Xfree(p.tls, cs)
	id := globalActionMap.Register(fn)
	cmd := tcl.XTcl_CreateObjCommand(
		p.tls,
		p.interp,
		cs,
		*(*uintptr)(unsafe.Pointer(&struct {
			f func(tls *libc.TLS, clientData, interp uintptr, objc int32, objv uintptr) int32
		}{tcl_actioncmd_proc})),
		id,
		*(*uintptr)(unsafe.Pointer(&struct {
			f func(tls *libc.TLS, clientData uintptr)
		}{tcl_deleteaction_proc})),
	)
	if cmd == 0 {
		err := fmt.Errorf("CreateAction %v failed", name)
		return 0, err
	}
	return id, nil
}

func (p *Interp) InvokeAction(id uintptr, args []string) error {
	return globalActionMap.Invoke(id, args)
}

func (p *Interp) GetVar(name string, global bool) *Obj {
	cname := toCString(name)
	defer libc.Xfree(p.tls, cname)
	flag := int32(tcl.TCL_LEAVE_ERR_MSG)
	if global {
		flag |= tcl.TCL_GLOBAL_ONLY
	}
	obj := tcl.XTcl_GetVar2Ex(p.tls, p.interp, cname, 0, flag)
	if obj == 0 {
		return nil
	}

	return &Obj{obj, p}
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
	tls := p.tls
	cname := toCString(name)
	defer libc.Xfree(tls, cname)
	flag := int32(tcl.TCL_LEAVE_ERR_MSG | tcl.TCL_APPEND_VALUE | tcl.TCL_LIST_ELEMENT)
	if global {
		flag |= tcl.TCL_GLOBAL_ONLY
	}
	for _, value := range list {
		cvalue := toCString(value)
		tcl.XTcl_SetVar(tls, p.interp, cname, cvalue, flag)
		libc.Xfree(tls, cvalue)
	}
	return nil
}

func (p *Interp) AppendStringList(name string, value string, global bool) error {
	tls := p.tls
	cname := toCString(name)
	defer libc.Xfree(tls, cname)
	cvalue := toCString(value)
	defer libc.Xfree(tls, cvalue)
	flag := int32(tcl.TCL_LEAVE_ERR_MSG | tcl.TCL_APPEND_VALUE | tcl.TCL_LIST_ELEMENT)
	if global {
		flag |= tcl.TCL_GLOBAL_ONLY
	}
	r := tcl.XTcl_SetVar(tls, p.interp, cname, cvalue, flag)
	if r == 0 {
		return p.GetErrorResult()
	}

	return nil
}

func (p *Interp) SetVarObj(name string, obj *Obj, global bool) error {
	if obj == nil {
		return os.ErrInvalid
	}

	cname := toCString(name)
	defer libc.Xfree(p.tls, cname)
	flag := int32(tcl.TCL_LEAVE_ERR_MSG)
	if global {
		flag |= tcl.TCL_GLOBAL_ONLY
	}
	r := tcl.XTcl_SetVar2Ex(p.tls, p.interp, cname, 0, obj.obj, flag)
	if r == 0 {
		return p.GetErrorResult()
	}

	return nil
}

func (p *Interp) SetVarListObj(name string, obj *ListObj, global bool) error {
	return p.SetVarObj(name, (*Obj)(obj), global)
}

func (p *Interp) SetStringVar(name string, value string, global bool) error {
	cname := toCString(name)
	defer libc.Xfree(p.tls, cname)
	cvalue := toCString(value)
	defer libc.Xfree(p.tls, cvalue)
	flag := int32(tcl.TCL_LEAVE_ERR_MSG)
	if global {
		flag |= tcl.TCL_GLOBAL_ONLY
	}
	r := tcl.XTcl_SetVar(p.tls, p.interp, cname, cvalue, flag)
	if r == 0 {
		return p.GetErrorResult()
	}

	return nil
}

func (p *Interp) AppendStringVar(name string, value string, global bool) error {
	cname := toCString(name)
	defer libc.Xfree(p.tls, cname)
	cvalue := toCString(value)
	defer libc.Xfree(p.tls, cvalue)
	flag := int32(tcl.TCL_LEAVE_ERR_MSG | tcl.TCL_APPEND_VALUE)
	if global {
		flag |= tcl.TCL_GLOBAL_ONLY
	}
	r := tcl.XTcl_SetVar(p.tls, p.interp, cname, cvalue, flag)
	if r == 0 {
		return p.GetErrorResult()
	}

	return nil
}

func (p *Interp) UnsetVar(name string, global bool) error {
	cname := toCString(name)
	defer libc.Xfree(p.tls, cname)
	flag := int32(tcl.TCL_LEAVE_ERR_MSG)
	if global {
		flag |= tcl.TCL_GLOBAL_ONLY
	}
	r := tcl.XTcl_UnsetVar(p.tls, p.interp, cname, flag)
	if r != tcl.TCL_OK {
		return p.GetErrorResult()
	}

	return nil
}

type Obj struct {
	obj    uintptr
	interp *Interp
}

func (o *Obj) ToFloat64() float64 {
	tls := o.interp.tls
	p := tls.Alloc(8)
	defer tls.Free(8)
	status := tcl.XTcl_GetDoubleFromObj(tls, o.interp.interp, o.obj, p)
	if status == tcl.TCL_OK {
		return *(*float64)(unsafe.Pointer(p))
	}

	return 0
}

func (o *Obj) ToInt64() int64 {
	tls := o.interp.tls
	p := tls.Alloc(8)
	defer tls.Free(8)
	status := tcl.XTcl_GetWideIntFromObj(tls, o.interp.interp, o.obj, p)
	if status == tcl.TCL_OK {
		return *(*int64)(unsafe.Pointer(p))
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
	tls := o.interp.tls
	p := tls.Alloc(8)
	defer tls.Free(8)
	status := tcl.XTcl_GetBooleanFromObj(tls, o.interp.interp, o.obj, p)
	if status == tcl.TCL_OK {
		return *(*int32)(unsafe.Pointer(p)) == 1
	}

	return false
}

func (o *Obj) ToString() string {
	tls := o.interp.tls
	p := tls.Alloc(8)
	defer tls.Free(8)
	out := tcl.XTcl_GetStringFromObj(tls, o.obj, p)
	return string(libc.GoBytes(out, int(*(*int32)(unsafe.Pointer(p)))))
}

func NewStringObj(value string, p *Interp) *Obj {
	tls := p.tls
	cs := toCString(value)
	defer libc.Xfree(tls, cs)
	return &Obj{tcl.XTcl_NewStringObj(tls, cs, int32(len(value))), p}
}

func NewFloat64Obj(value float64, p *Interp) *Obj {
	return &Obj{tcl.XTcl_NewDoubleObj(p.tls, value), p}
}

func NewInt64Obj(value int64, p *Interp) *Obj {
	return &Obj{tcl.XTcl_NewWideIntObj(p.tls, value), p}
}

func NewIntObj(value int, p *Interp) *Obj {
	return &Obj{tcl.XTcl_NewWideIntObj(p.tls, tcl.Tcl_WideInt(value)), p}
}

func NewBoolObj(value bool, p *Interp) *Obj {
	if value {
		return &Obj{tcl.XTcl_NewBooleanObj(p.tls, 1), p}
	} else {
		return &Obj{tcl.XTcl_NewBooleanObj(p.tls, 0), p}
	}
}

func objToString(tls *libc.TLS, interp uintptr, obj uintptr) string {
	p := tls.Alloc(8)
	defer tls.Free(8)
	out := tcl.XTcl_GetStringFromObj(tls, obj, p)
	return string(libc.GoBytes(out, int(*(*int32)(unsafe.Pointer(p)))))
}

func stringToObj(tls *libc.TLS, value string) uintptr {
	cs := toCString(value)
	defer libc.Xfree(tls, cs)
	return tcl.XTcl_NewStringObj(tls, cs, int32(len(value)))
}

type ListObj Obj

func NewListObj(p *Interp) *ListObj {
	o := tcl.XTcl_NewListObj(p.tls, 0, 0)
	return &ListObj{o, p}
}

func (o *ListObj) Length() int {
	tls := o.interp.tls
	p := tls.Alloc(8)
	defer tls.Free(8)
	tcl.XTcl_ListObjLength(tls, o.interp.interp, o.obj, p)
	return int(*(*int32)(unsafe.Pointer(p)))
}

func (o *ListObj) IndexObj(index int) *Obj {
	tls := o.interp.tls
	p := tls.Alloc(8)
	defer tls.Free(8)
	r := tcl.XTcl_ListObjIndex(tls, o.interp.interp, o.obj, int32(index), p)
	if r != tcl.TCL_OK || p == 0 {
		return nil
	}
	obj := *(*uintptr)(unsafe.Pointer(p))
	return &Obj{obj, o.interp}
}

func (o *ListObj) IndexString(index int) string {
	tls := o.interp.tls
	p := tls.Alloc(8)
	defer tls.Free(8)
	r := tcl.XTcl_ListObjIndex(tls, o.interp.interp, o.obj, int32(index), p)
	if r != tcl.TCL_OK || p == 0 {
		return ""
	}

	return objToString(tls, o.interp.interp, *(*uintptr)(unsafe.Pointer(p)))
}

func (o *ListObj) ToObjList() (list []*Obj) {
	tls := o.interp.tls
	p := tls.Alloc(16)
	defer tls.Free(16)
	tcl.XTcl_ListObjGetElements(tls, o.interp.interp, o.obj, p, p+8)
	objnum := *(*int32)(unsafe.Pointer(p))
	objs := *(*uintptr)(unsafe.Pointer(p + 8))
	for i := int32(0); i < objnum; i++ {
		obj := *(*uintptr)(unsafe.Pointer(objs))
		objs += unsafe.Sizeof(uintptr(0))
		list = append(list, &Obj{obj, o.interp})
	}
	return list
}

func (o *ListObj) ToStringList() (list []string) {
	tls := o.interp.tls
	p := tls.Alloc(24)
	defer tls.Free(24)
	tcl.XTcl_ListObjGetElements(tls, o.interp.interp, o.obj, p, p+8)
	objnum := *(*int32)(unsafe.Pointer(p))
	objs := *(*uintptr)(unsafe.Pointer(p + 8))
	for i := int32(0); i < objnum; i++ {
		obj := *(*uintptr)(unsafe.Pointer(objs))
		objs += unsafe.Sizeof(uintptr(0))
		out := tcl.XTcl_GetStringFromObj(tls, obj, p+16)
		list = append(list, string(libc.GoBytes(out, int(*(*int32)(unsafe.Pointer(p + 16))))))
	}
	return list
}

func (o *ListObj) ToIntList() (list []int) {
	tls := o.interp.tls
	p := tls.Alloc(24)
	defer tls.Free(24)
	tcl.XTcl_ListObjGetElements(tls, o.interp.interp, o.obj, p, p+8)
	objnum := *(*int32)(unsafe.Pointer(p))
	objs := *(*uintptr)(unsafe.Pointer(p + 8))
	for i := int32(0); i < objnum; i++ {
		obj := *(*uintptr)(unsafe.Pointer(objs))
		objs += unsafe.Sizeof(uintptr(0))
		tcl.XTcl_GetWideIntFromObj(tls, o.interp.interp, obj, p+16)
		out := *(*int64)(unsafe.Pointer(p + 16))
		list = append(list, int(out))
	}
	return list
}

func (o *ListObj) SetStringList(list []string) {
	tcl.XTcl_SetListObj(o.interp.tls, o.obj, 0, 0)
	o.AppendStringList(list)
}

func (o *ListObj) AppendStringList(list []string) {
	tls := o.interp.tls
	for _, v := range list {
		cs := toCString(v)
		obj := tcl.XTcl_NewStringObj(tls, cs, int32(len(v)))
		tcl.XTcl_ListObjAppendElement(tls, o.interp.interp, o.obj, obj)
		libc.Xfree(tls, cs)
	}
}

func (o *ListObj) AppendObj(obj *Obj) bool {
	if obj == nil {
		return false
	}
	tls := o.interp.tls
	tcl.XTcl_ListObjAppendElement(tls, o.interp.interp, o.obj, obj.obj)
	return true
}

func (o *ListObj) AppendString(s string) {
	tls := o.interp.tls
	tcl.XTcl_ListObjAppendElement(tls, o.interp.interp, o.obj, stringToObj(tls, s))
}

func (o *ListObj) InsertObj(index int, obj *Obj) {
	tls := o.interp.tls
	tcl.XTcl_ListObjReplace(tls, o.interp.interp, o.obj, int32(index), 0, 1, obj.obj)
}

func (o *ListObj) InsertString(index int, s string) {
	tls := o.interp.tls
	p := tls.Alloc(8)
	defer tls.Free(8)
	obj := stringToObj(tls, s)
	*(*uintptr)(unsafe.Pointer(p)) = obj
	tcl.XTcl_ListObjReplace(tls, o.interp.interp, o.obj, int32(index), 0, 1, p)
}

func (o *ListObj) SetIndexObj(index int, obj *Obj) bool {
	if obj == nil {
		return false
	}

	tls := o.interp.tls
	p := tls.Alloc(8)
	defer tls.Free(8)
	*(*uintptr)(unsafe.Pointer(p)) = obj.obj
	tcl.XTcl_ListObjReplace(tls, o.interp.interp, o.obj, int32(index), 1, 1, p)
	return true
}

func (o *ListObj) SetIndexString(index int, s string) {
	tls := o.interp.tls
	p := tls.Alloc(8)
	defer tls.Free(8)
	obj := stringToObj(tls, s)
	*(*uintptr)(unsafe.Pointer(p)) = obj
	tcl.XTcl_ListObjReplace(tls, o.interp.interp, o.obj, int32(index), 1, 1, p)
}

func (o *ListObj) Remove(first int, count int) {
	tls := o.interp.tls
	tcl.XTcl_ListObjReplace(tls, o.interp.interp, o.obj, int32(first), int32(count), 0, 0)
}

type Photo struct {
	handle uintptr
	interp *Interp
}

func FindPhoto(interp *Interp, imageName string) *Photo {
	tls := interp.tls
	cs := toCString(imageName)
	defer libc.Xfree(tls, cs)
	handle := tklib.XTk_FindPhoto(tls, interp.interp, cs)
	if handle == 0 {
		return nil
	}

	return &Photo{handle, interp}
}

func (p *Photo) Blank() {
	tls := p.interp.tls
	tklib.XTk_PhotoBlank(tls, p.handle)
}

func (p *Photo) SetSize(width int, height int) error {
	status := tklib.XTk_PhotoSetSize(p.interp.tls, p.interp.interp, p.handle, int32(width), int32(height))
	if status != tcl.TCL_OK {
		return p.interp.GetErrorResult()
	}

	return nil
}

func (p *Photo) Size() (int, int) {
	tls := p.interp.tls
	q := tls.Alloc(16)
	defer tls.Free(16)
	tklib.XTk_PhotoGetSize(tls, p.handle, q, q+8)
	return int(*(*int32)(unsafe.Pointer(q))), int(*(*int32)(unsafe.Pointer(q + 8)))
}

func (p *Photo) Expand(width int, height int) error {
	tls := p.interp.tls
	status := tklib.XTk_PhotoExpand(tls, p.interp.interp, p.handle, int32(width), int32(height))
	if status != tcl.TCL_OK {
		return p.interp.GetErrorResult()
	}
	return nil
}

func (p *Photo) ToImage() image.Image {
	tls := p.interp.tls
	q := tls.Alloc(int(unsafe.Sizeof(tklib.Tk_PhotoImageBlock{})))
	defer tls.Free(int(unsafe.Sizeof(tklib.Tk_PhotoImageBlock{})))
	tklib.XTk_PhotoGetImage(tls, p.handle, q)
	block := *(*tklib.Tk_PhotoImageBlock)(unsafe.Pointer(q))
	if block.Fwidth == 0 || block.Fheight == 0 {
		return nil
	}

	r := image.Rect(0, 0, int(block.Fwidth), int(block.Fheight))
	pix := libc.GoBytes(block.FpixelPtr, int(4*block.Fwidth*block.Fheight))
	return &image.NRGBA{pix, 4 * int(block.Fwidth), r}
}

func (p *Photo) PutImage(img image.Image, tk85alphacolor color.Color) error {
	if img == nil || img.Bounds().Empty() {
		return os.ErrInvalid
	}

	tls := p.interp.tls
	dstImage, ok := img.(*image.NRGBA)
	if !ok {
		dstImage = image.NewNRGBA(img.Bounds())
		draw.Draw(dstImage, dstImage.Bounds(), img, img.Bounds().Min, draw.Src)
	}
	stride := dstImage.Stride
	pixelPtr := toCBytes(tls, dstImage.Pix)
	defer libc.Xfree(tls, pixelPtr)
	width := img.Bounds().Dx()
	height := img.Bounds().Dy()
	offset := [4]int32{0, 1, 2, 3}
	block := tklib.Tk_PhotoImageBlock{
		pixelPtr,
		int32(width),
		int32(height),
		int32(stride),
		4,
		offset,
	}
	q := tls.Alloc(int(unsafe.Sizeof(tklib.Tk_PhotoImageBlock{})))
	defer tls.Free(int(unsafe.Sizeof(tklib.Tk_PhotoImageBlock{})))
	*(*tklib.Tk_PhotoImageBlock)(unsafe.Pointer(q)) = block
	status := tklib.XTk_PhotoPutBlock(tls, p.interp.interp, p.handle, q,
		0, 0, int32(width), int32(height),
		tklib.TK_PHOTO_COMPOSITE_SET)
	if status != tcl.TCL_OK {
		return p.interp.GetErrorResult()
	}

	return nil
}

func (p *Photo) PutZoomedImage(img image.Image, zoomX, zoomY, subsampleX, subsampleY int, tk85alphacolor color.Color) error {
	if img == nil || img.Bounds().Empty() {
		return os.ErrInvalid
	}

	tls := p.interp.tls
	dstImage, ok := img.(*image.NRGBA)
	if !ok {
		dstImage = image.NewNRGBA(img.Bounds())
		draw.Draw(dstImage, dstImage.Bounds(), img, img.Bounds().Min, draw.Src)
	}
	stride := dstImage.Stride
	pixelPtr := toCBytes(tls, dstImage.Pix)
	defer libc.Xfree(tls, pixelPtr)
	width := img.Bounds().Dx()
	height := img.Bounds().Dy()
	offset := [4]int32{0, 1, 2, 3}
	block := tklib.Tk_PhotoImageBlock{
		pixelPtr,
		int32(width),
		int32(height),
		int32(stride),
		4,
		offset,
	}
	q := tls.Alloc(int(unsafe.Sizeof(tklib.Tk_PhotoImageBlock{})))
	defer tls.Free(int(unsafe.Sizeof(tklib.Tk_PhotoImageBlock{})))
	*(*tklib.Tk_PhotoImageBlock)(unsafe.Pointer(q)) = block
	status := tklib.XTk_PhotoPutZoomedBlock(tls, p.interp.interp, p.handle, q,
		0, 0, int32(width), int32(height),
		int32(zoomX), int32(zoomY), int32(subsampleX), int32(subsampleY),
		tklib.TK_PHOTO_COMPOSITE_SET)
	if status != tcl.TCL_OK {
		return p.interp.GetErrorResult()
	}

	return nil
}
