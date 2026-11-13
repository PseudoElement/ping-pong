package constants

import (
	"cmp"
	"slices"

	"github.com/pseudoelement/ping-pong/sdk/constants"
)

const (
	ThemeTenis  constants.AppThemeName = "Tenis"
	ThemeCosmos constants.AppThemeName = "Cosmos"
	ThemeHockey constants.AppThemeName = "Hockey"
)

var APP_THEMES = map[constants.AppThemeName]constants.AppTheme{
	ThemeTenis:  {ImgMainBackground: ImgTenisBackground, ImgBall: ImgBallTenis},
	ThemeCosmos: {ImgMainBackground: ImgCosmosBackground, ImgBall: ImgBallCosmos},
	ThemeHockey: {ImgMainBackground: ImgHockeyBackground, ImgBall: ImgBallHockey},
}

func ThemeNamesList() []constants.AppThemeName {
	appThemes := make([]constants.AppThemeName, 0, len(APP_THEMES))
	for themeName, _ := range APP_THEMES {
		appThemes = append(appThemes, themeName)
	}
	slices.SortFunc(appThemes, func(a, b constants.AppThemeName) int {
		return cmp.Compare(a, b)
	})
	return appThemes
}
