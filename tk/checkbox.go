// Copyright 2018 visualfc. All rights reserved.

package tk

import "fmt"

// check button
type CheckButton struct {
	BaseWidget
	command *Command
}

func NewCheckButton(parent Widget, text string, attributes ...*WidgetAttr) *CheckButton {
	theme := checkInitUseTheme(attributes)
	iid := makeNamedWidgetId(parent, "atk_checkbutton")
	attributes = append(attributes, &WidgetAttr{"text", text})
	attributes = append(attributes, &WidgetAttr{"variable", variableId(iid)})
	info := CreateWidgetInfo(iid, WidgetTypeCheckButton, theme, attributes)
	if info == nil {
		return nil
	}
	w := &CheckButton{}
	w.id = iid
	w.info = info
	evalSetValue(variableId(iid), "")
	RegisterWidget(w)
	return w
}

func (w *CheckButton) Attach(id string) error {
	info, err := CheckWidgetInfo(id, WidgetTypeCheckButton)
	if err != nil {
		return err
	}
	w.id = id
	w.info = info
	RegisterWidget(w)
	return nil
}

func (w *CheckButton) SetText(text string) error {
	setObjText("atk_tmp_text", text)
	return eval(fmt.Sprintf("%v configure -text $atk_tmp_text", w.id))
}

func (w *CheckButton) Text() string {
	r, _ := evalAsString(fmt.Sprintf("%v cget -text", w.id))
	return r
}

func (w *CheckButton) SetWidth(width int) error {
	return eval(fmt.Sprintf("%v configure -width {%v}", w.id, width))
}

func (w *CheckButton) Width() int {
	r, _ := evalAsInt(fmt.Sprintf("%v cget -width", w.id))
	return r
}

func (w *CheckButton) SetImage(image *Image) error {
	if image == nil {
		return ErrInvalid
	}
	return eval(fmt.Sprintf("%v configure -image {%v}", w.id, image.Id()))
}

func (w *CheckButton) Image() *Image {
	r, err := evalAsString(fmt.Sprintf("%v cget -image", w.id))
	return parserImageResult(r, err)
}

func (w *CheckButton) SetCompound(compound Compound) error {
	return eval(fmt.Sprintf("%v configure -compound {%v}", w.id, compound))
}

func (w *CheckButton) Compound() Compound {
	r, err := evalAsString(fmt.Sprintf("%v cget -compound", w.id))
	return parserCompoundResult(r, err)
}

func (w *CheckButton) SetPaddingN(padx int, pady int) error {
	if w.info.IsTtk {
		return eval(fmt.Sprintf("%v configure -padding {%v %v}", w.id, padx, pady))
	}
	return eval(fmt.Sprintf("%v configure -padx {%v} -pady {%v}", w.id, padx, pady))
}

func (w *CheckButton) PaddingN() (int, int) {
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

func (w *CheckButton) SetPadding(pad Pad) error {
	return w.SetPaddingN(pad.X, pad.Y)
}

func (w *CheckButton) Padding() Pad {
	x, y := w.PaddingN()
	return Pad{x, y}
}

func (w *CheckButton) SetState(state State) error {
	return eval(fmt.Sprintf("%v configure -state {%v}", w.id, state))
}

func (w *CheckButton) State() State {
	r, err := evalAsString(fmt.Sprintf("%v cget -state", w.id))
	return parserStateResult(r, err)
}

func (w *CheckButton) SetTakeFocus(takefocus bool) error {
	return eval(fmt.Sprintf("%v configure -takefocus {%v}", w.id, boolToInt(takefocus)))
}

func (w *CheckButton) IsTakeFocus() bool {
	r, _ := evalAsBool(fmt.Sprintf("%v cget -takefocus", w.id))
	return r
}

func (w *CheckButton) SetChecked(check bool) error {
	return eval(fmt.Sprintf("set %v {%v}", variableId(w.id), boolToInt(check)))
}

func (w *CheckButton) IsChecked() bool {
	r, _ := evalAsBool(fmt.Sprintf("set %v", variableId(w.id)))
	return r
}

func (w *CheckButton) OnCommand(fn func()) error {
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

func (w *CheckButton) Invoke() {
	eval(fmt.Sprintf("%v invoke", w.id))
}

func CheckButtonAttrText(text string) *WidgetAttr {
	return &WidgetAttr{"text", text}
}

func CheckButtonAttrWidth(width int) *WidgetAttr {
	return &WidgetAttr{"width", width}
}

func CheckButtonAttrImage(image *Image) *WidgetAttr {
	if image == nil {
		return nil
	}
	return &WidgetAttr{"image", image.Id()}
}

func CheckButtonAttrCompound(compound Compound) *WidgetAttr {
	return &WidgetAttr{"compound", compound}
}

func CheckButtonAttrPadding(padding Pad) *WidgetAttr {
	return &WidgetAttr{"padding", padding}
}

func CheckButtonAttrState(state State) *WidgetAttr {
	return &WidgetAttr{"state", state}
}

func CheckButtonAttrTakeFocus(takefocus bool) *WidgetAttr {
	return &WidgetAttr{"takefocus", boolToInt(takefocus)}
}
