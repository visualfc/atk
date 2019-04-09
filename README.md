# atk
Another Golang Tcl/Tk binding GUI ToolKit


### Install Tcl/Tk

http://www.tcl-lang.org

* MacOS X, Windows

	https://www.activestate.com/activetcl/downloads

* Ubuntu

	$ sudo apt install tk-dev

* CentOS

	$ sudo yum install tk-devel

### Demo

https://github.com/visualfc/atk_demo

### Sample
```go
package main

import (
	"github.com/visualfc/atk/tk"
)

type Window struct {
	*tk.Window
}

func NewWindow() *Window {
	mw := &Window{tk.RootWindow()}
	lbl := tk.NewLabel(mw, "Hello ATK")
	btn := tk.NewButton(mw, "Quit")
	btn.OnCommand(func() {
		tk.Quit()
	})
	tk.NewVPackLayout(mw).AddWidgets(lbl, tk.NewLayoutSpacer(mw, 0, true), btn)
	mw.ResizeN(300, 200)
	return mw
}

func main() {
	tk.MainLoop(func() {
		mw := NewWindow()
		mw.SetTitle("ATK Sample")
		mw.Center()
		mw.ShowNormal()
	})
}
```