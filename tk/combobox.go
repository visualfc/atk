// Copyright 2018 visualfc. All rights reserved.

package tk

import "fmt"

// combbox
type ComboBox struct {
	BaseWidget
}

func NewComboBox(parent Widget, attributes ...*WidgetAttr) *ComboBox {
	theme := checkInitUseTheme(attributes)
	iid := makeNamedWidgetId(parent, "atk_combobox")
	attributes = append(attributes, &WidgetAttr{"textvariable", variableId(iid)})
	info := CreateWidgetInfo(iid, WidgetTypeComboBox, theme, attributes)
	if info == nil {
		return nil
	}
	w := &ComboBox{}
	w.id = iid
	w.info = info
	evalSetValue(variableId(iid), "")
	RegisterWidget(w)
	return w
}

func (w *ComboBox) Attach(id string) error {
	info, err := CheckWidgetInfo(id, WidgetTypeComboBox)
	if err != nil {
		return err
	}
	w.id = id
	w.info = info
	RegisterWidget(w)
	return nil
}

func (w *ComboBox) SetFont(font Font) error {
	if font == nil {
		return ErrInvalid
	}
	return eval(fmt.Sprintf("%v configure -font {%v}", w.id, font.Id()))
}

func (w *ComboBox) Font() Font {
	r, err := evalAsString(fmt.Sprintf("%v cget -font", w.id))
	return parserFontResult(r, err)
}

func (w *ComboBox) SetBackground(color string) error {
	setObjText("atk_tmp_text", color)
	return eval(fmt.Sprintf("%v configure -background $atk_tmp_text", w.id))
}

func (w *ComboBox) Background() string {
	r, _ := evalAsString(fmt.Sprintf("%v cget -background", w.id))
	return r
}

func (w *ComboBox) SetForground(color string) error {
	setObjText("atk_tmp_text", color)
	return eval(fmt.Sprintf("%v configure -foreground $atk_tmp_text", w.id))
}

func (w *ComboBox) Forground() string {
	r, _ := evalAsString(fmt.Sprintf("%v cget -foreground", w.id))
	return r
}

func (w *ComboBox) SetJustify(justify Justify) error {
	return eval(fmt.Sprintf("%v configure -justify {%v}", w.id, justify))
}

func (w *ComboBox) Justify() Justify {
	r, err := evalAsString(fmt.Sprintf("%v cget -justify", w.id))
	return parserJustifyResult(r, err)
}

func (w *ComboBox) SetWidth(width int) error {
	return eval(fmt.Sprintf("%v configure -width {%v}", w.id, width))
}

func (w *ComboBox) Width() int {
	r, _ := evalAsInt(fmt.Sprintf("%v cget -width", w.id))
	return r
}

func (w *ComboBox) SetHeight(height int) error {
	return eval(fmt.Sprintf("%v configure -height {%v}", w.id, height))
}

func (w *ComboBox) Height() int {
	r, _ := evalAsInt(fmt.Sprintf("%v cget -height", w.id))
	return r
}

func (w *ComboBox) SetEcho(echo string) error {
	setObjText("atk_tmp_text", echo)
	return eval(fmt.Sprintf("%v configure -show $atk_tmp_text", w.id))
}

func (w *ComboBox) Echo() string {
	r, _ := evalAsString(fmt.Sprintf("%v cget -show", w.id))
	return r
}

func (w *ComboBox) SetState(state State) error {
	return eval(fmt.Sprintf("%v configure -state {%v}", w.id, state))
}

func (w *ComboBox) State() State {
	r, err := evalAsString(fmt.Sprintf("%v cget -state", w.id))
	return parserStateResult(r, err)
}

func (w *ComboBox) SetTakeFocus(takefocus bool) error {
	return eval(fmt.Sprintf("%v configure -takefocus {%v}", w.id, boolToInt(takefocus)))
}

func (w *ComboBox) IsTakeFocus() bool {
	r, _ := evalAsBool(fmt.Sprintf("%v cget -takefocus", w.id))
	return r
}

func (w *ComboBox) SetValues(values []string) error {
	setObjTextList("atk_tmp_textlist", values)
	return eval(fmt.Sprintf("%v configure -values $atk_tmp_textlist", w.id))
}

func (w *ComboBox) Values() []string {
	r, _ := evalAsStringList(fmt.Sprintf("%v cget -values", w.id))
	return r
}

func (w *ComboBox) OnSelected(fn func()) error {
	if fn == nil {
		return ErrInvalid
	}
	w.BindEvent("<<ComboboxSelected>>", func(e *Event) {
		fn()
	})
	return nil
}

func (w *ComboBox) OnEditReturn(fn func()) error {
	if fn == nil {
		return ErrInvalid
	}
	w.BindEvent("<Return>", func(e *Event) {
		fn()
	})
	return nil
}

func (w *ComboBox) Entry() *Entry {
	return &Entry{w.BaseWidget, nil}
}

func (w *ComboBox) SetCurrentText(text string) *ComboBox {
	setObjText("atk_tmp_text", text)
	eval(fmt.Sprintf("%v set $atk_tmp_text", w.id))
	return w
}

func (w *ComboBox) CurrentText() string {
	r, _ := evalAsString(fmt.Sprintf("%v get", w.id))
	return r
}

func (w *ComboBox) SetCurrentIndex(index int) *ComboBox {
	eval(fmt.Sprintf("%v current {%v}", w.id, index))
	return w
}

func (w *ComboBox) CurrentIndex() int {
	r, _ := evalAsInt(fmt.Sprintf("%v current", w.id))
	return r
}

func ComboBoxAttrFont(font Font) *WidgetAttr {
	if font == nil {
		return nil
	}
	return &WidgetAttr{"font", font.Id()}
}

func ComboBoxAttrBackground(color string) *WidgetAttr {
	return &WidgetAttr{"background", color}
}

func ComboBoxAttrForground(color string) *WidgetAttr {
	return &WidgetAttr{"foreground", color}
}

func ComboBoxAttrJustify(justify Justify) *WidgetAttr {
	return &WidgetAttr{"justify", justify}
}

func ComboBoxAttrWidth(width int) *WidgetAttr {
	return &WidgetAttr{"width", width}
}

func ComboBoxAttrHeight(height int) *WidgetAttr {
	return &WidgetAttr{"height", height}
}

func ComboBoxAttrEcho(echo string) *WidgetAttr {
	return &WidgetAttr{"show", echo}
}

func ComboBoxAttrState(state State) *WidgetAttr {
	return &WidgetAttr{"state", state}
}

func ComboBoxAttrTakeFocus(takefocus bool) *WidgetAttr {
	return &WidgetAttr{"takefocus", boolToInt(takefocus)}
}

func ComboBoxAttrValues(values []string) *WidgetAttr {
	return &WidgetAttr{"values", values}
}
