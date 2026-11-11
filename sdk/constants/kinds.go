package constants

type ObjectKind int

/* object kinds */
const (
	KindText ObjectKind = iota
	KindRectangle
	KindImage
	KindTriangle
)
