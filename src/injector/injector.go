package injector

import (
	"github.com/pseudoelement/ping-pong/src/interfaces"
	"github.com/pseudoelement/ping-pong/src/settings"
)

type Injector struct {
	soundLoader  interfaces.ISoundLoader
	eventManager interfaces.IEventManager
	renderer     interfaces.IRenderer
	gameSettings *settings.GameSettings
}

func NewInjector() *Injector {
	return &Injector{}
}

func (i *Injector) SetSoundLoader(soundLoader interfaces.ISoundLoader) {
	i.soundLoader = soundLoader
}

func (i *Injector) SetEventManager(eventManager interfaces.IEventManager) {
	i.eventManager = eventManager
}

func (i *Injector) SetGameSettings(gameSettings *settings.GameSettings) {
	i.gameSettings = gameSettings
}

func (i *Injector) SetRenderer(renderer interfaces.IRenderer) {
	i.renderer = renderer
}

func (i *Injector) SoundLoader() interfaces.ISoundLoader {
	return i.soundLoader
}

func (i *Injector) EventManager() interfaces.IEventManager {
	return i.eventManager
}

func (i *Injector) Settings() *settings.GameSettings {
	return i.gameSettings
}
