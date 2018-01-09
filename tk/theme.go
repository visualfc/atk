// Copyright 2018 visualfc. All rights reserved.

package tk

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
