package models

import (
	"image/color"

	"github.com/pseudoelement/ping-pong/sdk/models"
)

type IBaseObject interface {
	Coords() models.Coords
	SetCoords(models.Coords)

	Color() color.RGBA
	SetColor(color.RGBA)

	Kind() ObjectKind
	Move(dX, dY int32)
	GetChildByID(id string) (IBaseObject, bool)
}

type BaseObject struct {
	id       string
	coords   models.Coords
	color    color.RGBA
	children []*BaseObject
	kind     ObjectKind
}

var _ IBaseObject = (*BaseObject)(nil)

type BaseObjectStructParams struct {
	Id            string
	Color         color.RGBA
	InitialCoords models.Coords
	Kind          ObjectKind
}

func NewBaseObject(p BaseObjectStructParams) *BaseObject {
	return &BaseObject{
		children: make([]*BaseObject, 0),
		id:       p.Id,
		coords:   p.InitialCoords,
		color:    p.Color,
	}
}

func (bo *BaseObject) Kind() ObjectKind {
	return bo.kind
}

func (bo *BaseObject) Coords() models.Coords {
	return bo.coords
}

func (bo *BaseObject) SetCoords(coords models.Coords) {
	bo.coords = coords
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

func (bo *BaseObject) AppendChild(child *BaseObject) {
	bo.children = append(bo.children, child)
}

func (bo *BaseObject) GetChildByID(id string) (IBaseObject, bool) {
	if id == bo.id {
		return bo, true
	}

	childrenQueue := append([]*BaseObject{}, bo.children...)
	for len(childrenQueue) > 0 {
		firstChild := childrenQueue[0]
		childrenQueue = childrenQueue[1:]

		if firstChild.id == id {
			return firstChild, true
		}

		childrenQueue = append(childrenQueue, firstChild.children...)
	}

	return nil, false
}

func (bo *BaseObject) Move(dX, dY int32) {
	bo.coords.X += dX
	bo.coords.Y += dY

	childrenQueue := append([]*BaseObject{}, bo.children...)
	for len(childrenQueue) > 0 {
		firstChild := childrenQueue[0]
		childrenQueue = childrenQueue[1:]

		firstChild.coords.X += dX
		firstChild.coords.Y += dY

		childrenQueue = append(childrenQueue, firstChild.children...)
	}
}
