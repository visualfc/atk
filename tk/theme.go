// Copyright 2017 visualfc. All rights reserved.

package tk

import (
	"fmt"
	"strings"
)

type WidgetType int

const (
	WidgetTypeButton WidgetType = iota
	WidgetTypeCanvas
	WidgetTypeCheckBox
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

type ThemeWidgetOpt struct {
	Key   string
	Value string
}

type Theme interface {
	Name() string
	IsTtk() bool
	WidgetOption(typ WidgetType) []ThemeWidgetOpt
}

func SetTheme(theme Theme) {
	mainTheme = theme
}

var (
	mainTheme Theme
)

func themeWidgetCommandByType(typ WidgetType) (cmd string, ttk bool) {
	mc, ok := typeMetaMap[typ]
	if !ok {
		panic(fmt.Errorf("error find metaclass type:%v", typ))
	}
	if mainTheme != nil && mainTheme.IsTtk() {
		if mc.TtkName != "" {
			return mc.TtkName, true
		}
		return mc.TkName, false

	}
	if mc.TkName != "" {
		return mc.TkName, false
	}
	return mc.TtkName, true
}

func customWidgetCommandByType(typ WidgetType) (cmd string, ttk bool) {
	mc, ok := typeMetaMap[typ]
	if !ok {
		panic(fmt.Errorf("error find metaclass type:%v", typ))
	}
	if mc.TkName != "" {
		return mc.TkName, false
	}
	return mc.TtkName, true
}

func themeWidgetConfigure(typ WidgetType) string {
	if mainTheme == nil {
		return ""
	}
	var list []string
	opts := mainTheme.WidgetOption(typ)
	for _, opt := range opts {
		list = append(list, fmt.Sprintf("-%v {%v}", opt.Key, opt.Value))
	}
	return strings.Join(list, " ")
}

type MetaClass struct {
	Name    string
	TkName  string
	TtkName string
}

var (
	typeMetaMap = make(map[WidgetType]*MetaClass)
)

func registerMeta(typ WidgetType, name string, tkname string, ttkname string) {
	typeMetaMap[typ] = &MetaClass{name, tkname, ttkname}
}

func init() {
	registerMeta(WidgetTypeButton, "Button", "tk::button", "ttk::button")
	registerMeta(WidgetTypeCanvas, "Canvas", "tk::canvas", "")
	registerMeta(WidgetTypeCheckBox, "CheckButton", "tk::checkbutton", "ttk::checkbutton")
	registerMeta(WidgetTypeComboBox, "ComboBox", "", "ttk::combobox")
	registerMeta(WidgetTypeEntry, "Entry", "tk::entry", "ttk::entry")
	registerMeta(WidgetTypeFrame, "Frame", "tk::frame", "ttk::frame")
	registerMeta(WidgetTypeLabel, "Label", "tk::label", "ttk::label")
	registerMeta(WidgetTypeLabelFrame, "LabelFrame", "tk::labelframe", "ttk::labelframe")
	registerMeta(WidgetTypeListBox, "ListBox", "tk::listbox", "")
	registerMeta(WidgetTypeMenu, "Menu", "menu", "")
	registerMeta(WidgetTypeMenuButton, "MenuButton", "tk::menubutton", "ttk::menubutton")
	registerMeta(WidgetTypeNoteBook, "NoteBook", "", "ttk::notebook")
	registerMeta(WidgetTypePanedWindow, "PanedWindow", "tk::panedwindow", "ttk::panedwindow")
	registerMeta(WidgetTypeProgressBar, "ProgressBar", "", "ttk::progressbar")
	registerMeta(WidgetTypeRadioButton, "RadioButton", "tk::radiobutton", "ttk::radiobutton")
	registerMeta(WidgetTypeScale, "Scale", "tk::scale", "ttk::scale")
	registerMeta(WidgetTypeScrollBar, "ScrollBar", "tk::scrollbar", "ttk::scrollbar")
	registerMeta(WidgetTypeSeparator, "Separator", "", "ttk::separator")
	registerMeta(WidgetTypeSizeGrip, "SizeGrip", "", "ttk::sizegrip")
	registerMeta(WidgetTypeSpinBox, "SpinBox", "tk::spinbox", "")
	registerMeta(WidgetTypeTextEdit, "TextEdit", "tk::text", "")
	registerMeta(WidgetTypeWindow, "Window", "toplevel", "")
	registerMeta(WidgetTypeTreeView, "TreeView", "", "ttk::treeview")
}
