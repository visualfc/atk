// Copyright 2017 visualfc. All rights reserved.

package tk

import (
	"fmt"
	"os"
	"strings"
)

var (
	fnGenFontId = NewGenIntFunc(1024)
)

func makeFontName() string {
	return fmt.Sprintf("go_font_%v", <-makeActionFunc())
}

type font_option struct {
	key   string
	value interface{}
}

//func FontOptFamily(family string) *font_option {
//	return &font_option{"family",family}
//}

//func FontOptSize(size int) *font_option {
//	return &font_option{"size",size}
//}

func FontOptBold() *font_option {
	return &font_option{"weight", "bold"}
}

func FontOptItalic() *font_option {
	return &font_option{"slant", "italic"}
}

func FontOptUnderline() *font_option {
	return &font_option{"underline", 1}
}

func FontOptOverstrike() *font_option {
	return &font_option{"overstrike", 1}
}

type FontDescription struct {
	info string
}

func (f *FontDescription) String() string {
	return f.info
}

type Font interface {
	Id() string
	IsValid() bool
	String() string
	Family() string
	Size() int
	IsBold() bool
	IsItalic() bool
	IsUnderline() bool
	IsOverstrike() bool
}

type BaseFont struct {
	id string
}

func (f *BaseFont) Id() string {
	return f.id
}

func (f *BaseFont) IsValid() bool {
	return f.id != ""
}

func (w *BaseFont) String() string {
	if w.id == "" {
		return "invalid"
	}
	r, _ := evalAsString(fmt.Sprintf("font actual %v", w.id))
	return r
}

func (w *BaseFont) Family() string {
	r, _ := evalAsString(fmt.Sprintf("font actual %v -family", w.id))
	return r
}

func (w *BaseFont) Size() int {
	r, _ := evalAsInt(fmt.Sprintf("font actual %v -size", w.id))
	return r
}

func (w *BaseFont) IsBold() bool {
	r, _ := evalAsString(fmt.Sprintf("font actual %v -weight", w.id))
	return r == "bold"
}

func (w *BaseFont) IsItalic() bool {
	r, _ := evalAsString(fmt.Sprintf("font actual %v -slant", w.id))
	return r == "italic"
}

func (w *BaseFont) IsUnderline() bool {
	r, _ := evalAsBool(fmt.Sprintf("font actual %v -underline", w.id))
	return r
}

func (w *BaseFont) IsOverstrike() bool {
	r, _ := evalAsBool(fmt.Sprintf("font actual %v -overstrike", w.id))
	return r
}

func (w *BaseFont) MeasureTextWidth(text string) int {
	r, _ := evalAsInt(fmt.Sprintf("font measure %v {%v}", w.id, text))
	return r
}

func (w *BaseFont) Clone() *UserFont {
	iid := makeFontName()
	script := fmt.Sprintf("font create %v %v", iid, w.String())
	if eval(script) != nil {
		return nil
	}
	return &UserFont{BaseFont{iid}}
}

type UserFont struct {
	BaseFont
}

func (f *UserFont) Destroy() error {
	if f.id == "" {
		return os.ErrInvalid
	}
	eval(fmt.Sprintf("font delete %v", f.id))
	f.id = ""
	return nil
}

func NewUserFont(family string, size int, options ...*font_option) *UserFont {
	var optList []string
	for _, opt := range options {
		if opt == nil {
			continue
		}
		optList = append(optList, fmt.Sprintf("-%v {%v}", opt.key, opt.value))
	}
	iid := makeFontName()
	script := fmt.Sprintf("font create %v -family {%v} -size %v", iid, family, size)
	if len(optList) > 0 {
		script += " " + strings.Join(optList, " ")
	}
	err := eval(script)
	if err != nil {
		return nil
	}
	return &UserFont{BaseFont{iid}}
}

func NewUserFontFromDescription(fd *FontDescription) *UserFont {
	iid := makeFontName()
	script := fmt.Sprintf("font create %v", iid)
	if fd != nil {
		script += " " + fd.String()
	}
	err := eval(script)
	if err != nil {
		return nil
	}
	return &UserFont{BaseFont{iid}}
}

func NewUserFontFromClone(font Font) *UserFont {
	if font == nil {
		return nil
	}
	iid := makeFontName()
	script := fmt.Sprintf("font create %v", iid)
	if font != nil {
		script += " " + font.String()
	}
	err := eval(script)
	if err != nil {
		return nil
	}
	return &UserFont{BaseFont{iid}}
}

func (w *UserFont) SetFamily(family string) *UserFont {
	eval(fmt.Sprintf("font configure %v -family {%v}", w.id, family))
	return w
}

func (w *UserFont) SetSize(size int) *UserFont {
	eval(fmt.Sprintf("font configure %v -size {%v}", w.id, size))
	return w
}

func (w *UserFont) SetBold(bold bool) *UserFont {
	var v string
	if bold {
		v = "bold"
	} else {
		v = "normal"
	}
	eval(fmt.Sprintf("font configure %v -weight {%v}", w.id, v))
	return w
}

func (w *UserFont) SetItalic(italic bool) *UserFont {
	var v string
	if italic {
		v = "italic"
	} else {
		v = "roman"
	}
	eval(fmt.Sprintf("font configure %v -slant {%v}", w.id, v))
	return w
}

func (w *UserFont) SetUnderline(underline bool) *UserFont {
	eval(fmt.Sprintf("font configure %v -underline {%v}", w.id, boolToInt(underline)))
	return w
}

func (w *UserFont) SetOverstrike(overstrike bool) *UserFont {
	eval(fmt.Sprintf("font configure %v -overstrike {%v}", w.id, boolToInt(overstrike)))
	return w
}

func FontFamilieList() []string {
	s, err := evalAsString("font families")
	if err != nil {
		return nil
	}
	return SplitTkList(s)
}

//tk system default font
type SysFont struct {
	BaseFont
}

type SysFontType int

const (
	SysDefaultFont SysFontType = 0
	SysTextFont
	SysFixedFont
	SysMenuFont
	SysHeadingFont
	SysCaptionFont
	SysSmallCaptionFont
	SysIconFont
	SysTooltipFont
)

var (
	sysFontNameList = []string{
		"TKDefaultFont",
		"TKTextFont",
		"TKFixedFont",
		"TKMenuFont",
		"TKHeadingFont",
		"TKCaptionFont",
		"TKSmallCaptionFont",
		"TKIconFont",
		"TKTooltipFont",
	}
	sysFontList []*SysFont
)

func init() {
	for _, name := range sysFontNameList {
		sysFontList = append(sysFontList, &SysFont{BaseFont{name}})
	}
}

func LoadSysFont(typ SysFontType) *SysFont {
	if int(typ) >= 0 && int(typ) < len(sysFontList) {
		return sysFontList[typ]
	}
	return nil
}
