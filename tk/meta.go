// Copyright 2018 visualfc. All rights reserved.

package tk

import (
	"fmt"
	"strings"
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
	Id        string
	Type      WidgetType
	IsTtk     bool
	MetaClass *MetaClass
}

func (typ WidgetType) MetaClass(theme bool) (meta *MetaClass, ttk bool) {
	mc, ok := typeMetaMap[typ]
	if !ok {
		panic(fmt.Errorf("error find metaclass type:%v", typ))
	}
	if theme && mainTheme != nil && mainTheme.IsTtk() {
		if mc.Ttk != nil {
			return mc.Ttk, true
		}
		return mc.Tk, false
	}
	if mc.Tk != nil {
		return mc.Tk, false
	}
	return mc.Ttk, true
}

func (typ WidgetType) ThemeConfigure() string {
	if mainTheme == nil {
		return ""
	}
	var list []string
	opts := mainTheme.WidgetOption(typ)
	meta, _ := typ.MetaClass(true)
	for _, opt := range opts {
		if !meta.HasOption(opt.Key) {
			continue
		}
		list = append(list, fmt.Sprintf("-%v {%v}", opt.Key, opt.Value))
	}
	return strings.Join(list, " ")
}

func CreateWidgetInfo(iid string, typ WidgetType, theme bool, extra string) *WidgetInfo {
	meta, isttk := typ.MetaClass(theme)
	script := fmt.Sprintf("%v %v", meta.Command, iid)
	if theme {
		cfg := typ.ThemeConfigure()
		if cfg != "" {
			script += " " + cfg
		}
	}
	if extra != "" {
		script += " " + extra
	}
	err := eval(script)
	if err != nil {
		return nil
	}
	return &WidgetInfo{iid, typ, isttk, meta}
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
