// Copyright 2018 visualfc. All rights reserved.

package tk

import (
	"fmt"
	"strings"
)

// entry
type Entry struct {
	BaseWidget
	xscrollcommand *CommandEx
}

func NewEntry(parent Widget, attributes ...*WidgetAttr) *Entry {
	theme := checkInitUseTheme(attributes)
	iid := makeNamedWidgetId(parent, "atk_entry")
	attributes = append(attributes, &WidgetAttr{"textvariable", variableId(iid)})
	info := CreateWidgetInfo(iid, WidgetTypeEntry, theme, attributes)
	if info == nil {
		return nil
	}
	w := &Entry{}
	w.id = iid
	w.info = info
	evalSetValue(variableId(iid), "")
	RegisterWidget(w)
	return w
}

func (w *Entry) Attach(id string) error {
	info, err := CheckWidgetInfo(id, WidgetTypeEntry)
	if err != nil {
		return err
	}
	w.id = id
	w.info = info
	RegisterWidget(w)
	return nil
}

func (w *Entry) SetForeground(color string) error {
	setObjText("atk_tmp_text", color)
	return eval(fmt.Sprintf("%v configure -foreground $atk_tmp_text", w.id))
}

func (w *Entry) Foreground() string {
	r, _ := evalAsString(fmt.Sprintf("%v cget -foreground", w.id))
	return r
}

func (w *Entry) SetBackground(color string) error {
	setObjText("atk_tmp_text", color)
	return eval(fmt.Sprintf("%v configure -background $atk_tmp_text", w.id))
}

func (w *Entry) Background() string {
	r, _ := evalAsString(fmt.Sprintf("%v cget -background", w.id))
	return r
}

func (w *Entry) SetWidth(width int) error {
	return eval(fmt.Sprintf("%v configure -width {%v}", w.id, width))
}

func (w *Entry) Width() int {
	r, _ := evalAsInt(fmt.Sprintf("%v cget -width", w.id))
	return r
}

func (w *Entry) SetFont(font Font) error {
	if font == nil {
		return ErrInvalid
	}
	return eval(fmt.Sprintf("%v configure -font {%v}", w.id, font.Id()))
}

func (w *Entry) Font() Font {
	r, err := evalAsString(fmt.Sprintf("%v cget -font", w.id))
	return parserFontResult(r, err)
}

func (w *Entry) SetJustify(justify Justify) error {
	return eval(fmt.Sprintf("%v configure -justify {%v}", w.id, justify))
}

func (w *Entry) Justify() Justify {
	r, err := evalAsString(fmt.Sprintf("%v cget -justify", w.id))
	return parserJustifyResult(r, err)
}

func (w *Entry) SetShow(show string) error {
	setObjText("atk_tmp_text", show)
	return eval(fmt.Sprintf("%v configure -show $atk_tmp_text", w.id))
}

func (w *Entry) Show() string {
	r, _ := evalAsString(fmt.Sprintf("%v cget -show", w.id))
	return r
}

func (w *Entry) SetState(state State) error {
	return eval(fmt.Sprintf("%v configure -state {%v}", w.id, state))
}

func (w *Entry) State() State {
	r, err := evalAsString(fmt.Sprintf("%v cget -state", w.id))
	return parserStateResult(r, err)
}

func (w *Entry) SetTakeFocus(takefocus bool) error {
	return eval(fmt.Sprintf("%v configure -takefocus {%v}", w.id, boolToInt(takefocus)))
}

func (w *Entry) IsTakeFocus() bool {
	r, _ := evalAsBool(fmt.Sprintf("%v cget -takefocus", w.id))
	return r
}

func (w *Entry) SetExportSelection(export bool) error {
	return eval(fmt.Sprintf("%v configure -exportselection {%v}", w.id, boolToInt(export)))
}

func (w *Entry) IsExportSelection() bool {
	r, _ := evalAsBool(fmt.Sprintf("%v cget -exportselection", w.id))
	return r
}

func (w *Entry) SetText(text string) error {
	setObjText("atk_tmp_text", text)
	return eval(fmt.Sprintf("set %v $atk_tmp_text", variableId(w.id)))
}

func (w *Entry) Text() string {
	r, _ := evalAsString(fmt.Sprintf("set %v", variableId(w.id)))
	return r
}

func (w *Entry) SetXViewArgs(args []string) error {
	return eval(fmt.Sprintf("%v xview %v", w.id, strings.Join(args, " ")))
}

func (w *Entry) OnXScrollEx(fn func([]string) error) error {
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

func (w *Entry) BindXScrollBar(bar *ScrollBar) error {
	if !IsValidWidget(bar) {
		return ErrInvalid
	}
	w.OnXScrollEx(bar.SetScrollArgs)
	bar.OnCommandEx(w.SetXViewArgs)
	return nil
}

func (w *Entry) SetCursorPosition(pos int) *Entry {
	eval(fmt.Sprintf("%v icursor %v", w.id, pos))
	return w
}

func (w *Entry) CursorPosition() int {
	return w.index("insert")
}

func (w *Entry) OnUpdate(fn func()) error {
	return traceVariable(variableId(w.id), fn)
}

func (w *Entry) OnEditReturn(fn func()) error {
	if fn == nil {
		return ErrInvalid
	}
	w.BindEvent("<Return>", func(e *Event) {
		fn()
	})
	return nil
}

func (w *Entry) Copy() {
	SendEvent(w, "<<Copy>>")
}

func (w *Entry) Paste() {
	SendEvent(w, "<<Paste")
}

func (w *Entry) Cut() {
	SendEvent(w, "<<Cut>>")
}

func (w *Entry) Clear() {
	w.SetText("")
}

func (w *Entry) Delete(index int) {
	eval(fmt.Sprintf("%v delete %v", w.id, index))
}

func (w *Entry) DeleteRange(start int, end int) {
	eval(fmt.Sprintf("%v delete %v %v", w.id, start, end))
}

func (w *Entry) TextLength() int {
	return w.index("end")
}

func (w *Entry) index(index string) int {
	r, _ := evalAsInt(fmt.Sprintf("%v index %v", w.id, index))
	return r
}

func (w *Entry) Index(index int) int {
	r, _ := evalAsInt(fmt.Sprintf("%v index %v", w.id, index))
	return r
}

func (w *Entry) Insert(index int, str string) error {
	setObjText("atk_entry_text", str)
	return eval(fmt.Sprintf("%v insert %v $atk_entry_text", w.id, index))
}

func (w *Entry) Append(str string) error {
	setObjText("atk_entry_text", str)
	return eval(fmt.Sprintf("%v insert end $atk_entry_text", w.id))
}

func (w *Entry) ClearSelection() {
	eval(fmt.Sprintf("%v selection clear", w.id))
}

func (w *Entry) HasSelectedText() bool {
	r, _ := evalAsBool(fmt.Sprintf("%v selection present", w.id))
	return r
}

func (w *Entry) SelectedText() string {
	if !w.HasSelectedText() {
		return ""
	}
	return SubString(w.Text(), w.SelectionStart(), w.SelectionEnd())
}

func (w *Entry) SetSelection(start int, end int) {
	eval(fmt.Sprintf("%v selection range %v %v", w.id, start, end))
}

func (w *Entry) SelectAll() {
	w.SetSelection(0, w.TextLength())
}

func (w *Entry) SelectionStart() int {
	if !w.HasSelectedText() {
		return -1
	}
	return w.index("sel.first")
}

func (w *Entry) SelectionEnd() int {
	if !w.HasSelectedText() {
		return -1
	}
	return w.index("sel.last")
}

func EntryAttrForeground(color string) *WidgetAttr {
	return &WidgetAttr{"foreground", color}
}

func EntryAttrBackground(color string) *WidgetAttr {
	return &WidgetAttr{"background", color}
}

func EntryAttrWidth(width int) *WidgetAttr {
	return &WidgetAttr{"width", width}
}

func EntryAttrFont(font Font) *WidgetAttr {
	if font == nil {
		return nil
	}
	return &WidgetAttr{"font", font.Id()}
}

func EntryAttrJustify(justify Justify) *WidgetAttr {
	return &WidgetAttr{"justify", justify}
}

func EntryAttrShow(show string) *WidgetAttr {
	return &WidgetAttr{"show", show}
}

func EntryAttrState(state State) *WidgetAttr {
	return &WidgetAttr{"state", state}
}

func EntryAttrTakeFocus(takefocus bool) *WidgetAttr {
	return &WidgetAttr{"takefocus", boolToInt(takefocus)}
}

func EntryAttrExportSelection(export bool) *WidgetAttr {
	return &WidgetAttr{"exportselection", boolToInt(export)}
}
