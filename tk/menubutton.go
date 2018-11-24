// Copyright 2018 visualfc. All rights reserved.

package tk

import "fmt"

// menubutton
type MenuButton struct {
	BaseWidget
}

func NewMenuButton(parent Widget, text string, attributes ...*WidgetAttr) *MenuButton {
	theme := checkInitUseTheme(attributes)
	iid := makeNamedWidgetId(parent, "atk_menubutton")
	attributes = append(attributes, &WidgetAttr{"text", text})
	info := CreateWidgetInfo(iid, WidgetTypeMenuButton, theme, attributes)
	if info == nil {
		return nil
	}
	w := &MenuButton{}
	w.id = iid
	w.info = info
	RegisterWidget(w)
	return w
}

func (w *MenuButton) Attach(id string) error {
	info, err := CheckWidgetInfo(id, WidgetTypeMenuButton)
	if err != nil {
		return err
	}
	w.id = id
	w.info = info
	RegisterWidget(w)
	return nil
}

func (w *MenuButton) SetText(text string) error {
	setObjText("atk_tmp_text", text)
	return eval(fmt.Sprintf("%v configure -text $atk_tmp_text", w.id))
}

func (w *MenuButton) Text() string {
	r, _ := evalAsString(fmt.Sprintf("%v cget -text", w.id))
	return r
}

func (w *MenuButton) SetWidth(width int) error {
	return eval(fmt.Sprintf("%v configure -width {%v}", w.id, width))
}

func (w *MenuButton) Width() int {
	r, _ := evalAsInt(fmt.Sprintf("%v cget -width", w.id))
	return r
}

func (w *MenuButton) SetImage(image *Image) error {
	if image == nil {
		return ErrInvalid
	}
	return eval(fmt.Sprintf("%v configure -image {%v}", w.id, image.Id()))
}

func (w *MenuButton) Image() *Image {
	r, err := evalAsString(fmt.Sprintf("%v cget -image", w.id))
	return parserImageResult(r, err)
}

func (w *MenuButton) SetCompound(compound Compound) error {
	return eval(fmt.Sprintf("%v configure -compound {%v}", w.id, compound))
}

func (w *MenuButton) Compound() Compound {
	r, err := evalAsString(fmt.Sprintf("%v cget -compound", w.id))
	return parserCompoundResult(r, err)
}

func (w *MenuButton) SetPaddingN(padx int, pady int) error {
	if w.info.IsTtk {
		return eval(fmt.Sprintf("%v configure -padding {%v %v}", w.id, padx, pady))
	}
	return eval(fmt.Sprintf("%v configure -padx {%v} -pady {%v}", w.id, padx, pady))
}

func (w *MenuButton) PaddingN() (int, int) {
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

func (w *MenuButton) SetPadding(pad Pad) error {
	return w.SetPaddingN(pad.X, pad.Y)
}

func (w *MenuButton) Padding() Pad {
	x, y := w.PaddingN()
	return Pad{x, y}
}

func (w *MenuButton) SetState(state State) error {
	return eval(fmt.Sprintf("%v configure -state {%v}", w.id, state))
}

func (w *MenuButton) State() State {
	r, err := evalAsString(fmt.Sprintf("%v cget -state", w.id))
	return parserStateResult(r, err)
}

func (w *MenuButton) SetTakeFocus(takefocus bool) error {
	return eval(fmt.Sprintf("%v configure -takefocus {%v}", w.id, boolToInt(takefocus)))
}

func (w *MenuButton) IsTakeFocus() bool {
	r, _ := evalAsBool(fmt.Sprintf("%v cget -takefocus", w.id))
	return r
}

func (w *MenuButton) SetDirection(direction Direction) error {
	return eval(fmt.Sprintf("%v configure -direction {%v}", w.id, direction))
}

func (w *MenuButton) Direction() Direction {
	r, err := evalAsString(fmt.Sprintf("%v cget -direction", w.id))
	return parserDirectionResult(r, err)
}

func (w *MenuButton) SetMenu(menu *Menu) error {
	if menu == nil {
		return ErrInvalid
	}
	return eval(fmt.Sprintf("%v configure -menu {%v}", w.id, menu.Id()))
}

func (w *MenuButton) Menu() *Menu {
	r, err := evalAsString(fmt.Sprintf("%v cget -menu", w.id))
	return parserMenuResult(r, err)
}

func MenuButtonAttrText(text string) *WidgetAttr {
	return &WidgetAttr{"text", text}
}

func MenuButtonAttrWidth(width int) *WidgetAttr {
	return &WidgetAttr{"width", width}
}

func MenuButtonAttrImage(image *Image) *WidgetAttr {
	if image == nil {
		return nil
	}
	return &WidgetAttr{"image", image.Id()}
}

func MenuButtonAttrCompound(compound Compound) *WidgetAttr {
	return &WidgetAttr{"compound", compound}
}

func MenuButtonAttrPadding(padding Pad) *WidgetAttr {
	return &WidgetAttr{"padding", padding}
}

func MenuButtonAttrState(state State) *WidgetAttr {
	return &WidgetAttr{"state", state}
}

func MenuButtonAttrTakeFocus(takefocus bool) *WidgetAttr {
	return &WidgetAttr{"takefocus", boolToInt(takefocus)}
}

func MenuButtonAttrDirection(direction Direction) *WidgetAttr {
	return &WidgetAttr{"direction", direction}
}

func MenuButtonAttrMenu(menu *Menu) *WidgetAttr {
	if menu == nil {
		return nil
	}
	return &WidgetAttr{"menu", menu.Id()}
}
