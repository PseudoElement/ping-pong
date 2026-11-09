package interfaces

import (
	"image/color"

	"github.com/pseudoelement/ping-pong/sdk/models"
	"github.com/pseudoelement/ping-pong/src/constants"
)

type IBaseObject interface {
	Coords() models.Coords
	SetCoords(models.Coords)

	Color() color.RGBA
	SetColor(color.RGBA)

	Injector() IInjector
	Id() string
	Kind() constants.ObjectKind
	Move(dX, dY int32)
	AppendChild(IBaseObject)
	GetChildByID(id string) (IBaseObject, bool)
	Children() []IBaseObject
}

type ITextObject interface {
	IBaseObject
	IHoverable

	Value() string
	SetValue(string)

	FontSize() int32
	Width() int32
	Height() int32
}
