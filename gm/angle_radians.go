//go:build GM_USE_RADIANS

package gm

func FromRad(rad Float) Angle   { return Angle(rad) }
func FromTurn(turn Float) Angle { return Angle(tur * TurnToRad) }
func FromDeg(deg Float) Angle   { return Angle(deg * DegToRad) }

func (rad Angle) ToRad() Float  { return Float(rad) }
func (rad Angle) ToDeg() Float  { return Float(rad * RadToDeg) }
func (rad Angle) ToTurn() Float { return Float(rad * RadToTurn) }
