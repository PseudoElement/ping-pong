package constants

type MovementStrategyName int

const (
	MovementStrategyArrow = iota
	MovementStrategyWasd
	MovementStrategyBallMirroring
	MovementStrategyTopBottom
)
