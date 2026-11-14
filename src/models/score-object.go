package models

import (
	"fmt"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/pseudoelement/ping-pong/sdk/common"
	"github.com/pseudoelement/ping-pong/sdk/interfaces"
	"github.com/pseudoelement/ping-pong/sdk/styles"
	"github.com/pseudoelement/ping-pong/src/constants"
	"github.com/pseudoelement/ping-pong/src/injector"
	modelsconstants "github.com/pseudoelement/ping-pong/src/models/constants"
	gameservices "github.com/pseudoelement/ping-pong/src/models/services"
)

type ScoreObject struct {
	*RectangleObject

	gameStorage *gameservices.GameStorage
	unsubs      []common.UnsubscribeFn
}

var _ interfaces.IRectangleObject = (*ScoreObject)(nil)
var _ interfaces.IDestroyable = (*ScoreObject)(nil)

func NewScoreObject(
	gameStorage *gameservices.GameStorage,
	injector *injector.Injector,
) *ScoreObject {
	fpsObject := &ScoreObject{
		RectangleObject: NewRectangleObject(
			RectangleObjectStructParams{
				Styles: &common.ObjectStyles{
					Color:        constants.ColorTransparentBlue,
					ZIndex:       math.MaxUint16 - 1,
					BorderRadius: 0.2,
				},
				Width:         150,
				Height:        60,
				InitialCoords: common.Coords{X: injector.GameSettings().WindowWidth() / 2, Y: 5},
			},
			injector,
		),
		unsubs:      make([]common.UnsubscribeFn, 0),
		gameStorage: gameStorage,
	}
	fpsObject._init()
	return fpsObject
}

func (fo *ScoreObject) OnDestroy() {
	for _, unsub := range fo.unsubs {
		unsub()
	}
}

func (fo *ScoreObject) _init() {
	scoreTextObj := NewTextObject(
		TextObjectStructParams{
			InitialCoords: common.Coords{
				X: fo.Coords().X + fo.Width()/2,
				Y: fo.Coords().Y + fo.Height()/2,
			},
			Styles: &common.ObjectStyles{
				Color:  rl.Red,
				ZIndex: math.MaxUint16,
			},
			InitialValue: fmt.Sprintf(
				"%d : %d\n",
				fo.gameStorage.GoalEventEmitter().Value().LeftPlayer,
				fo.gameStorage.GoalEventEmitter().Value().RightPlayer,
			),
			FontSize: 50,
			Id:       modelsconstants.IdScoreText,
		},
		fo.injector,
	)

	goalEventUnsubFn := fo.gameStorage.
		GoalEventEmitter().
		Subscribe(func(score gameservices.Score) {
			_goalEventHandler(score, scoreTextObj, fo.injector)
		})

	styles.StyleTranslate(scoreTextObj, -0.5, -0.5)

	fo.AppendChild(scoreTextObj)
	fo.unsubs = append(fo.unsubs, goalEventUnsubFn)
}

func _goalEventHandler(
	score gameservices.Score,
	scoreObj interfaces.ITextObject,
	injector *injector.Injector,
) {
	rl.PlaySound(
		injector.SoundLoader().GetSound(constants.SoundGoal),
	)
	scoreObj.SetValue(
		fmt.Sprintf("%d : %d\n", score.LeftPlayer, score.RightPlayer),
	)
	scoreObj.SetCoords(common.Coords{
		X: injector.GameSettings().WindowWidth()/2 - (scoreObj.Width() / 2),
		Y: 5,
	})
}
