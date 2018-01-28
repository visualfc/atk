// Copyright 2018 visualfc. All rights reserved.

package event

// EVENT TYPES
const (
	//These two events are sent to every sub-window of a toplevel when they change state. In addition to the focus Window, the Macintosh platform and Windows platforms have a notion of an active window (which often has but is not required to have the focus). On the Macintosh, widgets in the active window have a different appearance than widgets in deactive windows. The Activate event is sent to all the sub-windows in a toplevel when it changes from being deactive to active. Likewise, the Deactive event is sent when the window's state changes from active to deactive. There are no useful percent substitutions you would make when binding to these events.
	Activate   = "<Activate>"
	Deactivate = "<Deactivate>"

	//Many contemporary mice support a mouse wheel, which is used for scrolling documents without using the scrollbars. By rolling the wheel, the system will generate MouseWheel events that the application can use to scroll. Like Key events the event is always routed to the window that currently has focus. When the event is received you can use the %D substitution to get the delta field for the event, which is a integer value describing how the mouse wheel has moved. The smallest value for which the system will report is defined by the OS. The sign of the value determines which direction your widget should scroll. Positive values should scroll up and negative values should scroll down.
	MouseWheel = "<MouseWheel>"

	//The KeyPress and KeyRelease events are generated whenever a key is pressed or released. KeyPress and KeyRelease events are sent to the window which currently has the keyboard focus.
	KeyPress   = "<KeyPress>"
	KeyRelease = "<KeyRelease>"

	//The ButtonPress and ButtonRelease events are generated when the user presses or releases a mouse button. Motion events are generated whenever the pointer is moved. ButtonPress, ButtonRelease, and Motion events are normally sent to the window containing the pointer.
	//
	//When a mouse button is pressed, the window containing the pointer automatically obtains a temporary pointer grab. Subsequent ButtonPress, ButtonRelease, and Motion events will be sent to that window, regardless of which window contains the pointer, until all buttons have been released.
	ButtonPress   = "<ButtonPress>"
	ButtonRelease = "<ButtonRelease>"
	Motion        = "<Motion>"

	//A Configure event is sent to a window whenever its size, position, or border width changes, and sometimes when it has changed position in the stacking order.
	Configure = "<Configure>"

	//The Map and Unmap events are generated whenever the mapping state of a window changes.
	//Windows are created in the unmapped state. Top-level windows become mapped when they transition to the normal state, and are unmapped in the withdrawn and iconic states. Other windows become mapped when they are placed under control of a geometry manager (for example pack or grid).
	//
	//A window is viewable only if it and all of its ancestors are mapped. Note that geometry managers typically do not map their children until they have been mapped themselves, and unmap all children when they become unmapped; hence in Tk Map and Unmap events indicate whether or not a window is viewable.
	Map   = "<Map>"
	Unmap = "<Unmap>"

	//A window is said to be obscured when another window above it in the stacking order fully or partially overlaps it. Visibility events are generated whenever a window's obscurity state changes; the state field (%s) specifies the new state.
	Visibility = "<Visibility>"

	//An Expose event is generated whenever all or part of a window should be redrawn (for example, when a window is first mapped or if it becomes unobscured). It is normally not necessary for client applications to handle Expose events, since Tk handles them internally.
	Expose = "<Expose>"

	Destroy = "<Destroy>"
	//A Destroy event is delivered to a window when it is destroyed.
	//
	//When the Destroy event is delivered to a widget, it is in a “half-dead” state: the widget still exists, but most operations on it will fail.

	//The FocusIn and FocusOut events are generated whenever the keyboard focus changes. A FocusOut event is sent to the old focus window, and a FocusIn event is sent to the new one.
	//In addition, if the old and new focus windows do not share a common parent, “virtual crossing” focus events are sent to the intermediate windows in the hierarchy. Thus a FocusIn event indicates that the target window or one of its descendants has acquired the focus, and a FocusOut event indicates that the focus has been changed to a window outside the target window's hierarchy.
	//
	//The keyboard focus may be changed explicitly by a call to focus, or implicitly by the window manager.
	FocusIn  = "<FocusIn>"
	FocusOut = "<FocusOut>"

	//An Enter event is sent to a window when the pointer enters that window, and a Leave event is sent when the pointer leaves it.
	//If there is a pointer grab in effect, Enter and Leave events are only delivered to the window owning the grab.
	//
	//In addition, when the pointer moves between two windows, Enter and Leave “virtual crossing” events are sent to intermediate windows in the hierarchy in the same manner as for FocusIn and FocusOut events.
	Enter = "<Enter>"
	Leave = "<Leave>"

	//A Property event is sent to a window whenever an X property belonging to that window is changed or deleted. Property events are not normally delivered to Tk applications as they are handled by the Tk core.
	Property = "<Property>"

	//A Colormap event is generated whenever the colormap associated with a window has been changed, installed, or uninstalled.
	//
	//Widgets may be assigned a private colormap by specifying a -colormap option; the window manager is responsible for installing and uninstalling colormaps as necessary.
	//
	//Note that Tk provides no useful details for this event type.
	Colormap = "<Colormap>"

	//These events are not normally delivered to Tk applications. They are included for completeness, to make it possible to write X11 window managers in Tk. (These events are only delivered when a client has selected SubstructureRedirectMask on a window; the Tk core does not use this mask.)
	MapRequest       = "<MapRequest>"
	CirculateRequest = "<CirculateRequest>"
	ResizeRequest    = "<ResizeRequest>"
	ConfigureRequest = "<ConfigureRequest>"
	Create           = "<Create>"

	//The events Gravity and Reparent are not normally delivered to Tk applications. They are included for completeness.
	//
	//A Circulate event indicates that the window has moved to the top or to the bottom of the stacking order as a result of an XCirculateSubwindows protocol request. Note that the stacking order may be changed for other reasons which do not generate a Circulate event, and that Tk does not use XCirculateSubwindows() internally. This event type is included only for completeness; there is no reliable way to track changes to a window's position in the stacking order.)
	Gravity   = "<Gravity>"
	Reparent  = "<Reparent>"
	Circulate = "<Circulate>"
)

var (
	EventList = []string{
		"<Activate>", "<Destroy>", "<Map>", "<ButtonPress>", "<Button>",
		"<Enter>", "<MapRequest>", "<ButtonRelease>", "<Expose>", "<Motion>",
		"<Circulate>", "<FocusIn>", "<MouseWheel>", "<CirculateRequest>", "<FocusOut>",
		"<Property>", "<Colormap>", "<Gravity>", "<Reparent>", "<Configure>",
		"<KeyPress>", "<Key>", "<ResizeRequest>", "<ConfigureRequest>", "<KeyRelease>",
		"<Unmap>", "<Create>", "<Leave>", "<Visibility>", "<Deactivate>",
	}
)

//PREDEFINED VIRTUAL EVENTS
const (
	AltUnderlined   = "<<AltUnderlined>>"
	Invoke          = "<<Invoke>>"
	ListboxSelect   = "<<ListboxSelect>>"
	MenuSelect      = "<<MenuSelect>>"
	Modified        = "<<Modified>>"
	Selection       = "<<Selection>>"
	ThemeChanged    = "<<ThemeChanged>>"
	TraverseIn      = "<<TraverseIn>>"
	TraverseOut     = "<<TraverseOut>>"
	UndoStack       = "<<UndoStack>>"
	WidgetViewSync  = "<<WidgetViewSync>>"
	Clear           = "<<Clear>>"
	Copy            = "<<Copy>>"
	Cut             = "<<Cut>>"
	LineEnd         = "<<LineEnd>>"
	LineStart       = "<<LineStart>>"
	NextChar        = "<<NextChar>>"
	NextLine        = "<<NextLine>>"
	NextPara        = "<<NextPara>>"
	NextWord        = "<<NextWord>>"
	Paste           = "<<Paste>>"
	PasteSelection  = "<<PasteSelection>>"
	PrevChar        = "<<PrevChar>>"
	PrevLine        = "<<PrevLine>>"
	PrevPara        = "<<PrevPara>>"
	PrevWindow      = "<<PrevWindow>>"
	PrevWord        = "<<PrevWord>>"
	Redo            = "<<Redo>>"
	SelectAll       = "<<SelectAll>>"
	SelectLineEnd   = "<<SelectLineEnd>>"
	SelectLineStart = "<<SelectLineStart>>"
	SelectNextChar  = "<<SelectNextChar>>"
	SelectNextLine  = "<<SelectNextLine>>"
	SelectNextPara  = "<<SelectNextPara>>"
	SelectNextWord  = "<<SelectNextWord>>"
	SelectNone      = "<<SelectNone>>"
	SelectPrevChar  = "<<SelectPrevChar>>"
	SelectPrevLine  = "<<SelectPrevLine>>"
	SelectPrevPara  = "<<SelectPrevPara>>"
	SelectPrevWord  = "<<SelectPrevWord>>"
	ToggleSelection = "<<ToggleSelection>>"
	Undo            = "<<Undo>>"
)
