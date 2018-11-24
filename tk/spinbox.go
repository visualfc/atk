// Copyright 2018 visualfc. All rights reserved.

package tk

import "fmt"

// spinbox
type SpinBox struct {
	BaseWidget
	command        *Command
	xscrollcommand *CommandEx
}

func NewSpinBox(parent Widget, attributes ...*WidgetAttr) *SpinBox {
	theme := checkInitUseTheme(attributes)
	iid := makeNamedWidgetId(parent, "atk_spinbox")
	info := CreateWidgetInfo(iid, WidgetTypeSpinBox, theme, attributes)
	if info == nil {
		return nil
	}
	w := &SpinBox{}
	w.id = iid
	w.info = info
	RegisterWidget(w)
	return w
}

func (w *SpinBox) Attach(id string) error {
	info, err := CheckWidgetInfo(id, WidgetTypeSpinBox)
	if err != nil {
		return err
	}
	w.id = id
	w.info = info
	RegisterWidget(w)
	return nil
}

func (w *SpinBox) SetTakeFocus(takefocus bool) error {
	return eval(fmt.Sprintf("%v configure -takefocus {%v}", w.id, boolToInt(takefocus)))
}

func (w *SpinBox) IsTakeFocus() bool {
	r, _ := evalAsBool(fmt.Sprintf("%v cget -takefocus", w.id))
	return r
}

func (w *SpinBox) SetFrom(from float64) error {
	return eval(fmt.Sprintf("%v configure -from {%v}", w.id, from))
}

func (w *SpinBox) From() float64 {
	r, _ := evalAsFloat64(fmt.Sprintf("%v cget -from", w.id))
	return r
}

func (w *SpinBox) SetTo(to float64) error {
	return eval(fmt.Sprintf("%v configure -to {%v}", w.id, to))
}

func (w *SpinBox) To() float64 {
	r, _ := evalAsFloat64(fmt.Sprintf("%v cget -to", w.id))
	return r
}

func (w *SpinBox) SetIncrement(increment float64) error {
	return eval(fmt.Sprintf("%v configure -increment {%v}", w.id, increment))
}

func (w *SpinBox) Increment() float64 {
	r, _ := evalAsFloat64(fmt.Sprintf("%v cget -increment", w.id))
	return r
}

func (w *SpinBox) SetWrap(wrap bool) error {
	return eval(fmt.Sprintf("%v configure -wrap {%v}", w.id, boolToInt(wrap)))
}

func (w *SpinBox) IsWrap() bool {
	r, _ := evalAsBool(fmt.Sprintf("%v cget -wrap", w.id))
	return r
}

func (w *SpinBox) SetTextValues(values []string) error {
	setObjTextList("atk_tmp_textlist", values)
	return eval(fmt.Sprintf("%v configure -values $atk_tmp_textlist", w.id))
}

func (w *SpinBox) TextValues() []string {
	r, _ := evalAsStringList(fmt.Sprintf("%v cget -values", w.id))
	return r
}

func (w *SpinBox) OnCommand(fn func()) error {
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

func (w *SpinBox) OnXScrollEx(fn func([]string) error) error {
	if fn == nil {
		return ErrInvalid
	}
	if w.xscrollcommand == nil {
		w.xscrollcommand = &CommandEx{}
		bindCommandEx(w.id, "xscrollcommand", w.xscrollcommand.Invoke)
	}
	w.xscrollcommand.Bind(fn)
	return nil
}

func (w *SpinBox) OnEditReturn(fn func()) error {
	if fn == nil {
		return ErrInvalid
	}
	w.BindEvent("<Return>", func(e *Event) {
		fn()
	})
	return nil
}

func (w *SpinBox) Entry() *Entry {
	return &Entry{w.BaseWidget, nil}
}

func (w *SpinBox) Value() float64 {
	r, _ := evalAsFloat64(fmt.Sprintf("%v get", w.id))
	return r
}

func (w *SpinBox) SetValue(value float64) error {
	return eval(fmt.Sprintf("%v set %v", w.id, value))
}

func (w *SpinBox) SetTextValue(value string) error {
	return eval(fmt.Sprintf("%v set %v", w.id, value))
}

func (w *SpinBox) TextValue() string {
	r, _ := evalAsString(fmt.Sprintf("%v get", w.id))
	return r
}

func SpinBoxAttrTakeFocus(takefocus bool) *WidgetAttr {
	return &WidgetAttr{"takefocus", boolToInt(takefocus)}
}

func SpinBoxAttrFrom(from float64) *WidgetAttr {
	return &WidgetAttr{"from", from}
}

func SpinBoxAttrTo(to float64) *WidgetAttr {
	return &WidgetAttr{"to", to}
}

func SpinBoxAttrIncrement(increment float64) *WidgetAttr {
	return &WidgetAttr{"increment", increment}
}

func SpinBoxAttrWrap(wrap bool) *WidgetAttr {
	return &WidgetAttr{"wrap", boolToInt(wrap)}
}

func SpinBoxAttrTextValues(values []string) *WidgetAttr {
	return &WidgetAttr{"values", values}
}
