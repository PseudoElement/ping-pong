package injector

import (
	"github.com/pseudoelement/ping-pong/sdk/interfaces"
)

type Injector struct {
	soundLoader  interfaces.ISoundLoader
	imageLoader  interfaces.IImageLoader
	fontLoader   interfaces.IFontLoader
	eventManager interfaces.IEventManager
	renderer     interfaces.IRenderer
	gameSettings interfaces.IGameSettings
}

var _ interfaces.IInjector = (*Injector)(nil)

func NewInjector() *Injector {
	return &Injector{}
}

func (i *Injector) SetSoundLoader(soundLoader interfaces.ISoundLoader) {
	i.soundLoader = soundLoader
}

func (i *Injector) SetImageLoader(imageLoader interfaces.IImageLoader) {
	i.imageLoader = imageLoader
}

func (i *Injector) SetFontLoader(fontLoader interfaces.IFontLoader) {
	i.fontLoader = fontLoader
}

func (i *Injector) SetEventManager(eventManager interfaces.IEventManager) {
	i.eventManager = eventManager
}

func (i *Injector) SetGameSettings(gameSettings interfaces.IGameSettings) {
	i.gameSettings = gameSettings
}

func (i *Injector) SetRenderer(renderer interfaces.IRenderer) {
	i.renderer = renderer
}

func (i *Injector) SoundLoader() interfaces.ISoundLoader {
	return i.soundLoader
}

func (i *Injector) ImageLoader() interfaces.IImageLoader {
	return i.imageLoader
}

func (i *Injector) FontLoader() interfaces.IFontLoader {
	return i.fontLoader
}

func (i *Injector) EventManager() interfaces.IEventManager {
	return i.eventManager
}

func (i *Injector) GameSettings() interfaces.IGameSettings {
	return i.gameSettings
}

func (i *Injector) Renderer() interfaces.IRenderer {
	return i.renderer
}
