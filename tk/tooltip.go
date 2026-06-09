// a simple tooltip

// 参考：http://tmml.sourceforge.net/doc/tcllib/tooltip.html

// 例子：widget.Tooltip("This is a widget")

package tk

import "fmt"

var isPackReqTip = false // 保证 "package require tooltip" 只进行一次

func (w *BaseWidget) Tooltip(tip string) error {
    
    // ::tooltip::tooltip pathName ? option arg ? message 
    // tooltip::tooltip .l "This is a label widget"
    if !isPackReqTip {
        
        _ = eval(fmt.Sprintf("package require tooltip"))
        // fmt.Printf("package require tooltip ok ? %v\n", ok)
    }
    // s1 := fmt.Sprintf("tooltip::tooltip %v {%v}", w.id, tip)
    // fmt.Printf("s1 is %v\n",s1)
    return eval(fmt.Sprintf("tooltip::tooltip %v {%v}", w.id, tip))
}


