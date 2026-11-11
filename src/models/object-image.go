package models

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/pseudoelement/ping-pong/sdk/interfaces"
	"github.com/pseudoelement/ping-pong/sdk/models"
	"github.com/pseudoelement/ping-pong/src/constants"
	"github.com/pseudoelement/ping-pong/src/injector"
)

type ImageObjectStructParams struct {
	Styles        *models.ObjectStyles
	InitialCoords models.Coords
	Id            string
	Texture       rl.Texture2D
}

var _ interfaces.IImageObject = (*ImageObject)(nil)

type ImageObject struct {
	*BaseObject
	texture rl.Texture2D
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
		texture: p.Texture,
	}
}

func (to *ImageObject) Texture() rl.Texture2D {
	return to.texture
}

func (to *ImageObject) Width() int32 {
	return to.texture.Width
}

func (to *ImageObject) Height() int32 {
	return to.texture.Height
}
