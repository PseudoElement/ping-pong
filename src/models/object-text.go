package models

import (
	"image/color"

	"github.com/pseudoelement/ping-pong/sdk/models"
)

type ITextObject interface {
	IBaseObject

	Value() string
	SetValue(string)

	FontSize() int32
	Width() int32 // equals int32((float64(object.FontSize()) / 1.75)) * len(object.Value())
	Height() int32
}

type TextObjectStructParams struct {
	Id            string
	InitialValue  string
	FontSize      int32
	Color         color.RGBA
	InitialCoords models.Coords
}

var _ ITextObject = (*TextObject)(nil)

type TextObject struct {
	*BaseObject
	value    string
	fontSize int32
}

func NewTextObject(p TextObjectStructParams) *TextObject {
	return &TextObject{
		BaseObject: NewBaseObject(BaseObjectStructParams{
			Id:            p.Id,
			Color:         p.Color,
			InitialCoords: p.InitialCoords,
			Kind:          KindText,
		}),
		value:    p.InitialValue,
		fontSize: p.FontSize,
	}
}

func (to *TextObject) Value() string {
	return to.value
}

func (to *TextObject) SetValue(val string) {
	to.value = val
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
