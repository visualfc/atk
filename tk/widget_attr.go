// Copyright 2018 visualfc. All rights reserved.

package tk

import (
	"fmt"
)

// widget attribute
type WidgetAttr struct {
	Key   string
	Value interface{}
}

// setup widget init enable/disable theme
func WidgetAttrInitUseTheme(use bool) *WidgetAttr {
	return &WidgetAttr{"init_use_theme", use}
}

// setup widget font
func WidgetAttrFont(font Font) *WidgetAttr {
	if font == nil {
		return nil
	}
	return &WidgetAttr{"font", font.Id()}
}

// setup widget width
func WidgetAttrWidth(width int) *WidgetAttr {
	return &WidgetAttr{"width", width}
}

// setup widget height pixel or line
func WidgetAttrHeight(height int) *WidgetAttr {
	return &WidgetAttr{"height", height}
}

// setup widget text
func WidgetAttrText(text string) *WidgetAttr {
	return &WidgetAttr{"text", text}
}

// setup widget image
func WidgetAttrImage(image *Image) *WidgetAttr {
	if image == nil {
		return nil
	}
	return &WidgetAttr{"image", image.Id()}
}

// setup widget border style (tk relief attribute)
func WidgetAttrReliefStyle(style ReliefStyle) *WidgetAttr {
	return &WidgetAttr{"relief", style}
}

// setup widget border width
func WidgetAttrBorderWidth(width int) *WidgetAttr {
	return &WidgetAttr{"borderwidth", width}
}

// setup widget padding (ttk padding or tk padx/pady)
func WidgetAttrPadding(pad Pad) *WidgetAttr {
	return &WidgetAttr{"padding", pad}
}

// setup widget padding (ttk padding or tk padx/pady)
func WidgetAttrPaddingN(padx int, pady int) *WidgetAttr {
	return &WidgetAttr{"padding", Pad{padx, pady}}
}

func checkPaddingScript(ttk bool, attr *WidgetAttr) string {
	if pad, ok := attr.Value.(Pad); ok {
		if ttk {
			return fmt.Sprintf("-padding {%v %v}", pad.X, pad.Y)
		} else {
			return fmt.Sprintf("-padx {%v} -pady {%v}", pad.X, pad.Y)
		}
	}
	return ""
}

func checkInitUseTheme(attributes []*WidgetAttr) bool {
	for _, attr := range attributes {
		if attr != nil && attr.Key == "init_use_theme" {
			if use, ok := attr.Value.(bool); ok {
				return use
			}
		}
	}
	return mainTheme != nil
}
