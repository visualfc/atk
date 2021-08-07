// Copyright 2018 visualfc. All rights reserved.

package tk

import "fmt"

func ClearClipboard() error {
	return eval("clipboard clear")
}

func AppendToClipboard(text string) error {
	pname := "atk_tmp_clip"
	setObjText(pname, text)
	return eval(fmt.Sprintf("clipboard append $%v", pname))
}

func GetClipboardText() string {
	text, _ := evalAsString("clipboard get -type UTF8_STRING")
	return text
}
