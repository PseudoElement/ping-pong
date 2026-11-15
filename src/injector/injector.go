package injector

import (
	appinterfaces "github.com/pseudoelement/ping-pong/src/interfaces"
)

type Injector struct {
	soundLoader  appinterfaces.ISoundLoader
	imageLoader  appinterfaces.IImageLoader
	fontLoader   appinterfaces.IFontLoader
	eventManager appinterfaces.IEventManager
	renderer     appinterfaces.IRenderer
	gameSettings appinterfaces.IGameSettings
}

var _ appinterfaces.IInjector = (*Injector)(nil)

func NewInjector() *Injector {
	return &Injector{}
}

func (i *Injector) SetSoundLoader(soundLoader appinterfaces.ISoundLoader) {
	i.soundLoader = soundLoader
}

func (i *Injector) SetImageLoader(imageLoader appinterfaces.IImageLoader) {
	i.imageLoader = imageLoader
}

func (i *Injector) SetFontLoader(fontLoader appinterfaces.IFontLoader) {
	i.fontLoader = fontLoader
}

func (i *Injector) SetEventManager(eventManager appinterfaces.IEventManager) {
	i.eventManager = eventManager
}

func (i *Injector) SetGameSettings(gameSettings appinterfaces.IGameSettings) {
	i.gameSettings = gameSettings
}

func (i *Injector) SetRenderer(renderer appinterfaces.IRenderer) {
	i.renderer = renderer
}

func (i *Injector) SoundLoader() appinterfaces.ISoundLoader {
	return i.soundLoader
}

func (i *Injector) ImageLoader() appinterfaces.IImageLoader {
	return i.imageLoader
}

func (i *Injector) FontLoader() appinterfaces.IFontLoader {
	return i.fontLoader
}

func (i *Injector) EventManager() appinterfaces.IEventManager {
	return i.eventManager
}

func (i *Injector) GameSettings() appinterfaces.IGameSettings {
	return i.gameSettings
}

func (i *Injector) Renderer() appinterfaces.IRenderer {
	return i.renderer
}
