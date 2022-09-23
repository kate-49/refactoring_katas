package parrot

import (
	"math"
)

type parrotType int

const (
	TypeEuropean      parrotType = 1
	TypeAfrican       parrotType = 2
	TypeNorwegianBlue parrotType = 3
)

// Parrot has a Speed.
type Parrot interface {
	Speed() (float64, error)
}

type europeanParrot struct {
}

type africanParrot struct {
	numberOfCoconuts int
}

type norwegianBlueParrot struct {
	voltage float64
	nailed  bool
}

func CreateParrot(t parrotType, numberOfCoconuts int, voltage float64, nailed bool) Parrot {
	switch t {
	case 1:
		return europeanParrot{}
	case 2:
		return africanParrot{numberOfCoconuts}
	case 3:
		return norwegianBlueParrot{voltage, nailed}
	}
	return nil
}

func (parrot europeanParrot) Speed() (float64, error) {
	return baseSpeed(), nil
}

func (parrot africanParrot) Speed() (float64, error) {
	return math.Max(0, baseSpeed()-parrot.loadFactor()*float64(parrot.numberOfCoconuts)), nil
}

func (parrot norwegianBlueParrot) Speed() (float64, error) {
	if parrot.nailed {
		return 0, nil
	}
	return parrot.computeBaseSpeedForVoltage(parrot.voltage), nil
}

func (parrot norwegianBlueParrot) computeBaseSpeedForVoltage(voltage float64) float64 {
	return math.Min(24.0, voltage*baseSpeed())
}

func (parrot africanParrot) loadFactor() float64 {
	return 9.0
}

func baseSpeed() float64 {
	return 12.0
}
