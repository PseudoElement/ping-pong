package game

import (
	"slices"

	rl "github.com/gen2brain/raylib-go/raylib"
	sdkmodels "github.com/pseudoelement/ping-pong/sdk/models"
	"github.com/pseudoelement/ping-pong/sdk/styles"
	"github.com/pseudoelement/ping-pong/src/constants"
	"github.com/pseudoelement/ping-pong/src/injector"
	"github.com/pseudoelement/ping-pong/src/interfaces"
	"github.com/pseudoelement/ping-pong/src/models"
	"github.com/pseudoelement/ping-pong/src/models/ui"
)

type Renderer struct {
	injector          *injector.Injector
	openedScreen      interfaces.IScreen
	staticTextObjects []interfaces.ITextObject
}

var _ interfaces.IRenderer = (*Renderer)(nil)

func NewRenderer(injector *injector.Injector) *Renderer {
	return &Renderer{
		injector:          injector,
		openedScreen:      nil,
		staticTextObjects: make([]interfaces.ITextObject, 0),
	}
}

func (r *Renderer) Render() {
	rl.BeginDrawing()

	capacity :=
		len(r.openedScreen.TextObjects()) +
			len(r.openedScreen.RectObjects()) +
			len(r.staticTextObjects)
	allObjects := make([]interfaces.IBaseObject, 0, capacity)

	screenTextObjects := make([]interfaces.ITextObject, len(r.openedScreen.TextObjects()))
	screenRectObjects := make([]interfaces.IBaseObject, len(r.openedScreen.RectObjects()))
	staticTextObjects := make([]interfaces.ITextObject, len(r.staticTextObjects))

	copy(screenTextObjects, r.openedScreen.TextObjects())
	copy(screenRectObjects, r.openedScreen.RectObjects())
	copy(staticTextObjects, r.staticTextObjects)

	slices.SortFunc(screenTextObjects, func(prev, next interfaces.ITextObject) int {
		return int(prev.Styles().ZIndex - next.Styles().ZIndex)
	})
	slices.SortFunc(screenRectObjects, func(prev, next interfaces.IBaseObject) int {
		return int(prev.Styles().ZIndex - next.Styles().ZIndex)
	})
	slices.SortFunc(staticTextObjects, func(prev, next interfaces.ITextObject) int {
		return int(prev.Styles().ZIndex - next.Styles().ZIndex)
	})

	// imgPath := utils.AppPath("../assets/images/pad.png")
	// img := rl.LoadImage(imgPath.Normalize())
	// texture := rl.LoadTextureFromImage(img)
	// rl.ImageResize(img, 100, 100)
	// rl.DrawTexture(texture, 100, 100, rl.White)

	for _, textObject := range screenTextObjects {
		rl.DrawText(
			textObject.Value(),
			textObject.Coords().X,
			textObject.Coords().Y,
			textObject.FontSize(),
			textObject.Styles().Color,
		)
		allObjects = append(allObjects, textObject)
	}
	for _, rectObject := range screenRectObjects {
		rl.DrawRectangle(
			rectObject.Coords().X,
			rectObject.Coords().Y,
			100,
			100,
			rectObject.Styles().Color,
		)
		allObjects = append(allObjects, rectObject)
	}
	for _, textObject := range staticTextObjects {
		rl.DrawText(
			textObject.Value(),
			textObject.Coords().X,
			textObject.Coords().Y,
			textObject.FontSize(),
			textObject.Styles().Color,
		)
		allObjects = append(allObjects, textObject)
	}

	for _, obj := range allObjects {
		updateableObj, ok := obj.(interfaces.IUpdateable)
		if ok {
			updateableObj.Update()
		}
	}

	rl.EndDrawing()

}

func (r *Renderer) CurrentScreen() interfaces.IScreen {
	return r.openedScreen
}

func (r *Renderer) SetScreen(screen interfaces.IScreen) {
	r.openedScreen = screen
}

func (r *Renderer) init() {
	r._setStaticObjects()

	r.injector.EventManager().
		FpsToggleEventEmitter().
		Subscribe(func(event constants.FpsToggleEvent) {
			_fpsToggleEventHandler(event, r)
		})
}

func (r *Renderer) _setStaticObjects() {
	fpsObject := ui.NewFpsObject(
		models.TextObjectStructParams{
			InitialValue:  "0 FPS",
			FontSize:      32,
			Styles:        &styles.ObjectStyles{Color: rl.Green},
			InitialCoords: sdkmodels.Coords{X: 10, Y: 5},
			Id:            FpsObjectId,
		},
		r.injector,
	)

	r.staticTextObjects = append(r.staticTextObjects, fpsObject)
}

func _fpsToggleEventHandler(
	event constants.FpsToggleEvent,
	renderer *Renderer,
) {
	if event.ShowFps {
		fpsObject := ui.NewFpsObject(
			models.TextObjectStructParams{
				InitialValue:  "0 FPS",
				FontSize:      32,
				Styles:        &styles.ObjectStyles{Color: rl.Green},
				InitialCoords: sdkmodels.Coords{X: 10, Y: 5},
				Id:            FpsObjectId,
			},
			renderer.injector,
		)
		renderer.staticTextObjects = append(renderer.staticTextObjects, fpsObject)
	} else {
		nonNullableTextObjects := make([]interfaces.ITextObject, 0, len(renderer.staticTextObjects))
		for _, obj := range renderer.staticTextObjects {
			if obj.Id() == FpsObjectId {
				continue
			}
			nonNullableTextObjects = append(nonNullableTextObjects, obj)
		}
		renderer.staticTextObjects = nonNullableTextObjects
	}
}
