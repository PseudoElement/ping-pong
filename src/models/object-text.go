package models

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/pseudoelement/ping-pong/sdk/common"
	"github.com/pseudoelement/ping-pong/sdk/constants"
	"github.com/pseudoelement/ping-pong/sdk/interfaces"
	"github.com/pseudoelement/ping-pong/src/injector"
)

type TextObjectStructParams struct {
	Styles        *common.ObjectStyles
	InitialValue  string
	FontSize      int32
	InitialCoords common.Coords
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
				Styles:        p.Styles,
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
	return rl.MeasureText(to.value, to.fontSize)
	// return int32((float64(to.FontSize()) / 1.75)) * int32(len(to.Value()))
}

func (to *TextObject) Height() int32 {
	return to.FontSize()
}
