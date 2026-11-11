package models

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/pseudoelement/ping-pong/sdk/common"
	"github.com/pseudoelement/ping-pong/sdk/interfaces"
	"github.com/pseudoelement/ping-pong/sdk/styles"
	"github.com/pseudoelement/ping-pong/src/injector"
)

type ScreenGame struct {
	injector        *injector.Injector
	textObjectsPool []interfaces.ITextObject
	imgObjectsPool  []interfaces.IImageObject
	unsubs          []common.UnsubscribeFn
}

var _ interfaces.IScreen = (*ScreenGame)(nil)

func NewScreenGame(injector *injector.Injector) *ScreenGame {
	screen := &ScreenGame{
		injector:        injector,
		textObjectsPool: make([]interfaces.ITextObject, 0),
		imgObjectsPool:  make([]interfaces.IImageObject, 0),
	}
	screen._init()
	return screen
}

func (sg *ScreenGame) TextObjects() []interfaces.ITextObject {
	return sg.textObjectsPool
}

func (sg *ScreenGame) RectObjects() []interfaces.IRectangleObject {
	return []interfaces.IRectangleObject{}
}

func (sg *ScreenGame) ImageObjects() []interfaces.IImageObject {
	return sg.imgObjectsPool
}

func (sg *ScreenGame) OnDestroy() {
	for _, unsubFn := range sg.unsubs {
		unsubFn()
	}
}

func (sg *ScreenGame) _init() {
	windowWidth := sg.injector.GameSettings().WindowWidth()

	scoreObj := NewTextObject(
		TextObjectStructParams{
			Styles:        &common.ObjectStyles{Color: rl.Red, ZIndex: math.MaxUint16},
			InitialValue:  "0 : 0",
			FontSize:      36,
			InitialCoords: common.Coords{X: windowWidth / 2, Y: 5},
		},
		sg.injector,
	)
	styles.StyleTranslate(scoreObj, -0.5, 0)

	sg.textObjectsPool = append(sg.textObjectsPool, scoreObj)
	sg.imgObjectsPool = append(sg.imgObjectsPool)
}
