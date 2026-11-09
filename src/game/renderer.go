package game

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	sdkmodels "github.com/pseudoelement/ping-pong/sdk/models"
	"github.com/pseudoelement/ping-pong/src/injector"
	"github.com/pseudoelement/ping-pong/src/interfaces"
	"github.com/pseudoelement/ping-pong/src/models"
	"github.com/pseudoelement/ping-pong/src/models/ui"
)

type Renderer struct {
	injector          *injector.Injector
	openedScreen      interfaces.IScreen
	staticTextObjects []interfaces.ITextObject

	fpsShown bool
}

var _ interfaces.IRenderer = (*Renderer)(nil)

func NewRenderer(injector *injector.Injector) *Renderer {
	return &Renderer{
		injector:          injector,
		openedScreen:      nil,
		staticTextObjects: make([]interfaces.ITextObject, 0),
		fpsShown:          true,
	}
}

func (r *Renderer) Render() {
	rl.BeginDrawing()

	capacity :=
		len(r.openedScreen.TextObjects()) +
			len(r.openedScreen.RectObjects()) +
			len(r.staticTextObjects)
	allObjects := make([]interfaces.IBaseObject, 0, capacity)
	settings := r.injector.GameSettings()

	for _, textObject := range r.openedScreen.TextObjects() {
		rl.DrawText(
			textObject.Value(),
			textObject.Coords().X,
			textObject.Coords().Y,
			textObject.FontSize(),
			textObject.Color(),
		)
		allObjects = append(allObjects, textObject)
	}
	for _, rectObject := range r.openedScreen.RectObjects() {
		rl.DrawRectangle(
			rectObject.Coords().X,
			rectObject.Coords().Y,
			100,
			100,
			rectObject.Color(),
		)
		allObjects = append(allObjects, rectObject)
	}
	for _, textObject := range r.staticTextObjects {
		rl.DrawText(
			textObject.Value(),
			textObject.Coords().X,
			textObject.Coords().Y,
			textObject.FontSize(),
			textObject.Color(),
		)
		allObjects = append(allObjects, textObject)
	}

	for _, obj := range allObjects {
		updateableObj, ok := obj.(interfaces.IUpdateable)
		if ok {
			updateableObj.Update()
		}
	}

	nonNullableTextObjects := make([]interfaces.ITextObject, 0, len(r.staticTextObjects))
	for _, obj := range r.staticTextObjects {
		if r.fpsShown && !settings.ShowFps && obj.Id() == FpsObjectId {
			continue
		}
		nonNullableTextObjects = append(nonNullableTextObjects, obj)
	}
	r.staticTextObjects = nonNullableTextObjects

	rl.EndDrawing()

}

func (r *Renderer) CurrentScreen() interfaces.IScreen {
	return r.openedScreen
}

func (r *Renderer) SetScreen(screen interfaces.IScreen) {
	r.openedScreen = screen
}

func (r *Renderer) setStaticObjects() {
	fpsObject := ui.NewFpsObject(
		models.TextObjectStructParams{
			InitialValue:  "0 FPS",
			FontSize:      32,
			Color:         rl.Green,
			InitialCoords: sdkmodels.Coords{X: 10, Y: 5},
			Id:            FpsObjectId,
		},
		r.injector,
	)

	r.staticTextObjects = append(r.staticTextObjects, fpsObject)
}
