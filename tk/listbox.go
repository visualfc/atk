// Copyright 2018 visualfc. All rights reserved.

package tk

import (
	"fmt"
	"strings"
)

// listbox
type ListBox struct {
	BaseWidget
	xscrollcommand *CommandEx
	yscrollcommand *CommandEx
}

func NewListBox(parent Widget, attributes ...*WidgetAttr) *ListBox {
	theme := checkInitUseTheme(attributes)
	iid := makeNamedWidgetId(parent, "atk_listbox")
	attributes = append(attributes, &WidgetAttr{"listvariable", variableId(iid)})
	info := CreateWidgetInfo(iid, WidgetTypeListBox, theme, attributes)
	if info == nil {
		return nil
	}
	w := &ListBox{}
	w.id = iid
	w.info = info
	evalSetValue(variableId(iid), "")
	RegisterWidget(w)
	return w
}

func (w *ListBox) Attach(id string) error {
	info, err := CheckWidgetInfo(id, WidgetTypeListBox)
	if err != nil {
		return err
	}
	w.id = id
	w.info = info
	RegisterWidget(w)
	return nil
}

func (w *ListBox) SetBackground(color string) error {
	setObjText("atk_tmp_text", color)
	return eval(fmt.Sprintf("%v configure -background $atk_tmp_text", w.id))
}

func (w *ListBox) Background() string {
	r, _ := evalAsString(fmt.Sprintf("%v cget -background", w.id))
	return r
}

func (w *ListBox) SetBorderWidth(width int) error {
	return eval(fmt.Sprintf("%v configure -borderwidth {%v}", w.id, width))
}

func (w *ListBox) BorderWidth() int {
	r, _ := evalAsInt(fmt.Sprintf("%v cget -borderwidth", w.id))
	return r
}

func (w *ListBox) SetForground(color string) error {
	setObjText("atk_tmp_text", color)
	return eval(fmt.Sprintf("%v configure -foreground $atk_tmp_text", w.id))
}

func (w *ListBox) Forground() string {
	r, _ := evalAsString(fmt.Sprintf("%v cget -foreground", w.id))
	return r
}

func (w *ListBox) SetReliefStyle(relief ReliefStyle) error {
	return eval(fmt.Sprintf("%v configure -relief {%v}", w.id, relief))
}

func (w *ListBox) ReliefStyle() ReliefStyle {
	r, err := evalAsString(fmt.Sprintf("%v cget -relief", w.id))
	return parserReliefStyleResult(r, err)
}

func (w *ListBox) SetFont(font Font) error {
	if font == nil {
		return ErrInvalid
	}
	return eval(fmt.Sprintf("%v configure -font {%v}", w.id, font.Id()))
}

func (w *ListBox) Font() Font {
	r, err := evalAsString(fmt.Sprintf("%v cget -font", w.id))
	return parserFontResult(r, err)
}

func (w *ListBox) SetJustify(justify Justify) error {
	if !mainInterp.SupportVer86() {
		return ErrUnsupport
	}
	return eval(fmt.Sprintf("%v configure -justify {%v}", w.id, justify))
}

func (w *ListBox) Justify() Justify {
	if !mainInterp.SupportVer86() {
		return JustifyLeft
	}
	r, err := evalAsString(fmt.Sprintf("%v cget -justify", w.id))
	return parserJustifyResult(r, err)
}

func (w *ListBox) SetWidth(width int) error {
	return eval(fmt.Sprintf("%v configure -width {%v}", w.id, width))
}

func (w *ListBox) Width() int {
	r, _ := evalAsInt(fmt.Sprintf("%v cget -width", w.id))
	return r
}

func (w *ListBox) SetHeight(height int) error {
	return eval(fmt.Sprintf("%v configure -height {%v}", w.id, height))
}

func (w *ListBox) Height() int {
	r, _ := evalAsInt(fmt.Sprintf("%v cget -height", w.id))
	return r
}

func (w *ListBox) SetPaddingN(padx int, pady int) error {
	if w.info.IsTtk {
		return eval(fmt.Sprintf("%v configure -padding {%v %v}", w.id, padx, pady))
	}
	return eval(fmt.Sprintf("%v configure -padx {%v} -pady {%v}", w.id, padx, pady))
}

func (w *ListBox) PaddingN() (int, int) {
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

func (w *ListBox) SetPadding(pad Pad) error {
	return w.SetPaddingN(pad.X, pad.Y)
}

func (w *ListBox) Padding() Pad {
	x, y := w.PaddingN()
	return Pad{x, y}
}

func (w *ListBox) SetState(state State) error {
	return eval(fmt.Sprintf("%v configure -state {%v}", w.id, state))
}

func (w *ListBox) State() State {
	r, err := evalAsString(fmt.Sprintf("%v cget -state", w.id))
	return parserStateResult(r, err)
}

func (w *ListBox) SetSelectMode(mode ListSelectMode) error {
	return eval(fmt.Sprintf("%v configure -selectmode {%v}", w.id, mode))
}

func (w *ListBox) SelectMode() ListSelectMode {
	r, err := evalAsString(fmt.Sprintf("%v cget -selectmode", w.id))
	return parserListSelectModeResult(r, err)
}

func (w *ListBox) SetTakeFocus(takefocus bool) error {
	return eval(fmt.Sprintf("%v configure -takefocus {%v}", w.id, boolToInt(takefocus)))
}

func (w *ListBox) IsTakeFocus() bool {
	r, _ := evalAsBool(fmt.Sprintf("%v cget -takefocus", w.id))
	return r
}

func (w *ListBox) SetItems(items []string) *ListBox {
	mainInterp.SetStringList(variableId(w.id), items, true)
	return w
}

func (w *ListBox) Items() []string {
	return mainInterp.GetList(variableId(w.id), true).ToStringList()
}

func (w *ListBox) ItemCount() int {
	r, _ := evalAsInt(fmt.Sprintf("%v size", w.id))
	return r
}

func (w *ListBox) InsertItem(index int, item string) *ListBox {
	setObjText("atk_tmp_item", item)
	eval(fmt.Sprintf("%v insert %v $atk_tmp_item", w.id, index))
	return w
}

func (w *ListBox) AppendItem(index int, item string) *ListBox {
	mainInterp.AppendStringList(variableId(w.id), item, true)
	return w
}

func (w *ListBox) AppendItems(items []string) *ListBox {
	mainInterp.AppendStringListList(variableId(w.id), items, true)
	return w
}

func (w *ListBox) SetItemText(index int, item string) *ListBox {
	if index >= 0 && index < w.ItemCount() {
		setObjText("atk_tmp_item", item)
		eval(fmt.Sprintf("lset %v %v $atk_tmp_item", variableId(w.id), index))
	}
	return w
}

func (w *ListBox) ItemText(index int) string {
	r, _ := evalAsString(fmt.Sprintf("%v get %v", w.id, index))
	return r
}

func (w *ListBox) RemoveItem(index int) error {
	err := eval(fmt.Sprintf("%v delete %v %v", w.id, index, index))
	return err
}

func (w *ListBox) RemoveItemRange(start int, end int) error {
	err := eval(fmt.Sprintf("%v delete %v %v", w.id, start, end))
	return err
}

func (w *ListBox) SetSelectionRange(start int, end int) *ListBox {
	eval(fmt.Sprintf("%v selection set %v %v", w.id, start, end))
	return w
}

func (w *ListBox) SelectedIndexs() []int {
	r, _ := evalAsIntList(fmt.Sprintf("%v curselection", w.id))
	return r
}

func (w *ListBox) SelectedItems() (items []string) {
	indexs := w.SelectedIndexs()
	if indexs == nil {
		return nil
	}
	for _, index := range indexs {
		items = append(items, w.ItemText(index))
	}
	return
}

func (w *ListBox) ClearSelection() *ListBox {
	eval(fmt.Sprintf("%v selection clear 0 end", w.id))
	return w
}

func (w *ListBox) OnSelectionChanged(fn func()) error {
	if fn == nil {
		return ErrInvalid
	}
	return w.BindEvent("<<ListboxSelect>>", func(e *Event) {
		fn()
	})
}

func (w *ListBox) SetXViewArgs(args []string) error {
	return eval(fmt.Sprintf("%v xview %v", w.id, strings.Join(args, " ")))
}

func (w *ListBox) SetYViewArgs(args []string) error {
	return eval(fmt.Sprintf("%v yview %v", w.id, strings.Join(args, " ")))
}

func (w *ListBox) OnXScrollEx(fn func([]string) error) error {
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

func (w *ListBox) OnYScrollEx(fn func([]string) error) error {
	if fn == nil {
		return ErrInvalid
	}
	if w.yscrollcommand == nil {
		w.yscrollcommand = &CommandEx{}
		bindCommandEx(w.id, "yscrollcommand", w.yscrollcommand.Invoke)
	}
	w.yscrollcommand.Bind(fn)
	return nil
}

func (w *ListBox) BindXScrollBar(bar *ScrollBar) error {
	if !IsValidWidget(bar) {
		return ErrInvalid
	}
	w.OnXScrollEx(bar.SetScrollArgs)
	bar.OnCommandEx(w.SetXViewArgs)
	return nil
}

func (w *ListBox) BindYScrollBar(bar *ScrollBar) error {
	if !IsValidWidget(bar) {
		return ErrInvalid
	}
	w.OnYScrollEx(bar.SetScrollArgs)
	bar.OnCommandEx(w.SetYViewArgs)
	return nil
}

type ListBoxEx struct {
	*ScrollLayout
	*ListBox
}

func NewListBoxEx(parent Widget, attributs ...*WidgetAttr) *ListBoxEx {
	w := &ListBoxEx{}
	w.ScrollLayout = NewScrollLayout(parent)
	w.ListBox = NewListBox(parent, attributs...)
	w.SetWidget(w.ListBox)
	w.ListBox.BindXScrollBar(w.XScrollBar)
	w.ListBox.BindYScrollBar(w.YScrollBar)
	RegisterWidget(w)
	return w
}

func ListBoxAttrBackground(color string) *WidgetAttr {
	return &WidgetAttr{"background", color}
}

func ListBoxAttrBorderWidth(width int) *WidgetAttr {
	return &WidgetAttr{"borderwidth", width}
}

func ListBoxAttrForground(color string) *WidgetAttr {
	return &WidgetAttr{"foreground", color}
}

func ListBoxAttrReliefStyle(relief ReliefStyle) *WidgetAttr {
	return &WidgetAttr{"relief", relief}
}

func ListBoxAttrFont(font Font) *WidgetAttr {
	if font == nil {
		return nil
	}
	return &WidgetAttr{"font", font.Id()}
}

func ListBoxAttrJustify(justify Justify) *WidgetAttr {
	if !mainInterp.SupportVer86() {
		return nil
	}
	return &WidgetAttr{"justify", justify}
}

func ListBoxAttrWidth(width int) *WidgetAttr {
	return &WidgetAttr{"width", width}
}

func ListBoxAttrHeight(height int) *WidgetAttr {
	return &WidgetAttr{"height", height}
}

func ListBoxAttrPadding(padding Pad) *WidgetAttr {
	return &WidgetAttr{"padding", padding}
}

func ListBoxAttrState(state State) *WidgetAttr {
	return &WidgetAttr{"state", state}
}

func ListBoxAttrSelectMode(mode ListSelectMode) *WidgetAttr {
	return &WidgetAttr{"selectmode", mode}
}

func ListBoxAttrTakeFocus(takefocus bool) *WidgetAttr {
	return &WidgetAttr{"takefocus", boolToInt(takefocus)}
}
