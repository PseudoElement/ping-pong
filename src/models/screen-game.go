package models

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/pseudoelement/ping-pong/sdk/common"
	"github.com/pseudoelement/ping-pong/sdk/interfaces"
	"github.com/pseudoelement/ping-pong/sdk/styles"
	"github.com/pseudoelement/ping-pong/src/constants"
	appinterfaces "github.com/pseudoelement/ping-pong/src/interfaces"
	modelsconstants "github.com/pseudoelement/ping-pong/src/models/constants"
	gameservices "github.com/pseudoelement/ping-pong/src/models/services"
)

type ScreenGame struct {
	*ScreenBase

	gameStorage                 *gameservices.GameStorage
	leftPlayerMovementStrategy  appinterfaces.IMovementStrategy
	rightPlayerMovementStrategy appinterfaces.IMovementStrategy
	ballMovementStrategy        appinterfaces.IMovementStrategy
	gameEventEmitter            *common.EventEmitter[any]
}

var _ interfaces.IScreen = (*ScreenGame)(nil)

func NewScreenGame(injector appinterfaces.IInjector) *ScreenGame {
	screen := &ScreenGame{
		ScreenBase:  NewScreenBase(injector),
		gameStorage: gameservices.NewGameStorage(),
	}
	screen._init()
	return screen
}

func (sg *ScreenGame) OnDestroy() {
	sg.onDestroyBase()
	rl.ShowCursor()

	if sg.leftPlayerMovementStrategy != nil {
		sg.leftPlayerMovementStrategy.Stop()
	}
	if sg.rightPlayerMovementStrategy != nil {
		sg.rightPlayerMovementStrategy.Stop()
	}
	if sg.ballMovementStrategy != nil {
		sg.ballMovementStrategy.Stop()
	}
}

func (sg *ScreenGame) SetMovementStrategies(
	leftPlayerMovementStrategy appinterfaces.IMovementStrategy,
	rightPlayerMovementStrategy appinterfaces.IMovementStrategy,
	ballMovementStrategy appinterfaces.IMovementStrategy,
) {
	sg.leftPlayerMovementStrategy = leftPlayerMovementStrategy
	sg.rightPlayerMovementStrategy = rightPlayerMovementStrategy
	sg.ballMovementStrategy = ballMovementStrategy

	sg.leftPlayerMovementStrategy.Setup()
	sg.rightPlayerMovementStrategy.Setup()
	sg.ballMovementStrategy.Setup()
}

func (sg *ScreenGame) _init() {
	imgLoader := sg.injector.ImageLoader()
	settings := sg.injector.GameSettings()

	windowWidth := settings.WindowWidth()
	windowHeight := settings.WindowHeight()

	scoreObj := NewScoreObject(sg.gameStorage, sg.injector)
	leftPlayerObj := NewImageObject(
		ImageObjectStructParams{
			Styles:        &common.ObjectStyles{ZIndex: 1},
			InitialCoords: common.Coords{X: 0, Y: windowHeight / 2},
			Texture:       imgLoader.GetTexture(constants.ImgPlayerRocketLeft),
			Id:            modelsconstants.IdLeftPlayer,
		},
		sg.injector,
	)
	rightPlayerObj := NewImageObject(
		ImageObjectStructParams{
			Styles:        &common.ObjectStyles{ZIndex: 1},
			InitialCoords: common.Coords{X: windowWidth, Y: windowHeight / 2},
			Texture:       imgLoader.GetTexture(constants.ImgPlayerRocketRight),
			Id:            modelsconstants.IdRightPlayer,
		},
		sg.injector,
	)
	ballObj := NewImageObject(
		ImageObjectStructParams{
			Styles:        &common.ObjectStyles{ZIndex: 2},
			InitialCoords: common.Coords{X: windowWidth / 2, Y: windowHeight / 2},
			Texture:       imgLoader.GetTexture(sg.injector.GameSettings().Theme().ImgBall),
			Id:            modelsconstants.IdBall,
		},
		sg.injector,
	)

	styles.StyleTranslate(scoreObj, -0.5, 0)
	styles.StyleTranslate(leftPlayerObj, 0, -0.5)
	styles.StyleTranslate(rightPlayerObj, -1, -0.5)
	styles.StyleTranslate(ballObj, -0.5, -0.5)

	ballKickedEventUnsubFn := sg.gameStorage.
		BallKickedEventEmitter().
		Subscribe(func(s struct{}) {
			_ballKickedEventHandler(sg.injector)
		})

	sg.appendRectObjects(scoreObj)
	sg.appendImageObjects(leftPlayerObj, rightPlayerObj, ballObj)
	sg.appendUnsubs(ballKickedEventUnsubFn)

	rl.HideCursor()
}

func _ballKickedEventHandler(
	injector appinterfaces.IInjector,
) {
	rl.PlaySound(
		injector.SoundLoader().GetSound(constants.SoundKick),
	)
}
