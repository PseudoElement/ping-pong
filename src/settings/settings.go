package settings

type GameSettings struct {
	FontSize     int32
	WindowWidth  int32
	WindowHeight int32
	Fps          int32
	ShowFps      bool
}

func NewGameSettings() *GameSettings {
	return &GameSettings{
		FontSize:     0,
		WindowWidth:  0,
		WindowHeight: 0,
		Fps:          0,
		ShowFps:      true,
	}
}

func (gs *GameSettings) SetFontSize(fontSize int32) {
	gs.FontSize = fontSize
}

func (gs *GameSettings) SetWindowWidth(windowWidth int32) {
	gs.WindowWidth = windowWidth
}

func (gs *GameSettings) SetWindowHeight(windowHeight int32) {
	gs.WindowHeight = windowHeight
}

func (gs *GameSettings) SetFps(fps int32) {
	gs.Fps = fps
}

func (gs *GameSettings) SetShowFps(showFps bool) {
	gs.ShowFps = showFps
}
