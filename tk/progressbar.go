// Copyright 2018 visualfc. All rights reserved.

package tk

import "fmt"

// progressbar
type ProgressBar struct {
	BaseWidget
}

func NewProgressBar(parent Widget, orient Orient, attributes ...*WidgetAttr) *ProgressBar {
	iid := makeNamedWidgetId(parent, "atk_progressbar")
	attributes = append(attributes, &WidgetAttr{"orient", orient})
	info := CreateWidgetInfo(iid, WidgetTypeProgressBar, true, attributes)
	if info == nil {
		return nil
	}
	w := &ProgressBar{}
	w.id = iid
	w.info = info
	RegisterWidget(w)
	return w
}

func (w *ProgressBar) Attach(id string) error {
	info, err := CheckWidgetInfo(id, WidgetTypeProgressBar)
	if err != nil {
		return err
	}
	w.id = id
	w.info = info
	RegisterWidget(w)
	return nil
}

func (w *ProgressBar) SetOrient(orient Orient) error {
	return eval(fmt.Sprintf("%v configure -orient {%v}", w.id, orient))
}

func (w *ProgressBar) Orient() Orient {
	r, err := evalAsString(fmt.Sprintf("%v cget -orient", w.id))
	return parserOrientResult(r, err)
}

func (w *ProgressBar) SetTakeFocus(takefocus bool) error {
	return eval(fmt.Sprintf("%v configure -takefocus {%v}", w.id, boolToInt(takefocus)))
}

func (w *ProgressBar) IsTakeFocus() bool {
	r, _ := evalAsBool(fmt.Sprintf("%v cget -takefocus", w.id))
	return r
}

func (w *ProgressBar) SetLength(length int) error {
	return eval(fmt.Sprintf("%v configure -length {%v}", w.id, length))
}

func (w *ProgressBar) Length() int {
	r, _ := evalAsInt(fmt.Sprintf("%v cget -length", w.id))
	return r
}

func (w *ProgressBar) SetMaximum(maximum float64) error {
	return eval(fmt.Sprintf("%v configure -maximum {%v}", w.id, maximum))
}

func (w *ProgressBar) Maximum() float64 {
	r, _ := evalAsFloat64(fmt.Sprintf("%v cget -maximum", w.id))
	return r
}

func (w *ProgressBar) SetValue(value float64) error {
	return eval(fmt.Sprintf("%v configure -value {%v}", w.id, value))
}

func (w *ProgressBar) Value() float64 {
	r, _ := evalAsFloat64(fmt.Sprintf("%v cget -value", w.id))
	return r
}

func (w *ProgressBar) Phase() int {
	r, _ := evalAsInt(fmt.Sprintf("%v cget -phase", w.id))
	return r
}

func (w *ProgressBar) SetDeterminateMode(b bool) error {
	var mode string
	if b {
		mode = "determinate"
	} else {
		mode = "indeterminate"
	}
	return eval(fmt.Sprintf("%v configure -mode %v", w.id, mode))
}

func (w *ProgressBar) IsDeterminateMode() bool {
	r, _ := evalAsString(fmt.Sprintf("%v cget -mode", w.id))
	return r == "determinate"
}

func (w *ProgressBar) Start() error {
	return w.StartEx(50)
}

func (w *ProgressBar) StartEx(ms int) error {
	return eval(fmt.Sprintf("%v start %v", w.id, ms))
}

func (w *ProgressBar) Stop() error {
	return eval(fmt.Sprintf("%v stop", w.id))
}

func (w *ProgressBar) Pause() error {
	cur := w.Value()
	w.Stop()
	return w.SetValue(cur)
}

func ProgressBarAttrOrient(orient Orient) *WidgetAttr {
	return &WidgetAttr{"orient", orient}
}

func ProgressBarAttrTakeFocus(takefocus bool) *WidgetAttr {
	return &WidgetAttr{"takefocus", boolToInt(takefocus)}
}

func ProgressBarAttrLength(length int) *WidgetAttr {
	return &WidgetAttr{"length", length}
}

func ProgressBarAttrMaximum(maximum float64) *WidgetAttr {
	return &WidgetAttr{"maximum", maximum}
}

func ProgressBarAttrValue(value float64) *WidgetAttr {
	return &WidgetAttr{"value", value}
}
