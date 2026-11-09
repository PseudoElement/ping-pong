package interfaces

type IScreen interface {
	IDestroyable
	TextObjects() []ITextObject
	RectObjects() []IBaseObject
}
