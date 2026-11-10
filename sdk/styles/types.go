package styles

import (
	"image/color"
)

type ObjectStyles struct {
	Color  color.RGBA
	ZIndex int16
}

func (os *ObjectStyles) SetColor(newColor color.RGBA) {
	os.Color = newColor
}

func (os *ObjectStyles) SetZIndex(zIndex int16) {
	os.ZIndex = zIndex
}
