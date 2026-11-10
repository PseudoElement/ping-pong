package models

import (
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/pseudoelement/ping-pong/sdk/models"
	"github.com/pseudoelement/ping-pong/sdk/styles"
	"github.com/pseudoelement/ping-pong/sdk/utils"
	"github.com/pseudoelement/ping-pong/src/constants"
	"github.com/pseudoelement/ping-pong/src/injector"
	"github.com/pseudoelement/ping-pong/src/interfaces"
)

type ScreenMenu struct {
	injector        *injector.Injector
	textObjectsPool []interfaces.ITextObject
	rectObjectsPool []interfaces.IBaseObject
	unsubs          []models.UnsubscribeFn
}

var _ interfaces.IScreen = (*ScreenMenu)(nil)

func NewScreenMenu(injector *injector.Injector) *ScreenMenu {
	screen := &ScreenMenu{
		injector:        injector,
		textObjectsPool: make([]interfaces.ITextObject, 0),
		rectObjectsPool: make([]interfaces.IBaseObject, 0),
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

func (sm *ScreenMenu) OnDestroy() {
	for _, unsubFn := range sm.unsubs {
		unsubFn()
	}
}

func (sm *ScreenMenu) _init() {
	windowWidth := sm.injector.GameSettings().WindowWidth
	windowHeight := sm.injector.GameSettings().WindowHeight

	playWithBotBtn := NewTextObject(
		TextObjectStructParams{
			InitialValue:  "Play with bot",
			FontSize:      24,
			Styles:        &styles.ObjectStyles{Color: rl.Black},
			InitialCoords: models.Coords{X: windowWidth/2 - 50, Y: windowHeight/2 - 50},
		},
		sm.injector,
	)
	playWithPlayerBtn := NewTextObject(
		TextObjectStructParams{
			InitialValue:  "Play with player",
			FontSize:      24,
			Styles:        &styles.ObjectStyles{Color: rl.Black},
			InitialCoords: models.Coords{X: windowWidth/2 - 50, Y: windowHeight/2 - 50 + 40},
		},
		sm.injector,
	)
	optionsBtn := NewTextObject(
		TextObjectStructParams{
			InitialValue:  "Game settings",
			FontSize:      24,
			Styles:        &styles.ObjectStyles{Color: rl.Black},
			InitialCoords: models.Coords{X: windowWidth/2 - 50, Y: windowHeight/2 - 50 + 80},
		},
		sm.injector,
	)
	exitBtn := NewTextObject(
		TextObjectStructParams{
			InitialValue:  "Exit",
			FontSize:      24,
			Styles:        &styles.ObjectStyles{Color: rl.Black},
			InitialCoords: models.Coords{X: windowWidth/2 - 50, Y: windowHeight/2 - 50 + 120},
		},
		sm.injector,
	)

	mouseClickEventUnsubFn := sm.injector.EventManager().
		MouseClickEventEmitter().
		Subscribe(func(event constants.MouseClickEvent) {
			_playWithBotBtnEventHandler(event, playWithBotBtn, sm.injector)
			_playWithPlayerBtnEventHandler(event, playWithPlayerBtn, sm.injector)
			_optionsBtnEventHandler(event, optionsBtn, sm.injector)
			_exitBtnEventHandler(event, exitBtn)
		})
	mouseMoveEventUnsubFn := sm.injector.EventManager().
		MouseMoveEventEmitter().
		Subscribe(func(event constants.MouseMoveEvent) {
			_mouseMoveEventHandler(event, sm.textObjectsPool, sm.injector)
		})

	sm.unsubs = append(sm.unsubs, mouseClickEventUnsubFn, mouseMoveEventUnsubFn)
	sm.textObjectsPool = append(sm.textObjectsPool, playWithBotBtn, playWithPlayerBtn, optionsBtn, exitBtn)
}

func _playWithBotBtnEventHandler(
	mouseEvent constants.MouseClickEvent,
	textObj interfaces.ITextObject,
	injector *injector.Injector,
) {
	mouseCoords := models.Coords{X: mouseEvent.PosX, Y: mouseEvent.PosY}
	if utils.IsTextObjectHovered(mouseCoords, textObj.Coords(), textObj.Width(), textObj.Height()) {
		if mouseEvent.Kind == rl.MouseButtonLeft {
			injector.Renderer().SetScreen(NewScreenGame(injector))
		}
	}
}

func _playWithPlayerBtnEventHandler(
	mouseEvent constants.MouseClickEvent,
	textObj interfaces.ITextObject,
	injector *injector.Injector,
) {
	mouseCoords := models.Coords{X: mouseEvent.PosX, Y: mouseEvent.PosY}
	if utils.IsTextObjectHovered(mouseCoords, textObj.Coords(), textObj.Width(), textObj.Height()) {
		if mouseEvent.Kind == rl.MouseButtonLeft {
			injector.Renderer().CurrentScreen().OnDestroy()
			injector.Renderer().SetScreen(NewScreenGame(injector))
		}
	}
}

func _optionsBtnEventHandler(
	mouseEvent constants.MouseClickEvent,
	textObj interfaces.ITextObject,
	injector *injector.Injector,
) {
	mouseCoords := models.Coords{X: mouseEvent.PosX, Y: mouseEvent.PosY}
	if utils.IsTextObjectHovered(mouseCoords, textObj.Coords(), textObj.Width(), textObj.Height()) {
		if mouseEvent.Kind == rl.MouseButtonLeft {
			injector.Renderer().CurrentScreen().OnDestroy()
			injector.Renderer().SetScreen(NewScreenOptions(injector))
		}
	}
}

func _exitBtnEventHandler(mouseEvent constants.MouseClickEvent, textObj interfaces.ITextObject) {
	mouseCoords := models.Coords{X: mouseEvent.PosX, Y: mouseEvent.PosY}
	if utils.IsTextObjectHovered(mouseCoords, textObj.Coords(), textObj.Width(), textObj.Height()) {
		if mouseEvent.Kind == rl.MouseButtonLeft {
			rl.CloseWindow()
			os.Exit(0)
		}
	}
}

func _mouseMoveEventHandler(
	mouseEvent constants.MouseMoveEvent,
	textObjects []interfaces.ITextObject,
	injector *injector.Injector,
) {
	mouseCoords := models.Coords{X: mouseEvent.PosX, Y: mouseEvent.PosY}
	objHovered := false

	for _, textObj := range textObjects {
		if utils.IsTextObjectHovered(mouseCoords, textObj.Coords(), textObj.Width(), textObj.Height()) {
			textObj.SetHover(true)
			textObj.Styles().SetColor(rl.Red)
			objHovered = true
		} else {
			if textObj.IsHovered() {
				textObj.SetHover(false)
				textObj.Styles().SetColor(rl.Black)
			}
		}
	}

	if objHovered {
		rl.SetMouseCursor(rl.MouseCursorPointingHand)
	} else {
		rl.SetMouseCursor(rl.MouseCursorDefault)
	}
}
