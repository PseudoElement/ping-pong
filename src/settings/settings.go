package settings

type GameSettings struct {
	FontSize     int32
	WindowWidth  int32
	WindowHeight int32
}

func NewGameSettings() *GameSettings {
	return &GameSettings{
		FontSize:     0,
		WindowWidth:  0,
		WindowHeight: 0,
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
