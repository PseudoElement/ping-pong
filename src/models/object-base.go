package models

import (
	"image/color"

	"github.com/google/uuid"
	"github.com/pseudoelement/ping-pong/sdk/models"
	"github.com/pseudoelement/ping-pong/src/constants"
	"github.com/pseudoelement/ping-pong/src/injector"
	"github.com/pseudoelement/ping-pong/src/interfaces"
)

type BaseObject struct {
	injector *injector.Injector

	id       string
	coords   models.Coords
	color    color.RGBA
	children []interfaces.IBaseObject
	kind     constants.ObjectKind
}

var _ interfaces.IBaseObject = (*BaseObject)(nil)

type BaseObjectStructParams struct {
	Color         color.RGBA
	InitialCoords models.Coords
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
		color:    p.Color,
		kind:     p.Kind,
	}
}

func (bo *BaseObject) Injector() interfaces.IInjector {
	return bo.injector
}

func (bo *BaseObject) Kind() constants.ObjectKind {
	return bo.kind
}

func (bo *BaseObject) Coords() models.Coords {
	return bo.coords
}

func (bo *BaseObject) SetCoords(newCoords models.Coords) {
	prevCoords := bo.coords
	bo.coords = newCoords

	dX, dY := newCoords.X-prevCoords.X, newCoords.Y-prevCoords.Y
	bo._forEach(func(child *BaseObject) {
		child.coords.X += dX
		child.coords.Y += dY
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

	bo._forEach(func(child *BaseObject) {
		child.coords.X += dX
		child.coords.Y += dY
	})
}

func (bo *BaseObject) _forEach(callback func(child *BaseObject)) {
	childrenQueue := append([]interfaces.IBaseObject{}, bo.children...)
	for len(childrenQueue) > 0 {
		firstChild := childrenQueue[0].(*BaseObject)
		childrenQueue = childrenQueue[1:]

		callback(firstChild)
		childrenQueue = append(childrenQueue, firstChild.children...)
	}
}
