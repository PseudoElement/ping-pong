package movement

import (
	"github.com/pseudoelement/ping-pong/sdk/common"
	"github.com/pseudoelement/ping-pong/sdk/constants"
	"github.com/pseudoelement/ping-pong/sdk/interfaces"
)

type Direction int

const (
	DirectionTop Direction = iota
	DirectionBottom
)

type TopBottomMovementStrategy struct {
	target    interfaces.IObjectWithWidthAndHeight
	injector  interfaces.IInjector
	unsubs    []common.UnsubscribeFn
	direction Direction
}

var _ interfaces.IMovementStrategy = (*TopBottomMovementStrategy)(nil)

func NewTopBottomMovementStrategy(
	target interfaces.IObjectWithWidthAndHeight,
	injector interfaces.IInjector,
) *TopBottomMovementStrategy {
	strategy := &TopBottomMovementStrategy{
		target:    target,
		injector:  injector,
		unsubs:    make([]common.UnsubscribeFn, 0),
		direction: DirectionTop,
	}
	return strategy
}

func (akms *TopBottomMovementStrategy) Setup() {
	keyPressedUnsubFn := akms.injector.EventManager().
		NextFrameEventEmitter().
		Subscribe(func(event constants.NextFrameEvent) {
			_topBottomEventHandler(akms.target, akms.injector, &akms.direction)
		})

	akms.unsubs = append(akms.unsubs, keyPressedUnsubFn)
}

func (akms *TopBottomMovementStrategy) Stop() {
	for _, unsub := range akms.unsubs {
		unsub()
	}
}

func _topBottomEventHandler(
	target interfaces.IObjectWithWidthAndHeight,
	injector interfaces.IInjector,
	direction *Direction,
) {
	stepY := int32(8)

	switch *direction {
	case DirectionTop:
		newPosY := target.Coords().Y - stepY
		if newPosY < 0 {
			newPosY = 0
			*direction = DirectionBottom
		}
		target.SetCoords(common.Coords{
			Y: newPosY,
			X: target.Coords().X,
		})
	case DirectionBottom:
		newPosY := target.Coords().Y + stepY
		if newPosY+target.Height() > injector.GameSettings().WindowHeight() {
			newPosY = injector.GameSettings().WindowHeight() - target.Height()
			*direction = DirectionTop
		}
		target.SetCoords(common.Coords{
			Y: newPosY,
			X: target.Coords().X,
		})
	}
}
