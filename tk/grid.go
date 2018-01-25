// Copyright 2018 visualfc. All rights reserved.

package tk

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type GridAttr struct {
	key   string
	value interface{}
}

type GridIndexAttr struct {
	key   string
	value interface{}
}

func GridAttrColumn(n int) *GridAttr {
	return &GridAttr{"column", n}
}

func GridAttrColumnSpan(n int) *GridAttr {
	return &GridAttr{"columnspan", n}
}

func GridAttrRow(n int) *GridAttr {
	return &GridAttr{"row", n}
}

func GridAttrRowSpan(n int) *GridAttr {
	return &GridAttr{"rowspan", n}
}

func GridAttrInMaster(w Widget) *GridAttr {
	if !IsValidWidget(w) {
		return nil
	}
	return &GridAttr{"in", w.Id()}
}

func GridAttrIpadx(padx int) *GridAttr {
	return &GridAttr{"ipadx", padx}
}

func GridAttrIpady(pady int) *GridAttr {
	return &GridAttr{"ipady", pady}
}

func GridAttrPadx(padx int) *GridAttr {
	return &GridAttr{"padx", padx}
}

func GridAttrPady(pady int) *GridAttr {
	return &GridAttr{"pady", pady}
}

func GridAttrSticky(v Sticky) *GridAttr {
	return &GridAttr{"sticky", v}
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

func GridIndexUniform(groupname string) *GridIndexAttr {
	return &GridIndexAttr{"uniform", groupname}
}

func Grid(widget Widget, attributes ...*GridAttr) error {
	return GridList([]Widget{widget}, attributes...)
}

func GridList(widgets []Widget, attributes ...*GridAttr) error {
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
		if attr == nil {
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
