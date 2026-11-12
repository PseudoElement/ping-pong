package models

import (
	"image/color"

	"github.com/google/uuid"
	"github.com/pseudoelement/ping-pong/sdk/common"
	"github.com/pseudoelement/ping-pong/sdk/constants"
	"github.com/pseudoelement/ping-pong/sdk/interfaces"
	"github.com/pseudoelement/ping-pong/sdk/utils"
	"github.com/pseudoelement/ping-pong/src/injector"
)

type BaseObject struct {
	injector *injector.Injector

	id       string
	coords   common.Coords
	styles   *common.ObjectStyles
	color    color.RGBA
	children []interfaces.IBaseObject
	kind     constants.ObjectKind
	zIndex   int16
}

var _ interfaces.IBaseObject = (*BaseObject)(nil)

type BaseObjectStructParams struct {
	Styles        *common.ObjectStyles
	InitialCoords common.Coords
	Kind          constants.ObjectKind
	Id            string
}

func NewBaseObject(p BaseObjectStructParams, injector *injector.Injector) *BaseObject {
	id := p.Id
	if id == "" {
		uid, _ := uuid.NewUUID()
		id = uid.String()
	}

	return &BaseObject{
		injector: injector,
		children: make([]interfaces.IBaseObject, 0),
		id:       id,
		coords:   p.InitialCoords,
		styles:   p.Styles,
	}
}

func (bo *BaseObject) Injector() interfaces.IInjector {
	return bo.injector
}

func (bo *BaseObject) Styles() *common.ObjectStyles {
	return bo.styles
}

func (bo *BaseObject) Kind() constants.ObjectKind {
	return bo.kind
}

func (bo *BaseObject) Coords() common.Coords {
	return bo.coords
}

func (bo *BaseObject) SetCoords(newCoords common.Coords) {
	prevCoords := bo.coords
	bo.coords = newCoords

	dX, dY := newCoords.X-prevCoords.X, newCoords.Y-prevCoords.Y
	utils.ForEachInObject(bo, func(child interfaces.IBaseObject) {
		child.SetCoords(common.Coords{
			X: child.Coords().X + dX,
			Y: child.Coords().Y + dY,
		})
	})
}

func (bo *BaseObject) Color() color.RGBA {
	return bo.color
}

func (bo *BaseObject) SetColor(newColor color.RGBA) {
	bo.color = newColor
}

func (bo *BaseObject) Id() string {
	return bo.id
}

func (bo *BaseObject) Children() []interfaces.IBaseObject {
	return bo.children
}

func (bo *BaseObject) AppendChild(child interfaces.IBaseObject) {
	bo.children = append(bo.children, child)
}

func (bo *BaseObject) GetChildByID(id string) (interfaces.IBaseObject, bool) {
	if id == bo.id {
		return bo, true
	}

	var childrenQueue []interfaces.IBaseObject = make([]interfaces.IBaseObject, 0)
	for _, child := range bo.children {
		childrenQueue = append(childrenQueue, child)
	}

	for len(childrenQueue) > 0 {
		firstChild := childrenQueue[0].(*BaseObject)
		childrenQueue = childrenQueue[1:]

		if firstChild.id == id {
			return firstChild, true
		}

		for _, child := range firstChild.children {
			childrenQueue = append(childrenQueue, child)
		}
	}

	return nil, false
}

func (bo *BaseObject) Move(dX, dY int32) {
	bo.coords.X += dX
	bo.coords.Y += dY

	utils.ForEachInObject(bo, func(child interfaces.IBaseObject) {
		child.SetCoords(common.Coords{
			X: child.Coords().X + dX,
			Y: child.Coords().Y + dY,
		})
	})
}
