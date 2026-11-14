package gameservices

import "github.com/pseudoelement/ping-pong/sdk/common"

type GameStorage struct {
	goalEventEmitter       *common.EventEmitter[Score]
	ballKickedEventEmitter *common.EventEmitter[struct{}]
}

func NewGameStorage() *GameStorage {
	return &GameStorage{
		goalEventEmitter:       common.NewEventEmitter(Score{}),
		ballKickedEventEmitter: common.NewEventEmitter(struct{}{}),
	}
}

func (gs *GameStorage) GoalEventEmitter() *common.EventEmitter[Score] {
	return gs.goalEventEmitter
}

func (gs *GameStorage) BallKickedEventEmitter() *common.EventEmitter[struct{}] {
	return gs.ballKickedEventEmitter
}
