// Copyright 2018 visualfc. All rights reserved.

package tk

import (
	"fmt"
)

// widget option
type WidgetOpt struct {
	key   string
	value interface{}
}

// setup widget init enable/disable theme
func WidgetOptInitUseTheme(use bool) *WidgetOpt {
	return &WidgetOpt{"init_use_theme", use}
}

// setup widget font
func WidgetOptFont(font Font) *WidgetOpt {
	if font == nil {
		return nil
	}
	return &WidgetOpt{"font", font.Id()}
}

// setup widget width
func WidgetOptWidth(width int) *WidgetOpt {
	return &WidgetOpt{"width", width}
}

// setup widget height pixel or line
func WidgetOptHeight(height int) *WidgetOpt {
	return &WidgetOpt{"height", height}
}

// setup widget text
func WidgetOptText(text string) *WidgetOpt {
	return &WidgetOpt{"text", text}
}

// setup widget image
func WidgetOptImage(image *Image) *WidgetOpt {
	if image == nil {
		return nil
	}
	return &WidgetOpt{"image", image.Id()}
}

// setup widget border style (tk relief option)
func WidgetOptBorderStyle(style BorderStyle) *WidgetOpt {
	return &WidgetOpt{"relief", style}
}

// setup widget border width
func WidgetOptBorderWidth(width int) *WidgetOpt {
	return &WidgetOpt{"borderwidth", width}
}

// setup widget padding (ttk padding or tk padx/pady)
func WidgetOptPadding(pad Pad) *WidgetOpt {
	return &WidgetOpt{"padding", pad}
}

// setup widget padding (ttk padding or tk padx/pady)
func WidgetOptPaddingN(padx int, pady int) *WidgetOpt {
	return &WidgetOpt{"padding", Pad{padx, pady}}
}

func checkPaddingScript(ttk bool, opt *WidgetOpt) string {
	if pad, ok := opt.value.(Pad); ok {
		if ttk {
			return fmt.Sprintf("-padding {%v %v}", pad.X, pad.Y)
		} else {
			return fmt.Sprintf("-padx {%v} -pady {%v}", pad.X, pad.Y)
		}
	}
	return ""
}

func checkInitUseTheme(options []*WidgetOpt) bool {
	for _, opt := range options {
		if opt != nil && opt.key == "init_use_theme" {
			if use, ok := opt.value.(bool); ok {
				return use
			}
		}
	}
	return mainTheme != nil
}
