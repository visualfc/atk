package tk

import "fmt"

func ClearClipboard() {
	evalAsString("clipboard clear")
}

func AppendToClipboard(text string) {
	pname := "atk_tmp_clip"
	setObjText(pname, text)
	eval(fmt.Sprintf("clipboard append $%v", pname))
}
