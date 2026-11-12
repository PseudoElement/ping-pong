package constants

type AppSoundName string

type AppImageName string

type ImageFactory func(windowWidth, windowHeight int32) AppImage

type AppImage struct {
	/**
	 * Href should start from "../" then it normalized for build dirname
	 */
	Href          string
	Width, Height int32
	Name          AppImageName
	RotationDeg   int32
}
