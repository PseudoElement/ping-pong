package events

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/pseudoelement/ping-pong/sdk/models"
	"github.com/pseudoelement/ping-pong/src/constants"
)

type EventManager struct {
	eventsStrategy IEventStrategy

	mouseEventEmitter    *models.EventEmitter[constants.MouseClickEvent]
	keyboardEventEmitter *models.EventEmitter[constants.KeyPressedEvent]
}

func NewEventManager() *EventManager {
	rl.GetMouseX()
	return &EventManager{
		eventsStrategy:       NewMenuEventStrategy(),
		mouseEventEmitter:    models.NewEventEmitter(constants.MouseClickEvent{}),
		keyboardEventEmitter: models.NewEventEmitter(constants.KeyPressedEvent{}),
	}
}

func (em *EventManager) MouseEventEmitter() *models.EventEmitter[constants.MouseClickEvent] {
	return em.mouseEventEmitter
}

func (em *EventManager) KeyboardEventEmitter() *models.EventEmitter[constants.KeyPressedEvent] {
	return em.keyboardEventEmitter
}

func (em *EventManager) UseEventHandlers() {
	/* MOUSE EVENTS */
	if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
		em.mouseEventEmitter.Emit(constants.MouseClickEvent{
			PosX: rl.GetMouseX(),
			PosY: rl.GetMouseY(),
			Kind: rl.MouseButtonLeft,
		})
	}
	if rl.IsMouseButtonPressed(rl.MouseRightButton) {
		em.mouseEventEmitter.Emit(constants.MouseClickEvent{
			PosX: rl.GetMouseX(),
			PosY: rl.GetMouseY(),
			Kind: rl.MouseButtonRight,
		})
	}

	/* KEYBOARD EVENTS */
	for _, keyPressed := range []int32{rl.KeyW, rl.KeyS, rl.KeyA, rl.KeyD} {
		if rl.IsKeyPressed((keyPressed)) {
			em.keyboardEventEmitter.Emit(constants.KeyPressedEvent{
				KeyCode: keyPressed,
			})
		}
	}
}
