package constants

import rl "github.com/gen2brain/raylib-go/raylib"

type MouseClickEvent struct {
	PosX int32
	PosY int32
	Kind rl.MouseButton
}

type MouseMoveEvent struct {
	PosX int32
	PosY int32
}

type KeyPressedEvent struct {
	/* int value from raylib-go lib */
	KeyCode int32
}

type FpsToggleEvent struct {
	ShowFps bool
}

type NextFrameEvent struct{}

type BackgroundChangeEvent struct {
	Name string
	Href string
}
