package interfaces

type IScreenBase interface {
	TextObjects() []ITextObject
	RectObjects() []IRectangleObject
	ImageObjects() []IImageObject
}

type IScreen interface {
	IScreenBase
	IDestroyable
}
