package loaders

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	sdkconstants "github.com/pseudoelement/ping-pong/sdk/constants"
	"github.com/pseudoelement/ping-pong/sdk/interfaces"
	"github.com/pseudoelement/ping-pong/sdk/utils"
	"github.com/pseudoelement/ping-pong/src/constants"
)

type FontLoader struct {
	fonts map[sdkconstants.AppFontName]rl.Font
}

var _ interfaces.IFontLoader = (*FontLoader)(nil)

func NewFontLoader() *FontLoader {
	loader := &FontLoader{
		fonts: make(map[sdkconstants.AppFontName]rl.Font, 1),
	}
	loader._load()
	return loader
}

func (fl *FontLoader) GetFont(fontName sdkconstants.AppFontName) rl.Font {
	return fl.fonts[fontName]
}

func (fl *FontLoader) _load() {
	for _, fontName := range constants.APP_FONTS {
		fontPath := utils.AppPath(fontName)
		font := rl.LoadFontEx(fontPath.Normalize(), 50, nil)
		fl.fonts[fontName] = font
	}
}
