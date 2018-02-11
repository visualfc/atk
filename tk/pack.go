// Copyright 2018 visualfc. All rights reserved.

package tk

import (
	"fmt"
	"strings"
)

func PackAttrSide(side Side) *LayoutAttr {
	return &LayoutAttr{"side", side}
}

func PackAttrSideLeft() *LayoutAttr {
	return &LayoutAttr{"side", "left"}
}

func PackAttrSideRight() *LayoutAttr {
	return &LayoutAttr{"side", "right"}
}

func PackAttrSideTop() *LayoutAttr {
	return &LayoutAttr{"side", "top"}
}

func PackAttrSideBottom() *LayoutAttr {
	return &LayoutAttr{"side", "bottom"}
}

func PackAttrPadx(padx int) *LayoutAttr {
	return &LayoutAttr{"padx", padx}
}

func PackAttrPady(pady int) *LayoutAttr {
	return &LayoutAttr{"pady", pady}
}

func PackAttrIpadx(padx int) *LayoutAttr {
	return &LayoutAttr{"ipadx", padx}
}

func PackAttrIpady(pady int) *LayoutAttr {
	return &LayoutAttr{"ipady", pady}
}

func PackAttrAnchor(anchor Anchor) *LayoutAttr {
	v := anchor.String()
	if v == "" {
		return nil
	}
	return &LayoutAttr{"anchor", v}
}

func PackAttrExpand(b bool) *LayoutAttr {
	return &LayoutAttr{"expand", boolToInt(b)}
}

func PackAttrFill(fill Fill) *LayoutAttr {
	return &LayoutAttr{"fill", fill}
}

func PackAttrFillX() *LayoutAttr {
	return &LayoutAttr{"fill", "x"}
}

func PackAttrFillY() *LayoutAttr {
	return &LayoutAttr{"fill", "y"}
}

func PackAttrFillBoth() *LayoutAttr {
	return &LayoutAttr{"fill", "both"}
}

func PackAttrFillNone() *LayoutAttr {
	return &LayoutAttr{"fill", "none"}
}

func PackAttrBefore(w Widget) *LayoutAttr {
	if !IsValidWidget(w) {
		return nil
	}
	return &LayoutAttr{"before", w.Id()}
}

func PackAttrAfter(w Widget) *LayoutAttr {
	if !IsValidWidget(w) {
		return nil
	}
	return &LayoutAttr{"after", w.Id()}
}

func PackAttrInMaster(w Widget) *LayoutAttr {
	if !IsValidWidget(w) {
		return nil
	}
	return &LayoutAttr{"in", w.Id()}
}

var (
	packAttrKeys = []string{
		"side",
		"padx", "pady",
		"ipadx", "ipady",
		"anchor",
		"expand",
		"fill",
		"before",
		"after",
		"in",
	}
)

func Pack(widget Widget, attributes ...*LayoutAttr) error {
	return PackList([]Widget{widget}, attributes...)
}

func PackRemove(widget Widget) error {
	if !IsValidWidget(widget) {
		return ErrInvalid
	}
	widget = checkLayoutWidget(widget)
	return eval("pack forget " + widget.Id())
}

func PackList(widgets []Widget, attributes ...*LayoutAttr) error {
	var idList []string
	for _, w := range widgets {
		if IsValidWidget(w) {
			w = checkLayoutWidget(w)
			idList = append(idList, w.Id())
		}
	}
	if len(idList) == 0 {
		return ErrInvalid
	}
	var attrList []string
	for _, attr := range attributes {
		if attr == nil || !isValidKey(attr.key, packAttrKeys) {
			continue
		}
		attrList = append(attrList, fmt.Sprintf("-%v {%v}", attr.key, attr.value))
	}
	script := fmt.Sprintf("pack %v", strings.Join(idList, " "))
	if len(attrList) > 0 {
		script += " " + strings.Join(attrList, " ")
	}
	return eval(script)
}
