package styles

import (
	"image/color"

	"github.com/pseudoelement/ping-pong/sdk/models"
)

type EmptyTextObject struct {
}

func (eto *EmptyTextObject) Coords() models.Coords {
	return models.Coords{}
}
func (eto *EmptyTextObject) FontSize() int32 {
	return 0
}
func (eto *EmptyTextObject) Value() string {
	return ""
}
func (eto *EmptyTextObject) SetCoords(models.Coords) {}
func (eto *EmptyTextObject) Width() int32 {
	return 0
}
func (eto *EmptyTextObject) Height() int32 {
	return eto.FontSize()
}
func (eto *EmptyTextObject) Color() color.RGBA {
	return color.RGBA{}
}

type ITextObject interface {
	Coords() models.Coords
	SetCoords(models.Coords)
	FontSize() int32
	Color() color.RGBA
	Width() int32
	Height() int32
	Value() string
}
