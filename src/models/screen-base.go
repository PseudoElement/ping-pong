package models

import (
	"github.com/pseudoelement/ping-pong/sdk/common"
	"github.com/pseudoelement/ping-pong/sdk/interfaces"
	"github.com/pseudoelement/ping-pong/sdk/utils"
	"github.com/pseudoelement/ping-pong/src/injector"
)

type ScreenBase struct {
	injector        *injector.Injector
	textObjectsPool []interfaces.ITextObject
	imgObjectsPool  []interfaces.IImageObject
	rectObjectsPool []interfaces.IRectangleObject
	unsubs          []common.UnsubscribeFn
}

var _ interfaces.IScreenBase = (*ScreenBase)(nil)

func NewScreenBase(injector *injector.Injector) *ScreenBase {
	return &ScreenBase{
		injector:        injector,
		textObjectsPool: make([]interfaces.ITextObject, 0),
		imgObjectsPool:  make([]interfaces.IImageObject, 0),
		unsubs:          make([]common.UnsubscribeFn, 0),
	}
}

func (sg *ScreenBase) TextObjects() []interfaces.ITextObject {
	return sg.textObjectsPool
}

func (sg *ScreenBase) RectObjects() []interfaces.IRectangleObject {
	return sg.rectObjectsPool
}

func (sg *ScreenBase) ImageObjects() []interfaces.IImageObject {
	return sg.imgObjectsPool
}

func (sg *ScreenBase) appendTextObjects(objects ...interfaces.ITextObject) {
	sg.textObjectsPool = append(sg.textObjectsPool, objects...)
}

func (sg *ScreenBase) appendImageObjects(objects ...interfaces.IImageObject) {
	sg.imgObjectsPool = append(sg.imgObjectsPool, objects...)
}

func (sg *ScreenBase) appendRectObjects(objects ...interfaces.IRectangleObject) {
	sg.rectObjectsPool = append(sg.rectObjectsPool, objects...)
}

func (sg *ScreenBase) appendUnsubs(unsubFns ...common.UnsubscribeFn) {
	sg.unsubs = append(sg.unsubs, unsubFns...)
}

/*
 * NOTE: this method should be called at OnDestroy() in all IScreen instances
 */
func (sg *ScreenBase) onDestroyBase() {
	for _, unsubFn := range sg.unsubs {
		unsubFn()
	}

	utils.ForEachInScreen(sg, func(child interfaces.IBaseObject) {
		destroyableObj, ok := child.(interfaces.IDestroyable)
		if ok {
			destroyableObj.OnDestroy()
		}
	})
}
