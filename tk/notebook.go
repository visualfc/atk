// Copyright 2018 visualfc. All rights reserved.

package tk

import (
	"fmt"
	"strings"
)

// notebook
type Notebook struct {
	BaseWidget
}

func NewNotebook(parent Widget, attributes ...*WidgetAttr) *Notebook {
	theme := checkInitUseTheme(attributes)
	iid := makeNamedWidgetId(parent, "atk_notebook")
	info := CreateWidgetInfo(iid, WidgetTypeNotebook, theme, attributes)
	if info == nil {
		return nil
	}
	w := &Notebook{}
	w.id = iid
	w.info = info
	RegisterWidget(w)
	return w
}

func (w *Notebook) Attach(id string) error {
	info, err := CheckWidgetInfo(id, WidgetTypeNotebook)
	if err != nil {
		return err
	}
	w.id = id
	w.info = info
	RegisterWidget(w)
	return nil
}

func (w *Notebook) SetWidth(width int) error {
	return eval(fmt.Sprintf("%v configure -width {%v}", w.id, width))
}

func (w *Notebook) Width() int {
	r, _ := evalAsInt(fmt.Sprintf("%v cget -width", w.id))
	return r
}

func (w *Notebook) SetHeight(height int) error {
	return eval(fmt.Sprintf("%v configure -height {%v}", w.id, height))
}

func (w *Notebook) Height() int {
	r, _ := evalAsInt(fmt.Sprintf("%v cget -height", w.id))
	return r
}

func (w *Notebook) SetTakeFocus(takefocus bool) error {
	return eval(fmt.Sprintf("%v configure -takefocus {%v}", w.id, boolToInt(takefocus)))
}

func (w *Notebook) IsTakeFocus() bool {
	r, _ := evalAsBool(fmt.Sprintf("%v cget -takefocus", w.id))
	return r
}

func (w *Notebook) SetPaddingN(padx int, pady int) error {
	if w.info.IsTtk {
		return eval(fmt.Sprintf("%v configure -padding {%v %v}", w.id, padx, pady))
	}
	return eval(fmt.Sprintf("%v configure -padx {%v} -pady {%v}", w.id, padx, pady))
}

func (w *Notebook) PaddingN() (int, int) {
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

func (w *Notebook) SetPadding(pad Pad) error {
	return w.SetPaddingN(pad.X, pad.Y)
}

func (w *Notebook) Padding() Pad {
	x, y := w.PaddingN()
	return Pad{x, y}
}

var (
	tabAttributes = []string{
		"padding",
		"state",
		"sticky",
		"text",
		"image",
		"compound",
	}
)

func isTabAttributes(attr string) bool {
	for _, v := range tabAttributes {
		if v == attr {
			return true
		}
	}
	return false
}

func buildTabAttributeScript(ttk bool, attributes ...*WidgetAttr) string {
	var list []string
	for _, attr := range attributes {
		if attr == nil {
			continue
		}
		if attr.Key == "padding" {
			list = append(list, checkPaddingScript(ttk, attr))
			continue
		}
		if !isTabAttributes(attr.Key) {
			continue
		}
		if s, ok := attr.Value.(string); ok {
			pname := "atk_tmp_" + attr.Key
			setObjText(pname, s)
			list = append(list, fmt.Sprintf("-%v $%v", attr.Key, pname))
			continue
		}
		list = append(list, fmt.Sprintf("-%v {%v}", attr.Key, attr.Value))
	}
	return strings.Join(list, " ")
}

func (w *Notebook) AddTab(widget Widget, text string, attributes ...*WidgetAttr) error {
	if !IsValidWidget(widget) {
		return ErrInvalid
	}
	attributes = append(attributes, &WidgetAttr{"text", text})
	extra := buildTabAttributeScript(w.info.IsTtk, attributes...)
	script := fmt.Sprintf("%v add %v", w.id, widget.Id())
	if extra != "" {
		script += " " + extra
	}
	return eval(script)
}

func (w *Notebook) InsertTab(pos int, widget Widget, text string, attributes ...*WidgetAttr) error {
	if !IsValidWidget(widget) {
		return ErrInvalid
	}
	attributes = append(attributes, &WidgetAttr{"text", text})
	extra := buildTabAttributeScript(w.info.IsTtk, attributes...)
	script := fmt.Sprintf("%v insert %v %v", w.id, pos, widget.Id())
	if extra != "" {
		script += " " + extra
	}
	return eval(script)
}

func (w *Notebook) SetTab(widget Widget, text string, attributes ...*WidgetAttr) error {
	if !IsValidWidget(widget) {
		return ErrInvalid
	}
	attributes = append(attributes, &WidgetAttr{"text", text})
	extra := buildTabAttributeScript(w.info.IsTtk, attributes...)
	script := fmt.Sprintf("%v tab %v", w.id, widget.Id())
	if extra != "" {
		script += " " + extra
	}
	return eval(script)
}

func (w *Notebook) RemoveTab(widget Widget) error {
	if !IsValidWidget(widget) {
		return ErrInvalid
	}
	return eval(fmt.Sprintf("%v forget %v", w.id, widget.Id()))
}

func (w *Notebook) SetCurrentTab(widget Widget) error {
	if !IsValidWidget(widget) {
		return ErrInvalid
	}
	return eval(fmt.Sprintf("%v select %v", w.id, widget.Id()))
}

func (w *Notebook) CurrentTab() Widget {
	r, _ := evalAsString(fmt.Sprintf("%v select", w.id))
	return FindWidget(r)
}

func (w *Notebook) CurrentTabIndex() int {
	r, _ := evalAsInt(fmt.Sprintf("%v index current", w.id))
	return r
}

func (w *Notebook) TabCount() int {
	r, _ := evalAsInt(fmt.Sprintf("%v index end", w.id))
	return r
}

func (w *Notebook) TabIndex(widget Widget) int {
	if !IsValidWidget(widget) {
		return -1
	}
	r, _ := evalAsInt(fmt.Sprintf("%v index %v", w.id, widget.Id()))
	return r
}

func TabAttrState(state State) *WidgetAttr {
	return &WidgetAttr{"state", state}
}

func TabAttrSticky(sticky Sticky) *WidgetAttr {
	return &WidgetAttr{"sticky", sticky}
}

func TabAttrPadding(padding Pad) *WidgetAttr {
	return &WidgetAttr{"padding", padding}
}

func TabAttrText(text string) *WidgetAttr {
	return &WidgetAttr{"text", text}
}

func TabAttrImage(image *Image) *WidgetAttr {
	if image == nil {
		return nil
	}
	return &WidgetAttr{"image", image.Id()}
}

func TabAttrCompound(compound Compound) *WidgetAttr {
	return &WidgetAttr{"compound", compound}
}

func NotebookAttrWidth(width int) *WidgetAttr {
	return &WidgetAttr{"width", width}
}

func NotebookAttrHeight(height int) *WidgetAttr {
	return &WidgetAttr{"height", height}
}

func NotebookAttrTakeFocus(takefocus bool) *WidgetAttr {
	return &WidgetAttr{"takefocus", boolToInt(takefocus)}
}

func NotebookAttrPadding(padding Pad) *WidgetAttr {
	return &WidgetAttr{"padding", padding}
}
