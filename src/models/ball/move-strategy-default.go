package ball

import (
	"context"
	"math/rand"
	"time"

	"github.com/pseudoelement/ping-pong/sdk/common"
	"github.com/pseudoelement/ping-pong/sdk/constants"
	"github.com/pseudoelement/ping-pong/sdk/interfaces"
	"github.com/pseudoelement/ping-pong/sdk/utils"
	gameservices "github.com/pseudoelement/ping-pong/src/models/services"
)

type DefaultMovementStrategy struct {
	injector    interfaces.IInjector
	gameStorage *gameservices.GameStorage

	unsubs   []common.UnsubscribeFn
	velocity common.Coords

	isDelayedAfterGoal bool

	ball        interfaces.IObjectWithWidthAndHeight
	leftPlayer  interfaces.IObjectWithWidthAndHeight
	rightPlayer interfaces.IObjectWithWidthAndHeight
}

var _ interfaces.IMovementStrategy = (*DefaultMovementStrategy)(nil)

func NewDefaultMovementStrategy(
	ball interfaces.IObjectWithWidthAndHeight,
	leftPlayer interfaces.IObjectWithWidthAndHeight,
	rightPlayer interfaces.IObjectWithWidthAndHeight,
	gameStorage *gameservices.GameStorage,
	injector interfaces.IInjector,
) *DefaultMovementStrategy {
	return &DefaultMovementStrategy{
		injector:           injector,
		gameStorage:        gameStorage,
		unsubs:             make([]common.UnsubscribeFn, 0),
		velocity:           _randomizeVelocity(),
		isDelayedAfterGoal: false,
		ball:               ball,
		leftPlayer:         leftPlayer,
		rightPlayer:        rightPlayer,
	}
}

func (dms *DefaultMovementStrategy) Setup() {
	keyPressedUnsubFn := dms.injector.EventManager().
		NextFrameEventEmitter().
		Subscribe(func(event constants.NextFrameEvent) {
			dms._defaultMovementEventHandler()
		})

	dms.unsubs = append(dms.unsubs, keyPressedUnsubFn)
}

func (dms *DefaultMovementStrategy) Stop() {
	for _, unsub := range dms.unsubs {
		unsub()
	}
}

func (dms *DefaultMovementStrategy) ChangeTarget(newTarget interfaces.IObjectWithWidthAndHeight) {
	dms.Stop()
	dms.ball = newTarget
	dms.Setup()
}

func (dms *DefaultMovementStrategy) _defaultMovementEventHandler() {
	if dms.isDelayedAfterGoal {
		return
	}

	windowWidth := dms.injector.GameSettings().WindowWidth()
	windowHeight := dms.injector.GameSettings().WindowHeight()

	next := common.Coords{
		X: dms.ball.Coords().X + dms.velocity.X,
		Y: dms.ball.Coords().Y + dms.velocity.Y,
	}

	kicksBallByLeftPlayer :=
		dms.ball.Coords().Y <= dms.leftPlayer.Coords().Y+dms.leftPlayer.Height() &&
			dms.ball.Coords().Y+dms.ball.Height() >= dms.leftPlayer.Coords().Y &&
			dms.ball.Coords().X <= dms.leftPlayer.Coords().X+dms.leftPlayer.Width() &&
			dms.ball.Coords().X+dms.ball.Width() >= dms.leftPlayer.Coords().X

	kicksBallByRightPlayer :=
		dms.ball.Coords().Y <= dms.rightPlayer.Coords().Y+dms.rightPlayer.Height() &&
			dms.ball.Coords().Y+dms.ball.Height() >= dms.rightPlayer.Coords().Y &&
			dms.ball.Coords().X+dms.ball.Width() >= dms.rightPlayer.Coords().X &&
			dms.ball.Coords().X <= dms.rightPlayer.Coords().X+dms.rightPlayer.Width()

	if next.X < 0-dms.leftPlayer.Width() {
		dms.gameStorage.ScoreEventEmmiter().Emit(gameservices.Score{
			LeftPlayer:  dms.gameStorage.ScoreEventEmmiter().Value().LeftPlayer,
			RightPlayer: dms.gameStorage.ScoreEventEmmiter().Value().RightPlayer + 1,
		})

		dms.isDelayedAfterGoal = true
		utils.Timeout(
			context.Background(),
			func() { dms.isDelayedAfterGoal = false },
			3*time.Second,
		)
		dms.velocity = _randomizeVelocityRightDirection()

		dms.ball.SetCoords(common.Coords{
			X: dms.injector.GameSettings().WindowWidth()/2 - (dms.ball.Width() / 2),
			Y: dms.injector.GameSettings().WindowHeight()/2 - (dms.ball.Height() / 2),
		})
		return
	}
	if next.X > windowWidth+dms.ball.Width() {
		dms.gameStorage.ScoreEventEmmiter().Emit(gameservices.Score{
			LeftPlayer:  dms.gameStorage.ScoreEventEmmiter().Value().LeftPlayer + 1,
			RightPlayer: dms.gameStorage.ScoreEventEmmiter().Value().RightPlayer,
		})

		dms.isDelayedAfterGoal = true
		utils.Timeout(
			context.Background(),
			func() { dms.isDelayedAfterGoal = false },
			3*time.Second,
		)
		dms.velocity = _randomizeVelocityLeftDirection()

		dms.ball.SetCoords(common.Coords{
			X: dms.injector.GameSettings().WindowWidth()/2 - (dms.ball.Width() / 2),
			Y: dms.injector.GameSettings().WindowHeight()/2 - (dms.ball.Height() / 2),
		})
		return
	}

	if kicksBallByLeftPlayer || kicksBallByRightPlayer {
		dms.velocity.X = -dms.velocity.X
		next.X = dms.ball.Coords().X + dms.velocity.X
	}
	if next.Y <= 0 || next.Y+dms.ball.Height() >= windowHeight {
		dms.velocity.Y = -dms.velocity.Y
		next.Y = dms.ball.Coords().Y + dms.velocity.Y
	}

	dms.ball.SetCoords(next)
}

func _randomizeVelocity() common.Coords {
	var velocity common.Coords
	randY := rand.Intn(2)
	randX := rand.Intn(2)

	if randX == 0 {
		velocity.X = 6
	} else {
		velocity.X = -6
	}

	if randY == 0 {
		velocity.Y = 6
	} else {
		velocity.Y = -6
	}

	return velocity
}

func _randomizeVelocityLeftDirection() common.Coords {
	var velocity common.Coords
	velocity.X = -6

	randY := rand.Intn(2)
	if randY == 0 {
		velocity.Y = 6
	} else {
		velocity.Y = -6
	}

	return velocity
}

func _randomizeVelocityRightDirection() common.Coords {
	var velocity common.Coords
	velocity.X = 6

	randY := rand.Intn(2)
	if randY == 0 {
		velocity.Y = 6
	} else {
		velocity.Y = -6
	}

	return velocity
}
