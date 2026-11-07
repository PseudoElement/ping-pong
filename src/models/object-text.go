package models

import (
	"image/color"

	"github.com/pseudoelement/ping-pong/sdk/models"
)

type ITextObject interface {
	Coords() models.Coords
	SetCoords(models.Coords)

	Value() string
	SetValue(string)

	Color() color.RGBA
	SetColor(color.RGBA)

	FontSize() int32
	Width() int32 // equals int32((float64(object.FontSize()) / 1.75)) * len(object.Value())
	Height() int32
}

type TextObjectStructParams struct {
	InitialValue  string
	FontSize      int32
	Color         color.RGBA
	InitialCoords models.Coords
}

var _ ITextObject = (*TextObject)(nil)

type TextObject struct {
	value    string
	fontSize int32
	color    color.RGBA
	coords   models.Coords
}

func NewTextObject(p TextObjectStructParams) *TextObject {
	return &TextObject{
		value:    p.InitialValue,
		fontSize: p.FontSize,
		color:    p.Color,
		coords:   p.InitialCoords,
	}
}

func (to *TextObject) Coords() models.Coords {
	return to.coords
}

func (to *TextObject) SetCoords(coords models.Coords) {
	to.coords = coords
}

func (to *TextObject) Value() string {
	return to.value
}

func (to *TextObject) SetValue(val string) {
	to.value = val
}

func (to *TextObject) Color() color.RGBA {
	return to.color
}

func (to *TextObject) SetColor(newColor color.RGBA) {
	to.color = newColor
}

func (to *TextObject) FontSize() int32 {
	return to.fontSize
}

func (to *TextObject) Width() int32 {
	return int32((float64(to.FontSize()) / 1.75)) * int32(len(to.Value()))
}

func (to *TextObject) Height() int32 {
	return to.FontSize()
}
