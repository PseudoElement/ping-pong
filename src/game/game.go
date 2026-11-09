package game

import (
	"context"
	"fmt"
	"time"

	"github.com/fstanis/screenresolution"
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/pseudoelement/ping-pong/sdk/utils"
	"github.com/pseudoelement/ping-pong/src/events"
	"github.com/pseudoelement/ping-pong/src/injector"
	"github.com/pseudoelement/ping-pong/src/loaders"
	"github.com/pseudoelement/ping-pong/src/models"
	"github.com/pseudoelement/ping-pong/src/settings"
)

type Game struct {
	injector *injector.Injector
}

func NewGame() *Game {
	return &Game{
		injector: injector.NewInjector(),
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

	rl.InitAudioDevice()
	rl.InitWindow(windowWidth, windowHeight, "PingPong(raylib)")
	rl.SetTargetFPS(60)

	eventManager := events.NewEventManager()
	soundLoader := loaders.NewSoundLoader()
	gameSettings := settings.NewGameSettings()
	renderer := NewRenderer(g.injector)

	g.injector.SetEventManager(eventManager)
	g.injector.SetSoundLoader(soundLoader)
	g.injector.SetGameSettings(gameSettings)
	g.injector.SetRenderer(renderer)

	g.injector.GameSettings().SetFontSize(fontSize)
	g.injector.GameSettings().SetWindowHeight(windowHeight)
	g.injector.GameSettings().SetWindowWidth(windowWidth)

	renderer.SetScreen(models.NewScreenMenu(g.injector))
	renderer.setStaticObjects()

	var frameCount int32 = 1
	fpsFn := func() {
		g.injector.GameSettings().SetFps(frameCount)
		frameCount = 1
	}
	utils.Interval(context.Background(), fpsFn, 1*time.Second)

	// imgPath := utils.AppPath("../assets/images/pad.png")
	// img := rl.LoadImage(imgPath.Normalize())
	// texture := rl.LoadTextureFromImage(img)

	for !rl.WindowShouldClose() {
		frameCount++
		// rl.DrawTexture(texture, 100, 100, rl.White)
		rl.ClearBackground(rl.White)
		eventManager.UseEventHandlers()
		renderer.Render()
	}

	rl.CloseWindow()
}
