//go:build MG_USE_TURNS || !(MG_USE_RADIANS || MG_USE_DEGREES)

package mg

func FromTurn(turn Float) Angle { return Angle(turn) }
func FromRad(rad Float) Angle   { return Angle(rad * RadToTurn) }
func FromDeg(deg Float) Angle   { return Angle(deg * DegToTurn) }

func (turn Angle) ToTurn() Float { return Float(turn) }
func (turn Angle) ToRad() Float  { return Float(turn * TurnToRad) }
func (turn Angle) ToDeg() Float  { return Float(turn * TurnToDeg) }
