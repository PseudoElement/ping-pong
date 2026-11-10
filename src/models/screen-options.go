package models

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/pseudoelement/ping-pong/sdk/models"
	"github.com/pseudoelement/ping-pong/sdk/styles"
	"github.com/pseudoelement/ping-pong/sdk/utils"
	"github.com/pseudoelement/ping-pong/src/constants"
	"github.com/pseudoelement/ping-pong/src/injector"
	"github.com/pseudoelement/ping-pong/src/interfaces"
	"github.com/pseudoelement/ping-pong/src/loaders"
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
	mainBgCode := so.injector.GameSettings().MainBackground
	bgImg, _ := so.injector.ImageLoader().GetAppBackground(mainBgCode)

	toggleFpsBtn := NewTextObject(
		TextObjectStructParams{
			InitialValue:  "Toggle FPS counter",
			FontSize:      24,
			Styles:        &styles.ObjectStyles{Color: rl.Black},
			InitialCoords: models.Coords{X: windowWidth/2 - 50, Y: windowHeight/2 - 50},
		},
		so.injector,
	)
	changeBackgroundBtn := NewTextObject(
		TextObjectStructParams{
			InitialValue:  fmt.Sprintf("Change theme: <%s>\n", bgImg.Name),
			FontSize:      24,
			Styles:        &styles.ObjectStyles{Color: rl.Black},
			InitialCoords: models.Coords{X: windowWidth/2 - 50, Y: windowHeight/2 - 50 + 40},
			Id:            "change-bg-btn-id",
		},
		so.injector,
	)
	backBtn := NewTextObject(
		TextObjectStructParams{
			InitialValue:  "Back",
			FontSize:      24,
			Styles:        &styles.ObjectStyles{Color: rl.Black},
			InitialCoords: models.Coords{X: windowWidth/2 - 50, Y: windowHeight/2 - 50 + 80},
		},
		so.injector,
	)

	mouseClickEventUnsubFn := so.injector.EventManager().
		MouseClickEventEmitter().
		Subscribe(func(event constants.MouseClickEvent) {
			_toggleFpsBtnEventHandler(event, toggleFpsBtn, so.injector)
			_changeBackgroundBtnEventHandler(event, changeBackgroundBtn, so.injector, so)
			_backBtnEventHandler(event, backBtn, so.injector)
		})
	mouseMoveEventUnsubFn := so.injector.EventManager().
		MouseMoveEventEmitter().
		Subscribe(func(event constants.MouseMoveEvent) {
			_mouseMoveEventHandler(event, so.textObjectsPool, so.injector)
		})

	so.unsubs = append(so.unsubs, mouseClickEventUnsubFn, mouseMoveEventUnsubFn)
	so.textObjectsPool = append(so.textObjectsPool, toggleFpsBtn, changeBackgroundBtn, backBtn)
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

func _changeBackgroundBtnEventHandler(
	mouseEvent constants.MouseClickEvent,
	textObj interfaces.ITextObject,
	injector *injector.Injector,
	screen *ScreenOptions,
) {
	mouseCoords := models.Coords{X: mouseEvent.PosX, Y: mouseEvent.PosY}
	if utils.IsTextObjectHovered(mouseCoords, textObj.Coords(), textObj.Width(), textObj.Height()) {
		if mouseEvent.Kind == rl.MouseButtonLeft {
			currentBgName := injector.GameSettings().MainBackground
			appBackgrounds := injector.ImageLoader().GetAppBackgrounds()

			var nextBg loaders.AppImage
			for idx, bgImg := range appBackgrounds {
				if currentBgName == bgImg.Name {
					if idx == len(appBackgrounds)-1 {
						nextBg = appBackgrounds[0]
					} else {
						nextBg = appBackgrounds[idx+1]
					}
					break
				}
			}

			injector.GameSettings().SetMainBackground(nextBg.Name)
			injector.EventManager().
				BackgroundChangeEventEmitter().
				Emit(constants.BackgroundChangeEvent{
					Name: string(nextBg.Name),
					Href: nextBg.Href,
				})

			for _, obj := range screen.textObjectsPool {
				if obj.Id() == "change-bg-btn-id" {
					obj.SetValue(fmt.Sprintf("Change theme: <%s>\n", nextBg.Name))
				}
			}
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
