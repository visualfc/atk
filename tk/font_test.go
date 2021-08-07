// Copyright 2018 visualfc. All rights reserved.

package tk

import (
	"testing"
)

func init() {
	registerTest("Font", testFont)
}

func testFont(t *testing.T) {
	font := NewUserFont("Courier", 18, FontAttrBold(), FontAttrItalic(), FontAttrUnderline(), FontAttrOverstrike())
	defer font.Destroy()

	fname := font.Family()
	font.SetFamily("Courier")
	if v := font.Family(); v != fname {
		t.Fatal(v)
	}

	font.SetSize(20)
	if v := font.Size(); v != 20 {
		t.Fatal(v, 20)
	}

	font.SetBold(true)
	if v := font.IsBold(); v != true {
		t.Fatal(v)
	}

	font.SetItalic(true)
	if v := font.IsItalic(); v != true {
		t.Fatal(v)
	}

	font.SetUnderline(true)
	if v := font.IsUnderline(); v != true {
		t.Fatal(v)
	}

	font.SetOverstrike(true)
	if v := font.IsOverstrike(); v != true {
		t.Fatal(v)
	}
}
