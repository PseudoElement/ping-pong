package models

import (
	"github.com/pseudoelement/ping-pong/sdk/models"
	"github.com/pseudoelement/ping-pong/src/injector"
	"github.com/pseudoelement/ping-pong/src/interfaces"
)

type ScreenGame struct {
	injector        *injector.Injector
	textObjectsPool []interfaces.ITextObject
	rectObjectsPool []interfaces.IBaseObject
	unsubs          []models.UnsubscribeFn
}

var _ interfaces.IScreen = (*ScreenGame)(nil)

func NewScreenGame(injector *injector.Injector) *ScreenGame {
	screen := &ScreenGame{
		injector:        injector,
		textObjectsPool: make([]interfaces.ITextObject, 0),
		rectObjectsPool: make([]interfaces.IBaseObject, 0),
	}
	// screen._init()
	return screen
}

func (sm *ScreenGame) TextObjects() []interfaces.ITextObject {
	return sm.textObjectsPool
}

func (sm *ScreenGame) RectObjects() []interfaces.IBaseObject {
	return sm.rectObjectsPool
}

func (sm *ScreenGame) OnDestroy() {
	for _, unsubFn := range sm.unsubs {
		unsubFn()
	}
}
