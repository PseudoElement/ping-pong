package models

import (
	"github.com/pseudoelement/ping-pong/sdk/models"
	"github.com/pseudoelement/ping-pong/sdk/styles"
	"github.com/pseudoelement/ping-pong/src/constants"
	"github.com/pseudoelement/ping-pong/src/injector"
	"github.com/pseudoelement/ping-pong/src/interfaces"
)

type ImageObjectStructParams struct {
	Styles        *styles.ObjectStyles
	InitialCoords models.Coords
	Id            string
}

var _ interfaces.IImageObject = (*ImageObject)(nil)

type ImageObject struct {
	*BaseObject
}

func NewImageObject(p ImageObjectStructParams, injector *injector.Injector) *ImageObject {
	return &ImageObject{
		BaseObject: NewBaseObject(
			BaseObjectStructParams{
				Styles:        p.Styles,
				InitialCoords: p.InitialCoords,
				Kind:          constants.KindText,
				Id:            p.Id,
			},
			injector,
		),
	}
}

func (to *ImageObject) Width() int32 {
	return 100
}

func (to *ImageObject) Height() int32 {
	return 100
}
