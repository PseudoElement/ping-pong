package models

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/pseudoelement/ping-pong/sdk/models"
	"github.com/pseudoelement/ping-pong/sdk/styles"
	"github.com/pseudoelement/ping-pong/sdk/utils"
	"github.com/pseudoelement/ping-pong/src/constants"
	"github.com/pseudoelement/ping-pong/src/injector"
	"github.com/pseudoelement/ping-pong/src/interfaces"
)

type ScreenOptions struct {
	injector        *injector.Injector
	textObjectsPool []interfaces.ITextObject
	rectObjectsPool []interfaces.IBaseObject
	unsubs          []models.UnsubscribeFn
}

var _ interfaces.IScreen = (*ScreenOptions)(nil)

func NewScreenOptions(injector *injector.Injector) *ScreenOptions {
	screen := &ScreenOptions{
		injector:        injector,
		textObjectsPool: make([]interfaces.ITextObject, 0),
		rectObjectsPool: make([]interfaces.IBaseObject, 0),
	}
	screen._init()
	return screen
}

func (so *ScreenOptions) TextObjects() []interfaces.ITextObject {
	return so.textObjectsPool
}

func (so *ScreenOptions) RectObjects() []interfaces.IBaseObject {
	return so.rectObjectsPool
}

func (so *ScreenOptions) OnDestroy() {
	for _, unsubFn := range so.unsubs {
		unsubFn()
	}
}

func (so *ScreenOptions) _init() {
	windowWidth := so.injector.GameSettings().WindowWidth
	windowHeight := so.injector.GameSettings().WindowHeight

	toggleFpsBtn := NewTextObject(
		TextObjectStructParams{
			InitialValue:  "Toggle FPS counter",
			FontSize:      24,
			Styles:        &styles.ObjectStyles{Color: rl.Black},
			InitialCoords: models.Coords{X: windowWidth/2 - 50, Y: windowHeight/2 - 50},
		},
		so.injector,
	)
	backBtn := NewTextObject(
		TextObjectStructParams{
			InitialValue:  "Back",
			FontSize:      24,
			Styles:        &styles.ObjectStyles{Color: rl.Black},
			InitialCoords: models.Coords{X: windowWidth/2 - 50, Y: windowHeight/2 - 50 + 40},
		},
		so.injector,
	)

	mouseClickEventUnsubFn := so.injector.EventManager().
		MouseClickEventEmitter().
		Subscribe(func(event constants.MouseClickEvent) {
			_toggleFpsBtnEventHandler(event, toggleFpsBtn, so.injector)
			_backBtnEventHandler(event, backBtn, so.injector)
		})
	mouseMoveEventUnsubFn := so.injector.EventManager().
		MouseMoveEventEmitter().
		Subscribe(func(event constants.MouseMoveEvent) {
			_mouseMoveEventHandler(event, so.textObjectsPool, so.injector)
		})

	so.unsubs = append(so.unsubs, mouseClickEventUnsubFn, mouseMoveEventUnsubFn)
	so.textObjectsPool = append(so.textObjectsPool, toggleFpsBtn, backBtn)
}

func _toggleFpsBtnEventHandler(
	mouseEvent constants.MouseClickEvent,
	textObj interfaces.ITextObject,
	injector *injector.Injector,
) {
	mouseCoords := models.Coords{X: mouseEvent.PosX, Y: mouseEvent.PosY}
	if utils.IsTextObjectHovered(mouseCoords, textObj.Coords(), textObj.Width(), textObj.Height()) {
		if mouseEvent.Kind == rl.MouseButtonLeft {
			injector.GameSettings().SetShowFps(!injector.GameSettings().ShowFps)
			injector.EventManager().FpsToggleEventEmitter().Emit(constants.FpsToggleEvent{
				ShowFps: injector.GameSettings().ShowFps,
			})
		}
	}
}

func _backBtnEventHandler(
	mouseEvent constants.MouseClickEvent,
	textObj interfaces.ITextObject,
	injector *injector.Injector,
) {
	mouseCoords := models.Coords{X: mouseEvent.PosX, Y: mouseEvent.PosY}
	if utils.IsTextObjectHovered(mouseCoords, textObj.Coords(), textObj.Width(), textObj.Height()) {
		if mouseEvent.Kind == rl.MouseButtonLeft {
			injector.Renderer().CurrentScreen().OnDestroy()
			injector.Renderer().SetScreen(NewScreenMenu(injector))
		}
	}
}
