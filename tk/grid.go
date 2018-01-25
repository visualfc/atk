// Copyright 2018 visualfc. All rights reserved.

package tk

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GridAttrColumn(n int) *LayoutAttr {
	return &LayoutAttr{"column", n}
}

func GridAttrColumnSpan(n int) *LayoutAttr {
	return &LayoutAttr{"columnspan", n}
}

func GridAttrRow(n int) *LayoutAttr {
	return &LayoutAttr{"row", n}
}

func GridAttrRowSpan(n int) *LayoutAttr {
	return &LayoutAttr{"rowspan", n}
}

func GridAttrInMaster(w Widget) *LayoutAttr {
	if !IsValidWidget(w) {
		return nil
	}
	return &LayoutAttr{"in", w.Id()}
}

func GridAttrIpadx(padx int) *LayoutAttr {
	return &LayoutAttr{"ipadx", padx}
}

func GridAttrIpady(pady int) *LayoutAttr {
	return &LayoutAttr{"ipady", pady}
}

func GridAttrPadx(padx int) *LayoutAttr {
	return &LayoutAttr{"padx", padx}
}

func GridAttrPady(pady int) *LayoutAttr {
	return &LayoutAttr{"pady", pady}
}

func GridAttrSticky(v Sticky) *LayoutAttr {
	return &LayoutAttr{"sticky", v}
}

type GridIndexAttr struct {
	key   string
	value interface{}
}

func GridIndexAttrMinSize(amount int) *GridIndexAttr {
	return &GridIndexAttr{"minsize", amount}
}

func GridIndexAttrPad(amount int) *GridIndexAttr {
	return &GridIndexAttr{"pad", amount}
}

func GridIndexAttrWeight(value int) *GridIndexAttr {
	return &GridIndexAttr{"weight", value}
}

func GridIndexAttrUniform(groupname string) *GridIndexAttr {
	return &GridIndexAttr{"uniform", groupname}
}

func Grid(widget Widget, attributes ...*LayoutAttr) error {
	return GridList([]Widget{widget}, attributes...)
}

func GridRemove(widget Widget) error {
	if !IsValidWidget(widget) {
		return os.ErrInvalid
	}
	return eval("grid forget " + widget.Id())
}

var (
	gridAttrKeys = []string{
		"column", "columnspan",
		"row", "rowspan",
		"in",
		"ipadx", "ipady",
		"padx", "pady",
		"sticky",
	}
	gridIndexAttrKeys = []string{
		"minsize",
		"pad",
		"weight",
		"uniform",
	}
)

func GridList(widgets []Widget, attributes ...*LayoutAttr) error {
	var idList []string
	for _, w := range widgets {
		if IsValidWidget(w) {
			idList = append(idList, w.Id())
		} else {
			idList = append(idList, "x")
		}
	}
	if len(idList) == 0 {
		return os.ErrInvalid
	}
	var attrList []string
	for _, attr := range attributes {
		if attr == nil || !isValidKey(attr.key, gridAttrKeys) {
			continue
		}
		attrList = append(attrList, fmt.Sprintf("-%v {%v}", attr.key, attr.value))
	}
	script := fmt.Sprintf("grid %v", strings.Join(idList, " "))
	if len(attrList) > 0 {
		script += " " + strings.Join(attrList, " ")
	}
	return eval(script)
}

// row index from 0; -1=all
func GridRowIndex(master Widget, index int, attributes ...*GridIndexAttr) error {
	return gridIndex(master, true, index, attributes)
}

// column index from 0; -1=all
func GridColumnIndex(master Widget, index int, attributes ...*GridIndexAttr) error {
	return gridIndex(master, false, index, attributes)
}

func gridIndex(master Widget, row bool, index int, attributes []*GridIndexAttr) error {
	if master == nil {
		master = mainWindow
	}
	var sindex string
	if index < 0 {
		sindex = "all"
	} else {
		sindex = strconv.Itoa(index)
	}
	var attrList []string
	for _, attr := range attributes {
		if attr == nil || !isValidKey(attr.key, gridIndexAttrKeys) {
			continue
		}
		attrList = append(attrList, fmt.Sprintf("-%v {%v}", attr.key, attr.value))
	}
	var script string
	if row {
		script = fmt.Sprintf("grid rowconfigure %v %v", master.Id(), sindex)
	} else {
		script = fmt.Sprintf("grid columnconfigure %v %v", master.Id(), sindex)
	}
	if len(attrList) > 0 {
		script += " " + strings.Join(attrList, " ")
	}
	return eval(script)
}
