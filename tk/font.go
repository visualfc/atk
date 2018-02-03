// Copyright 2018 visualfc. All rights reserved.

package tk

import (
	"fmt"
	"strings"
)

type Font interface {
	Id() string
	IsValid() bool
	String() string
	Description() string
	Family() string
	Size() int
	IsBold() bool
	IsItalic() bool
	IsUnderline() bool
	IsOverstrike() bool
}

type FontAttr struct {
	key   string
	value interface{}
}

func FontAttrBold() *FontAttr {
	return &FontAttr{"weight", "bold"}
}

func FontAttrItalic() *FontAttr {
	return &FontAttr{"slant", "italic"}
}

func FontAttrUnderline() *FontAttr {
	return &FontAttr{"underline", 1}
}

func FontAttrOverstrike() *FontAttr {
	return &FontAttr{"overstrike", 1}
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
	return fmt.Sprintf("Font{%v}", w.id)
}

func (w *BaseFont) Description() string {
	if w.id == "" {
		return ""
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
	r, _ := evalAsInt(fmt.Sprintf("font measure %v {%v}", w.id, buildTkString(text)))
	return r
}

func (w *BaseFont) Clone() *UserFont {
	iid := makeNamedId("atk_font")
	script := fmt.Sprintf("font create %v %v", iid, w.Description())
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
		return ErrInvalid
	}
	eval(fmt.Sprintf("font delete %v", f.id))
	f.id = ""
	return nil
}

func NewUserFont(family string, size int, attributes ...*FontAttr) *UserFont {
	var attrList []string
	for _, attr := range attributes {
		if attr == nil {
			continue
		}
		if s, ok := attr.value.(string); ok {
			attrList = append(attrList, fmt.Sprintf("-%v {%v}", attr.key, buildTkString(s)))
			continue
		}
		attrList = append(attrList, fmt.Sprintf("-%v {%v}", attr.key, attr.value))
	}
	iid := makeNamedId("atk_font")
	script := fmt.Sprintf("font create %v -family {%v} -size %v", iid, buildTkString(family), size)
	if len(attrList) > 0 {
		script += " " + strings.Join(attrList, " ")
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
	iid := makeNamedId("atk_font")
	script := fmt.Sprintf("font create %v", iid)
	if font != nil {
		script += " " + font.Description()
	}
	err := eval(script)
	if err != nil {
		return nil
	}
	return &UserFont{BaseFont{iid}}
}

func (w *UserFont) SetFamily(family string) *UserFont {
	eval(fmt.Sprintf("font configure %v -family {%v}", w.id, buildTkString(family)))
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
	return FromTkList(s)
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

func parserFontResult(r string, err error) Font {
	if err != nil || r == "" {
		return nil
	}
	for _, f := range sysFontList {
		if f.Id() == r {
			return f
		}
	}
	return &UserFont{BaseFont{r}}
}
