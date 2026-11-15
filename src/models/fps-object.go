package models

import (
	"fmt"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/pseudoelement/ping-pong/sdk/common"
	sdkconstants "github.com/pseudoelement/ping-pong/sdk/constants"
	"github.com/pseudoelement/ping-pong/sdk/interfaces"
	"github.com/pseudoelement/ping-pong/sdk/styles"
	"github.com/pseudoelement/ping-pong/src/constants"
	"github.com/pseudoelement/ping-pong/src/injector"
)

type FpsObject struct {
	*RectangleObject

	unsubs []common.UnsubscribeFn
}

var _ interfaces.IRectangleObject = (*FpsObject)(nil)
var _ interfaces.IDestroyable = (*FpsObject)(nil)

func NewFpsObject(id string, injector *injector.Injector) *FpsObject {
	fpsObject := &FpsObject{
		RectangleObject: NewRectangleObject(
			RectangleObjectStructParams{
				Styles: &common.ObjectStyles{
					Color:        constants.ColorTransparentBlue,
					ZIndex:       math.MaxUint16 - 1,
					BorderRadius: 0.2,
				},
				Width:         140,
				Height:        30,
				InitialCoords: common.Coords{X: 5, Y: 5},
				Id:            id,
			},
			injector,
		),
		unsubs: make([]common.UnsubscribeFn, 0),
	}
	fpsObject._init()
	return fpsObject
}

func (fo *FpsObject) OnDestroy() {
	for _, unsub := range fo.unsubs {
		unsub()
	}
}

func (fo *FpsObject) _init() {
	fpsTextObj := NewTextObject(
		TextObjectStructParams{
			InitialCoords: common.Coords{
				X: fo.Coords().X + fo.Width()/2,
				Y: fo.Coords().Y + fo.Height()/2,
			},
			Styles: &common.ObjectStyles{
				Color:  rl.Red,
				ZIndex: math.MaxUint16,
			},
			InitialValue: "0 FPS",
			FontSize:     30,
		},
		fo.injector,
	)

	styles.StyleTranslate(fpsTextObj, -0.5, -0.5)

	nextFrameUnsubFn := fo.injector.EventManager().
		NextFrameEventEmitter().
		Subscribe(func(nfe sdkconstants.NextFrameEvent) {
			fo._nextFrameEventHandler(fpsTextObj)
		})

	fo.AppendChild(fpsTextObj)
	fo.unsubs = append(fo.unsubs, nextFrameUnsubFn)
}

func (fo *FpsObject) _nextFrameEventHandler(fpsTextObj interfaces.ITextObject) {
	settings := fo.injector.GameSettings()

	value := fmt.Sprintf("%d FPS", settings.Fps())
	fpsTextObj.SetValue(value)

	fpsTextObj.SetCoords(common.Coords{
		X: fo.Coords().X + fo.Width()/2,
		Y: fo.Coords().Y + fo.Height()/2,
	})
	styles.StyleTranslate(fpsTextObj, -0.5, -0.5)
}
