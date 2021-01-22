package user32

import (
	"testing"
)

func TestGetWindowScreen(t *testing.T) {
	cxLogical, cyLogical, cxPhysical, cyPhysical := GetWindowScreen()
	t.Log(cxLogical, cyLogical, cxPhysical, cyPhysical)
}
