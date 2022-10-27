package unit

import (
	beerproto "github.com/beerproto/beerproto_go"
)

func DiastaticPower(value string, options OptionsFunc[beerproto.DiastaticPowerUnitType]) *beerproto.DiastaticPowerRangeType {
	rangeType := &RangeType[beerproto.DiastaticPowerUnitType]{}

	parse(value, rangeType, options)

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
