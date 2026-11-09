package models

import (
	"image/color"

	"github.com/pseudoelement/ping-pong/sdk/models"
	"github.com/pseudoelement/ping-pong/src/constants"
	"github.com/pseudoelement/ping-pong/src/injector"
	"github.com/pseudoelement/ping-pong/src/interfaces"
)

type TextObjectStructParams struct {
	InitialValue  string
	FontSize      int32
	Color         color.RGBA
	InitialCoords models.Coords
	Id            string
}

var _ interfaces.ITextObject = (*TextObject)(nil)

type TextObject struct {
	*BaseObject
	value    string
	fontSize int32
	hovered  bool
}

func NewTextObject(p TextObjectStructParams, injector *injector.Injector) *TextObject {
	return &TextObject{
		BaseObject: NewBaseObject(
			BaseObjectStructParams{
				Color:         p.Color,
				InitialCoords: p.InitialCoords,
				Kind:          constants.KindText,
				Id:            p.Id,
			},
			injector,
		),
		value:    p.InitialValue,
		fontSize: p.FontSize,
		hovered:  false,
	}
}

func (to *TextObject) Value() string {
	return to.value
}

func (to *TextObject) SetValue(val string) {
	to.value = val
}

func (to *TextObject) IsHovered() bool {
	return to.hovered
}

func (to *TextObject) SetHover(hovered bool) {
	to.hovered = hovered
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
