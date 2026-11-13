package models

import (
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/pseudoelement/ping-pong/sdk/common"
	sdkconstants "github.com/pseudoelement/ping-pong/sdk/constants"
	"github.com/pseudoelement/ping-pong/sdk/interfaces"
	"github.com/pseudoelement/ping-pong/sdk/styles"
	"github.com/pseudoelement/ping-pong/sdk/utils"
	"github.com/pseudoelement/ping-pong/src/constants"
	"github.com/pseudoelement/ping-pong/src/injector"
	"github.com/pseudoelement/ping-pong/src/models/ball"
	"github.com/pseudoelement/ping-pong/src/models/player"
)

type ScreenMenu struct {
	*ScreenBase
}

var _ interfaces.IScreen = (*ScreenMenu)(nil)

func NewScreenMenu(injector *injector.Injector) *ScreenMenu {
	screen := &ScreenMenu{
		ScreenBase: NewScreenBase(injector),
	}
	screen._init()
	return screen
}

func (sg *ScreenMenu) OnDestroy() {
	sg.onDestroyBase()
}

func (sm *ScreenMenu) _init() {
	windowWidth := sm.injector.GameSettings().WindowWidth()
	windowHeight := sm.injector.GameSettings().WindowHeight()

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
		sm.injector,
	)
	playWithBotBtn := NewTextObject(
		TextObjectStructParams{
			InitialValue:  "Play with bot",
			FontSize:      24,
			Styles:        &common.ObjectStyles{Color: rl.Black, ZIndex: 2},
			InitialCoords: common.Coords{X: windowWidth / 2, Y: windowHeight / 2},
		},
		sm.injector,
	)
	playWithPlayerBtn := NewTextObject(
		TextObjectStructParams{
			InitialValue:  "Play with player",
			FontSize:      24,
			Styles:        &common.ObjectStyles{Color: rl.Black, ZIndex: 2},
			InitialCoords: common.Coords{X: windowWidth / 2, Y: windowHeight/2 + 40},
		},
		sm.injector,
	)
	optionsBtn := NewTextObject(
		TextObjectStructParams{
			InitialValue:  "Game settings",
			FontSize:      24,
			Styles:        &common.ObjectStyles{Color: rl.Black, ZIndex: 2},
			InitialCoords: common.Coords{X: windowWidth / 2, Y: windowHeight/2 + 80},
		},
		sm.injector,
	)
	exitBtn := NewTextObject(
		TextObjectStructParams{
			InitialValue:  "Exit",
			FontSize:      24,
			Styles:        &common.ObjectStyles{Color: rl.Black, ZIndex: 2},
			InitialCoords: common.Coords{X: windowWidth / 2, Y: windowHeight/2 + 120},
		},
		sm.injector,
	)

	styles.StyleTranslate(wrapper, -0.5, -0.5)
	styles.StyleTranslate(playWithBotBtn, -0.5, -3)
	styles.StyleTranslate(playWithPlayerBtn, -0.5, -3)
	styles.StyleTranslate(optionsBtn, -0.5, -3)
	styles.StyleTranslate(exitBtn, -0.5, -3)

	mouseClickEventUnsubFn := sm.injector.EventManager().
		MouseClickEventEmitter().
		Subscribe(func(event sdkconstants.MouseClickEvent) {
			_playWithBotBtnEventHandler(event, playWithBotBtn, sm.injector)
			_playWithPlayerBtnEventHandler(event, playWithPlayerBtn, sm.injector)
			_optionsBtnEventHandler(event, optionsBtn, sm.injector)
			_exitBtnEventHandler(event, exitBtn)
		})
	mouseMoveEventUnsubFn := sm.injector.EventManager().
		MouseMoveEventEmitter().
		Subscribe(func(event sdkconstants.MouseMoveEvent) {
			_mouseMoveEventHandler(event, sm.textObjectsPool)
		})

	sm.appendUnsubs(mouseClickEventUnsubFn, mouseMoveEventUnsubFn)

	sm.appendTextObjects(playWithBotBtn, playWithPlayerBtn, optionsBtn, exitBtn)
	sm.appendRectObjects(wrapper)
}

func _playWithBotBtnEventHandler(
	mouseEvent sdkconstants.MouseClickEvent,
	textObj interfaces.ITextObject,
	injector *injector.Injector,
) {
	mouseCoords := common.Coords{X: mouseEvent.PosX, Y: mouseEvent.PosY}
	if utils.IsTextObjectHovered(mouseCoords, textObj.Coords(), textObj.Width(), textObj.Height()) {
		if mouseEvent.Kind == rl.MouseButtonLeft {
			injector.Renderer().SetScreen(NewScreenPlayWithBot(injector))
		}
	}
}

func _playWithPlayerBtnEventHandler(
	mouseEvent sdkconstants.MouseClickEvent,
	textObj interfaces.ITextObject,
	injector *injector.Injector,
) {
	mouseCoords := common.Coords{X: mouseEvent.PosX, Y: mouseEvent.PosY}
	if utils.IsTextObjectHovered(mouseCoords, textObj.Coords(), textObj.Width(), textObj.Height()) {
		if mouseEvent.Kind == rl.MouseButtonLeft {
			screenGame := NewScreenGame(injector)

			leftPlayerObj, ok := utils.GetChildByIdScreen(screenGame, IdLeftPlayer)
			if !ok {
				panic("[_playWithPlayerBtnEventHandler] leftPlayerObj not found")
			}
			rightPlayerObj, ok := utils.GetChildByIdScreen(screenGame, IdRightPlayer)
			if !ok {
				panic("[_playWithPlayerBtnEventHandler] rightPlayerObj not found")
			}
			ballObj, ok := utils.GetChildByIdScreen(screenGame, IdBall)
			if !ok {
				panic("[_playWithPlayerBtnEventHandler] ballObj not found")
			}

			leftPlayerMovementStrategy := player.NewWasdKeysMovementStrategy(
				leftPlayerObj.(interfaces.IObjectWithWidthAndHeight),
				injector,
			)
			rightPlayerMovementStrategy := player.NewArrowsKeysMovementStrategy(
				rightPlayerObj.(interfaces.IObjectWithWidthAndHeight),
				injector,
			)
			ballMovementStrategy := ball.NewDefaultMovementStrategy(
				ballObj.(interfaces.IObjectWithWidthAndHeight),
				leftPlayerObj.(interfaces.IObjectWithWidthAndHeight),
				rightPlayerObj.(interfaces.IObjectWithWidthAndHeight),
				screenGame.gameStorage,
				injector,
			)

			screenGame.SetMovementStrategies(
				leftPlayerMovementStrategy,
				rightPlayerMovementStrategy,
				ballMovementStrategy,
			)
			injector.Renderer().SetScreen(screenGame)
		}
	}
}

func _optionsBtnEventHandler(
	mouseEvent sdkconstants.MouseClickEvent,
	textObj interfaces.ITextObject,
	injector *injector.Injector,
) {
	mouseCoords := common.Coords{X: mouseEvent.PosX, Y: mouseEvent.PosY}
	if utils.IsTextObjectHovered(mouseCoords, textObj.Coords(), textObj.Width(), textObj.Height()) {
		if mouseEvent.Kind == rl.MouseButtonLeft {
			injector.Renderer().SetScreen(NewScreenOptions(injector))
		}
	}
}

func _exitBtnEventHandler(mouseEvent sdkconstants.MouseClickEvent, textObj interfaces.ITextObject) {
	mouseCoords := common.Coords{X: mouseEvent.PosX, Y: mouseEvent.PosY}
	if utils.IsTextObjectHovered(mouseCoords, textObj.Coords(), textObj.Width(), textObj.Height()) {
		if mouseEvent.Kind == rl.MouseButtonLeft {
			rl.CloseWindow()
			os.Exit(0)
		}
	}
}

func _mouseMoveEventHandler(
	mouseEvent sdkconstants.MouseMoveEvent,
	textObjects []interfaces.ITextObject,
) {
	mouseCoords := common.Coords{X: mouseEvent.PosX, Y: mouseEvent.PosY}
	objHovered := false

	for _, textObj := range textObjects {
		if utils.IsTextObjectHovered(mouseCoords, textObj.Coords(), textObj.Width(), textObj.Height()) {
			textObj.SetHover(true)
			textObj.Styles().SetColor(rl.Red)
			objHovered = true
		} else {
			if textObj.IsHovered() {
				textObj.SetHover(false)
				textObj.Styles().SetColor(rl.Black)
			}
		}
	}

	if objHovered {
		rl.SetMouseCursor(rl.MouseCursorPointingHand)
	} else {
		rl.SetMouseCursor(rl.MouseCursorDefault)
	}
}
