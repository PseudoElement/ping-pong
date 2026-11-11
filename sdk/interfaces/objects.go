package interfaces

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/pseudoelement/ping-pong/sdk/models"
	"github.com/pseudoelement/ping-pong/src/constants"
)

type IBaseObject interface {
	Coords() models.Coords
	SetCoords(coords models.Coords)

	Id() string
	Kind() constants.ObjectKind
	Styles() *models.ObjectStyles

	Move(dX, dY int32)
	AppendChild(child IBaseObject)
	GetChildByID(id string) (IBaseObject, bool)

	Injector() IInjector
	Children() []IBaseObject
}

type IObjectWithWidthAndHeight interface {
	IBaseObject

	Width() int32
	Height() int32
}

type ITextObject interface {
	IObjectWithWidthAndHeight
	IHoverable

	Value() string
	SetValue(val string)

	FontSize() int32
}

type IImageObject interface {
	IObjectWithWidthAndHeight

	Texture() rl.Texture2D
}

type IRectangleObject interface {
	IObjectWithWidthAndHeight

	SetWidth(width int32)
	SetHeight(height int32)
}
