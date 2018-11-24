// Copyright 2018 visualfc. All rights reserved.

package tk

import "fmt"

// label frame
type LabelFrame struct {
	BaseWidget
}

func NewLabelFrame(parent Widget, attributes ...*WidgetAttr) *LabelFrame {
	theme := checkInitUseTheme(attributes)
	iid := makeNamedWidgetId(parent, "atk_labelframe")
	info := CreateWidgetInfo(iid, WidgetTypeLabelFrame, theme, attributes)
	if info == nil {
		return nil
	}
	w := &LabelFrame{}
	w.id = iid
	w.info = info
	RegisterWidget(w)
	return w
}

func (w *LabelFrame) Attach(id string) error {
	info, err := CheckWidgetInfo(id, WidgetTypeLabelFrame)
	if err != nil {
		return err
	}
	w.id = id
	w.info = info
	RegisterWidget(w)
	return nil
}

func (w *LabelFrame) SetLabelText(text string) error {
	setObjText("atk_tmp_text", text)
	return eval(fmt.Sprintf("%v configure -text $atk_tmp_text", w.id))
}

func (w *LabelFrame) LabelText() string {
	r, _ := evalAsString(fmt.Sprintf("%v cget -text", w.id))
	return r
}

func (w *LabelFrame) SetLabelAnchor(anchor Anchor) error {
	return eval(fmt.Sprintf("%v configure -labelanchor {%v}", w.id, anchor))
}

func (w *LabelFrame) LabelAnchor() Anchor {
	r, err := evalAsString(fmt.Sprintf("%v cget -labelanchor", w.id))
	return parserAnchorResult(r, err)
}

func (w *LabelFrame) SetBorderWidth(width int) error {
	return eval(fmt.Sprintf("%v configure -borderwidth {%v}", w.id, width))
}

func (w *LabelFrame) BorderWidth() int {
	r, _ := evalAsInt(fmt.Sprintf("%v cget -borderwidth", w.id))
	return r
}

func (w *LabelFrame) SetReliefStyle(relief ReliefStyle) error {
	return eval(fmt.Sprintf("%v configure -relief {%v}", w.id, relief))
}

func (w *LabelFrame) ReliefStyle() ReliefStyle {
	r, err := evalAsString(fmt.Sprintf("%v cget -relief", w.id))
	return parserReliefStyleResult(r, err)
}

func (w *LabelFrame) SetWidth(width int) error {
	return eval(fmt.Sprintf("%v configure -width {%v}", w.id, width))
}

func (w *LabelFrame) Width() int {
	r, _ := evalAsInt(fmt.Sprintf("%v cget -width", w.id))
	return r
}

func (w *LabelFrame) SetHeight(height int) error {
	return eval(fmt.Sprintf("%v configure -height {%v}", w.id, height))
}

func (w *LabelFrame) Height() int {
	r, _ := evalAsInt(fmt.Sprintf("%v cget -height", w.id))
	return r
}

func (w *LabelFrame) SetPaddingN(padx int, pady int) error {
	if w.info.IsTtk {
		return eval(fmt.Sprintf("%v configure -padding {%v %v}", w.id, padx, pady))
	}
	return eval(fmt.Sprintf("%v configure -padx {%v} -pady {%v}", w.id, padx, pady))
}

func (w *LabelFrame) PaddingN() (int, int) {
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

func (w *LabelFrame) SetPadding(pad Pad) error {
	return w.SetPaddingN(pad.X, pad.Y)
}

func (w *LabelFrame) Padding() Pad {
	x, y := w.PaddingN()
	return Pad{x, y}
}

func (w *LabelFrame) SetTakeFocus(takefocus bool) error {
	return eval(fmt.Sprintf("%v configure -takefocus {%v}", w.id, boolToInt(takefocus)))
}

func (w *LabelFrame) IsTakeFocus() bool {
	r, _ := evalAsBool(fmt.Sprintf("%v cget -takefocus", w.id))
	return r
}

func LabelFrameAttrLabelText(text string) *WidgetAttr {
	return &WidgetAttr{"text", text}
}

func LabelFrameAttrLabelAnchor(anchor Anchor) *WidgetAttr {
	return &WidgetAttr{"labelanchor", anchor}
}

func LabelFrameAttrBorderWidth(width int) *WidgetAttr {
	return &WidgetAttr{"borderwidth", width}
}

func LabelFrameAttrReliefStyle(relief ReliefStyle) *WidgetAttr {
	return &WidgetAttr{"relief", relief}
}

func LabelFrameAttrWidth(width int) *WidgetAttr {
	return &WidgetAttr{"width", width}
}

func LabelFrameAttrHeight(height int) *WidgetAttr {
	return &WidgetAttr{"height", height}
}

func LabelFrameAttrPadding(pad Pad) *WidgetAttr {
	return &WidgetAttr{"pad", pad}
}

func LabelFrameAttrTakeFocus(takefocus bool) *WidgetAttr {
	return &WidgetAttr{"takefocus", boolToInt(takefocus)}
}
