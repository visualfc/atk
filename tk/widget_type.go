// Copyright 2018 visualfc. All rights reserved.

package tk

import (
	"fmt"
	"strings"
)

var (
	ErrorInvalidWidgetInfo  = fmt.Errorf("invalid widget info")
	ErrorNotMatchWidgetInfo = fmt.Errorf("widget info not match")
)

type WidgetType int

const (
	WidgetTypeNone WidgetType = iota
	WidgetTypeButton
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
	WidgetTypeLast
)

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
	attrs := mainTheme.InitAttributes(typ)
	_, meta, _ := typ.MetaClass(true)
	for _, attr := range attrs {
		if !meta.HasAttribute(attr.Key) {
			continue
		}
		list = append(list, fmt.Sprintf("-%v {%v}", attr.Key, attr.Value))
	}
	return strings.Join(list, " ")
}

type WidgetInfo struct {
	Type      WidgetType
	TypeName  string
	IsTtk     bool
	MetaClass *MetaClass
}

func buildWidgetAttributeScript(meta *MetaClass, ttk bool, attributes []*WidgetAttr) string {
	var list []string
	for _, attr := range attributes {
		if attr == nil {
			continue
		}
		if attr.key == "padding" {
			list = append(list, checkPaddingScript(ttk, attr))
			continue
		}
		if !meta.HasAttribute(attr.key) {
			continue
		}
		list = append(list, fmt.Sprintf("-%v {%v}", attr.key, attr.value))
	}
	return strings.Join(list, " ")
}

func CreateWidgetInfo(iid string, typ WidgetType, theme bool, attributes []*WidgetAttr) *WidgetInfo {
	typName, meta, isttk := typ.MetaClass(theme)
	script := fmt.Sprintf("%v %v", meta.Command, iid)
	if theme {
		cfg := typ.ThemeConfigure()
		if cfg != "" {
			script += " " + cfg
		}
	}
	if len(attributes) > 0 {
		extra := buildWidgetAttributeScript(meta, isttk, attributes)
		if len(extra) > 0 {
			script += " " + extra
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
	if id == "" {
		return nil
	}
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
