//go:build GM_USE_DEGREES

package gm

func FromDeg(deg Float) Angle   { return Angle(deg) }
func FromRad(rad Float) Angle   { return Angle(rad * RadToDeg) }
func FromTurn(turn Float) Angle { return Angle(tur * TurnToDeg) }

func (deg Angle) ToDeg() Float  { return Float(deg) }
func (deg Angle) ToRad() Float  { return Float(deg * DegToRad) }
func (deg Angle) ToTurn() Float { return Float(deg * DegToTurn) }
