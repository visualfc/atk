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

var (
	typeMetaMap = make(map[WidgetType]*MetaType)
)
