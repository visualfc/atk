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

func MakeFontName() string {
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

func (w *BaseFont) Clone() *Font {
	iid := MakeFontName()
	script := fmt.Sprintf("font create %v %v", iid, w.String())
	if eval(script) != nil {
		return nil
	}
	return &Font{BaseFont{iid}}
}

type Font struct {
	BaseFont
}

func (f *Font) Destroy() error {
	if f.id == "" {
		return os.ErrInvalid
	}
	eval(fmt.Sprintf("font delete %v", f.id))
	f.id = ""
	return nil
}

func NewFont(family string, size int, options ...*font_option) *Font {
	var optList []string
	for _, opt := range options {
		if opt == nil {
			continue
		}
		optList = append(optList, fmt.Sprintf("-%v {%v}", opt.key, opt.value))
	}
	iid := MakeFontName()
	script := fmt.Sprintf("font create %v -family {%v} -size %v", iid, family, size)
	if len(optList) > 0 {
		script += " " + strings.Join(optList, " ")
	}
	if eval(script) != nil {
		return nil
	}
	return &Font{BaseFont{iid}}
}

func (w *Font) SetFamily(family string) *Font {
	eval(fmt.Sprintf("font configure %v -family {%v}", w.id, family))
	return w
}

func (w *Font) SetSize(size int) *Font {
	eval(fmt.Sprintf("font configure %v -size {%v}", w.id, size))
	return w
}

func (w *Font) SetBold(bold bool) *Font {
	var v string
	if bold {
		v = "bold"
	} else {
		v = "normal"
	}
	eval(fmt.Sprintf("font configure %v -weight {%v}", w.id, v))
	return w
}

func (w *Font) SetItalic(italic bool) *Font {
	var v string
	if italic {
		v = "italic"
	} else {
		v = "roman"
	}
	eval(fmt.Sprintf("font configure %v -slant {%v}", w.id, v))
	return w
}

func (w *Font) SetUnderline(underline bool) *Font {
	eval(fmt.Sprintf("font configure %v -underline {%v}", w.id, boolToInt(underline)))
	return w
}

func (w *Font) SetOverstrike(overstrike bool) *Font {
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

type DefaultFontType int

const (
	DefaultGuiFont DefaultFontType = 0
	DefaultTextFont
	DefaultFixedFont
	DefaultMenuFont
	DefaultHeadingFont
	DefaultCaptionFont
	DefaultSmallCaptionFont
	DefaultIconFont
	DefaultTooltipFont
)

var (
	defaultFontList = []string{
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
)

func DefaultFont(typ DefaultFontType) *BaseFont {
	if int(typ) >= 0 && int(typ) < len(defaultFontList) {
		return &BaseFont{defaultFontList[typ]}
	}
	return nil
}
