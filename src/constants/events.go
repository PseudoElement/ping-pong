package constants

import rl "github.com/gen2brain/raylib-go/raylib"

type MouseClickEvent struct {
	PosX int32
	PosY int32
	Kind rl.MouseButton
}

type KeyPressedEvent struct {
	/* int value from raylib-go lib */
	KeyCode int32
}
