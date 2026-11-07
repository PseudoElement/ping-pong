package utils

const (
	INVISIBLE_HEIGHT int = 70
)

func GetRealY(offsetY int) int {
	// @TODO check if INVISIBLE_HEIGHT needed for windows and ubuntu
	return offsetY - INVISIBLE_HEIGHT
}

func GetRealX(offsetX int) int {
	return offsetX
}
