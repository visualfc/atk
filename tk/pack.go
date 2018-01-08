// Copyright 2018 visualfc. All rights reserved.

package tk

import (
	"fmt"
	"os"
	"strings"
)

type PackOpt struct {
	key   string
	value interface{}
}

func PackOptPadx(padx int) *PackOpt {
	return &PackOpt{"padx", padx}
}

func PackOptPady(pady int) *PackOpt {
	return &PackOpt{"pady", pady}
}

func PackOptIpadx(padx int) *PackOpt {
	return &PackOpt{"ipadx", padx}
}

func PackOptIpady(pady int) *PackOpt {
	return &PackOpt{"ipady", pady}
}

func PackOptSideTop() *PackOpt {
	return &PackOpt{"side", "top"}
}

func PackOptSideBottom() *PackOpt {
	return &PackOpt{"side", "bottom"}
}

func PackOptSideLeft() *PackOpt {
	return &PackOpt{"side", "left"}
}

func PackOptSideRight() *PackOpt {
	return &PackOpt{"side", "right"}
}

func PackOptAnchor(anchor Anchor) *PackOpt {
	v := anchor.String()
	if v == "" {
		return nil
	}
	return &PackOpt{"anchor", v}
}

func PackOptExpand(b bool) *PackOpt {
	return &PackOpt{"expand", boolToInt(b)}
}

func PackOptFillVertical() *PackOpt {
	return &PackOpt{"fill", "x"}
}

func PackOptFillHorizontal() *PackOpt {
	return &PackOpt{"fill", "y"}
}

func PackOptFillBoth() *PackOpt {
	return &PackOpt{"fill", "both"}
}

func PackOptBefore(w Widget) *PackOpt {
	if !IsValidWidget(w) {
		return nil
	}
	return &PackOpt{"before", w.Id()}
}

func PackOptAfter(w Widget) *PackOpt {
	if !IsValidWidget(w) {
		return nil
	}
	return &PackOpt{"after", w.Id()}
}

func PackOptInMaster(w Widget) *PackOpt {
	if !IsValidWidget(w) {
		return nil
	}
	return &PackOpt{"in", w.Id()}
}

func Pack(widget Widget, options ...*PackOpt) error {
	return PackList([]Widget{widget}, options...)
}

func PackList(widgets []Widget, options ...*PackOpt) error {
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
