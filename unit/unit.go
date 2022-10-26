package unit

import (
	beerproto "github.com/beerproto/beerproto_go"
	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Integer | constraints.Float
}

type RangeType[TUnit Unit, TNumber Number] struct {
	Minimum UnitType[TUnit, TNumber]
	Maximum UnitType[TUnit, TNumber]
}

type UnitType[TUnit Unit, TNumber Number] struct {
	Value TNumber
	Unit  TUnit
}

type Unit interface {
	beerproto.DiastaticPowerUnitType | beerproto.ColorUnitType
}
