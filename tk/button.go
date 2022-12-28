// Copyright 2018 visualfc. All rights reserved.

package tk

import "fmt"

// button
type Button struct {
	BaseWidget
	command *Command
}

func NewButton(parent Widget, text string, attributes ...*WidgetAttr) *Button {
	theme := checkInitUseTheme(attributes)
	iid := makeNamedWidgetId(parent, "atk_button")
	attributes = append(attributes, &WidgetAttr{"text", text})
	info := CreateWidgetInfo(iid, WidgetTypeButton, theme, attributes)
	if info == nil {
		return nil
	}
	w := &Button{}
	w.id = iid
	w.info = info
	RegisterWidget(w)
	return w
}

func (w *Button) Attach(id string) error {
	info, err := CheckWidgetInfo(id, WidgetTypeButton)
	if err != nil {
		return err
	}
	w.id = id
	w.info = info
	RegisterWidget(w)
	return nil
}

func (w *Button) SetText(text string) error {
	setObjText("atk_tmp_text", text)
	return eval(fmt.Sprintf("%v configure -text $atk_tmp_text", w.id))
}

func (w *Button) Text() string {
	r, _ := evalAsString(fmt.Sprintf("%v cget -text", w.id))
	return r
}

func (w *Button) SetUnder(pos int) error {
        return eval(fmt.Sprintf("%v configure -under %d",w.id,pos))
}

func (w *Button) SetWidth(width int) error {
	return eval(fmt.Sprintf("%v configure -width {%v}", w.id, width))
}

func (w *Button) Width() int {
	r, _ := evalAsInt(fmt.Sprintf("%v cget -width", w.id))
	return r
}

func (w *Button) SetImage(image *Image) error {
	if image == nil {
		return ErrInvalid
	}
	return eval(fmt.Sprintf("%v configure -image {%v}", w.id, image.Id()))
}

func (w *Button) Image() *Image {
	r, err := evalAsString(fmt.Sprintf("%v cget -image", w.id))
	return parserImageResult(r, err)
}

func (w *Button) SetCompound(compound Compound) error {
	return eval(fmt.Sprintf("%v configure -compound {%v}", w.id, compound))
}

func (w *Button) Compound() Compound {
	r, err := evalAsString(fmt.Sprintf("%v cget -compound", w.id))
	return parserCompoundResult(r, err)
}

func (w *Button) SetPaddingN(padx int, pady int) error {
	if w.info.IsTtk {
		return eval(fmt.Sprintf("%v configure -padding {%v %v}", w.id, padx, pady))
	}
	return eval(fmt.Sprintf("%v configure -padx {%v} -pady {%v}", w.id, padx, pady))
}

func (w *Button) PaddingN() (int, int) {
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

func (w *Button) SetPadding(pad Pad) error {
	return w.SetPaddingN(pad.X, pad.Y)
}

func (w *Button) Padding() Pad {
	x, y := w.PaddingN()
	return Pad{x, y}
}

func (w *Button) SetState(state State) error {
	return eval(fmt.Sprintf("%v configure -state {%v}", w.id, state))
}

func (w *Button) State() State {
	r, err := evalAsString(fmt.Sprintf("%v cget -state", w.id))
	return parserStateResult(r, err)
}

func (w *Button) SetTakeFocus(takefocus bool) error {
	return eval(fmt.Sprintf("%v configure -takefocus {%v}", w.id, boolToInt(takefocus)))
}

func (w *Button) IsTakeFocus() bool {
	r, _ := evalAsBool(fmt.Sprintf("%v cget -takefocus", w.id))
	return r
}

func (w *Button) OnCommand(fn func()) error {
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

func (w *Button) Invoke() {
	eval(fmt.Sprintf("%v invoke", w.id))
}

func ButtonAttrText(text string) *WidgetAttr {
	return &WidgetAttr{"text", text}
}

func ButtonAttrWidth(width int) *WidgetAttr {
	return &WidgetAttr{"width", width}
}

func ButtonAttrImage(image *Image) *WidgetAttr {
	if image == nil {
		return nil
	}
	return &WidgetAttr{"image", image.Id()}
}

func ButtonAttrCompound(compound Compound) *WidgetAttr {
	return &WidgetAttr{"compound", compound}
}

func ButtonAttrPadding(padding Pad) *WidgetAttr {
	return &WidgetAttr{"padding", padding}
}

func ButtonAttrState(state State) *WidgetAttr {
	return &WidgetAttr{"state", state}
}

func ButtonAttrTakeFocus(takefocus bool) *WidgetAttr {
	return &WidgetAttr{"takefocus", boolToInt(takefocus)}
}
