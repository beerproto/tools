package unit

import (
	beerproto "github.com/beerproto/beerproto_go"
)

func DiastaticPower(value string, options ...OptionsFunc[beerproto.DiastaticPowerUnitType]) *beerproto.DiastaticPowerRangeType {
	rangeType := &RangeType[beerproto.DiastaticPowerUnitType, float64]{}

	options = append(options,
		WithMinContains[beerproto.DiastaticPowerUnitType]([]string{"min"}),
		WithMinTrim[beerproto.DiastaticPowerUnitType]([]string{"wk", "min"}),
		WithMaxTrim[beerproto.DiastaticPowerUnitType]([]string{"max"}),
		WithMaxContains[beerproto.DiastaticPowerUnitType]([]string{"max"}),
		WithUnit(beerproto.DiastaticPowerUnitType_WK),
	)

	parse(value, rangeType, options...)

	diastaticPower := &beerproto.DiastaticPowerRangeType{}

	if rangeType.Maximum != nil {
		diastaticPower.Maximum = &beerproto.DiastaticPowerType{
			Unit:  rangeType.Maximum.Unit,
			Value: rangeType.Maximum.Value,
		}
	}

	if rangeType.Minimum != nil {
		diastaticPower.Minimum = &beerproto.DiastaticPowerType{
			Unit:  rangeType.Minimum.Unit,
			Value: rangeType.Minimum.Value,
		}
	}

	return diastaticPower
}
