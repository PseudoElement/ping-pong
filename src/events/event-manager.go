package events

import rl "github.com/gen2brain/raylib-go/raylib"

type EventManager struct {
	eventsStrategy IEventStrategy
}

func NewEventManager() *EventManager {
	return &EventManager{
		eventsStrategy: NewMenuEventStrategy(),
	}
}

func (em *EventManager) UseEventHandlers() {
	/* MOUSE EVENTS */
	if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
		if em.eventsStrategy.SupportsMouseEvent(rl.MouseLeftButton) {
			em.eventsStrategy.HandleMouseEvent(rl.MouseLeftButton)
		}
	}
	if rl.IsMouseButtonPressed(rl.MouseRightButton) {
		if em.eventsStrategy.SupportsMouseEvent(rl.MouseRightButton) {
			em.eventsStrategy.HandleMouseEvent(rl.MouseRightButton)
		}
	}

	/* KEYBOARD EVENTS */
	if rl.IsKeyPressed(rl.KeyW) {
		if em.eventsStrategy.SupportsKeyboardEvent(rl.KeyW) {
			em.eventsStrategy.HandleKeyboardEvent(rl.KeyW)
		}
	}
	if rl.IsKeyPressed(rl.KeyS) {
		if em.eventsStrategy.SupportsKeyboardEvent(rl.KeyS) {
			em.eventsStrategy.HandleKeyboardEvent(rl.KeyS)
		}
	}
}
