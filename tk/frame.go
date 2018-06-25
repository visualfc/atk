// Copyright 2018 visualfc. All rights reserved.

package tk

import "fmt"

// frame
type Frame struct {
	BaseWidget
}

func NewFrame(parent Widget, attributes ...*WidgetAttr) *Frame {
	theme := checkInitUseTheme(attributes)
	iid := makeNamedWidgetId(parent, "atk_frame")
	info := CreateWidgetInfo(iid, WidgetTypeFrame, theme, attributes)
	if info == nil {
		return nil
	}
	w := &Frame{}
	w.id = iid
	w.info = info
	RegisterWidget(w)
	return w
}

func (w *Frame) Attach(id string) error {
	info, err := CheckWidgetInfo(id, WidgetTypeFrame)
	if err != nil {
		return err
	}
	w.id = id
	w.info = info
	RegisterWidget(w)
	return nil
}

func (w *Frame) SetBorderWidth(width int) error {
	return eval(fmt.Sprintf("%v configure -borderwidth {%v}", w.id, width))
}

func (w *Frame) BorderWidth() int {
	r, _ := evalAsInt(fmt.Sprintf("%v cget -borderwidth", w.id))
	return r
}

func (w *Frame) SetReliefStyle(relief ReliefStyle) error {
	return eval(fmt.Sprintf("%v configure -relief {%v}", w.id, relief))
}

func (w *Frame) ReliefStyle() ReliefStyle {
	r, err := evalAsString(fmt.Sprintf("%v cget -relief", w.id))
	return parserReliefStyleResult(r, err)
}

func (w *Frame) SetWidth(width int) error {
	return eval(fmt.Sprintf("%v configure -width {%v}", w.id, width))
}

func (w *Frame) Width() int {
	r, _ := evalAsInt(fmt.Sprintf("%v cget -width", w.id))
	return r
}

func (w *Frame) SetHeight(height int) error {
	return eval(fmt.Sprintf("%v configure -height {%v}", w.id, height))
}

func (w *Frame) Height() int {
	r, _ := evalAsInt(fmt.Sprintf("%v cget -height", w.id))
	return r
}

func (w *Frame) SetPaddingN(padx int, pady int) error {
	if w.info.IsTtk {
		return eval(fmt.Sprintf("%v configure -padding {%v %v}", w.id, padx, pady))
	}
	return eval(fmt.Sprintf("%v configure -padx {%v} -pady {%v}", w.id, padx, pady))
}

func (w *Frame) PaddingN() (int, int) {
	var r string
	var err error
	if w.info.IsTtk {
		r, err = evalAsString(fmt.Sprintf("%v cget -padding", w.id))
	} else {
		r1, _ := evalAsString(fmt.Sprintf("%v cget -padx", w.id))
		r2, _ := evalAsString(fmt.Sprintf("%v cget -pady", w.id))
		r = r1 + " " + r2
	}
	return parserPaddingResult(r, err)
}

func (w *Frame) SetPadding(pad Pad) error {
	return w.SetPaddingN(pad.X, pad.Y)
}

func (w *Frame) Padding() Pad {
	x, y := w.PaddingN()
	return Pad{x, y}
}

func (w *Frame) SetTakeFocus(takefocus bool) error {
	return eval(fmt.Sprintf("%v configure -takefocus {%v}", w.id, boolToInt(takefocus)))
}

func (w *Frame) IsTakeFocus() bool {
	r, _ := evalAsBool(fmt.Sprintf("%v cget -takefocus", w.id))
	return r
}

func FrameAttrBorderWidth(width int) *WidgetAttr {
	return &WidgetAttr{"borderwidth", width}
}

func FrameAttrReliefStyle(relief ReliefStyle) *WidgetAttr {
	return &WidgetAttr{"relief", relief}
}

func FrameAttrWidth(width int) *WidgetAttr {
	return &WidgetAttr{"width", width}
}

func FrameAttrHeight(height int) *WidgetAttr {
	return &WidgetAttr{"height", height}
}

func FrameAttrPadding(pad Pad) *WidgetAttr {
	return &WidgetAttr{"pad", pad}
}

func FrameAttrTakeFocus(takefocus bool) *WidgetAttr {
	return &WidgetAttr{"takefocus", boolToInt(takefocus)}
}
