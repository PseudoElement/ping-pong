package events

import (
	"log"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type MenuEventStrategy struct {
	keyboardEventHandlers map[KeyboardKey]func()
	mouseEventHandlers    map[rl.MouseButton]func()
}

var _ IEventStrategy = (*MenuEventStrategy)(nil)

func NewMenuEventStrategy() *MenuEventStrategy {
	strategy := &MenuEventStrategy{
		keyboardEventHandlers: make(map[KeyboardKey]func()),
		mouseEventHandlers:    make(map[rl.MouseButton]func()),
	}
	strategy._init()
	return strategy
}

func (mes *MenuEventStrategy) HandleKeyboardEvent(keyPressed KeyboardKey) {
	eventHandler, ok := mes.keyboardEventHandlers[keyPressed]
	if ok {
		eventHandler()
	}
}

func (mes *MenuEventStrategy) SupportsKeyboardEvent(keyPressed KeyboardKey) bool {
	_, ok := mes.keyboardEventHandlers[keyPressed]
	return ok
}

func (mes *MenuEventStrategy) HandleMouseEvent(mousePressed rl.MouseButton) {
	eventHandler, ok := mes.mouseEventHandlers[mousePressed]
	if ok {
		eventHandler()
	}
}

func (mes *MenuEventStrategy) SupportsMouseEvent(mousePressed rl.MouseButton) bool {
	_, ok := mes.mouseEventHandlers[mousePressed]
	return ok
}

func (mse *MenuEventStrategy) _init() {
	earnSound := rl.LoadSound("/Users/paveldavidovich/Desktop/Web/Backend/ping-pong/builds/assets/audios/coin-earn.mp3")
	mse.mouseEventHandlers[rl.MouseLeftButton] = func() {
		log.Println("LEFT BUTTON clicked!!!")
		rl.PlaySound(earnSound)
	}
	mse.mouseEventHandlers[rl.MouseRightButton] = func() {
		log.Println("RIGHT BUTTON clicked!!!")
	}
}
