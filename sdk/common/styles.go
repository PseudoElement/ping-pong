package common

import (
	"image/color"
)

type ObjectStyles struct {
	Color        color.RGBA
	ZIndex       uint16
	BorderRadius float32
}

func (os *ObjectStyles) SetColor(newColor color.RGBA) {
	os.Color = newColor
}

func (os *ObjectStyles) SetZIndex(zIndex uint16) {
	os.ZIndex = zIndex
}

func (os *ObjectStyles) SetBorderRadius(radius float32) {
	os.BorderRadius = radius
}
