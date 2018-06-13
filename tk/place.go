// Copyright 2018 visualfc. All rights reserved.

package tk

import (
	"fmt"
	"strings"
)

func PlaceAttrAnchor(anchor Anchor) *LayoutAttr {
	v := anchor.String()
	if v == "" {
		return nil
	}
	return &LayoutAttr{"anchor", v}
}

func PlaceAttrBorderMode(mode BorderMode) *LayoutAttr {
	v := mode.String()
	if v == "" {
		return nil
	}
	return &LayoutAttr{"bordermode", v}
}

func PlaceAttrWidth(size int) *LayoutAttr {
	return &LayoutAttr{"width", size}
}

func PlaceAttrHeight(size int) *LayoutAttr {
	return &LayoutAttr{"height", size}
}

func PlaceAttrRelWidth(size float64) *LayoutAttr {
	return &LayoutAttr{"relwidth", size}
}

func PlaceAttrRelHeight(size float64) *LayoutAttr {
	return &LayoutAttr{"relheight", size}
}

func PlaceAttrX(location int) *LayoutAttr {
	return &LayoutAttr{"x", location}
}

func PlaceAttrY(location int) *LayoutAttr {
	return &LayoutAttr{"y", location}
}

func PlaceAttrRelX(location float64) *LayoutAttr {
	return &LayoutAttr{"relx", location}
}

func PlaceAttrRelY(location float64) *LayoutAttr {
	return &LayoutAttr{"rely", location}
}

func PlaceAttrInMaster(w Widget) *LayoutAttr {
	if !IsValidWidget(w) {
		return nil
	}
	return &LayoutAttr{"in", w.Id()}
}

var (
	placeAttrKeys = []string{
		"anchor",
		"bordermode",
		"x", "y",
		"relx", "rely",
		"width", "height",
		"relwidth", "relheight",
		"in",
	}
)

func Place(widget Widget, attributes ...*LayoutAttr) error {
	if !IsValidWidget(widget) {
		return ErrInvalid
	}
	var attrList []string
	for _, attr := range attributes {
		if attr == nil || !isValidKey(attr.key, placeAttrKeys) {
			continue
		}
		attrList = append(attrList, fmt.Sprintf("-%v {%v}", attr.key, attr.value))
	}
	script := fmt.Sprintf("place %v", widget.Id())
	if len(attrList) > 0 {
		script += " " + strings.Join(attrList, " ")
	}
	return eval(script)
}

func PlaceRemove(widget Widget) error {
	if !IsValidWidget(widget) {
		return ErrInvalid
	}
	widget = checkLayoutWidget(widget)
	return eval("place forget " + widget.Id())
}
