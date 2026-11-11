package interfaces

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/pseudoelement/ping-pong/sdk/models"
	"github.com/pseudoelement/ping-pong/src/constants"
	"github.com/pseudoelement/ping-pong/src/loaders"
	"github.com/pseudoelement/ping-pong/src/settings"
)

type IInjector interface {
	SoundLoader() ISoundLoader
	ImageLoader() IImageLoader
	EventManager() IEventManager
	GameSettings() *settings.GameSettings
	Renderer() IRenderer

	SetSoundLoader(soundLoader ISoundLoader)
	SetEventManager(eventManager IEventManager)
	SetGameSettings(gameSettings *settings.GameSettings)
	SetRenderer(renderer IRenderer)
}

type ISoundLoader interface {
	GetSound(soundName loaders.AppSoundName) rl.Sound
}

type IImageLoader interface {
	GetTexture(imgName loaders.AppImageName) rl.Texture2D
	GetAppBackgrounds() []loaders.AppImage
	GetAppBackground(bgName loaders.AppImageName) (loaders.AppImage, bool)
}

type IEventManager interface {
	MouseClickEventEmitter() *models.EventEmitter[constants.MouseClickEvent]
	MouseMoveEventEmitter() *models.EventEmitter[constants.MouseMoveEvent]
	KeyboardEventEmitter() *models.EventEmitter[constants.KeyPressedEvent]

	FpsToggleEventEmitter() *models.EventEmitter[constants.FpsToggleEvent]
	BackgroundChangeEventEmitter() *models.EventEmitter[constants.BackgroundChangeEvent]
	NextFrameEventEmitter() *models.EventEmitter[constants.NextFrameEvent]

	UseUserInputHandlers()
}

type IRenderer interface {
	Render()
	CurrentScreen() IScreen
	SetScreen(screen IScreen)
}
