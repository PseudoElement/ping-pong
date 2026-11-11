package interfaces

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/pseudoelement/ping-pong/sdk/common"
	"github.com/pseudoelement/ping-pong/sdk/constants"
)

type IBaseObject interface {
	Coords() common.Coords
	SetCoords(coords common.Coords)

	Id() string
	Kind() constants.ObjectKind
	Styles() *common.ObjectStyles

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
