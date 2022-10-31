package unit

type RangeType[TUnit Unit, TValue Value] struct {
	Minimum *UnitType[TUnit, TValue]
	Maximum *UnitType[TUnit, TValue]
}

type UnitType[TUnit Unit, TValue Value] struct {
	Value TValue
	Unit  TUnit
}

type Unit interface {
	~int32
}

type Value interface {
	~float64 | ~int64
}
