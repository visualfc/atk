// Copyright 2017 visualfc. All rights reserved.

package tk

import (
	"fmt"
)

type WindowInfo struct {
	X      int
	Y      int
	Width  int
	Height int
}

var (
	globalWindowInfoMap = make(map[string]*WindowInfo)
)

func init() {
	globalWindowInfoMap["."] = &WindowInfo{0, 0, 200, 200}
}

type Window struct {
	id string
}

func (w *Window) Id() string {
	return w.id
}

func (w *Window) SetTitle(title string) *Window {
	eval(fmt.Sprintf("wm title %v {%v}", w.id, title))
	return w
}

func (w *Window) Title() string {
	s, _ := evalAsString(fmt.Sprintf("wm title %v", w.id))
	return s
}

func (w *Window) SetAlpha(alpha float64) *Window {
	eval(fmt.Sprintf("wm attributes %v -alpha {%v}", w.id, alpha))
	return w
}

func (w *Window) Alpha() float64 {
	r, _ := evalAsFloat64(fmt.Sprintf("wm attributes %v -alpha", w.id))
	return r
}

func (w *Window) SetFullScreen(full bool) *Window {
	eval(fmt.Sprintf("wm attributes %v -fullscreen %v", w.id, boolToInt(full)))
	return w
}

func (w *Window) IsFullScreen() bool {
	r, _ := evalAsBool(fmt.Sprintf("wm attributes %v -fullscreen", w.id))
	return r
}

func (w *Window) SetTopmost(full bool) *Window {
	eval(fmt.Sprintf("wm attributes %v -topmost %v", w.id, boolToInt(full)))
	return w
}

func (w *Window) IsTopmost() bool {
	r, _ := evalAsBool(fmt.Sprintf("wm attributes %v -topmost", w.id))
	return r
}

func (w *Window) SetGeometry(x int, y int, width int, height int) *Window {
	globalWindowInfoMap[w.id] = &WindowInfo{x, y, width, height}
	eval(fmt.Sprintf("wm geometry %v %vx%v+%v+%v", w.id, width, height, x, y))
	return w
}

func (w *Window) Geometry() (x int, y int, width int, height int) {
	if !w.IsVisible() {
		if info, ok := globalWindowInfoMap[w.id]; ok {
			return info.X, info.Y, info.Width, info.Height
		}
	}
	s, err := evalAsString(fmt.Sprintf("update\nwm geometry %v", w.id))
	if err != nil {
		return
	}
	var ar []*int = []*int{&width, &height, &x, &y}
	var n *int = ar[0]
	var index int
	for _, r := range s {
		if r == 'x' || r == '+' {
			index++
			n = ar[index]
		} else {
			*n = *n*10 + int(r-'0')
		}
	}
	return
}

func (w *Window) Move(x int, y int) *Window {
	return w.SetPos(x, y)
}

func (w *Window) SetPos(x int, y int) *Window {
	globalWindowInfoMap[w.id].X = x
	globalWindowInfoMap[w.id].Y = y
	eval(fmt.Sprintf("wm geometry %v +%v+%v", w.id, x, y))
	return w
}

func (w *Window) Pos() (x int, y int) {
	x, y, _, _ = w.Geometry()
	return
}

func (w *Window) Resize(width int, height int) *Window {
	return w.SetSize(width, height)
}

func (w *Window) SetSize(width int, height int) *Window {
	globalWindowInfoMap[w.id].Width = width
	globalWindowInfoMap[w.id].Height = height
	eval(fmt.Sprintf("wm geometry %v %vx%v", w.id, width, height))
	return w
}

func (w *Window) Size() (width int, height int) {
	_, _, width, height = w.Geometry()
	return
}

func (w *Window) SetWidth(width int) *Window {
	_, _, _, height := w.Geometry()
	w.SetSize(width, height)
	return w
}

func (w *Window) Width() (width int) {
	_, _, width, _ = w.Geometry()
	return
}

func (w *Window) SetHeight(height int) *Window {
	_, _, width, _ := w.Geometry()
	w.SetSize(width, height)
	return w
}

func (w *Window) Height() (height int) {
	_, _, _, height = w.Geometry()
	return
}

func (w *Window) SetNaturalSize() *Window {
	eval(fmt.Sprintf("wm geometry %v {}", w.id))
	return w
}

func (w *Window) SetResizable(enableWidth bool, enableHeight bool) *Window {
	eval(fmt.Sprintf("wm resizable %v %v %v", w.id, boolToInt(enableWidth), boolToInt(enableHeight)))
	return w
}

func (w *Window) IsResizable() (enableWidth bool, enableHeight bool) {
	s, err := evalAsString(fmt.Sprintf("wm resizable %v", w.id))
	if err == nil {
		n1, n2 := parserTwoInt(s)
		enableWidth = n1 != 0
		enableHeight = n2 != 0
	}
	return
}

func (w *Window) Iconify() *Window {
	eval(fmt.Sprintf("wm iconify %v", w.id))
	return w
}

func (w *Window) IsIconify() bool {
	r, _ := evalAsString(fmt.Sprintf("wm state %v", w.id))
	return r == "iconic"
}

func (w *Window) ShowNormal() *Window {
	if w.IsFullScreen() {
		w.SetFullScreen(false)
	}
	eval(fmt.Sprintf("wm state %v normal", w.id))
	return w
}

func (w *Window) ShowFullScreen() *Window {
	return w.SetFullScreen(true)
}

func (w *Window) ShowMinimized() *Window {
	return w.Iconify()
}

func (w *Window) IsMinimized() bool {
	return w.IsIconify()
}

func (w *Window) Hide() *Window {
	eval(fmt.Sprintf("wm state %v withdrawn", w.id))
	return w
}

func (w *Window) IsVisible() bool {
	s, _ := evalAsString(fmt.Sprintf("wm state %v", w.id))
	return s != "withdrawn"
}

func (w *Window) SetVisible(b bool) *Window {
	if w.IsVisible() != b {
		if b {
			w.ShowNormal()
		} else {
			w.Hide()
		}
	}
	return w
}

func (w *Window) Deiconify() *Window {
	eval(fmt.Sprintf("wm deiconify %v", w.id))
	return w
}

func (w *Window) SetMaximumSize(width int, height int) *Window {
	eval(fmt.Sprintf("wm maxsize %v %v %v", w.id, width, height))
	return w
}

func (w *Window) MaximumSize() (int, int) {
	s, _ := evalAsString(fmt.Sprintf("wm maxsize %v", w.id))
	return parserTwoInt(s)
}

func (w *Window) SetMinimumSize(width int, height int) *Window {
	eval(fmt.Sprintf("wm minsize %v %v %v", w.id, width, height))
	return w
}

func (w *Window) MinimumSize() (int, int) {
	s, _ := evalAsString(fmt.Sprintf("wm minsize %v", w.id))
	return parserTwoInt(s)
}

func (w *Window) ScreenSize() (width int, height int) {
	width, _ = evalAsInt(fmt.Sprintf("winfo screenwidth %v", w.id))
	height, _ = evalAsInt(fmt.Sprintf("winfo screenheight %v", w.id))
	return
}

func (w *Window) Center() *Window {
	sw, sh := w.ScreenSize()
	width, height := w.Size()
	x := (sw - width) / 2
	y := (sh - height) / 2
	return w.Move(x, y)
}

func (w *Window) OnClose(fn func()) error {
	fnName, err := mainInterp.CreateActionByType("window_close", func() {
		if fn != nil {
			fn()
		} else {
			w.Destroy()
		}
	})
	if err != nil {
		return err
	}
	return eval(fmt.Sprintf("wm protocol %v WM_DELETE_WINDOW %v", w.id, fnName))
}

func (w *Window) Destroy() error {
	return eval(fmt.Sprintf("destroy %v", w.id))
}

func MainWindow() *Window {
	return mainWindow
}

func NewWindow(id string) *Window {
	return &Window{id}
}
