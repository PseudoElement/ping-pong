package loaders

import (
	"github.com/pseudoelement/ping-pong/sdk/utils"
)

/* ========================================SOUNDS======================================== */
type AppSoundName string

/**
 * AppSoundName should start from "../" then it normalized for build dirname
 */
const (
	CoinEarn AppSoundName = "../assets/audios/coin-earn.mp3"
)

var app_sounds []AppSoundName = []AppSoundName{
	CoinEarn,
}

/* ========================================IMAGES======================================== */

type AppImageName string

// all image names
const (
	CosmosBackground AppImageName = "Cosmos"
	TenisBackground  AppImageName = "Tenis"
	HockeyBackground AppImageName = "Hockey"
)

type AppImage struct {
	/**
	 * Href should start from "../" then it normalized for build dirname
	 */
	Href          string
	Width, Height int32
	Name          AppImageName
}

type ImageFactory func(windowWidth, windowHeight int32) AppImage

func _imageCosmosBackground(windowWidth, windowHeight int32) AppImage {
	return AppImage{
		Href:   "../assets/images/cosmos-bg.png",
		Width:  windowWidth,
		Height: windowHeight,
		Name:   CosmosBackground,
	}
}

func _imageTenisBackground(windowWidth, windowHeight int32) AppImage {
	return AppImage{
		Href:   "../assets/images/tenis-bg.png",
		Width:  windowWidth,
		Height: windowHeight,
		Name:   TenisBackground,
	}
}

func _imageHockeyBackground(windowWidth, windowHeight int32) AppImage {
	return AppImage{
		Href:   "../assets/images/hockey-bg.png",
		Width:  windowWidth,
		Height: windowHeight,
		Name:   HockeyBackground,
	}
}

var app_background_factories []ImageFactory = []ImageFactory{
	_imageCosmosBackground,
	_imageTenisBackground,
	_imageHockeyBackground,
}

var app_image_factories []ImageFactory = utils.MergeSlices([][]ImageFactory{app_background_factories})
