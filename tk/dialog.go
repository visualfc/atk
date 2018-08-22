// Copyright 2018 visualfc. All rights reserved.

package tk

import (
	"fmt"
)

//tk_chooseColor — pops up a dialog box for the user to select a color.
func ChooseColor(parent Widget, title string, initcolor string) (string, error) {
	script := fmt.Sprintf("tk_chooseColor")
	if parent != nil {
		script += " -parent " + parent.Id()
	}
	if initcolor != "" {
		script += " -initialcolor " + initcolor
	}
	if title != "" {
		setObjText("atk_tmp_title", title)
		script += " -title $atk_tmp_title"
	}
	return evalAsString(script)
}

//tk_chooseDirectory — pops up a dialog box for the user to select a directory.
func ChooseDirectory(parent Widget, title string, initialdir string, mustexist bool) (string, error) {
	script := fmt.Sprintf("tk_chooseDirectory")
	if parent != nil {
		script += " -parent " + parent.Id()
	}
	if initialdir != "" {
		setObjText("atk_tmp_initialdir", initialdir)
		script += " -initialdir $atk_tmp_initialdir"
	}
	if mustexist {
		script += " -mustexist true"
	}
	if title != "" {
		setObjText("atk_tmp_title", title)
		script += " -title $atk_tmp_title"
	}
	return evalAsString(script)
}

type FileType struct {
	Info string
	Ext  string
}

func (v FileType) String() string {
	return fmt.Sprintf("{%v} {%v}", v.Info, v.Ext)
}

//tk_getOpenFile, tk_getSaveFile — pop up a dialog box for the user to select a file to open or save.
func GetOpenFile(parent Widget, title string, filetypes []FileType, initialdir string, initialfile string) (string, error) {
	script := fmt.Sprintf("tk_getOpenFile")
	if parent != nil {
		script += " -parent " + parent.Id()
	}
	if filetypes != nil {
		var info []string
		for _, v := range filetypes {
			info = append(info, v.String())
		}
		setObjTextList("atk_tmp_filetypes", info)
		script += " -filetypes $atk_tmp_filetypes"
	}
	if initialdir != "" {
		setObjText("atk_tmp_initialdir", initialdir)
		script += " -initialdir $atk_tmp_initialdir"
	}
	if initialfile != "" {
		setObjText("atk_tmp_initialfile", initialfile)
		script += " -initialfile $atk_tmp_initialfile"
	}
	if title != "" {
		setObjText("atk_tmp_title", title)
		script += " -title $atk_tmp_title"
	}
	return evalAsString(script)
}

func GetOpenMultipleFile(parent Widget, title string, filetypes []FileType, initialdir string, initialfile string) ([]string, error) {
	script := fmt.Sprintf("tk_getOpenFile")
	if parent != nil {
		script += " -parent " + parent.Id()
	}
	if filetypes != nil {
		var info []string
		for _, v := range filetypes {
			info = append(info, v.String())
		}
		setObjTextList("atk_tmp_filetypes", info)
		script += " -filetypes $atk_tmp_filetypes"
	}
	if initialdir != "" {
		setObjText("atk_tmp_initialdir", initialdir)
		script += " -initialdir $atk_tmp_initialdir"
	}
	if initialfile != "" {
		setObjText("atk_tmp_initialfile", initialfile)
		script += " -initialfile $atk_tmp_initialfile"
	}
	if title != "" {
		setObjText("atk_tmp_title", title)
		script += " -title $atk_tmp_title"
	}
	script += " -multiple true"
	return evalAsStringList(script)
}

//tk_getOpenFile, tk_getSaveFile — pop up a dialog box for the user to select a file to open or save.
func GetSaveFile(parent Widget, title string, confirmoverwrite bool, defaultextension string, filetypes []FileType, initialdir string, initialfile string) (string, error) {
	script := fmt.Sprintf("tk_getSaveFile")
	if parent != nil {
		script += " -parent " + parent.Id()
	}
	script += " -confirmoverwrite " + fmt.Sprint(confirmoverwrite)
	if defaultextension != "" {
		setObjText("atk_tmp_defaultextension", defaultextension)
		script += " -defaultextension $atk_tmp_defaultextension"
	}
	if filetypes != nil {
		var info []string
		for _, v := range filetypes {
			info = append(info, v.String())
		}
		setObjTextList("atk_tmp_filetypes", info)
		script += " -filetypes $atk_tmp_filetypes"
	}
	if initialdir != "" {
		setObjText("atk_tmp_initialdir", initialdir)
		script += " -initialdir $atk_tmp_initialdir"
	}
	if initialfile != "" {
		setObjText("atk_tmp_initialfile", initialfile)
		script += " -initialfile $atk_tmp_initialfile"
	}
	if title != "" {
		setObjText("atk_tmp_title", title)
		script += " -title $atk_tmp_title"
	}
	return evalAsString(script)
}

type MessageBoxIcon int

const (
	MessageBoxIconNone MessageBoxIcon = iota
	MessageBoxIconError
	MessageBoxIconInfo
	MessageBoxIconQuestion
	MessageBoxIconWarning
)

var (
	messageBoxIconName = []string{"", "error", "info", "question", "warning"}
)

func (v MessageBoxIcon) String() string {
	if v >= 0 && int(v) < len(messageBoxIconName) {
		return messageBoxIconName[v]
	}
	return ""
}

type MessageBoxType int

const (
	MessageBoxTypeOk MessageBoxType = iota
	MessageBoxTypeOkCancel
	MessageBoxTypeAbortRetryIgnore
	MessageBoxTypeRetryCancel
	MessageBoxTypeYesNo
	MessageBoxTypeYesNoCancel
)

var (
	messageBoxTypeName = []string{"ok", "okcancel", "abortretryignore", "retrycancel", "yesno", "yesnocancel"}
)

func (v MessageBoxType) String() string {
	if v >= 0 && int(v) < len(messageBoxTypeName) {
		return messageBoxTypeName[v]
	}
	return ""
}

//tk_messageBox — pops up a message window and waits for user response.
func MessageBox(parent Widget, title string, message string, detail string, defaultbutton string, icon MessageBoxIcon, typ MessageBoxType) (string, error) {
	script := fmt.Sprintf("tk_messageBox")
	if parent != nil {
		script += " -parent " + parent.Id()
	}
	if defaultbutton != "" {
		setObjText("atk_tmp_defaultbutton", defaultbutton)
		script += " -default $atk_tmp_defaultbutton"
	}
	if message != "" {
		setObjText("atk_tmp_message", message)
		script += " -message $atk_tmp_message"
	}
	sicon := icon.String()
	if sicon != "" {
		script += " -icon " + sicon
	}
	styp := typ.String()
	if styp != "" {
		script += " -type " + styp
	}
	if detail != "" {
		setObjText("atk_tmp_detail", detail)
		script += " -detail $atk_tmp_detail"
	}
	if title != "" {
		setObjText("atk_tmp_title", title)
		script += " -title $atk_tmp_title"
	}
	return evalAsString(script)
}
