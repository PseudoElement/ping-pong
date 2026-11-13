package player

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

func (tbms *TopBottomMovementStrategy) Setup() {
	keyPressedUnsubFn := tbms.injector.EventManager().
		NextFrameEventEmitter().
		Subscribe(func(event constants.NextFrameEvent) {
			tbms._topBottomEventHandler()
		})

	tbms.unsubs = append(tbms.unsubs, keyPressedUnsubFn)
}

func (tbms *TopBottomMovementStrategy) Stop() {
	for _, unsub := range tbms.unsubs {
		unsub()
	}
}

func (tbms *TopBottomMovementStrategy) ChangeTarget(newTarget interfaces.IObjectWithWidthAndHeight) {
	tbms.Stop()
	tbms.target = newTarget
	tbms.Setup()
}

func (tbms *TopBottomMovementStrategy) _topBottomEventHandler() {
	stepY := int32(6)

	switch tbms.direction {
	case DirectionTop:
		newPosY := tbms.target.Coords().Y - stepY
		if newPosY < 0 {
			newPosY = 0
			tbms.direction = DirectionBottom
		}
		tbms.target.SetCoords(common.Coords{
			Y: newPosY,
			X: tbms.target.Coords().X,
		})
	case DirectionBottom:
		newPosY := tbms.target.Coords().Y + stepY
		if newPosY+tbms.target.Height() > tbms.injector.GameSettings().WindowHeight() {
			newPosY = tbms.injector.GameSettings().WindowHeight() - tbms.target.Height()
			tbms.direction = DirectionTop
		}
		tbms.target.SetCoords(common.Coords{
			Y: newPosY,
			X: tbms.target.Coords().X,
		})
	}
}
