package loaders

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	sdkconstants "github.com/pseudoelement/ping-pong/sdk/constants"
	"github.com/pseudoelement/ping-pong/sdk/utils"
	"github.com/pseudoelement/ping-pong/src/constants"
	appinterfaces "github.com/pseudoelement/ping-pong/src/interfaces"
)

type ImageLoader struct {
	windowWidth  int32
	windowHeight int32
	textures     map[sdkconstants.AppImageName]rl.Texture2D
}

var _ appinterfaces.IImageLoader = (*ImageLoader)(nil)

func NewImageLoader(windowWidth, windowHeight int32) *ImageLoader {
	loader := &ImageLoader{
		textures:     make(map[sdkconstants.AppImageName]rl.Texture2D),
		windowHeight: windowHeight,
		windowWidth:  windowWidth,
	}
	return loader
}

func (il *ImageLoader) GetTexture(imgName sdkconstants.AppImageName) rl.Texture2D {
	return il.textures[imgName]
}

func (il *ImageLoader) GetAppBackground(bgName sdkconstants.AppImageName) (sdkconstants.AppImage, bool) {
	for _, bgFactory := range constants.APP_BACKGROUND_FACTORIES {
		bg := bgFactory(il.windowWidth, il.windowHeight)
		if bg.Name == bgName {
			return bg, true
		}
	}
	return sdkconstants.AppImage{}, false
}

func (il *ImageLoader) Load() {
	for _, imgFactory := range constants.APP_IMAGE_FACTORIES {
		img := imgFactory(il.windowWidth, il.windowHeight)
		imgPath := utils.AppPath(img.Href)

		rlImg := rl.LoadImage(imgPath.Normalize())
		rl.ImageResize(rlImg, img.Width, img.Height)
		rl.ImageRotate(rlImg, img.RotationDeg)

		texture := rl.LoadTextureFromImage(rlImg)
		rl.UnloadImage(rlImg)

		il.textures[img.Name] = texture
	}
}
