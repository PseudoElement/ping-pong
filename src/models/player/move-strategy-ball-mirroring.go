package player

import (
	"github.com/pseudoelement/ping-pong/sdk/common"
	"github.com/pseudoelement/ping-pong/sdk/constants"
	"github.com/pseudoelement/ping-pong/sdk/interfaces"
)

type BallMirroringMovementStrategy struct {
	target   interfaces.IObjectWithWidthAndHeight
	ball     interfaces.IObjectWithWidthAndHeight
	injector interfaces.IInjector
	unsubs   []common.UnsubscribeFn
}

var _ interfaces.IMovementStrategy = (*BallMirroringMovementStrategy)(nil)

func NewBallMirroringMovementStrategy(
	target interfaces.IObjectWithWidthAndHeight,
	ball interfaces.IObjectWithWidthAndHeight,
	injector interfaces.IInjector,
) *BallMirroringMovementStrategy {
	return &BallMirroringMovementStrategy{
		target:   target,
		ball:     ball,
		injector: injector,
	}
}

func (bmms *BallMirroringMovementStrategy) Setup() {
	keyPressedUnsubFn := bmms.injector.EventManager().
		NextFrameEventEmitter().
		Subscribe(func(event constants.NextFrameEvent) {
			bmms._nextFrameEventHandler()
		})

	bmms.unsubs = append(bmms.unsubs, keyPressedUnsubFn)
}

func (bmms *BallMirroringMovementStrategy) Stop() {
	for _, unsub := range bmms.unsubs {
		unsub()
	}
}

func (bmms *BallMirroringMovementStrategy) ChangeTarget(newTarget interfaces.IObjectWithWidthAndHeight) {
	bmms.Stop()
	bmms.target = newTarget
	bmms.Setup()
}

func (bmms *BallMirroringMovementStrategy) _nextFrameEventHandler() {
	windowHeight := bmms.injector.GameSettings().WindowHeight()

	nextCoords := common.Coords{
		X: bmms.target.Coords().X,
		Y: bmms.ball.Coords().Y,
	}

	if nextCoords.Y < 0 {
		nextCoords.Y = 0
	}
	if nextCoords.Y+bmms.target.Height() > windowHeight {
		nextCoords.Y = windowHeight - bmms.target.Height()
	}

	bmms.target.SetCoords(nextCoords)
}
