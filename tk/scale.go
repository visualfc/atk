// Copyright 2018 visualfc. All rights reserved.

package tk

import "fmt"

// scale
type Scale struct {
	BaseWidget
	command *Command
}

func NewScale(parent Widget, orient Orient, attributes ...*WidgetAttr) *Scale {
	iid := makeNamedWidgetId(parent, "atk_separator")
	attributes = append(attributes, &WidgetAttr{"orient", orient})
	info := CreateWidgetInfo(iid, WidgetTypeScale, true, attributes)
	if info == nil {
		return nil
	}
	w := &Scale{}
	w.id = iid
	w.info = info
	RegisterWidget(w)
	return w
}

func (w *Scale) Attach(id string) error {
	info, err := CheckWidgetInfo(id, WidgetTypeScale)
	if err != nil {
		return err
	}
	w.id = id
	w.info = info
	RegisterWidget(w)
	return nil
}

func (w *Scale) SetOrient(orient Orient) error {
	return eval(fmt.Sprintf("%v configure -orient {%v}", w.id, orient))
}

func (w *Scale) Orient() Orient {
	r, err := evalAsString(fmt.Sprintf("%v cget -orient", w.id))
	return parserOrientResult(r, err)
}

func (w *Scale) SetTakeFocus(takefocus bool) error {
	return eval(fmt.Sprintf("%v configure -takefocus {%v}", w.id, boolToInt(takefocus)))
}

func (w *Scale) IsTakeFocus() bool {
	r, _ := evalAsBool(fmt.Sprintf("%v cget -takefocus", w.id))
	return r
}

func (w *Scale) SetFrom(from float64) error {
	return eval(fmt.Sprintf("%v configure -from {%v}", w.id, from))
}

func (w *Scale) From() float64 {
	r, _ := evalAsFloat64(fmt.Sprintf("%v cget -from", w.id))
	return r
}

func (w *Scale) SetTo(to float64) error {
	return eval(fmt.Sprintf("%v configure -to {%v}", w.id, to))
}

func (w *Scale) To() float64 {
	r, _ := evalAsFloat64(fmt.Sprintf("%v cget -to", w.id))
	return r
}

func (w *Scale) SetValue(value float64) error {
	return eval(fmt.Sprintf("%v configure -value {%v}", w.id, value))
}

func (w *Scale) Value() float64 {
	r, _ := evalAsFloat64(fmt.Sprintf("%v cget -value", w.id))
	return r
}

func (w *Scale) SetLength(length int) error {
	return eval(fmt.Sprintf("%v configure -length {%v}", w.id, length))
}

func (w *Scale) Length() int {
	r, _ := evalAsInt(fmt.Sprintf("%v cget -length", w.id))
	return r
}

func (w *Scale) OnCommand(fn func()) error {
	if fn == nil {
		return ErrInvalid
	}
	if w.command == nil {
		w.command = &Command{}
		bindCommand(w.id, "command", w.command.Invoke)
	}
	w.command.Bind(fn)
	return nil
}

func (w *Scale) SetRange(from, to float64) error {
	return eval(fmt.Sprintf("%v configure -from {%v} -to {%v}", w.id, from, to))
}

func (w *Scale) Range() (float64, float64) {
	return w.From(), w.To()
}

func ScaleAttrOrient(orient Orient) *WidgetAttr {
	return &WidgetAttr{"orient", orient}
}

func ScaleAttrTakeFocus(takefocus bool) *WidgetAttr {
	return &WidgetAttr{"takefocus", boolToInt(takefocus)}
}

func ScaleAttrFrom(from float64) *WidgetAttr {
	return &WidgetAttr{"from", from}
}

func ScaleAttrTo(to float64) *WidgetAttr {
	return &WidgetAttr{"to", to}
}

func ScaleAttrValue(value float64) *WidgetAttr {
	return &WidgetAttr{"value", value}
}

func ScaleAttrLength(length int) *WidgetAttr {
	return &WidgetAttr{"length", length}
}
