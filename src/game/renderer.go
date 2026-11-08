package game

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/pseudoelement/ping-pong/src/injector"
	"github.com/pseudoelement/ping-pong/src/interfaces"
)

type Renderer struct {
	injector     *injector.Injector
	openedScreen interfaces.IScreen
}

var _ interfaces.IRenderer = (*Renderer)(nil)

func NewRenderer(injector *injector.Injector) *Renderer {
	return &Renderer{}
}

func (r *Renderer) Render() {
	rl.BeginDrawing()
	for _, textObject := range r.openedScreen.TextObjects() {
		rl.DrawText(
			textObject.Value(),
			textObject.Coords().X,
			textObject.Coords().Y,
			textObject.FontSize(),
			textObject.Color(),
		)
	}
	for _, rectObject := range r.openedScreen.RectObjects() {
		rl.DrawRectangle(
			rectObject.Coords().X,
			rectObject.Coords().Y,
			100,
			100,
			rectObject.Color(),
		)
	}
	rl.EndDrawing()
}

func (r *Renderer) CurrentScreen() interfaces.IScreen {
	return r.openedScreen
}

func (r *Renderer) setScreen(screen interfaces.IScreen) {
	r.openedScreen = screen
}
