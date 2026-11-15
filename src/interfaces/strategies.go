package appinterfaces

import (
	sdkinterfaces "github.com/pseudoelement/ping-pong/sdk/interfaces"
)

type IMovementStrategy interface {
	Setup()
	Stop()
	ChangeTarget(newTarget sdkinterfaces.IObjectWithWidthAndHeight)
}
