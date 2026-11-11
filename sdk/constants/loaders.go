package constants

/* ========================================SOUNDS======================================== */
type AppSoundName string

/**
 * AppSoundName should start from "../" then it normalized for build dirname
 */
const (
	SoundCoinEarn AppSoundName = "../assets/audios/coin-earn.mp3"
)

/* ========================================IMAGES======================================== */

type AppImageName string

// all image names
const (
	ImgCosmosBackground AppImageName = "Cosmos"
	ImgTenisBackground  AppImageName = "Tenis"
	ImgHockeyBackground AppImageName = "Hockey"
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
