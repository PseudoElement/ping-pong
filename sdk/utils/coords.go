package utils

import "github.com/pseudoelement/ping-pong/sdk/models"

func IsTextObjectHovered(mouseCoords, objectCoords models.Coords, width int32, height int32) bool {
	objCoordLeft, objCoordRight := objectCoords.X, objectCoords.X+width
	objCoordTop, objCoordBottom := objectCoords.Y, objectCoords.Y+height
	if (mouseCoords.X >= objCoordLeft && mouseCoords.X <= objCoordRight) &&
		(mouseCoords.Y >= objCoordTop && mouseCoords.Y <= objCoordBottom) {
		return true
	}
	return false
}
