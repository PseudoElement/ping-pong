package models

import (
	"github.com/pseudoelement/ping-pong/sdk/interfaces"
	sdkmodels "github.com/pseudoelement/ping-pong/sdk/models"
	"github.com/pseudoelement/ping-pong/src/constants"
	"github.com/pseudoelement/ping-pong/src/injector"
)

type RectangleObject struct {
	*BaseObject

	injector *injector.Injector
	width    int32
	height   int32
}

var _ interfaces.IRectangleObject = (*RectangleObject)(nil)

type RectangleObjectStructParams struct {
	Width, Height int32
	InitialCoords sdkmodels.Coords
	Styles        *sdkmodels.ObjectStyles
	Id            string
}

func NewRectangleObject(p RectangleObjectStructParams, injector *injector.Injector) *RectangleObject {
	return &RectangleObject{
		BaseObject: NewBaseObject(
			BaseObjectStructParams{
				Styles:        p.Styles,
				InitialCoords: p.InitialCoords,
				Kind:          constants.KindRectangle,
				Id:            p.Id,
			},
			injector,
		),
		injector: injector,
		width:    p.Width,
		height:   p.Height,
	}
}

func (ro *RectangleObject) Width() int32 {
	return ro.width
}

func (ro *RectangleObject) SetWidth(width int32) {
	ro.width = width
}

func (ro *RectangleObject) Height() int32 {
	return ro.height
}

func (ro *RectangleObject) SetHeight(height int32) {
	ro.height = height
}
