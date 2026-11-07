package game

import (
	"fmt"

	"github.com/fstanis/screenresolution"
	rl "github.com/gen2brain/raylib-go/raylib"
	sdkmodels "github.com/pseudoelement/ping-pong/sdk/models"
	sdkstyles "github.com/pseudoelement/ping-pong/sdk/styles"
	"github.com/pseudoelement/ping-pong/src/models"
	"github.com/pseudoelement/ping-pong/src/settings"
)

type Game struct {
	settings *settings.GameSettings
}

func NewGame() *Game {
	return &Game{
		settings: settings.NewGameSettings(),
	}
}

func (g *Game) Start() {
	resolution := screenresolution.GetPrimary()
	if resolution.Height <= 0 || resolution.Width <= 0 {
		fmt.Println("[main] failed to get screen resolution")
		resolution.Height = 1000
		resolution.Width = 750
	}
	fmt.Printf("[main] Width: %d, Height: %d.\n", resolution.Width, resolution.Height)

	windowWidth, windowHeight := int32(resolution.Width), int32(resolution.Height-70)
	fontSize := int32(35)

	g.settings.SetFontSize(fontSize)
	g.settings.SetWindowHeight(windowHeight)
	g.settings.SetWindowWidth(windowWidth)

	rl.InitWindow(windowWidth, windowHeight, "PingPong(raylib)")
	rl.SetTargetFPS(60)

	textObjects := make([]sdkstyles.ITextObject, 3)
	textObjects[0] = models.NewTextObject(models.TextObjectStructParams{
		InitialValue:  "First",
		FontSize:      20,
		Color:         rl.Black,
		InitialCoords: sdkmodels.Coords{X: 100, Y: 100},
	})
	textObjects[1] = models.NewTextObject(models.TextObjectStructParams{
		InitialValue:  "Second",
		FontSize:      20,
		Color:         rl.Black,
		InitialCoords: sdkmodels.Coords{X: 100, Y: 100},
	})
	textObjects[2] = models.NewTextObject(models.TextObjectStructParams{
		InitialValue:  "Third",
		FontSize:      20,
		Color:         rl.Black,
		InitialCoords: sdkmodels.Coords{X: 100, Y: 100},
	})

	for !rl.WindowShouldClose() {
		rl.ClearBackground(rl.Pink)
		rl.BeginDrawing()
		sdkstyles.Flex(textObjects, sdkstyles.FLEX_DIR_COLUMN, windowWidth/2, windowHeight/2, 25)
		// rl.DrawText("End!", 100, windowHeight-60, fontSize, rl.Black)
		rl.EndDrawing()
	}

	rl.CloseWindow()
}
