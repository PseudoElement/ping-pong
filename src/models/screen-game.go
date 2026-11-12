package models

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/pseudoelement/ping-pong/sdk/common"
	"github.com/pseudoelement/ping-pong/sdk/interfaces"
	"github.com/pseudoelement/ping-pong/sdk/styles"
	"github.com/pseudoelement/ping-pong/src/constants"
	"github.com/pseudoelement/ping-pong/src/injector"
)

type ScreenGame struct {
	*ScreenBase

	leftPlayerMovementStrategy  interfaces.IMovementStrategy
	rightPlayerMovementStrategy interfaces.IMovementStrategy
}

var _ interfaces.IScreen = (*ScreenGame)(nil)

func NewScreenGame(injector *injector.Injector) *ScreenGame {
	screen := &ScreenGame{
		ScreenBase: NewScreenBase(injector),
	}
	screen._init()
	return screen
}

func (sg *ScreenGame) OnDestroy() {
	sg.onDestroyBase()
	if sg.leftPlayerMovementStrategy != nil {
		sg.leftPlayerMovementStrategy.Stop()
	}
	if sg.rightPlayerMovementStrategy != nil {
		sg.rightPlayerMovementStrategy.Stop()
	}
}

func (sg *ScreenGame) SetMovementStrategies(
	leftPlayerMovementStrategy interfaces.IMovementStrategy,
	rightPlayerMovementStrategy interfaces.IMovementStrategy,
) {
	sg.leftPlayerMovementStrategy = leftPlayerMovementStrategy
	sg.rightPlayerMovementStrategy = rightPlayerMovementStrategy

	sg.leftPlayerMovementStrategy.Setup()
	sg.rightPlayerMovementStrategy.Setup()
}

func (sg *ScreenGame) _init() {
	imgLoader := sg.injector.ImageLoader()
	settings := sg.injector.GameSettings()

	windowWidth := settings.WindowWidth()
	windowHeight := settings.WindowHeight()

	scoreObj := NewTextObject(
		TextObjectStructParams{
			Styles:        &common.ObjectStyles{Color: rl.Red, ZIndex: math.MaxUint16},
			InitialValue:  "0 : 0",
			FontSize:      50,
			InitialCoords: common.Coords{X: windowWidth / 2, Y: 5},
		},
		sg.injector,
	)
	leftPlayerObj := NewImageObject(
		ImageObjectStructParams{
			Styles:        &common.ObjectStyles{ZIndex: 1},
			InitialCoords: common.Coords{X: 0, Y: windowHeight / 2},
			Texture:       imgLoader.GetTexture(constants.ImgPlayerRocketLeft),
			Id:            IdLeftPlayer,
		},
		sg.injector,
	)
	rightPlayerObj := NewImageObject(
		ImageObjectStructParams{
			Styles:        &common.ObjectStyles{ZIndex: 1},
			InitialCoords: common.Coords{X: windowWidth, Y: windowHeight / 2},
			Texture:       imgLoader.GetTexture(constants.ImgPlayerRocketRight),
			Id:            IdRightPlayer,
		},
		sg.injector,
	)
	ballObj := NewImageObject(
		ImageObjectStructParams{
			Styles:        &common.ObjectStyles{ZIndex: 2},
			InitialCoords: common.Coords{X: windowWidth / 2, Y: windowHeight / 2},
			Texture:       imgLoader.GetTexture(constants.ImgBallBaseball),
			Id:            IdBall,
		},
		sg.injector,
	)

	styles.StyleTranslate(scoreObj, -0.5, 0)
	styles.StyleTranslate(leftPlayerObj, 0, -0.5)
	styles.StyleTranslate(rightPlayerObj, -1, -0.5)
	styles.StyleTranslate(ballObj, -0.5, -0.5)

	sg.appendTextObjects(scoreObj)
	sg.appendImageObjects(leftPlayerObj, rightPlayerObj, ballObj)
}
