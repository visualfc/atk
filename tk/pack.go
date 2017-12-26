// Copyright 2017 visualfc. All rights reserved.

package tk

import (
	"fmt"
	"os"
	"strings"
)

type pack_option struct {
	key   string
	value interface{}
}

func PackOptPadx(padx int) *pack_option {
	return &pack_option{"padx", padx}
}

func PackOptPady(pady int) *pack_option {
	return &pack_option{"pady", pady}
}

func PackOptSideTop() *pack_option {
	return &pack_option{"side", "top"}
}

func PackOptSideBottom() *pack_option {
	return &pack_option{"side", "bottom"}
}

func PackOptSideLeft() *pack_option {
	return &pack_option{"side", "left"}
}

func PackOptSideRight() *pack_option {
	return &pack_option{"side", "right"}
}

func PackOptAnchorNorth() *pack_option {
	return &pack_option{"anchor", "n"}
}

func PackOptAnchorSouth() *pack_option {
	return &pack_option{"anchor", "s"}
}

func PackOptAnchorWest() *pack_option {
	return &pack_option{"anchor", "w"}
}

func PackOptAnchorEast() *pack_option {
	return &pack_option{"anchor", "e"}
}

func PackOptExpand(b bool) *pack_option {
	return &pack_option{"expand", b}
}

func PackOptFillVertically() *pack_option {
	return &pack_option{"fill", "x"}
}

func PackOptFillHorizontally() *pack_option {
	return &pack_option{"fill", "x"}
}

func PackOptFillBoth() *pack_option {
	return &pack_option{"fill", "both"}
}

func PackOptAfter(w Widget) *pack_option {
	if !IsValidWidget(w) {
		return nil
	}
	return &pack_option{"after", w.Id()}
}

func Pack(w Widget, options ...*pack_option) error {
	if !IsValidWidget(w) {
		return os.ErrInvalid
	}
	var optList []string
	for _, opt := range options {
		if opt == nil {
			continue
		}
		optList = append(optList, fmt.Sprintf("-%v {%v}", opt.key, opt.value))
	}
	if len(optList) > 0 {
		return eval(fmt.Sprintf("pack %v %v", w.Id(), strings.Join(optList, " ")))
	}
	return eval(fmt.Sprintf("pack %v", w.Id()))
}
