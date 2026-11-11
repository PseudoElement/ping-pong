package loaders

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/pseudoelement/ping-pong/sdk/constants"
	"github.com/pseudoelement/ping-pong/sdk/utils"
)

type ImageLoader struct {
	windowWidth  int32
	windowHeight int32
	textures     map[constants.AppImageName]rl.Texture2D
}

func NewImageLoader(windowWidth, windowHeight int32) *ImageLoader {
	loader := &ImageLoader{
		textures:     make(map[constants.AppImageName]rl.Texture2D),
		windowHeight: windowHeight,
		windowWidth:  windowWidth,
	}
	loader._load()
	return loader
}

func (il *ImageLoader) GetTexture(imgName constants.AppImageName) rl.Texture2D {
	return il.textures[imgName]
}

func (il *ImageLoader) GetAppBackgrounds() []constants.AppImage {
	bgs := make([]constants.AppImage, len(app_background_factories))
	for idx, bgFactory := range app_background_factories {
		bgs[idx] = bgFactory(il.windowWidth, il.windowHeight)
	}
	return bgs
}

func (il *ImageLoader) GetAppBackground(bgName constants.AppImageName) (constants.AppImage, bool) {
	for _, bgFactory := range app_background_factories {
		bg := bgFactory(il.windowWidth, il.windowHeight)
		if bg.Name == bgName {
			return bg, true
		}
	}
	return constants.AppImage{}, false
}

func (il *ImageLoader) _load() {
	for _, imgFactory := range app_image_factories {
		img := imgFactory(il.windowWidth, il.windowHeight)
		imgPath := utils.AppPath(img.Href)

		rlImg := rl.LoadImage(imgPath.Normalize())
		rl.ImageResize(rlImg, img.Width, img.Height)
		texture := rl.LoadTextureFromImage(rlImg)
		rl.UnloadImage(rlImg)

		il.textures[img.Name] = texture
	}
}
