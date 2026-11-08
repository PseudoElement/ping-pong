package injector

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/pseudoelement/ping-pong/src/sounds"
)

type ISoundLoader interface {
	GetSound(soundName sounds.AppSound) (rl.Sound, bool)
}

type IEventManager interface {
	UseEventHandlers()
}
