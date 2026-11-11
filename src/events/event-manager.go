package events

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/pseudoelement/ping-pong/sdk/common"
	"github.com/pseudoelement/ping-pong/sdk/constants"
	"github.com/pseudoelement/ping-pong/sdk/interfaces"
)

type EventManager struct {
	mouseClickEventEmitter *common.EventEmitter[constants.MouseClickEvent]
	mouseMoveEventEmitter  *common.EventEmitter[constants.MouseMoveEvent]
	keyboardEventEmitter   *common.EventEmitter[constants.KeyPressedEvent]
	fpsToggleEventEmitter  *common.EventEmitter[constants.FpsToggleEvent]
	bgChangedEventEmitter  *common.EventEmitter[constants.BackgroundChangeEvent]
	nextFrameEventEmitter  *common.EventEmitter[constants.NextFrameEvent]
}

var _ interfaces.IEventManager = (*EventManager)(nil)

func NewEventManager() *EventManager {
	rl.GetMouseX()
	return &EventManager{
		mouseClickEventEmitter: common.NewEventEmitter(constants.MouseClickEvent{}),
		mouseMoveEventEmitter:  common.NewEventEmitter(constants.MouseMoveEvent{}),
		keyboardEventEmitter:   common.NewEventEmitter(constants.KeyPressedEvent{}),
		fpsToggleEventEmitter:  common.NewEventEmitter(constants.FpsToggleEvent{}),
		bgChangedEventEmitter:  common.NewEventEmitter(constants.BackgroundChangeEvent{}),
		nextFrameEventEmitter:  common.NewEventEmitter(constants.NextFrameEvent{}),
	}
}

func (em *EventManager) MouseClickEventEmitter() *common.EventEmitter[constants.MouseClickEvent] {
	return em.mouseClickEventEmitter
}

func (em *EventManager) MouseMoveEventEmitter() *common.EventEmitter[constants.MouseMoveEvent] {
	return em.mouseMoveEventEmitter
}

func (em *EventManager) KeyboardEventEmitter() *common.EventEmitter[constants.KeyPressedEvent] {
	return em.keyboardEventEmitter
}

func (em *EventManager) FpsToggleEventEmitter() *common.EventEmitter[constants.FpsToggleEvent] {
	return em.fpsToggleEventEmitter
}

func (em *EventManager) BackgroundChangeEventEmitter() *common.EventEmitter[constants.BackgroundChangeEvent] {
	return em.bgChangedEventEmitter
}

func (em *EventManager) NextFrameEventEmitter() *common.EventEmitter[constants.NextFrameEvent] {
	return em.nextFrameEventEmitter
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
	for _, keyCode := range []int32{rl.KeyW, rl.KeyS, rl.KeyA, rl.KeyD} {
		if rl.IsKeyPressed((keyCode)) {
			em.keyboardEventEmitter.Emit(constants.KeyPressedEvent{
				KeyCode: keyCode,
			})
		}
	}

	/* OTHER EVENTS */
	em.nextFrameEventEmitter.Emit(constants.NextFrameEvent{})
}
