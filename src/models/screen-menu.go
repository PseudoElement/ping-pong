package models

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/pseudoelement/ping-pong/sdk/models"
	"github.com/pseudoelement/ping-pong/sdk/utils"
	"github.com/pseudoelement/ping-pong/src/constants"
	"github.com/pseudoelement/ping-pong/src/injector"
	"github.com/pseudoelement/ping-pong/src/interfaces"
)

type ScreenMenu struct {
	injector                 *injector.Injector
	textObjectsPool          []interfaces.ITextObject
	rectObjectsPool          []interfaces.IBaseObject
	mouseClickEventListeners map[string]interfaces.MouseClickEventHandler
}

// var _ IScreen = (*ScreenMenu)(nil)

func NewScreenMenu(injector *injector.Injector) *ScreenMenu {
	screen := &ScreenMenu{
		injector:                 injector,
		textObjectsPool:          make([]interfaces.ITextObject, 0),
		rectObjectsPool:          make([]interfaces.IBaseObject, 0),
		mouseClickEventListeners: make(map[string]interfaces.MouseClickEventHandler),
	}
	screen._init()
	return screen
}

func (sm *ScreenMenu) TextObjects() []interfaces.ITextObject {
	return sm.textObjectsPool
}

func (sm *ScreenMenu) RectObjects() []interfaces.IBaseObject {
	return sm.rectObjectsPool
}

func (sm *ScreenMenu) _init() {
	windowWidth := sm.injector.Settings().WindowWidth
	windowHeight := sm.injector.Settings().WindowHeight

	playWithBotBtn := NewTextObject(TextObjectStructParams{
		InitialValue:  "Play with bot",
		FontSize:      24,
		Color:         rl.Black,
		InitialCoords: models.Coords{X: windowWidth/2 - 50, Y: windowHeight/2 - 50},
	})
	playWithPlayerBtn := NewTextObject(TextObjectStructParams{
		InitialValue:  "Play with player",
		FontSize:      24,
		Color:         rl.Black,
		InitialCoords: models.Coords{X: windowWidth/2 - 50, Y: windowHeight/2 - 50 + 30},
	})
	exitBtn := NewTextObject(TextObjectStructParams{
		InitialValue:  "Exit",
		FontSize:      24,
		Color:         rl.Black,
		InitialCoords: models.Coords{X: windowWidth/2 - 50, Y: windowHeight/2 - 50 + 60},
	})

	sm.mouseClickEventListeners[playWithBotBtn.Id()] = playWithBotBtnEventHandler

	sm.injector.EventManager().MouseEventEmitter().Subscribe(func(event constants.MouseClickEvent) {
		playWithBotBtnEventHandler(event, playWithBotBtn)
	})

	sm.textObjectsPool = append(sm.textObjectsPool, playWithBotBtn, playWithPlayerBtn, exitBtn)
}

func playWithBotBtnEventHandler(mouseEvent constants.MouseClickEvent, textObj interfaces.ITextObject) {
	mouseCoords := models.Coords{X: mouseEvent.PosX, Y: mouseEvent.PosY}
	if utils.IsTextObjectClicked(mouseCoords, textObj.Coords(), textObj.Width(), textObj.Height()) {
		println("PLAY_WITH_BOT_BTN_CLICKED_!")
	}
}
