package events

import (
	"log"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type GameEventStrategy struct {
	keyboardEventHandlers map[KeyboardKey]func()
	mouseEventHandlers    map[rl.MouseButton]func()
}

var _ IEventStrategy = (*GameEventStrategy)(nil)

func NewGameEventStrategy() *GameEventStrategy {
	strategy := &GameEventStrategy{
		keyboardEventHandlers: make(map[KeyboardKey]func()),
		mouseEventHandlers:    make(map[rl.MouseButton]func()),
	}
	strategy._init()
	return strategy
}

func (ges *GameEventStrategy) HandleKeyboardEvent(keyPressed KeyboardKey) {
	eventHandler, ok := ges.keyboardEventHandlers[keyPressed]
	if ok {
		eventHandler()
	} else {
		log.Println("[GameEventStrategy_HandleKeyboardEvent] unexpected event: ", keyPressed)
	}
}

func (ges *GameEventStrategy) SupportsKeyboardEvent(keyPressed KeyboardKey) bool {
	_, ok := ges.keyboardEventHandlers[keyPressed]
	return ok
}

func (ges *GameEventStrategy) HandleMouseEvent(mousePressed rl.MouseButton) {
	eventHandler, ok := ges.mouseEventHandlers[mousePressed]
	if ok {
		eventHandler()
	} else {
		log.Println("[GameEventStrategy_HandleMouseEvent] unexpected event: ", mousePressed)
	}
}

func (ges *GameEventStrategy) SupportsMouseEvent(mousePressed rl.MouseButton) bool {
	_, ok := ges.mouseEventHandlers[mousePressed]
	return ok
}

func (mse *GameEventStrategy) _init() {
	mse.mouseEventHandlers[rl.MouseLeftButton] = func() {
	}
	mse.mouseEventHandlers[rl.MouseRightButton] = func() {
	}
}
