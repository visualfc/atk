// Copyright 2018 visualfc. All rights reserved.

package tk

type WidgetOpt struct {
	Key   string
	Value interface{}
}

func WidgetOptInitId(id string) *WidgetOpt {
	return &WidgetOpt{"init_id", id}
}

func WidgetOptInitTheme(use bool) *WidgetOpt {
	return &WidgetOpt{"init_theme", use}
}

func checkInitId(options []*WidgetOpt) string {
	for _, opt := range options {
		if opt != nil && opt.Key == "init_id" {
			if id, ok := opt.Value.(string); ok {
				return id
			}
		}
	}
	return ""
}

func checkInitTheme(options []*WidgetOpt) bool {
	for _, opt := range options {
		if opt != nil && opt.Key == "init_theme" {
			if use, ok := opt.Value.(bool); ok {
				return use
			}
		}
	}
	return mainTheme != nil
}
