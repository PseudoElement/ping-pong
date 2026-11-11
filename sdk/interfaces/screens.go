package interfaces

type IScreen interface {
	IDestroyable

	TextObjects() []ITextObject
	RectObjects() []IRectangleObject
	ImageObjects() []IImageObject
}
