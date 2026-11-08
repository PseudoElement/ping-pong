package interfaces

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/pseudoelement/ping-pong/sdk/models"
	"github.com/pseudoelement/ping-pong/src/constants"
	"github.com/pseudoelement/ping-pong/src/sounds"
)

type ISoundLoader interface {
	GetSound(soundName sounds.AppSound) (rl.Sound, bool)
}

type IEventManager interface {
	MouseEventEmitter() *models.EventEmitter[constants.MouseClickEvent]
	KeyboardEventEmitter() *models.EventEmitter[constants.KeyPressedEvent]
	UseEventHandlers()
}

type IRenderer interface {
	Render()
	CurrentScreen() IScreen
}
