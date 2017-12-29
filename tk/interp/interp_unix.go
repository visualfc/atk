// Copyright 2017 visualfc. All rights reserved.

// +build !windows

package interp

import (
	"errors"
	"fmt"
	"image"
	"image/draw"
	"os"
	"unsafe"
)

/*
#cgo darwin CFLAGS: -I/Library/Frameworks/Tcl.framework/Headers -I/Library/Frameworks/Tk.framework/Headers
#cgo darwin LDFLAGS: -lTcl -lTk
#cgo linux CFLAGS: -I/usr/include/tcl8.6
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
	objs := (*(*[1 << 20]*C.Tcl_Obj)(objv))[1:objc]
	var args []string
	for _, obj := range objs {
		args = append(args, ObjToString(interp, obj))
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
	interp        *C.Tcl_Interp
	fnErrorHandle func(error)
}

func NewInterp() (*Interp, error) {
	interp := C.Tcl_CreateInterp()
	if interp == nil {
		return nil, errors.New("Tcl_CreateInterp failed")
	}
	return &Interp{interp, nil}, nil
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

func (p *Interp) GetObjResult() *Obj {
	return &Obj{C.Tcl_GetObjResult(p.interp), p.interp}
}

func (p *Interp) Eval(script string) error {
	cs := C.CString(script)
	defer C.free(unsafe.Pointer(cs))
	if C.Tcl_EvalEx(p.interp, cs, C.int(len(script)), 0) != TCL_OK {
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

func (p *Interp) CreateAction(name string, fn func([]string)) (uintptr, error) {
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

func (p *Interp) InvokeAction(id uintptr, args []string) error {
	return globalActionMap.Invoke(id, args)
}

type Obj struct {
	obj    *C.Tcl_Obj
	interp *C.Tcl_Interp
}

func NewRawObj(obj *C.Tcl_Obj, interp *C.Tcl_Interp) *Obj {
	return &Obj{obj, interp}
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

func (p *Photo) PutImage(img image.Image) error {
	dstImage, ok := img.(*image.NRGBA)
	if !ok {
		dstImage = image.NewNRGBA(img.Bounds())
		draw.Draw(dstImage, dstImage.Bounds(), img, img.Bounds().Min, draw.Src)
	}
	if len(dstImage.Pix) == 0 {
		return os.ErrInvalid
	}
	pixelPtr := toCBytes(dstImage.Pix)
	defer C.free(pixelPtr)

	block := C.Tk_PhotoImageBlock{
		(*C.uchar)(pixelPtr),
		C.int(dstImage.Rect.Max.X),
		C.int(dstImage.Rect.Max.Y),
		C.int(dstImage.Stride),
		4,
		[...]C.int{0, 1, 2, 0},
	}
	status := C.Tk_PhotoPutBlock(p.interp.interp, p.handle, &block, 0, 0,
		C.int(dstImage.Rect.Max.X), C.int(dstImage.Rect.Max.Y),
		C.TK_PHOTO_COMPOSITE_SET)
	if status != C.TCL_OK {
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
	if len(dstImage.Pix) == 0 {
		return os.ErrInvalid
	}

	pixelPtr := toCBytes(dstImage.Pix)
	defer C.free(pixelPtr)

	block := C.Tk_PhotoImageBlock{
		(*C.uchar)(pixelPtr),
		C.int(dstImage.Rect.Max.X),
		C.int(dstImage.Rect.Max.Y),
		C.int(dstImage.Stride),
		4,
		[...]C.int{0, 1, 2, 0},
	}
	status := C.Tk_PhotoPutZoomedBlock(p.interp.interp, p.handle, &block,
		0, 0, C.int(dstImage.Rect.Max.X), C.int(dstImage.Rect.Max.Y),
		C.int(zoomX), C.int(zoomY), C.int(subsampleX), C.int(subsampleY),
		C.TK_PHOTO_COMPOSITE_SET)
	if status != C.TCL_OK {
		return p.interp.GetErrorResult()
	}
	return nil
}
