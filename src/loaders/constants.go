package loaders

type AppSound string

/**
 * AppSound should start from "../" then it normalized for build dirname
 */
const (
	CoinEarn AppSound = "../assets/audios/coin-earn.mp3"
)

var ALL_SOUNDS []AppSound = []AppSound{
	CoinEarn,
}

type AppImage string

/**
 * AppSound should start from "../" then it normalized for build dirname
 */
const (
	SkyBackground AppImage = "../assets/audios/coin-earn.mp3"
)

var ALL_IMAGES []AppImage = []AppImage{
	SkyBackground,
}
