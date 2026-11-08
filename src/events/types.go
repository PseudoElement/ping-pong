package events

import rl "github.com/gen2brain/raylib-go/raylib"

type UserEvent int

const (
	MOUSE_EVENT UserEvent = iota
	KEYBOARD_EVENT
	GAMEPAD_EVENT
)

type KeyboardKey int

type IEventStrategy interface {
	HandleKeyboardEvent(keyPressed KeyboardKey)
	SupportsKeyboardEvent(keyPressed KeyboardKey) bool

	HandleMouseEvent(mousePressed rl.MouseButton)
	SupportsMouseEvent(mousePressed rl.MouseButton) bool
}
