package ui

import (
	"fmt"

	"github.com/pseudoelement/ping-pong/src/injector"
	"github.com/pseudoelement/ping-pong/src/interfaces"
	"github.com/pseudoelement/ping-pong/src/models"
)

type FpsObject struct {
	*models.TextObject
}

var _ interfaces.ITextObject = (*FpsObject)(nil)
var _ interfaces.IUpdateable = (*FpsObject)(nil)

func NewFpsObject(p models.TextObjectStructParams, injector *injector.Injector) *FpsObject {
	return &FpsObject{
		TextObject: models.NewTextObject(p, injector),
	}
}

func (fo *FpsObject) Update() {
	settings := fo.Injector().GameSettings()
	value := fmt.Sprintf("%d FPS", settings.Fps)
	fo.SetValue(value)
}
