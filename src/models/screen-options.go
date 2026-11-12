package models

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/pseudoelement/ping-pong/sdk/common"
	sdkconstants "github.com/pseudoelement/ping-pong/sdk/constants"
	"github.com/pseudoelement/ping-pong/sdk/interfaces"
	"github.com/pseudoelement/ping-pong/sdk/styles"
	"github.com/pseudoelement/ping-pong/sdk/utils"
	"github.com/pseudoelement/ping-pong/src/constants"
	"github.com/pseudoelement/ping-pong/src/injector"
)

type ScreenOptions struct {
	*ScreenBase
}

var _ interfaces.IScreen = (*ScreenOptions)(nil)

func NewScreenOptions(injector *injector.Injector) *ScreenOptions {
	screen := &ScreenOptions{
		ScreenBase: NewScreenBase(injector),
	}
	screen._init()
	return screen
}

func (so *ScreenOptions) OnDestroy() {
	so.onDestroyBase()
}

func (so *ScreenOptions) _init() {
	windowWidth := so.injector.GameSettings().WindowWidth()
	windowHeight := so.injector.GameSettings().WindowHeight()
	mainBgCode := so.injector.GameSettings().MainBackground()
	bgImg, _ := so.injector.ImageLoader().GetAppBackground(mainBgCode)

	wrapper := NewRectangleObject(
		RectangleObjectStructParams{
			InitialCoords: common.Coords{X: windowWidth / 2, Y: windowHeight / 2},
			Width:         400,
			Height:        400,
			Styles: &common.ObjectStyles{
				Color:        constants.ColorTransparentBlue,
				ZIndex:       1,
				BorderRadius: 0.3,
			},
		},
		so.injector,
	)
	toggleFpsBtn := NewTextObject(
		TextObjectStructParams{
			InitialValue:  "Toggle FPS counter",
			FontSize:      24,
			Styles:        &common.ObjectStyles{Color: rl.Black, ZIndex: 2},
			InitialCoords: common.Coords{X: windowWidth / 2, Y: windowHeight / 2},
		},
		so.injector,
	)
	changeBackgroundBtn := NewTextObject(
		TextObjectStructParams{
			InitialValue:  fmt.Sprintf("Change theme: <%s>\n", bgImg.Name),
			FontSize:      24,
			Styles:        &common.ObjectStyles{Color: rl.Black, ZIndex: 2},
			InitialCoords: common.Coords{X: windowWidth / 2, Y: windowHeight/2 + 40},
			Id:            "change-bg-btn-id",
		},
		so.injector,
	)
	backBtn := NewTextObject(
		TextObjectStructParams{
			InitialValue:  "Back",
			FontSize:      24,
			Styles:        &common.ObjectStyles{Color: rl.Black, ZIndex: 2},
			InitialCoords: common.Coords{X: windowWidth / 2, Y: windowHeight/2 + 80},
		},
		so.injector,
	)

	styles.StyleTranslate(wrapper, -0.5, -0.5)
	styles.StyleTranslate(toggleFpsBtn, -0.5, -2)
	styles.StyleTranslate(changeBackgroundBtn, -0.5, -2)
	styles.StyleTranslate(backBtn, -0.5, -2)

	mouseClickEventUnsubFn := so.injector.EventManager().
		MouseClickEventEmitter().
		Subscribe(func(event sdkconstants.MouseClickEvent) {
			_toggleFpsBtnEventHandler(event, toggleFpsBtn, so.injector)
			_changeBackgroundBtnEventHandler(event, changeBackgroundBtn, so.injector, so)
			_backBtnEventHandler(event, backBtn, so.injector)
		})
	mouseMoveEventUnsubFn := so.injector.EventManager().
		MouseMoveEventEmitter().
		Subscribe(func(event sdkconstants.MouseMoveEvent) {
			_mouseMoveEventHandler(event, so.textObjectsPool, so.injector)
		})

	so.appendUnsubs(mouseClickEventUnsubFn, mouseMoveEventUnsubFn)

	so.appendTextObjects(toggleFpsBtn, changeBackgroundBtn, backBtn)
	so.appendRectObjects(wrapper)
}

func _toggleFpsBtnEventHandler(
	mouseEvent sdkconstants.MouseClickEvent,
	textObj interfaces.ITextObject,
	injector *injector.Injector,
) {
	mouseCoords := common.Coords{X: mouseEvent.PosX, Y: mouseEvent.PosY}
	if utils.IsTextObjectHovered(mouseCoords, textObj.Coords(), textObj.Width(), textObj.Height()) {
		if mouseEvent.Kind == rl.MouseButtonLeft {
			injector.GameSettings().SetShowFps(!injector.GameSettings().IsShownFps())
			injector.EventManager().FpsToggleEventEmitter().Emit(sdkconstants.FpsToggleEvent{
				ShowFps: injector.GameSettings().IsShownFps(),
			})
		}
	}
}

func _changeBackgroundBtnEventHandler(
	mouseEvent sdkconstants.MouseClickEvent,
	textObj interfaces.ITextObject,
	injector *injector.Injector,
	screen *ScreenOptions,
) {
	mouseCoords := common.Coords{X: mouseEvent.PosX, Y: mouseEvent.PosY}
	if utils.IsTextObjectHovered(mouseCoords, textObj.Coords(), textObj.Width(), textObj.Height()) {
		if mouseEvent.Kind == rl.MouseButtonLeft {
			currentBgName := injector.GameSettings().MainBackground()
			appBackgrounds := injector.ImageLoader().GetAppBackgrounds()

			var nextBg sdkconstants.AppImage
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
				Emit(sdkconstants.BackgroundChangeEvent{
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
	mouseEvent sdkconstants.MouseClickEvent,
	textObj interfaces.ITextObject,
	injector *injector.Injector,
) {
	mouseCoords := common.Coords{X: mouseEvent.PosX, Y: mouseEvent.PosY}
	if utils.IsTextObjectHovered(mouseCoords, textObj.Coords(), textObj.Width(), textObj.Height()) {
		if mouseEvent.Kind == rl.MouseButtonLeft {
			injector.Renderer().SetScreen(NewScreenMenu(injector))
		}
	}
}
