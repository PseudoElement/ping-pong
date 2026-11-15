package settings

import (
	"github.com/pseudoelement/ping-pong/sdk/common"
	sdkconstants "github.com/pseudoelement/ping-pong/sdk/constants"

	"github.com/pseudoelement/ping-pong/src/constants"
	appinterfaces "github.com/pseudoelement/ping-pong/src/interfaces"
)

type GameSettings struct {
	fontSize     int32
	windowWidth  int32
	windowHeight int32
	fps          int32
	showFps      bool
	theme        sdkconstants.AppThemeName
	ballVelocity common.Coords
}

var _ appinterfaces.IGameSettings = (*GameSettings)(nil)

func NewGameSettings() *GameSettings {
	return &GameSettings{
		fontSize:     0,
		windowWidth:  0,
		windowHeight: 0,
		fps:          0,
		showFps:      true,
		theme:        constants.ThemeCosmos,
		ballVelocity: common.Coords{X: 6, Y: 6},
	}
}

func (gs *GameSettings) FontSize() int32 {
	return gs.fontSize
}

func (gs *GameSettings) SetFontSize(fontSize int32) {
	gs.fontSize = fontSize
}

func (gs *GameSettings) WindowWidth() int32 {
	return gs.windowWidth
}

func (gs *GameSettings) SetWindowWidth(windowWidth int32) {
	gs.windowWidth = windowWidth
}

func (gs *GameSettings) WindowHeight() int32 {
	return gs.windowHeight
}

func (gs *GameSettings) SetWindowHeight(windowHeight int32) {
	gs.windowHeight = windowHeight
}

func (gs *GameSettings) Fps() int32 {
	return gs.fps
}

func (gs *GameSettings) SetFps(fps int32) {
	gs.fps = fps
}

func (gs *GameSettings) IsShownFps() bool {
	return gs.showFps
}

func (gs *GameSettings) SetShowFps(showFps bool) {
	gs.showFps = showFps
}

func (gs *GameSettings) Theme() sdkconstants.AppTheme {
	return constants.APP_THEMES[gs.theme]
}

func (gs *GameSettings) ThemeName() sdkconstants.AppThemeName {
	return gs.theme
}

func (gs *GameSettings) SetTheme(theme sdkconstants.AppThemeName) {
	gs.theme = theme
}

func (gs *GameSettings) BallVelocity() common.Coords {
	return gs.ballVelocity
}

func (gs *GameSettings) ChangeBallVelocity() {
	velocity := common.Coords{
		X: gs.ballVelocity.X + 1,
		Y: gs.ballVelocity.Y + 1,
	}
	if velocity.X > 12 {
		velocity.X = 6
	}
	if velocity.Y > 12 {
		velocity.Y = 6
	}

	gs.ballVelocity = velocity
}
