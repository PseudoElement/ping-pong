package models

import (
	"fmt"

	"github.com/pseudoelement/ping-pong/sdk/interfaces"
	"github.com/pseudoelement/ping-pong/src/injector"
)

type FpsObject struct {
	*TextObject
}

var _ interfaces.ITextObject = (*FpsObject)(nil)
var _ interfaces.IUpdateable = (*FpsObject)(nil)

func NewFpsObject(p TextObjectStructParams, injector *injector.Injector) *FpsObject {
	return &FpsObject{
		TextObject: NewTextObject(p, injector),
	}
}

func (fo *FpsObject) Update() {
	settings := fo.Injector().GameSettings()
	value := fmt.Sprintf("%d FPS", settings.Fps())
	fo.SetValue(value)
}
