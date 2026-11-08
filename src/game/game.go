package game

import (
	"fmt"

	"github.com/fstanis/screenresolution"
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/pseudoelement/ping-pong/src/events"
	"github.com/pseudoelement/ping-pong/src/injector"
	"github.com/pseudoelement/ping-pong/src/models"
	"github.com/pseudoelement/ping-pong/src/settings"
	"github.com/pseudoelement/ping-pong/src/sounds"
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
	soundLoader := sounds.NewSoundLoader()
	gameSettings := settings.NewGameSettings()
	renderer := NewRenderer(g.injector)

	g.injector.SetEventManager(eventManager)
	g.injector.SetSoundLoader(soundLoader)
	g.injector.SetGameSettings(gameSettings)
	g.injector.SetRenderer(renderer)

	g.injector.Settings().SetFontSize(fontSize)
	g.injector.Settings().SetWindowHeight(windowHeight)
	g.injector.Settings().SetWindowWidth(windowWidth)

	renderer.setScreen(models.NewScreenMenu(g.injector))

	// imgPath := utils.AppPath("../assets/images/pad.png")
	// img := rl.LoadImage(imgPath.Normalize())
	// texture := rl.LoadTextureFromImage(img)

	for !rl.WindowShouldClose() {
		// rl.DrawTexture(texture, 100, 100, rl.White)
		rl.ClearBackground(rl.White)
		eventManager.UseEventHandlers()
		renderer.Render()
	}

	rl.CloseWindow()
}
