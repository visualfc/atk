// Copyright 2018 visualfc. All rights reserved.

package tk

import "fmt"

// menu
type Menu struct {
	BaseWidget
}

func NewMenu(parent Widget, attributes ...*WidgetAttr) *Menu {
	theme := checkInitUseTheme(attributes)
	iid := makeNamedWidgetId(parent, "atk_menu")
	info := CreateWidgetInfo(iid, WidgetTypeMenu, theme, attributes)
	if info == nil {
		return nil
	}
	w := &Menu{}
	w.id = iid
	w.info = info
	RegisterWidget(w)
	return w
}

func (w *Menu) Attach(id string) error {
	info, err := CheckWidgetInfo(id, WidgetTypeMenu)
	if err != nil {
		return err
	}
	w.id = id
	w.info = info
	RegisterWidget(w)
	return nil
}

func (w *Menu) SetFont(font Font) error {
	if font == nil {
		return ErrInvalid
	}
	return eval(fmt.Sprintf("%v configure -font {%v}", w.id, font.Id()))
}

func (w *Menu) Font() Font {
	r, err := evalAsString(fmt.Sprintf("%v cget -font", w.id))
	return parserFontResult(r, err)
}

func (w *Menu) SetActiveBackground(color string) error {
	setObjText("atk_tmp_text", color)
	return eval(fmt.Sprintf("%v configure -activebackground $atk_tmp_text", w.id))
}

func (w *Menu) ActiveBackground() string {
	r, _ := evalAsString(fmt.Sprintf("%v cget -activebackground", w.id))
	return r
}

func (w *Menu) SetActiveForground(color string) error {
	setObjText("atk_tmp_text", color)
	return eval(fmt.Sprintf("%v configure -activeforeground $atk_tmp_text", w.id))
}

func (w *Menu) ActiveForground() string {
	r, _ := evalAsString(fmt.Sprintf("%v cget -activeforeground", w.id))
	return r
}

func (w *Menu) SetBackground(color string) error {
	setObjText("atk_tmp_text", color)
	return eval(fmt.Sprintf("%v configure -background $atk_tmp_text", w.id))
}

func (w *Menu) Background() string {
	r, _ := evalAsString(fmt.Sprintf("%v cget -background", w.id))
	return r
}

func (w *Menu) SetForground(color string) error {
	setObjText("atk_tmp_text", color)
	return eval(fmt.Sprintf("%v configure -foreground $atk_tmp_text", w.id))
}

func (w *Menu) Forground() string {
	r, _ := evalAsString(fmt.Sprintf("%v cget -foreground", w.id))
	return r
}

func (w *Menu) SetSelectColor(color string) error {
	setObjText("atk_tmp_text", color)
	return eval(fmt.Sprintf("%v configure -selectcolor $atk_tmp_text", w.id))
}

func (w *Menu) SelectColor() string {
	r, _ := evalAsString(fmt.Sprintf("%v cget -selectcolor", w.id))
	return r
}

func (w *Menu) SetDisabledForground(color string) error {
	setObjText("atk_tmp_text", color)
	return eval(fmt.Sprintf("%v configure -disabledforeground $atk_tmp_text", w.id))
}

func (w *Menu) DisabledForground() string {
	r, _ := evalAsString(fmt.Sprintf("%v cget -disabledforeground", w.id))
	return r
}

func (w *Menu) SetActiveBorderWidth(width int) error {
	return eval(fmt.Sprintf("%v configure -activeborderwidth {%v}", w.id, width))
}

func (w *Menu) ActiveBorderWidth() int {
	r, _ := evalAsInt(fmt.Sprintf("%v cget -activeborderwidth", w.id))
	return r
}

func (w *Menu) SetBorderWidth(width int) error {
	return eval(fmt.Sprintf("%v configure -borderwidth {%v}", w.id, width))
}

func (w *Menu) BorderWidth() int {
	r, _ := evalAsInt(fmt.Sprintf("%v cget -borderwidth", w.id))
	return r
}

func (w *Menu) SetReliefStyle(relief ReliefStyle) error {
	return eval(fmt.Sprintf("%v configure -relief {%v}", w.id, relief))
}

func (w *Menu) ReliefStyle() ReliefStyle {
	r, err := evalAsString(fmt.Sprintf("%v cget -relief", w.id))
	return parserReliefStyleResult(r, err)
}

func (w *Menu) SetTearoffTitle(title string) error {
	setObjText("atk_tmp_text", title)
	return eval(fmt.Sprintf("%v configure -title $atk_tmp_text", w.id))
}

func (w *Menu) TearoffTitle() string {
	r, _ := evalAsString(fmt.Sprintf("%v cget -title", w.id))
	return r
}

func (w *Menu) SetTearoff(tearoff bool) error {
	return eval(fmt.Sprintf("%v configure -tearoff {%v}", w.id, boolToInt(tearoff)))
}

func (w *Menu) IsTearoff() bool {
	r, _ := evalAsBool(fmt.Sprintf("%v cget -tearoff", w.id))
	return r
}

func (w *Menu) SetTakeFocus(takefocus bool) error {
	return eval(fmt.Sprintf("%v configure -takefocus {%v}", w.id, boolToInt(takefocus)))
}

func (w *Menu) IsTakeFocus() bool {
	r, _ := evalAsBool(fmt.Sprintf("%v cget -takefocus", w.id))
	return r
}

func (w *Menu) AddSubMenu(label string, sub *Menu) error {
	setObjText("atk_tmp_label", label)
	err := eval(fmt.Sprintf("%v add cascade -label $atk_tmp_label -menu {%v}",
		w.id, sub.id))
	return err
}

func (w *Menu) AddNewSubMenu(label string, attributes ...*WidgetAttr) *Menu {
	sub := NewMenu(w, attributes...)
	err := w.AddSubMenu(label, sub)
	if err != nil {
		return nil
	}
	return sub
}

func (w *Menu) InsertSubMenu(index int, label string, sub *Menu) error {
	if index < 0 {
		return w.AddSubMenu(label, sub)
	}
	setObjText("atk_tmp_label", label)
	err := eval(fmt.Sprintf("%v insert %v cascade -label $atk_tmp_label -menu {%v}",
		w.id, index, sub.id))
	return err
}

func (w *Menu) InsertNewSubMenu(index int, label string, attributes ...*WidgetAttr) *Menu {
	sub := NewMenu(w, attributes...)
	err := w.InsertSubMenu(index, label, sub)
	if err != nil {
		return nil
	}
	return sub
}

func (w *Menu) AddAction(act *Action) error {
	var script string
	if act.IsSeparator() {
		script = fmt.Sprintf("%v add separator", w.id)
	} else if act.IsRadioAction() {
		setObjText("atk_tmp_label", act.label)
		script = fmt.Sprintf("%v add radiobutton -label $atk_tmp_label -variable {%v} -value {%v} -command {%v}",
			w.id, act.groupid, act.radioid, act.actid)
	} else if act.IsCheckAction() {
		setObjText("atk_tmp_label", act.label)
		script = fmt.Sprintf("%v add checkbutton -label $atk_tmp_label -variable {%v} -command {%v}",
			w.id, act.checkid, act.actid)
	} else {
		setObjText("atk_tmp_label", act.label)
		script = fmt.Sprintf("%v add command -label $atk_tmp_label -command {%v}",
			w.id, act.actid)
	}
	return eval(script)
}

func (w *Menu) InsertAction(index int, act *Action) error {
	if index < 0 {
		return w.AddAction(act)
	}
	setObjText("atk_tmp_label", act.label)
	var script string
	if act.IsSeparator() {
		script = fmt.Sprintf("%v insert %v separator", w.id, index)
	} else if act.IsRadioAction() {
		script = fmt.Sprintf("%v insert %v radiobutton -label $atk_tmp_label -variable {%v} -value {%v} -command {%v}",
			w.id, index, act.groupid, act.radioid, act.actid)
	} else if act.IsCheckAction() {
		script = fmt.Sprintf("%v insert %v checkbutton -label $atk_tmp_label -variable {%v} -command {%v}",
			w.id, index, act.checkid, act.actid)
	} else {
		script = fmt.Sprintf("%v insert %v command -label $atk_tmp_label -command {%v}",
			w.id, index, act.actid)
	}
	return eval(script)
}

func (w *Menu) AddActions(actions []*Action) {
	for _, act := range actions {
		w.AddAction(act)
	}
}

func (w *Menu) AddSeparator() error {
	return eval(fmt.Sprintf("%v add separator", w.id))
}

func (w *Menu) InsertSeparator(index int) error {
	if index < 0 {
		return w.AddSeparator()
	}
	return eval(fmt.Sprintf("%v insert %v separator", w.id, index))
}

func parserMenuResult(r string, err error) *Menu {
	if err != nil {
		return nil
	}
	if r == "" {
		return nil
	}
	if i := FindWidget(r); i != nil {
		if m, ok := i.(*Menu); ok {
			return m
		}
		return nil
	}
	m := &Menu{}
	if m.Attach(r) != nil {
		return nil
	}
	return m
}

func SetMenuTearoff(enable bool) {
	eval(fmt.Sprintf("option add *Menu.tearOff %v", boolToInt(enable)))
}

func PopupMenu(menu *Menu, xpos int, ypos int) error {
	if !IsValidWidget(menu) {
		return ErrInvalid
	}
	return eval(fmt.Sprintf("tk_popup %v %v %v", menu.Id(), xpos, ypos))
}

func MenuAttrFont(font Font) *WidgetAttr {
	if font == nil {
		return nil
	}
	return &WidgetAttr{"font", font.Id()}
}

func MenuAttrActiveBackground(color string) *WidgetAttr {
	return &WidgetAttr{"activebackground", color}
}

func MenuAttrActiveForground(color string) *WidgetAttr {
	return &WidgetAttr{"activeforeground", color}
}

func MenuAttrBackground(color string) *WidgetAttr {
	return &WidgetAttr{"background", color}
}

func MenuAttrForground(color string) *WidgetAttr {
	return &WidgetAttr{"foreground", color}
}

func MenuAttrSelectColor(color string) *WidgetAttr {
	return &WidgetAttr{"selectcolor", color}
}

func MenuAttrDisabledForground(color string) *WidgetAttr {
	return &WidgetAttr{"disabledforeground", color}
}

func MenuAttrActiveBorderWidth(width int) *WidgetAttr {
	return &WidgetAttr{"activeborderwidth", width}
}

func MenuAttrBorderWidth(width int) *WidgetAttr {
	return &WidgetAttr{"borderwidth", width}
}

func MenuAttrReliefStyle(relief ReliefStyle) *WidgetAttr {
	return &WidgetAttr{"relief", relief}
}

func MenuAttrTearoffTitle(title string) *WidgetAttr {
	return &WidgetAttr{"title", title}
}

func MenuAttrTearoff(tearoff bool) *WidgetAttr {
	return &WidgetAttr{"tearoff", boolToInt(tearoff)}
}

func MenuAttrTakeFocus(takefocus bool) *WidgetAttr {
	return &WidgetAttr{"takefocus", boolToInt(takefocus)}
}
