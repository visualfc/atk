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

func (t *ttkTheme) WidgetOption(typ WidgetType) []WidgetOpt {
	return nil
}

func (t *ttkTheme) ThemeIdList() []string {
	return ttk_theme_list
}

func (t *ttkTheme) SetThemeId(id string) error {
	for _, v := range ttk_theme_list {
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
	ttk_theme_list []string
	TtkTheme       = &ttkTheme{}
)

func init() {
	registerInit(func() {
		s, _ := evalAsString("ttk::themes")
		ttk_theme_list = SplitTkList(s)
	})
}
