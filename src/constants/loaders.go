package constants

import (
	"github.com/pseudoelement/ping-pong/sdk/common"
	"github.com/pseudoelement/ping-pong/sdk/constants"
)

/* ========================================SOUNDS======================================== */

/**
 * AppSoundName should start from "../" then it normalized for build dirname
 */
const (
	SoundCoinEarn constants.AppSoundName = "../assets/audios/coin-earn.mp3"
)

var APP_SOUNDS []constants.AppSoundName = []constants.AppSoundName{
	SoundCoinEarn,
}

/* ========================================IMAGES======================================== */

// all image names
const (
	ImgCosmosBackground constants.AppImageName = "Cosmos"
	ImgTenisBackground  constants.AppImageName = "Tenis"
	ImgHockeyBackground constants.AppImageName = "Hockey"

	ImgPlayerRocketRight constants.AppImageName = "PlayerRocketRight"
	ImgPlayerRocketLeft  constants.AppImageName = "PlayerRocketLeft"
	ImgBallBaseball      constants.AppImageName = "ImgBallBaseball"
)

func _imageCosmosBackground(windowWidth, windowHeight int32) constants.AppImage {
	return constants.AppImage{
		Href:   "../assets/images/cosmos-bg.png",
		Width:  windowWidth,
		Height: windowHeight,
		Name:   ImgCosmosBackground,
	}
}

func _imageTenisBackground(windowWidth, windowHeight int32) constants.AppImage {
	return constants.AppImage{
		Href:   "../assets/images/tenis-bg.png",
		Width:  windowWidth,
		Height: windowHeight,
		Name:   ImgTenisBackground,
	}
}

func _imageHockeyBackground(windowWidth, windowHeight int32) constants.AppImage {
	return constants.AppImage{
		Href:   "../assets/images/hockey-bg.png",
		Width:  windowWidth,
		Height: windowHeight,
		Name:   ImgHockeyBackground,
	}
}

func _imagePlayerRocketRight(windowWidth, windowHeight int32) constants.AppImage {
	return constants.AppImage{
		Href:        "../assets/images/player-rocket.png",
		Width:       100,
		Height:      250,
		RotationDeg: 180,
		Name:        ImgPlayerRocketRight,
	}
}

func _imagePlayerRocketLeft(windowWidth, windowHeight int32) constants.AppImage {
	return constants.AppImage{
		Href:   "../assets/images/player-rocket.png",
		Width:  100,
		Height: 250,
		Name:   ImgPlayerRocketLeft,
	}
}

func _imageBaseballBall(windowWidth, windowHeight int32) constants.AppImage {
	return constants.AppImage{
		Href:   "../assets/images/ball-baseball.png",
		Width:  90,
		Height: 90,
		Name:   ImgBallBaseball,
	}
}

var APP_BACKGROUND_FACTORIES []constants.ImageFactory = []constants.ImageFactory{
	_imageCosmosBackground,
	_imageTenisBackground,
	_imageHockeyBackground,
}

var APP_IMAGE_FACTORIES []constants.ImageFactory = common.MergeSlices(
	[][]constants.ImageFactory{
		APP_BACKGROUND_FACTORIES,
		{
			_imagePlayerRocketRight,
			_imagePlayerRocketLeft,
			_imageBaseballBall,
		},
	},
)
