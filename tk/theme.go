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

func SetTheme(theme Theme) {
	mainTheme = theme
}

var (
	mainTheme Theme
)
