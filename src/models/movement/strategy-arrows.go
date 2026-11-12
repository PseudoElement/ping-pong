package movement

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/pseudoelement/ping-pong/sdk/common"
	"github.com/pseudoelement/ping-pong/sdk/constants"
	"github.com/pseudoelement/ping-pong/sdk/interfaces"
)

type ArrowsKeysMovementStrategy struct {
	target   interfaces.IObjectWithWidthAndHeight
	injector interfaces.IInjector
	unsubs   []common.UnsubscribeFn
}

var _ interfaces.IMovementStrategy = (*ArrowsKeysMovementStrategy)(nil)

func NewArrowsKeysMovementStrategy(
	target interfaces.IObjectWithWidthAndHeight,
	injector interfaces.IInjector,
) *ArrowsKeysMovementStrategy {
	strategy := &ArrowsKeysMovementStrategy{
		target:   target,
		injector: injector,
		unsubs:   make([]common.UnsubscribeFn, 0),
	}
	return strategy
}

func (akms *ArrowsKeysMovementStrategy) Setup() {
	keyPressedUnsubFn := akms.injector.EventManager().
		KeyDownEventEmitter().
		Subscribe(func(event constants.KeyDownEvent) {
			_arrowPressedEventHandler(event, akms.target, akms.injector)
		})

	akms.unsubs = append(akms.unsubs, keyPressedUnsubFn)
}

func (akms *ArrowsKeysMovementStrategy) Stop() {
	for _, unsub := range akms.unsubs {
		unsub()
	}
}

func _arrowPressedEventHandler(
	event constants.KeyDownEvent,
	target interfaces.IObjectWithWidthAndHeight,
	injector interfaces.IInjector,
) {
	stepY := int32(8)

	switch event.KeyCode {
	case rl.KeyUp:
		if target.Coords().Y >= 0+stepY {
			target.SetCoords(common.Coords{
				Y: target.Coords().Y - stepY,
				X: target.Coords().X,
			})
		}
	case rl.KeyDown:
		if target.Coords().Y+target.Height() <= injector.GameSettings().WindowHeight()-stepY {
			target.SetCoords(common.Coords{
				Y: target.Coords().Y + stepY,
				X: target.Coords().X,
			})
		}
	}
}
