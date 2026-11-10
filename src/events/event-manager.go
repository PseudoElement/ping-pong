package events

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/pseudoelement/ping-pong/sdk/models"
	"github.com/pseudoelement/ping-pong/src/constants"
	"github.com/pseudoelement/ping-pong/src/interfaces"
)

type EventManager struct {
	mouseClickEventEmitter *models.EventEmitter[constants.MouseClickEvent]
	mouseMoveEventEmitter  *models.EventEmitter[constants.MouseMoveEvent]
	keyboardEventEmitter   *models.EventEmitter[constants.KeyPressedEvent]
	fpsToggleEventEmitter  *models.EventEmitter[constants.FpsToggleEvent]
}

var _ interfaces.IEventManager = (*EventManager)(nil)

func NewEventManager() *EventManager {
	rl.GetMouseX()
	return &EventManager{
		mouseClickEventEmitter: models.NewEventEmitter(constants.MouseClickEvent{}),
		mouseMoveEventEmitter:  models.NewEventEmitter(constants.MouseMoveEvent{}),
		keyboardEventEmitter:   models.NewEventEmitter(constants.KeyPressedEvent{}),
		fpsToggleEventEmitter:  models.NewEventEmitter(constants.FpsToggleEvent{}),
	}
}

func (em *EventManager) MouseClickEventEmitter() *models.EventEmitter[constants.MouseClickEvent] {
	return em.mouseClickEventEmitter
}

func (em *EventManager) MouseMoveEventEmitter() *models.EventEmitter[constants.MouseMoveEvent] {
	return em.mouseMoveEventEmitter
}

func (em *EventManager) KeyboardEventEmitter() *models.EventEmitter[constants.KeyPressedEvent] {
	return em.keyboardEventEmitter
}

func (em *EventManager) FpsToggleEventEmitter() *models.EventEmitter[constants.FpsToggleEvent] {
	return em.fpsToggleEventEmitter
}

func (em *EventManager) UseUserInputHandlers() {
	/* MOUSE EVENTS */
	if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
		em.mouseClickEventEmitter.Emit(constants.MouseClickEvent{
			PosX: rl.GetMouseX(),
			PosY: rl.GetMouseY(),
			Kind: rl.MouseButtonLeft,
		})
	}
	if rl.IsMouseButtonPressed(rl.MouseRightButton) {
		em.mouseClickEventEmitter.Emit(constants.MouseClickEvent{
			PosX: rl.GetMouseX(),
			PosY: rl.GetMouseY(),
			Kind: rl.MouseButtonRight,
		})
	}
	em.mouseMoveEventEmitter.Emit(constants.MouseMoveEvent{
		PosX: rl.GetMouseX(),
		PosY: rl.GetMouseY(),
	})

	/* KEYBOARD EVENTS */
	for _, keyPressed := range []int32{rl.KeyW, rl.KeyS, rl.KeyA, rl.KeyD} {
		if rl.IsKeyPressed((keyPressed)) {
			em.keyboardEventEmitter.Emit(constants.KeyPressedEvent{
				KeyCode: keyPressed,
			})
		}
	}
}
