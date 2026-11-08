package interfaces

type IScreen interface {
	TextObjects() []ITextObject
	RectObjects() []IBaseObject
}
