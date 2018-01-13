// Copyright 2018 visualfc. All rights reserved.

package tk

import (
	"fmt"
	"strings"
)

var (
	ErrorInvalidWidgetInfo  = fmt.Errorf("invalid widget info")
	ErrorNotMatchWidgetInfo = fmt.Errorf("not match widget info")
)

type WidgetType int

const (
	WidgetTypeButton WidgetType = iota
	WidgetTypeCanvas
	WidgetTypeCheckButton
	WidgetTypeComboBox
	WidgetTypeEntry
	WidgetTypeFrame
	WidgetTypeLabel
	WidgetTypeLabelFrame
	WidgetTypeListBox
	WidgetTypeMenu
	WidgetTypeMenuButton
	WidgetTypeNoteBook
	WidgetTypePanedWindow
	WidgetTypeProgressBar
	WidgetTypeRadioButton
	WidgetTypeScale
	WidgetTypeScrollBar
	WidgetTypeSeparator
	WidgetTypeSizeGrip
	WidgetTypeSpinBox
	WidgetTypeTextEdit
	WidgetTypeWindow
	WidgetTypeTreeView
)

type MetaClass struct {
	Command string
	Class   string
	Options []string
}

func (m *MetaClass) HasOption(opt string) bool {
	if opt == "" {
		return false
	}
	for _, v := range m.Options {
		if v == opt {
			return true
		}
	}
	return false
}

type MetaType struct {
	Type string
	Tk   *MetaClass
	Ttk  *MetaClass
}

type WidgetInfo struct {
	Type      WidgetType
	TypeName  string
	IsTtk     bool
	MetaClass *MetaClass
}

func (typ WidgetType) MetaClass(theme bool) (typName string, meta *MetaClass, ttk bool) {
	mc, ok := typeMetaMap[typ]
	if !ok {
		panic(fmt.Errorf("error find metaclass type:%v", typ))
	}
	if theme && mainTheme != nil && mainTheme.IsTtk() {
		if mc.Ttk != nil {
			return mc.Type, mc.Ttk, true
		}
		return mc.Type, mc.Tk, false
	}
	if mc.Tk != nil {
		return mc.Type, mc.Tk, false
	}
	return mc.Type, mc.Ttk, true
}

func (typ WidgetType) ThemeConfigure() string {
	if mainTheme == nil {
		return ""
	}
	var list []string
	opts := mainTheme.WidgetOption(typ)
	_, meta, _ := typ.MetaClass(true)
	for _, opt := range opts {
		if !meta.HasOption(opt.Key) {
			continue
		}
		list = append(list, fmt.Sprintf("-%v {%v}", opt.Key, opt.Value))
	}
	return strings.Join(list, " ")
}

type WidgetOpt struct {
	Key   string
	Value interface{}
}

func lookupId(options []*WidgetOpt) string {
	for _, opt := range options {
		if opt != nil && opt.Key == "id" {
			if id, ok := opt.Value.(string); ok {
				return id
			}
		}
	}
	return ""
}

func CreateWidgetInfo(iid string, typ WidgetType, theme bool, options []*WidgetOpt) *WidgetInfo {
	typName, meta, isttk := typ.MetaClass(theme)
	script := fmt.Sprintf("%v %v", meta.Command, iid)
	if theme {
		cfg := typ.ThemeConfigure()
		if cfg != "" {
			script += " " + cfg
		}
	}
	if len(options) > 0 {
		var list []string
		for _, opt := range options {
			if opt == nil {
				continue
			}
			if !meta.HasOption(opt.Key) {
				continue
			}
			list = append(list, fmt.Sprintf("-%v {%v}", opt.Key, opt.Value))
		}
		if len(list) > 0 {
			script += " " + strings.Join(list, " ")
		}
	}
	err := eval(script)
	if err != nil {
		return nil
	}
	return &WidgetInfo{typ, typName, isttk, meta}
}

func findClassById(id string) string {
	if id == "." {
		return "Toplevel"
	}
	s, err := mainInterp.EvalAsString(fmt.Sprintf("winfo class {%v}", id))
	if err != nil {
		return ""
	}
	return s
}

func FindWidgetInfo(id string) *WidgetInfo {
	class := findClassById(id)
	if class == "" {
		return nil
	}
	for k, v := range typeMetaMap {
		if v.Tk != nil && v.Tk.Class == class {
			return &WidgetInfo{k, v.Type, false, v.Tk}
		}
		if v.Ttk != nil && v.Ttk.Class == class {
			return &WidgetInfo{k, v.Type, true, v.Ttk}
		}
	}
	return nil
}

func CheckWidgetInfo(id string, typ WidgetType) (*WidgetInfo, error) {
	info := FindWidgetInfo(id)
	if info == nil {
		return nil, ErrorInvalidWidgetInfo
	}
	if info.Type != WidgetTypeWindow {
		return nil, ErrorNotMatchWidgetInfo
	}
	return info, nil
}

var (
	typeMetaMap = make(map[WidgetType]*MetaType)
)

func IsTtkClass(class string) bool {
	for _, v := range ttkClassList {
		if v == class {
			return true
		}
	}
	return false
}

func IsTkClass(class string) bool {
	for _, v := range tkClassList {
		if v == class {
			return true
		}
	}
	return false
}

var (
	tkClassList = []string{
		"Button",
		"Canvas",
		"Checkbutton",
		"Entry",
		"Frame",
		"Label",
		"Labelframe",
		"Listbox",
		"Menu",
		"Menubutton",
		"Panedwindow",
		"Radiobutton",
		"Scale",
		"Scrollbar",
		"Spinbox",
		"Text",
		"Toplevel",
	}
	ttkClassList = []string{
		"TButton",
		"TCheckbutton",
		"TCombobox",
		"TEntry",
		"TFrame",
		"TLabel",
		"TLabelframe",
		"TMenubutton",
		"TNotebook",
		"TPanedwindow",
		"TProgressbar",
		"TRadiobutton",
		"TScale",
		"Scrollbar",
		"TSeparator",
		"TSizegrip",
		"Treeview",
	}
)
