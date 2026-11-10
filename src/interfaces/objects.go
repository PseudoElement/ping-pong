package interfaces

import (
	"github.com/pseudoelement/ping-pong/sdk/models"
	"github.com/pseudoelement/ping-pong/sdk/styles"
	"github.com/pseudoelement/ping-pong/src/constants"
)

type IBaseObject interface {
	Coords() models.Coords
	SetCoords(models.Coords)

	Id() string
	Kind() constants.ObjectKind
	Styles() *styles.ObjectStyles

	Move(dX, dY int32)
	AppendChild(IBaseObject)
	GetChildByID(id string) (IBaseObject, bool)

	Injector() IInjector
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

type IImageObject interface {
	IBaseObject

	Width() int32
	Height() int32
}
