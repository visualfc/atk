// Copyright 2018 visualfc. All rights reserved.

package tk

import (
	"fmt"
)

type ttkTheme struct {
}

func (t *ttkTheme) Name() string {
	return "ttk theme"
}

func (t *ttkTheme) IsTtk() bool {
	return true
}

func (t *ttkTheme) InitAttributes(typ WidgetType) []NativeAttr {
	return nil
}

func (t *ttkTheme) ThemeIdList() []string {
	ttk_theme_list, _ := evalAsStringList("ttk::themes")
	return ttk_theme_list
}

func (t *ttkTheme) SetThemeId(id string) error {
	for _, v := range t.ThemeIdList() {
		if v == id {
			err := eval(fmt.Sprintf("ttk::setTheme %v", id))
			return err
		}
	}
	err := fmt.Errorf("not found ttk_theme id:%v", id)
	dumpError(err)
	return err
}

func (t *ttkTheme) ThemeId() string {
	r, _ := evalAsString("ttk::style theme use")
	return r
}

var (
	TtkTheme = &ttkTheme{}
)

func init() {
	/*
	registerInit(func() {
		ttk_theme_list, _ = evalAsStringList("ttk::themes")
	})
	*/
	SetMainTheme(TtkTheme)
}
