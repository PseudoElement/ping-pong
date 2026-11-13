package gameservices

import "github.com/pseudoelement/ping-pong/sdk/common"

type GameStorage struct {
	scoreEventEmmiter *common.EventEmitter[Score]
}

func NewGameStorage() *GameStorage {
	return &GameStorage{
		scoreEventEmmiter: common.NewEventEmitter(Score{}),
	}
}

func (gs *GameStorage) ScoreEventEmmiter() *common.EventEmitter[Score] {
	return gs.scoreEventEmmiter
}
