// Copyright 2018 visualfc. All rights reserved.

// +build !windows

package interp

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"os"
	"unsafe"
)

/*
#cgo darwin CFLAGS: -I/Library/Frameworks/Tcl.framework/Headers -I/Library/Frameworks/Tk.framework/Headers
#cgo darwin LDFLAGS: -F/Library/Frameworks -framework tcl -framework tk
#cgo linux CFLAGS: -I/usr/include/tcl
#cgo linux LDFLAGS: -ltcl -ltk -lX11 -lm -lz -ldl

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
	objs := (*(*[1 << 20]*C.Tcl_Obj)(objv))[1:objc:objc]
	var args []string
	for _, obj := range objs {
		args = append(args, objToString(interp, obj))
	}
	result, err := globalCommandMap.Invoke(uintptr(clientData), args)
	if err != nil {
		cs := C.CString(err.Error())
		defer C.free(unsafe.Pointer(cs))
		C._c_wrong_num_args(interp, 1, objv, cs)
		return TCL_ERROR
	}
	if result != "" {
		C.Tcl_SetObjResult(interp, stringToObj(result))
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
	objs := (*(*[1 << 20]*C.Tcl_Obj)(objv))[1:objc:objc]
	var args []string
	for _, obj := range objs {
		args = append(args, objToString(interp, obj))
	}
	err := globalActionMap.Invoke(uintptr(clientData), args)
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
	interp      *C.Tcl_Interp
	supportTk86 bool
}

func NewInterp() (*Interp, error) {
	interp := C.Tcl_CreateInterp()
	if interp == nil {
		return nil, errors.New("Tcl_CreateInterp failed")
	}
	return &Interp{interp, false}, nil
}

func (p *Interp) SupportTk86() bool {
	return p.supportTk86
}

func (p *Interp) InitTcl(tcl_library string) error {
	if tcl_library != "" {
		p.Eval(fmt.Sprintf("set tcl_library {%s}", tcl_library))
	}
	if C.Tcl_Init(p.interp) != TCL_OK {
		err := errors.New("Tcl_Init failed")
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
		return err
	}
	p.supportTk86 = p.TkVersion() >= "8.6"
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

func (p *Interp) GetObjResult() *Obj {
	return &Obj{C.Tcl_GetObjResult(p.interp), p.interp}
}

func (p *Interp) GetListObjResult() *ListObj {
	return &ListObj{C.Tcl_GetObjResult(p.interp), p.interp}
}

func (p *Interp) Eval(script string) error {
	cs := C.CString(script)
	defer C.free(unsafe.Pointer(cs))
	if C.Tcl_EvalEx(p.interp, cs, C.int(len(script)), 0) != TCL_OK {
		err := errors.New(p.GetStringResult())
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
		return 0, err
	}
	return id, nil
}

func (p *Interp) InvokeCommand(id uintptr, args []string) (string, error) {
	return globalCommandMap.Invoke(id, args)
}

func (p *Interp) CreateAction(name string, fn func([]string)) (uintptr, error) {
	cs := C.CString(name)
	defer C.free(unsafe.Pointer(cs))

	id := globalActionMap.Register(fn)
	cmd := C._c_create_action_command(p.interp, cs, unsafe.Pointer(id))
	if cmd == nil {
		err := fmt.Errorf("CreateAction %v failed", name)
		return 0, err
	}
	return id, nil
}

func (p *Interp) InvokeAction(id uintptr, args []string) error {
	return globalActionMap.Invoke(id, args)
}

func (p *Interp) GetVar(name string, global bool) *Obj {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	var flag C.int = C.TCL_LEAVE_ERR_MSG
	if global {
		flag |= C.TCL_GLOBAL_ONLY
	}
	obj := C.Tcl_GetVar2Ex(p.interp, cname, nil, flag)
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
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	var flag C.int = C.TCL_LEAVE_ERR_MSG | C.TCL_APPEND_VALUE | C.TCL_LIST_ELEMENT
	if global {
		flag |= C.TCL_GLOBAL_ONLY
	}
	for _, value := range list {
		cvalue := C.CString(value)
		C.Tcl_SetVar(p.interp, cname, cvalue, flag)
		C.free(unsafe.Pointer(cvalue))
	}
	return nil
}

func (p *Interp) AppendStringList(name string, value string, global bool) error {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	cvalue := C.CString(value)
	defer C.free(unsafe.Pointer(cvalue))
	var flag C.int = C.TCL_LEAVE_ERR_MSG | C.TCL_APPEND_VALUE | C.TCL_LIST_ELEMENT
	if global {
		flag |= C.TCL_GLOBAL_ONLY
	}
	r := C.Tcl_SetVar(p.interp, cname, cvalue, flag)
	if r == nil {
		return p.GetErrorResult()
	}
	return nil
}

func (p *Interp) SetVarObj(name string, obj *Obj, global bool) error {
	if obj == nil {
		return os.ErrInvalid
	}
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	var flag C.int = C.TCL_LEAVE_ERR_MSG
	if global {
		flag |= C.TCL_GLOBAL_ONLY
	}
	r := C.Tcl_SetVar2Ex(p.interp, cname, nil, obj.obj, flag)
	if r == nil {
		return p.GetErrorResult()
	}
	return nil
}

func (p *Interp) SetVarListObj(name string, obj *ListObj, global bool) error {
	return p.SetVarObj(name, (*Obj)(obj), global)
}

func (p *Interp) SetStringVar(name string, value string, global bool) error {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	cvalue := C.CString(value)
	defer C.free(unsafe.Pointer(cvalue))
	var flag C.int = C.TCL_LEAVE_ERR_MSG
	if global {
		flag |= C.TCL_GLOBAL_ONLY
	}
	r := C.Tcl_SetVar(p.interp, cname, cvalue, flag)
	if r == nil {
		return p.GetErrorResult()
	}
	return nil
}

func (p *Interp) AppendStringVar(name string, value string, global bool) error {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	cvalue := C.CString(value)
	defer C.free(unsafe.Pointer(cvalue))
	var flag C.int = C.TCL_LEAVE_ERR_MSG | C.TCL_APPEND_VALUE
	if global {
		flag |= C.TCL_GLOBAL_ONLY
	}
	r := C.Tcl_SetVar(p.interp, cname, cvalue, flag)
	if r == nil {
		return p.GetErrorResult()
	}
	return nil
}

func (p *Interp) UnsetVar(name string, global bool) error {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	var flag C.int = C.TCL_LEAVE_ERR_MSG
	if global {
		flag |= C.TCL_GLOBAL_ONLY
	}
	r := C.Tcl_UnsetVar(p.interp, cname, flag)
	if r != C.TCL_OK {
		return p.GetErrorResult()
	}
	return nil
}

type Obj struct {
	obj    *C.Tcl_Obj
	interp *C.Tcl_Interp
}

func (o *Obj) ToFloat64() float64 {
	var out C.double
	status := C.Tcl_GetDoubleFromObj(o.interp, o.obj, &out)
	if status == C.TCL_OK {
		return float64(out)
	}
	return 0
}

func (o *Obj) ToInt64() int64 {
	var out C.Tcl_WideInt
	status := C.Tcl_GetWideIntFromObj(o.interp, o.obj, &out)
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
	var out C.int
	status := C.Tcl_GetBooleanFromObj(o.interp, o.obj, &out)
	if status == C.TCL_OK {
		return out == 1
	}
	return false
}

func (o *Obj) ToString() string {
	var n C.int
	out := C.Tcl_GetStringFromObj(o.obj, &n)
	return C.GoStringN(out, n)
}

func NewStringObj(value string, p *Interp) *Obj {
	cs := C.CString(value)
	defer C.free(unsafe.Pointer(cs))
	return &Obj{C.Tcl_NewStringObj(cs, C.int(len(value))), p.interp}
}

func NewFloat64Obj(value float64, p *Interp) *Obj {
	return &Obj{C.Tcl_NewDoubleObj(C.double(value)), p.interp}
}

func NewInt64Obj(value int64, p *Interp) *Obj {
	return &Obj{C.Tcl_NewWideIntObj(C.Tcl_WideInt(value)), p.interp}
}

func NewIntObj(value int, p *Interp) *Obj {
	return &Obj{C.Tcl_NewWideIntObj(C.Tcl_WideInt(value)), p.interp}
}

func NewBoolObj(value bool, p *Interp) *Obj {
	if value {
		return &Obj{C.Tcl_NewBooleanObj(1), p.interp}
	} else {
		return &Obj{C.Tcl_NewBooleanObj(0), p.interp}
	}
}

func objToString(interp *C.Tcl_Interp, obj *C.Tcl_Obj) string {
	var n C.int
	out := C.Tcl_GetStringFromObj(obj, &n)
	return C.GoStringN(out, n)
}

func stringToObj(value string) *C.Tcl_Obj {
	cs := C.CString(value)
	defer C.free(unsafe.Pointer(cs))
	return C.Tcl_NewStringObj(cs, C.int(len(value)))
}

type ListObj Obj

func NewListObj(p *Interp) *ListObj {
	o := C.Tcl_NewListObj(0, nil)
	return &ListObj{o, p.interp}
}

func (o *ListObj) Length() int {
	var length C.int
	C.Tcl_ListObjLength(o.interp, o.obj, &length)
	return int(length)
}

func (o *ListObj) IndexObj(index int) *Obj {
	var obj *C.Tcl_Obj
	r := C.Tcl_ListObjIndex(o.interp, o.obj, C.int(index), &obj)
	if r != C.TCL_OK || obj == nil {
		return nil
	}
	return &Obj{obj, o.interp}
}

func (o *ListObj) IndexString(index int) string {
	var obj *C.Tcl_Obj
	r := C.Tcl_ListObjIndex(o.interp, o.obj, C.int(index), &obj)
	if r != C.TCL_OK || obj == nil {
		return ""
	}
	return objToString(o.interp, obj)
}

func (o *ListObj) ToObjList() (list []*Obj) {
	var objs **C.Tcl_Obj
	var objnum C.int
	C.Tcl_ListObjGetElements(o.interp, o.obj, &objnum, &objs)
	if objnum == 0 {
		return
	}
	lst := (*[1 << 28]*C.Tcl_Obj)(unsafe.Pointer(objs))[:int(objnum):int(objnum)]
	for _, v := range lst {
		list = append(list, &Obj{v, o.interp})
	}
	return
}

func (o *ListObj) ToStringList() (list []string) {
	var objs **C.Tcl_Obj
	var objnum C.int
	C.Tcl_ListObjGetElements(o.interp, o.obj, &objnum, &objs)
	if objnum == 0 {
		return
	}
	lst := (*[1 << 28]*C.Tcl_Obj)(unsafe.Pointer(objs))[:int(objnum):int(objnum)]
	var n C.int
	for _, obj := range lst {
		out := C.Tcl_GetStringFromObj(obj, &n)
		list = append(list, C.GoStringN(out, n))
	}
	return
}

func (o *ListObj) ToIntList() (list []int) {
	var objs **C.Tcl_Obj
	var objnum C.int
	C.Tcl_ListObjGetElements(o.interp, o.obj, &objnum, &objs)
	if objnum == 0 {
		return
	}
	lst := (*[1 << 28]*C.Tcl_Obj)(unsafe.Pointer(objs))[:int(objnum):int(objnum)]
	var out C.Tcl_WideInt
	for _, obj := range lst {
		C.Tcl_GetWideIntFromObj(o.interp, obj, &out)
		list = append(list, int(out))
	}
	return
}

func (o *ListObj) SetStringList(list []string) {
	C.Tcl_SetListObj(o.obj, 0, nil)
	o.AppendStringList(list)
}

func (o *ListObj) AppendStringList(list []string) {
	for _, v := range list {
		cs := C.CString(v)
		obj := C.Tcl_NewStringObj(cs, C.int(len(v)))
		C.Tcl_ListObjAppendElement(o.interp, o.obj, obj)
		C.free(unsafe.Pointer(cs))
	}
}

func (o *ListObj) AppendObj(obj *Obj) bool {
	if obj == nil {
		return false
	}
	C.Tcl_ListObjAppendElement(o.interp, o.obj, obj.obj)
	return true
}

func (o *ListObj) AppendString(s string) {
	C.Tcl_ListObjAppendElement(o.interp, o.obj, stringToObj(s))
}

func (o *ListObj) InsertObj(index int, obj *Obj) {
	C.Tcl_ListObjReplace(o.interp, o.obj, C.int(index), 0, 1, &obj.obj)
}

func (o *ListObj) InsertString(index int, s string) {
	obj := stringToObj(s)
	C.Tcl_ListObjReplace(o.interp, o.obj, C.int(index), 0, 1, &obj)
}

func (o *ListObj) SetIndexObj(index int, obj *Obj) bool {
	if obj == nil {
		return false
	}
	C.Tcl_ListObjReplace(o.interp, o.obj, C.int(index), 1, 1, &obj.obj)
	return true
}

func (o *ListObj) SetIndexString(index int, s string) {
	obj := stringToObj(s)
	C.Tcl_ListObjReplace(o.interp, o.obj, C.int(index), 1, 1, &obj)
}

func (o *ListObj) Remove(first int, count int) {
	C.Tcl_ListObjReplace(o.interp, o.obj, C.int(first), C.int(count), 0, nil)
}

type Photo struct {
	handle C.Tk_PhotoHandle
	interp *Interp
}

func FindPhoto(interp *Interp, imageName string) *Photo {
	cs := C.CString(imageName)
	defer C.free(unsafe.Pointer(cs))
	handle := C.Tk_FindPhoto(interp.interp, cs)
	if handle == nil {
		return nil
	}
	return &Photo{handle, interp}
}

func (p *Photo) Blank() {
	C.Tk_PhotoBlank(p.handle)
}

func (p *Photo) SetSize(width int, height int) error {
	status := C.Tk_PhotoSetSize(p.interp.interp, p.handle, C.int(width), C.int(height))
	if status != C.TCL_OK {
		return p.interp.GetErrorResult()
	}
	return nil
}

func (p *Photo) Size() (int, int) {
	var width, height C.int
	C.Tk_PhotoGetSize(p.handle, &width, &height)
	return int(width), int(height)
}

func (p *Photo) Expand(width int, height int) error {
	status := C.Tk_PhotoExpand(p.interp.interp, p.handle, C.int(width), C.int(height))
	if status != C.TCL_OK {
		return p.interp.GetErrorResult()
	}
	return nil
}

func (p *Photo) ToImage() image.Image {
	var block C.Tk_PhotoImageBlock
	C.Tk_PhotoGetImage(p.handle, &block)
	if block.width == 0 || block.height == 0 {
		return nil
	}
	r := image.Rect(0, 0, int(block.width), int(block.height))
	pix := C.GoBytes(unsafe.Pointer(block.pixelPtr), C.int(4*block.width*block.height))
	return &image.NRGBA{pix, 4 * int(block.width), r}
}

func (p *Photo) PutImage(img image.Image, tk85alphacolor color.Color) error {
	if img == nil || img.Bounds().Empty() {
		return os.ErrInvalid
	}
	var pixelPtr unsafe.Pointer
	var stride int
	if p.interp.supportTk86 {
		dstImage, ok := img.(*image.NRGBA)
		if !ok {
			dstImage = image.NewNRGBA(img.Bounds())
			draw.Draw(dstImage, dstImage.Bounds(), img, img.Bounds().Min, draw.Src)
		}
		stride = dstImage.Stride
		pixelPtr = toCBytes(dstImage.Pix)
	} else {
		var r, g, b uint8
		if tk85alphacolor != nil {
			clr := color.RGBAModel.Convert(tk85alphacolor).(color.RGBA)
			r, g, b = clr.R, clr.G, clr.B
		}
		dstImage := image.NewRGBA(img.Bounds())
		for i := 0; i < len(dstImage.Pix); i += 4 {
			dstImage.Pix[i+0] = r
			dstImage.Pix[i+1] = g
			dstImage.Pix[i+2] = b
			dstImage.Pix[i+3] = 0xff
		}
		draw.Draw(dstImage, dstImage.Bounds(), img, img.Bounds().Min, draw.Over)
		stride = dstImage.Stride
		pixelPtr = toCBytes(dstImage.Pix)
	}
	defer C.free(pixelPtr)
	width := img.Bounds().Dx()
	height := img.Bounds().Dy()
	offset := [4]C.int{0, 1, 2, 3}
	block := C.Tk_PhotoImageBlock{
		(*C.uchar)(pixelPtr),
		C.int(width),
		C.int(height),
		C.int(stride),
		4,
		offset,
	}
	status := C.Tk_PhotoPutBlock(p.interp.interp, p.handle, &block,
		0, 0, C.int(width), C.int(height),
		C.TK_PHOTO_COMPOSITE_SET)
	if status != C.TCL_OK {
		return p.interp.GetErrorResult()
	}
	return nil
}

func (p *Photo) PutZoomedImage(img image.Image, zoomX, zoomY, subsampleX, subsampleY int, tk85alphacolor color.Color) error {
	if img == nil || img.Bounds().Empty() {
		return os.ErrInvalid
	}
	var pixelPtr unsafe.Pointer
	var stride int
	if p.interp.supportTk86 {
		dstImage, ok := img.(*image.NRGBA)
		if !ok {
			dstImage = image.NewNRGBA(img.Bounds())
			draw.Draw(dstImage, dstImage.Bounds(), img, img.Bounds().Min, draw.Src)
		}
		stride = dstImage.Stride
		pixelPtr = toCBytes(dstImage.Pix)
	} else {
		var r, g, b uint8
		if tk85alphacolor != nil {
			clr := color.RGBAModel.Convert(tk85alphacolor).(color.RGBA)
			r, g, b = clr.R, clr.G, clr.B
		}
		dstImage := image.NewRGBA(img.Bounds())
		for i := 0; i < len(dstImage.Pix); i += 4 {
			dstImage.Pix[i+0] = r
			dstImage.Pix[i+1] = g
			dstImage.Pix[i+2] = b
			dstImage.Pix[i+3] = 0xff
		}
		draw.Draw(dstImage, dstImage.Bounds(), img, img.Bounds().Min, draw.Over)
		stride = dstImage.Stride
		pixelPtr = toCBytes(dstImage.Pix)
	}
	defer C.free(pixelPtr)
	width := img.Bounds().Dx()
	height := img.Bounds().Dy()
	offset := [4]C.int{0, 1, 2, 3}
	block := C.Tk_PhotoImageBlock{
		(*C.uchar)(pixelPtr),
		C.int(width),
		C.int(height),
		C.int(stride),
		4,
		offset,
	}
	status := C.Tk_PhotoPutZoomedBlock(p.interp.interp, p.handle, &block,
		0, 0, C.int(width), C.int(height),
		C.int(zoomX), C.int(zoomY), C.int(subsampleX), C.int(subsampleY),
		C.TK_PHOTO_COMPOSITE_SET)
	if status != C.TCL_OK {
		return p.interp.GetErrorResult()
	}
	return nil
}
