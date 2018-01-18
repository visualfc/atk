// Copyright 2018 visualfc. All rights reserved.

package tk

import (
	"fmt"
	"os"
	"strings"
)

type PackAttr struct {
	key   string
	value interface{}
}

func PackAttrPadx(padx int) *PackAttr {
	return &PackAttr{"padx", padx}
}

func PackAttrPady(pady int) *PackAttr {
	return &PackAttr{"pady", pady}
}

func PackAttrIpadx(padx int) *PackAttr {
	return &PackAttr{"ipadx", padx}
}

func PackAttrIpady(pady int) *PackAttr {
	return &PackAttr{"ipady", pady}
}

func PackAttrSideTop() *PackAttr {
	return &PackAttr{"side", "top"}
}

func PackAttrSideBottom() *PackAttr {
	return &PackAttr{"side", "bottom"}
}

func PackAttrSideLeft() *PackAttr {
	return &PackAttr{"side", "left"}
}

func PackAttrSideRight() *PackAttr {
	return &PackAttr{"side", "right"}
}

func PackAttrAnchor(anchor Anchor) *PackAttr {
	v := anchor.String()
	if v == "" {
		return nil
	}
	return &PackAttr{"anchor", v}
}

func PackAttrExpand(b bool) *PackAttr {
	return &PackAttr{"expand", boolToInt(b)}
}

func PackAttrFillVertical() *PackAttr {
	return &PackAttr{"fill", "x"}
}

func PackAttrFillHorizontal() *PackAttr {
	return &PackAttr{"fill", "y"}
}

func PackAttrFillBoth() *PackAttr {
	return &PackAttr{"fill", "both"}
}

func PackAttrBefore(w Widget) *PackAttr {
	if !IsValidWidget(w) {
		return nil
	}
	return &PackAttr{"before", w.Id()}
}

func PackAttrAfter(w Widget) *PackAttr {
	if !IsValidWidget(w) {
		return nil
	}
	return &PackAttr{"after", w.Id()}
}

func PackAttrInMaster(w Widget) *PackAttr {
	if !IsValidWidget(w) {
		return nil
	}
	return &PackAttr{"in", w.Id()}
}

func Pack(widget Widget, attributes ...*PackAttr) error {
	return PackList([]Widget{widget}, attributes...)
}

func PackList(widgets []Widget, attributes ...*PackAttr) error {
	var idList []string
	for _, w := range widgets {
		if IsValidWidget(w) {
			idList = append(idList, w.Id())
		}
	}
	if len(idList) == 0 {
		return os.ErrInvalid
	}
	var attrList []string
	for _, attr := range attributes {
		if attr == nil {
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
