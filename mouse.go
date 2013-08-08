package we

import (
	"image"
)

// Button corresponds to a mouse button.
type Button int

// Mouse buttons.
const (
	Button1 Button = iota
	Button2
	Button3
	Button4
	Button5
	Button6
	Button7
	Button8

	ButtonLeft   = Button1
	ButtonRight  = Button2
	ButtonMiddle = Button3
)

// MousePress is triggered when a mouse button is pressed.
type MousePress struct {
	// Coordinates of the mouse cursor.
	image.Point
	// Mouse button that was pressed.
	Button Button
	// Bitfield of modifier keys.
	Mod Mod
}

// MouseRelease is triggered when a mouse button is released.
type MouseRelease struct {
	// Coordinates of the mouse cursor.
	image.Point
	// Mouse button that was released.
	Button Button
	// Bitfield of modifier keys.
	Mod Mod
}

// MouseMove is triggered when the mouse is moved from one location to another.
type MouseMove struct {
	// Coordinates of the mouse cursor.
	image.Point
	// Coordinates of the mouse cursor at the beginning of the move event.
	From image.Point
}

// MouseDrag is triggered when the mouse is moved from one location to another
// while a mouse button is held down.
type MouseDrag struct {
	// Coordinates of the mouse cursor.
	image.Point
	// Coordinates of the mouse cursor at the beginning of the drag event.
	From image.Point
	// Mouse button that was held down.
	Button Button
	// Bitfield of modifier keys.
	Mod Mod
}

// MouseEnter is triggered when the mouse enters or leaves the window. On mouse
// enter the value is true, otherwise it is false.
type MouseEnter bool

// ScrollUp is triggered when the mouse wheel is scrolled upwards.
type ScrollUp struct {
	// Scroll offset.
	Off int
	// Bitfield of modifier keys.
	Mod Mod
}

// ScrollDown is triggered when the mouse wheel is scrolled downwards.
type ScrollDown struct {
	// Scroll offset.
	Off int
	// Bitfield of modifier keys.
	Mod Mod
}
