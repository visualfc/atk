// Copyright 2017 visualfc. All rights reserved.

package tk

import (
	"testing"
)

func init() {
	InitTest()
}

func TestFont(t *testing.T) {
	font := NewUserFont("Courier", 18, FontOptBold(), FontOptItalic(), FontOptUnderline(), FontOptOverstrike())
	defer font.Destroy()

	fname := font.Family()
	if v := font.SetFamily("Courier").Family(); v != fname {
		t.Fatal(v)
	}

	if v := font.SetSize(20).Size(); v != 20 {
		t.Fatal(v, 20)
	}

	if v := font.SetBold(true).IsBold(); v != true {
		t.Fatal(v)
	}

	if v := font.SetItalic(true).IsItalic(); v != true {
		t.Fatal(v)
	}

	if v := font.SetUnderline(true).IsUnderline(); v != true {
		t.Fatal(v)
	}

	if v := font.SetOverstrike(true).IsOverstrike(); v != true {
		t.Fatal(v)
	}
}
