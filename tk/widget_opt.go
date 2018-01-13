// Copyright 2018 visualfc. All rights reserved.

package tk

import (
	"fmt"
)

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

func WidgetOptFont(font Font) *WidgetOpt {
	if font == nil {
		return nil
	}
	return &WidgetOpt{"font", font.Id()}
}

func WidgetOptWidth(width int) *WidgetOpt {
	return &WidgetOpt{"width", width}
}

func WidgetOptHeight(height int) *WidgetOpt {
	return &WidgetOpt{"height", height}
}

func WidgetOptText(text string) *WidgetOpt {
	return &WidgetOpt{"text", text}
}

func WidgetOptImage(image *Image) *WidgetOpt {
	if image == nil {
		return nil
	}
	return &WidgetOpt{"image", image.Id()}
}

func WidgetOptBorderStyle(style BorderStyle) *WidgetOpt {
	return &WidgetOpt{"relief", style}
}

func WidgetOptBorderWidth(width int) *WidgetOpt {
	return &WidgetOpt{"borderwidth", width}
}

func WidgetOptPadding(pad *Pad) *WidgetOpt {
	if pad == nil {
		return nil
	}
	return &WidgetOpt{"padding", pad}
}

func checkPaddingScript(ttk bool, opt *WidgetOpt) string {
	if pad, ok := opt.Value.(*Pad); ok {
		if ttk {
			return fmt.Sprintf("-padding {%v %v}", pad.X, pad.Y)
		} else {
			return fmt.Sprintf("-padx {%v} -pady {%v}", pad.X, pad.Y)
		}
	}
	return ""
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
