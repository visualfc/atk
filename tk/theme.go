// Copyright 2018 visualfc. All rights reserved.

package tk

type NativeAttr struct {
	Key   string
	Value string
}

type Theme interface {
	Name() string
	IsTtk() bool
	InitAttributes(typ WidgetType) []NativeAttr
}

func SetMainTheme(theme Theme) {
	mainTheme = theme
}

func MainTheme() Theme {
	return mainTheme
}

func HasTheme() bool {
	return mainTheme != nil
}

var (
	mainTheme Theme
)
