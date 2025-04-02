package gm

import (
	"math"
)

type Angle Float

const (
	Deg180   = 180.0
	TurnHalf = 0.5

	RadToDeg  = Deg180 / math.Pi
	RadToTurn = TurnHalf / math.Pi
	DegToRad  = math.Pi / Deg180
	DegToTurn = TurnHalf / Deg180
	TurnToRad = math.Pi / TurnHalf
	TurnToDeg = Deg180 / TurnHalf
)

func (a Angle) Mod(b Angle) Angle {
	return Angle(Mod(Float(a), Float(b)))
}

func (a Angle) Sin() Angle {
	return FromRad(Sin(a.ToRad()))
}

func (a Angle) Cos() Angle {
	return FromRad(Cos(a.ToRad()))
}

func (a Angle) Tan() Angle {
	return FromRad(Tan(a.ToRad()))
}

func (a Angle) Asin() Angle {
	return FromRad(Asin(a.ToRad()))
}

func (a Angle) Acos() Angle {
	return FromRad(Acos(a.ToRad()))
}

func (a Angle) Atan() Angle {
	return FromRad(Atan(a.ToRad()))
}
