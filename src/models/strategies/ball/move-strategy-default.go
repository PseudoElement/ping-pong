package ball

import (
	"context"
	"math"
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

	stopBall                bool
	doesBallIntersectRacket bool

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
	strategy := &DefaultMovementStrategy{
		injector:    injector,
		gameStorage: gameStorage,
		unsubs:      make([]common.UnsubscribeFn, 0),
		velocity:    _randomizeVelocity(injector.GameSettings().BallVelocity()),
		stopBall:    true,
		ball:        ball,
		leftPlayer:  leftPlayer,
		rightPlayer: rightPlayer,
	}
	utils.Timeout(
		context.Background(),
		func() { strategy.stopBall = false },
		time.Second*2,
	)
	return strategy
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
	if dms.stopBall {
		return
	}

	windowWidth := dms.injector.GameSettings().WindowWidth()
	windowHeight := dms.injector.GameSettings().WindowHeight()

	next := common.Coords{
		X: dms.ball.Coords().X + dms.velocity.X,
		Y: dms.ball.Coords().Y + dms.velocity.Y,
	}

	if next.X < 0-dms.leftPlayer.Width() {
		dms.gameStorage.GoalEventEmitter().Emit(gameservices.Score{
			LeftPlayer:  dms.gameStorage.GoalEventEmitter().Value().LeftPlayer,
			RightPlayer: dms.gameStorage.GoalEventEmitter().Value().RightPlayer + 1,
		})

		dms.stopBall = true
		utils.Timeout(
			context.Background(),
			func() { dms.stopBall = false },
			time.Second*2,
		)
		dms.velocity = _randomizeVelocityRightDirection(
			dms.injector.GameSettings().BallVelocity(),
		)

		dms.ball.SetCoords(common.Coords{
			X: dms.injector.GameSettings().WindowWidth()/2 - (dms.ball.Width() / 2),
			Y: dms.injector.GameSettings().WindowHeight()/2 - (dms.ball.Height() / 2),
		})
		return
	}
	if next.X > windowWidth+dms.ball.Width() {
		dms.gameStorage.GoalEventEmitter().Emit(gameservices.Score{
			LeftPlayer:  dms.gameStorage.GoalEventEmitter().Value().LeftPlayer + 1,
			RightPlayer: dms.gameStorage.GoalEventEmitter().Value().RightPlayer,
		})

		dms.stopBall = true
		utils.Timeout(
			context.Background(),
			func() { dms.stopBall = false },
			time.Second*2,
		)
		dms.velocity = _randomizeVelocityLeftDirection(
			dms.injector.GameSettings().BallVelocity(),
		)

		dms.ball.SetCoords(common.Coords{
			X: dms.injector.GameSettings().WindowWidth()/2 - (dms.ball.Width() / 2),
			Y: dms.injector.GameSettings().WindowHeight()/2 - (dms.ball.Height() / 2),
		})
		return
	}

	if next.Y <= 0 || next.Y+dms.ball.Height() >= windowHeight {
		dms.gameStorage.BallKickedEventEmitter().Emit(struct{}{})
		dms.velocity.Y = _randomizeVelocityWithOppositeSign(dms.velocity).Y
		next.Y = dms.ball.Coords().Y + dms.velocity.Y
	}

	kicksBallByLeftPlayer :=
		dms.ball.Coords().Y <= dms.leftPlayer.Coords().Y+dms.leftPlayer.Height() &&
			dms.ball.Coords().Y+dms.ball.Height() >= dms.leftPlayer.Coords().Y &&
			dms.ball.Coords().X < dms.leftPlayer.Coords().X+dms.leftPlayer.Width() &&
			dms.ball.Coords().X+dms.ball.Width() > dms.leftPlayer.Coords().X+dms.leftPlayer.Width()

	kicksBallByRightPlayer :=
		dms.ball.Coords().Y <= dms.rightPlayer.Coords().Y+dms.rightPlayer.Height() &&
			dms.ball.Coords().Y+dms.ball.Height() >= dms.rightPlayer.Coords().Y &&
			dms.ball.Coords().X+dms.ball.Width() > dms.rightPlayer.Coords().X &&
			dms.ball.Coords().X < dms.rightPlayer.Coords().X

	if kicksBallByLeftPlayer || kicksBallByRightPlayer {
		/* Prevent infinite coordX changes in next frames until the ball leaves area of racket */
		if dms.doesBallIntersectRacket {
			dms.ball.SetCoords(next)
		} else {
			dms.gameStorage.BallKickedEventEmitter().Emit(struct{}{})
			dms.velocity.X = _randomizeVelocityWithOppositeSign(dms.velocity).X
			next.X = dms.ball.Coords().X + dms.velocity.X
			dms.doesBallIntersectRacket = true
			dms.ball.SetCoords(next)
		}
	} else {
		dms.doesBallIntersectRacket = false
		dms.ball.SetCoords(next)
	}
}

func _randomizeVelocity(baseVelocity common.Coords) common.Coords {
	var velocity common.Coords
	randY := rand.Intn(2)
	randX := rand.Intn(2)

	if randX == 0 {
		velocity.X = baseVelocity.X
	} else {
		velocity.X = -baseVelocity.X
	}

	if randY == 0 {
		velocity.Y = baseVelocity.Y
	} else {
		velocity.Y = -baseVelocity.Y
	}

	return velocity
}

func _randomizeVelocityWithOppositeSign(prevVelocity common.Coords) common.Coords {
	velocity := _randomizeVelocity(prevVelocity)

	if prevVelocity.X > 0 {
		velocity.X = -int32(math.Abs(float64(velocity.X)))
	} else {
		velocity.X = int32(math.Abs(float64(velocity.X)))
	}

	if prevVelocity.Y > 0 {
		velocity.Y = -int32(math.Abs(float64(velocity.Y)))
	} else {
		velocity.Y = int32(math.Abs(float64(velocity.X)))
	}

	return velocity
}

func _randomizeVelocityLeftDirection(baseVelocity common.Coords) common.Coords {
	var velocity common.Coords
	velocity.X = -baseVelocity.X

	topOrBottom := rand.Intn(2)
	randVelocityY := int32(6 + rand.Intn(int(baseVelocity.Y)-5))
	if topOrBottom == 0 {
		velocity.Y = randVelocityY
	} else {
		velocity.Y = -randVelocityY
	}

	return velocity
}

func _randomizeVelocityRightDirection(baseVelocity common.Coords) common.Coords {
	var velocity common.Coords
	velocity.X = baseVelocity.X

	topOrBottom := rand.Intn(2)
	randVelocityY := int32(6 + rand.Intn(int(baseVelocity.Y)-5))
	if topOrBottom == 0 {
		velocity.Y = randVelocityY
	} else {
		velocity.Y = -randVelocityY
	}

	return velocity
}
