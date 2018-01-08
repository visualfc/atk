// Copyright 2018 visualfc. All rights reserved.

package tk

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

type MetaType struct {
	Type string
	Tk   *MetaClass
	Ttk  *MetaClass
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
