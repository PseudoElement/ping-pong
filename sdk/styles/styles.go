package styles

import (
	"github.com/pseudoelement/ping-pong/sdk/common"
	"github.com/pseudoelement/ping-pong/sdk/interfaces"
)

func StyleTranslate[T interfaces.IObjectWithWidthAndHeight](
	obj T,
	translateX float32,
	translateY float32,
) {
	dX := float32(obj.Width()) * translateX
	dY := float32(obj.Height()) * translateY

	obj.SetCoords(common.Coords{
		X: obj.Coords().X + int32(dX),
		Y: obj.Coords().Y + int32(dY),
	})
}
