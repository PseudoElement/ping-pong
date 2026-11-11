package interfaces

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/pseudoelement/ping-pong/sdk/common"
	"github.com/pseudoelement/ping-pong/sdk/constants"
)

type IInjector interface {
	SoundLoader() ISoundLoader
	ImageLoader() IImageLoader
	EventManager() IEventManager
	GameSettings() IGameSettings
	Renderer() IRenderer

	SetSoundLoader(soundLoader ISoundLoader)
	SetEventManager(eventManager IEventManager)
	SetGameSettings(gameSettings IGameSettings)
	SetRenderer(renderer IRenderer)
}

type ISoundLoader interface {
	GetSound(soundName constants.AppSoundName) rl.Sound
}

type IImageLoader interface {
	GetTexture(imgName constants.AppImageName) rl.Texture2D
	GetAppBackgrounds() []constants.AppImage
	GetAppBackground(bgName constants.AppImageName) (constants.AppImage, bool)
}

type IEventManager interface {
	MouseClickEventEmitter() *common.EventEmitter[constants.MouseClickEvent]
	MouseMoveEventEmitter() *common.EventEmitter[constants.MouseMoveEvent]
	KeyboardEventEmitter() *common.EventEmitter[constants.KeyPressedEvent]

	FpsToggleEventEmitter() *common.EventEmitter[constants.FpsToggleEvent]
	BackgroundChangeEventEmitter() *common.EventEmitter[constants.BackgroundChangeEvent]
	NextFrameEventEmitter() *common.EventEmitter[constants.NextFrameEvent]

	UseUserInputHandlers()
}

type IRenderer interface {
	Render()
	CurrentScreen() IScreen
	SetScreen(screen IScreen)
}

type IGameSettings interface {
	FontSize() int32
	SetFontSize(fontSize int32)

	WindowWidth() int32
	SetWindowWidth(windowWidth int32)

	WindowHeight() int32
	SetWindowHeight(windowHeight int32)

	Fps() int32
	SetFps(fps int32)

	IsShownFps() bool
	SetShowFps(showFps bool)

	MainBackground() constants.AppImageName
	SetMainBackground(bg constants.AppImageName)
}
