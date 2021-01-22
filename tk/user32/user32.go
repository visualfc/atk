package user32

import (
	"syscall"
	"unsafe"
)

func GetWindowScreen() (cxLogical, cyLogical, cxPhysical, cyPhysical int) {
	hWnd := GetDesktopWindow()
	hMonitor := MonitorFromWindow(hWnd, MONITOR_DEFAULTTONEAREST)
	var miex MONITORINFOEX
	miex.cbSize = uint32(unsafe.Sizeof(miex))
	GetMonitorInfo(hMonitor, &miex)
	cxLogical = int(miex.rcMonitor.Right - miex.rcMonitor.Left)
	cyLogical = int(miex.rcMonitor.Bottom - miex.rcMonitor.Top)
	var dm DEVMODE
	dm.dmSize = uint16(unsafe.Sizeof(dm))
	dm.dmDriverExtra = 0
	EnumDisplaySettings(&miex.szDevice[0], -1, &dm)
	cxPhysical = int(dm.dmPelsWidth)
	cyPhysical = int(dm.dmPelsHeight)
	return
}

type HWND uintptr
type HMONITOR uintptr

const (
	MONITOR_DEFAULTTONULL    = 0x00000000
	MONITOR_DEFAULTTOPRIMARY = 0x00000001
	MONITOR_DEFAULTTONEAREST = 0x00000002
)

func GetDesktopWindow() HWND {
	ret, _, _ := syscall.Syscall(getDesktopWindow.Addr(), 0,
		0,
		0,
		0)

	return HWND(ret)
}

func MonitorFromWindow(hwnd HWND, dwFlags uint32) HMONITOR {
	ret, _, _ := syscall.Syscall(monitorFromWindow.Addr(), 2,
		uintptr(hwnd),
		uintptr(dwFlags),
		0)

	return HMONITOR(ret)
}

func GetMonitorInfo(hMonitor HMONITOR, lpmi *MONITORINFOEX) bool {
	ret, _, _ := syscall.Syscall(getMonitorInfo.Addr(), 2,
		uintptr(hMonitor),
		uintptr(unsafe.Pointer(lpmi)),
		0)

	return ret != 0
}

func EnumDisplaySettings(driverName *uint16, imode int32, lpmode *DEVMODE) bool {
	ret, _, _ := syscall.Syscall(enumDisplaySettings.Addr(), 3,
		uintptr(unsafe.Pointer(driverName)),
		uintptr(imode),
		uintptr(unsafe.Pointer(lpmode)))

	return ret != 0
}

var (
	user32              = syscall.NewLazyDLL("user32.dll")
	getDesktopWindow    = user32.NewProc("GetDesktopWindow")
	monitorFromWindow   = user32.NewProc("MonitorFromWindow")
	getMonitorInfo      = user32.NewProc("GetMonitorInfoW")
	enumDisplaySettings = user32.NewProc("EnumDisplaySettingsW")
)

const (
	CCHDEVICENAME = 32
	CCHFORMNAME   = 32
)

type POINT struct {
	X, Y int32
}

type RECT struct {
	Left, Top, Right, Bottom int32
}

type MONITORINFO struct {
	cbSize    uint32
	rcMonitor RECT
	rcWork    RECT
	dwFlags   uint32
}

type MONITORINFOEX struct {
	MONITORINFO
	szDevice [CCHDEVICENAME]uint16
}

type DEVMODE struct {
	dmDeviceName       [CCHDEVICENAME]uint16
	dmSpecVersion      uint16
	dmDriverVersion    uint16
	dmSize             uint16
	dmDriverExtra      uint16
	dmFields           uint32
	dmOrientation      int16
	dmPaperSize        int16
	dmPaperLength      int16
	dmPaperWidth       int16
	dmScale            int16
	dmCopies           int16
	dmDefaultSource    int16
	dmPrintQuality     int16
	dmColor            int16
	dmDuplex           int16
	dmYResolution      int16
	dmTTOption         int16
	dmCollate          int16
	dmFormName         [CCHFORMNAME]uint16
	dmLogPixels        uint16
	dmBitsPerPel       uint32
	dmPelsWidth        uint32
	dmPelsHeight       uint32
	dmDisplayFlags     uint32
	dmDisplayFrequency uint32
	dmICMMethod        uint32
	dmICMIntent        uint32
	dmMediaType        uint32
	dmDitherType       uint32
	dmReserved1        uint32
	dmReserved2        uint32
	dmPanningWidth     uint32
	dmPanningHeight    uint32
}
