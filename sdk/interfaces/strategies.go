package interfaces

type IMovementStrategy interface {
	Setup()
	Stop()
	ChangeTarget(newTarget IObjectWithWidthAndHeight)
}
