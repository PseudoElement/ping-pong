package loaders

import (
	"github.com/pseudoelement/ping-pong/sdk/constants"
	"github.com/pseudoelement/ping-pong/sdk/utils"
)

/* ========================================SOUNDS======================================== */

var app_sounds []constants.AppSoundName = []constants.AppSoundName{
	constants.SoundCoinEarn,
}

/* ========================================IMAGES======================================== */

func _imageCosmosBackground(windowWidth, windowHeight int32) constants.AppImage {
	return constants.AppImage{
		Href:   "../assets/images/cosmos-bg.png",
		Width:  windowWidth,
		Height: windowHeight,
		Name:   constants.ImgCosmosBackground,
	}
}

func _imageTenisBackground(windowWidth, windowHeight int32) constants.AppImage {
	return constants.AppImage{
		Href:   "../assets/images/tenis-bg.png",
		Width:  windowWidth,
		Height: windowHeight,
		Name:   constants.ImgTenisBackground,
	}
}

func _imageHockeyBackground(windowWidth, windowHeight int32) constants.AppImage {
	return constants.AppImage{
		Href:   "../assets/images/hockey-bg.png",
		Width:  windowWidth,
		Height: windowHeight,
		Name:   constants.ImgHockeyBackground,
	}
}

var app_background_factories []constants.ImageFactory = []constants.ImageFactory{
	_imageCosmosBackground,
	_imageTenisBackground,
	_imageHockeyBackground,
}

var app_image_factories []constants.ImageFactory = utils.MergeSlices(
	[][]constants.ImageFactory{app_background_factories},
)
