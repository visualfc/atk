// Copyright 2018 visualfc. All rights reserved.

package event

const (
	TypeKeyPress         = 2
	TypeKeyRelease       = 3
	TypeButtonPress      = 4
	TypeButtonRelease    = 5
	TypeMotionNotify     = 6
	TypeEnterNotify      = 7
	TypeLeaveNotify      = 8
	TypeFocusIn          = 9
	TypeFocusOut         = 10
	TypeKeymapNotify     = 11
	TypeExpose           = 12
	TypeGraphicsExpose   = 13
	TypeNoExpose         = 14
	TypeVisibilityNotify = 15
	TypeCreateNotify     = 16
	TypeDestroyNotify    = 17
	TypeUnmapNotify      = 18
	TypeMapNotify        = 19
	TypeMapRequest       = 20
	TypeReparentNotify   = 21
	TypeConfigureNotify  = 22
	TypeConfigureRequest = 23
	TypeGravityNotify    = 24
	TypeResizeRequest    = 25
	TypeCirculateNotify  = 26
	TypeCirculateRequest = 27
	TypePropertyNotify   = 28
	TypeSelectionClear   = 29
	TypeSelectionRequest = 30
	TypeSelectionNotify  = 31
	TypeColormapNotify   = 32
	TypeClientMessage    = 33
	TypeMappingNotify    = 34
	LASTEvent            = 35
)

const (
	TypeVirtualEvent     = (TypeMappingNotify + 1)
	TypeActivateNotify   = (TypeMappingNotify + 2)
	TypeDeactivateNotify = (TypeMappingNotify + 3)
	TypeMouseWheelEvent  = (TypeMappingNotify + 4)
	TK_LASTEVENT         = (TypeMappingNotify + 5)
)
