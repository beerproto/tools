package unit

import (
	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Integer | constraints.Float
}

type RangeType[TUnit Unit] struct {
	Minimum UnitType[TUnit]
	Maximum UnitType[TUnit]
}

type UnitType[TUnit Unit] struct {
	Value float64
	Unit  TUnit
}

type Unit interface {
	~int32
}
