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

func PackOptIpadx(padx int) *pack_option {
	return &pack_option{"ipadx", padx}
}

func PackOptIpady(pady int) *pack_option {
	return &pack_option{"ipady", pady}
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

func PackOptAnchor(anchor Anchor) *pack_option {
	v := anchor.String()
	if v == "" {
		return nil
	}
	return &pack_option{"anchor", v}
}

func PackOptExpand(b bool) *pack_option {
	return &pack_option{"expand", boolToInt(b)}
}

func PackOptFillVertical() *pack_option {
	return &pack_option{"fill", "x"}
}

func PackOptFillHorizontal() *pack_option {
	return &pack_option{"fill", "y"}
}

func PackOptFillBoth() *pack_option {
	return &pack_option{"fill", "both"}
}

func PackOptBefore(w Widget) *pack_option {
	if !IsValidWidget(w) {
		return nil
	}
	return &pack_option{"before", w.Id()}
}

func PackOptAfter(w Widget) *pack_option {
	if !IsValidWidget(w) {
		return nil
	}
	return &pack_option{"after", w.Id()}
}

func PackOptInMaster(w Widget) *pack_option {
	if !IsValidWidget(w) {
		return nil
	}
	return &pack_option{"in", w.Id()}
}

func Pack(widget Widget, options ...*pack_option) error {
	return PackList([]Widget{widget}, options...)
}

func PackList(widgets []Widget, options ...*pack_option) error {
	var idList []string
	for _, w := range widgets {
		if IsValidWidget(w) {
			idList = append(idList, w.Id())
		}
	}
	if len(idList) == 0 {
		return os.ErrInvalid
	}
	var optList []string
	for _, opt := range options {
		if opt == nil {
			continue
		}
		optList = append(optList, fmt.Sprintf("-%v {%v}", opt.key, opt.value))
	}
	script := fmt.Sprintf("pack %v", strings.Join(idList, " "))
	if len(optList) > 0 {
		script += " " + strings.Join(optList, " ")
	}
	return eval(script)
}
