package interfaces

type IDestroyable interface {
	/* Call this method every time when element removed  */
	OnDestroy()
}

type IHoverable interface {
	SetHover(bool)
	IsHovered() bool
}

type IUpdateable interface {
	Update()
}
