package player

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/pseudoelement/ping-pong/sdk/common"
	"github.com/pseudoelement/ping-pong/sdk/constants"
	"github.com/pseudoelement/ping-pong/sdk/interfaces"
)

type WasdKeysMovementStrategy struct {
	target   interfaces.IObjectWithWidthAndHeight
	injector interfaces.IInjector
	unsubs   []common.UnsubscribeFn
}

var _ interfaces.IMovementStrategy = (*WasdKeysMovementStrategy)(nil)

func NewWasdKeysMovementStrategy(
	target interfaces.IObjectWithWidthAndHeight,
	injector interfaces.IInjector,
) *WasdKeysMovementStrategy {
	strategy := &WasdKeysMovementStrategy{
		target:   target,
		injector: injector,
		unsubs:   make([]common.UnsubscribeFn, 0),
	}
	return strategy
}

func (wkms *WasdKeysMovementStrategy) Setup() {
	keyPressedUnsubFn := wkms.injector.EventManager().
		KeyDownEventEmitter().
		Subscribe(func(event constants.KeyDownEvent) {
			wkms._wasdPressedEventHandler(event, wkms.target)
		})

	wkms.unsubs = append(wkms.unsubs, keyPressedUnsubFn)
}

func (wkms *WasdKeysMovementStrategy) Stop() {
	for _, unsub := range wkms.unsubs {
		unsub()
	}
}

func (wkms *WasdKeysMovementStrategy) ChangeTarget(newTarget interfaces.IObjectWithWidthAndHeight) {
	wkms.Stop()
	wkms.target = newTarget
	wkms.Setup()
}

func (wkms *WasdKeysMovementStrategy) _wasdPressedEventHandler(
	event constants.KeyDownEvent,
	target interfaces.IObjectWithWidthAndHeight,
) {
	stepY := int32(8)

	switch event.KeyCode {
	case rl.KeyW:
		if target.Coords().Y >= 0+stepY {
			target.SetCoords(common.Coords{
				Y: target.Coords().Y - stepY,
				X: target.Coords().X,
			})
		}
	case rl.KeyS:
		if target.Coords().Y+target.Height() <= wkms.injector.GameSettings().WindowHeight()-stepY {
			target.SetCoords(common.Coords{
				Y: target.Coords().Y + stepY,
				X: target.Coords().X,
			})
		}
	}
}
