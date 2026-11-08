package injector

import "github.com/pseudoelement/ping-pong/src/settings"

type Injector struct {
	soundLoader  ISoundLoader
	eventManager IEventManager
	gameSettings *settings.GameSettings
}

func NewInjector() *Injector {
	return &Injector{}
}

func (i *Injector) SetSoundLoader(soundLoader ISoundLoader) {
	i.soundLoader = soundLoader
}

func (i *Injector) SetEventManager(eventManager IEventManager) {
	i.eventManager = eventManager
}

func (i *Injector) SetGameSettings(gameSettings *settings.GameSettings) {
	i.gameSettings = gameSettings
}

func (i *Injector) SoundLoader() ISoundLoader {
	return i.soundLoader
}

func (i *Injector) EventManager() IEventManager {
	return i.eventManager
}

func (i *Injector) Settings() *settings.GameSettings {
	return i.gameSettings
}
