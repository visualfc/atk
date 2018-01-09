// Copyright 2018 visualfc. All rights reserved.

package tk

import (
	"fmt"
	"strings"
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

func themeWidgetCommand(typ WidgetType) (cmd string, ttk bool) {
	mc, ok := typeMetaMap[typ]
	if !ok {
		panic(fmt.Errorf("error find metaclass type:%v", typ))
	}
	if mainTheme != nil && mainTheme.IsTtk() {
		if mc.Ttk != nil {
			return mc.Ttk.Command, true
		}
		return mc.Tk.Command, false
	}
	if mc.Tk != nil {
		return mc.Tk.Command, false
	}
	return mc.Ttk.Command, true
}

func GetWidgetMetaClass(typ WidgetType, theme bool) (meta *MetaClass, ttk bool) {
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

func GetWidgetThemeConfigure(typ WidgetType) string {
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

func CreateWidgetInfo(iid string, typ WidgetType, theme bool) *WidgetInfo {
	meta, isttk := GetWidgetMetaClass(typ, theme)
	script := fmt.Sprintf("%v %v", meta.Command, iid)
	if theme {
		cfg := GetWidgetThemeConfigure(typ)
		if cfg != "" {
			script += " " + cfg
		}
	}
	err := eval(script)
	if err != nil {
		return nil
	}
	return &WidgetInfo{iid, typ, isttk, meta}
}
