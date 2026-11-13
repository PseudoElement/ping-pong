package models

import (
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

type ScreenPlayWithBot struct {
	*ScreenBase
}

var _ interfaces.IScreen = (*ScreenPlayWithBot)(nil)

func NewScreenPlayWithBot(injector *injector.Injector) *ScreenPlayWithBot {
	screen := &ScreenPlayWithBot{
		ScreenBase: NewScreenBase(injector),
	}
	screen._init()
	return screen
}

func (sg *ScreenPlayWithBot) OnDestroy() {
	sg.onDestroyBase()
}

func (sm *ScreenPlayWithBot) _init() {
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
	playWithEasyBotBtn := NewTextObject(
		TextObjectStructParams{
			InitialValue:  "Easy",
			FontSize:      24,
			Styles:        &common.ObjectStyles{Color: rl.Black, ZIndex: 2},
			InitialCoords: common.Coords{X: windowWidth / 2, Y: windowHeight / 2},
		},
		sm.injector,
	)
	playWithProBotBtn := NewTextObject(
		TextObjectStructParams{
			InitialValue:  "Professional",
			FontSize:      24,
			Styles:        &common.ObjectStyles{Color: rl.Black, ZIndex: 2},
			InitialCoords: common.Coords{X: windowWidth / 2, Y: windowHeight/2 + 40},
		},
		sm.injector,
	)

	styles.StyleTranslate(wrapper, -0.5, -0.5)
	styles.StyleTranslate(playWithEasyBotBtn, -0.5, -3)
	styles.StyleTranslate(playWithProBotBtn, -0.5, -3)

	mouseClickEventUnsubFn := sm.injector.EventManager().
		MouseClickEventEmitter().
		Subscribe(func(event sdkconstants.MouseClickEvent) {
			_playWithEasyBotBtnEventHandler(event, playWithEasyBotBtn, sm.injector)
			_playWithProBotBtnEventHandler(event, playWithProBotBtn, sm.injector)
		})
	mouseMoveEventUnsubFn := sm.injector.EventManager().
		MouseMoveEventEmitter().
		Subscribe(func(event sdkconstants.MouseMoveEvent) {
			_mouseMoveEventHandler(event, sm.textObjectsPool)
		})

	sm.appendUnsubs(mouseClickEventUnsubFn, mouseMoveEventUnsubFn)

	sm.appendTextObjects(playWithEasyBotBtn, playWithProBotBtn)
	sm.appendRectObjects(wrapper)
}

func _playWithProBotBtnEventHandler(
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
				panic("[_playWithProBotBtnEventHandler] leftPlayerObj not found")
			}
			rightPlayerObj, ok := utils.GetChildByIdScreen(screenGame, IdRightPlayer)
			if !ok {
				panic("[_playWithProBotBtnEventHandler] rightPlayerObj not found")
			}
			ballObj, ok := utils.GetChildByIdScreen(screenGame, IdBall)
			if !ok {
				panic("[_playWithProBotBtnEventHandler] ballObj not found")
			}

			leftPlayerMovementStrategy := player.NewWasdKeysMovementStrategy(
				leftPlayerObj.(interfaces.IObjectWithWidthAndHeight),
				injector,
			)
			rightPlayerMovementStrategy := player.NewBallMirroringMovementStrategy(
				rightPlayerObj.(interfaces.IObjectWithWidthAndHeight),
				ballObj.(interfaces.IObjectWithWidthAndHeight),
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

func _playWithEasyBotBtnEventHandler(
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
				panic("[_playWithEasyBotBtnEventHandler] leftPlayerObj not found")
			}
			rightPlayerObj, ok := utils.GetChildByIdScreen(screenGame, IdRightPlayer)
			if !ok {
				panic("[_playWithEasyBotBtnEventHandler] rightPlayerObj not found")
			}
			ballObj, ok := utils.GetChildByIdScreen(screenGame, IdBall)
			if !ok {
				panic("[_playWithEasyBotBtnEventHandler] ballObj not found")
			}

			leftPlayerMovementStrategy := player.NewWasdKeysMovementStrategy(
				leftPlayerObj.(interfaces.IObjectWithWidthAndHeight),
				injector,
			)
			rightPlayerMovementStrategy := player.NewTopBottomMovementStrategy(
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
