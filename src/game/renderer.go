package game

import (
	"math"
	"slices"

	rl "github.com/gen2brain/raylib-go/raylib"
	sdkmodels "github.com/pseudoelement/ping-pong/sdk/models"
	"github.com/pseudoelement/ping-pong/sdk/styles"
	"github.com/pseudoelement/ping-pong/src/constants"
	"github.com/pseudoelement/ping-pong/src/injector"
	"github.com/pseudoelement/ping-pong/src/interfaces"
	"github.com/pseudoelement/ping-pong/src/loaders"
	"github.com/pseudoelement/ping-pong/src/models"
	"github.com/pseudoelement/ping-pong/src/models/ui"
)

type Renderer struct {
	injector      *injector.Injector
	openedScreen  interfaces.IScreen
	staticObjects []interfaces.IBaseObject
}

var _ interfaces.IRenderer = (*Renderer)(nil)

func NewRenderer(injector *injector.Injector) *Renderer {
	return &Renderer{
		injector:      injector,
		openedScreen:  nil,
		staticObjects: make([]interfaces.IBaseObject, 0),
	}
}

func (r *Renderer) Render() {
	rl.ClearBackground(rl.White)
	rl.BeginDrawing()

	capacity :=
		len(r.openedScreen.TextObjects()) +
			len(r.openedScreen.RectObjects()) +
			len(r.staticObjects)
	allObjects := make([]interfaces.IBaseObject, 0, capacity)

	for _, textObject := range r.openedScreen.TextObjects() {
		allObjects = append(allObjects, textObject)
	}
	for _, rectObject := range r.openedScreen.RectObjects() {
		allObjects = append(allObjects, rectObject)
	}
	for _, textObject := range r.staticObjects {
		allObjects = append(allObjects, textObject)
	}

	slices.SortFunc(allObjects, func(prev, next interfaces.IBaseObject) int {
		return int(next.Styles().ZIndex - prev.Styles().ZIndex)
	})

	for _, obj := range allObjects {
		updateableObj, ok := obj.(interfaces.IUpdateable)
		if ok {
			updateableObj.Update()
		}

		switch typedObj := obj.(type) {
		case interfaces.ITextObject:
			rl.DrawText(
				typedObj.Value(),
				typedObj.Coords().X,
				typedObj.Coords().Y,
				typedObj.FontSize(),
				typedObj.Styles().Color,
			)
		case interfaces.IImageObject:
			rl.DrawTexture(
				typedObj.Texture(),
				typedObj.Coords().X,
				typedObj.Coords().Y,
				rl.White,
			)
		case interfaces.IRectangleObject:
			rl.DrawRectangle(
				typedObj.Coords().X,
				typedObj.Coords().Y,
				100,
				100,
				typedObj.Styles().Color,
			)
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
	r.injector.EventManager().
		BackgroundChangeEventEmitter().
		Subscribe(func(event constants.BackgroundChangeEvent) {
			_backgroundChangeEventHandler(event, r)
		})
}

func (r *Renderer) _setStaticObjects() {
	settings := r.injector.GameSettings()
	imgLoader := r.injector.ImageLoader()

	fpsObject := ui.NewFpsObject(
		models.TextObjectStructParams{
			InitialValue:  "0 FPS",
			FontSize:      32,
			Styles:        &styles.ObjectStyles{Color: rl.Green, ZIndex: math.MaxInt16},
			InitialCoords: sdkmodels.Coords{X: 10, Y: 5},
			Id:            FpsObjectId,
		},
		r.injector,
	)
	appBackgroundObject := models.NewImageObject(
		models.ImageObjectStructParams{
			Styles:        &styles.ObjectStyles{ZIndex: math.MinInt16},
			InitialCoords: sdkmodels.Coords{X: 0, Y: 0},
			Texture:       imgLoader.GetImage(settings.MainBackground),
			Id:            MainBgObjectId,
		},
		r.injector,
	)

	r.staticObjects = append(r.staticObjects, fpsObject, appBackgroundObject)
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
				Styles:        &styles.ObjectStyles{Color: rl.Green, ZIndex: math.MaxInt16},
				InitialCoords: sdkmodels.Coords{X: 10, Y: 5},
				Id:            FpsObjectId,
			},
			renderer.injector,
		)
		renderer.staticObjects = append(renderer.staticObjects, fpsObject)
	} else {
		newStaticObjects := make([]interfaces.IBaseObject, 0, len(renderer.staticObjects))
		for _, obj := range renderer.staticObjects {
			if obj.Id() == FpsObjectId {
				continue
			}
			newStaticObjects = append(newStaticObjects, obj)
		}
		renderer.staticObjects = newStaticObjects
	}
}

func _backgroundChangeEventHandler(
	event constants.BackgroundChangeEvent,
	renderer *Renderer,
) {
	imgLoader := renderer.injector.ImageLoader()

	newStaticObjects := make([]interfaces.IBaseObject, 0, len(renderer.staticObjects))
	for _, obj := range renderer.staticObjects {
		if obj.Id() == MainBgObjectId {
			obj = models.NewImageObject(
				models.ImageObjectStructParams{
					Styles:        &styles.ObjectStyles{ZIndex: math.MinInt16},
					InitialCoords: sdkmodels.Coords{X: 0, Y: 0},
					Texture:       imgLoader.GetImage(loaders.AppImageName(event.Name)),
					Id:            MainBgObjectId,
				},
				renderer.injector,
			)
		}
		newStaticObjects = append(newStaticObjects, obj)
	}

	renderer.staticObjects = newStaticObjects
}
