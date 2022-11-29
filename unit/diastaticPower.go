package unit

import (
	beerproto "github.com/beerproto/beerproto_go"
)

func DiastaticPower(value string, options ...OptionsFunc[beerproto.DiastaticPowerUnit]) *beerproto.DiastaticPowerRangeType {
	rangeType := &RangeType[beerproto.DiastaticPowerUnit, float64]{}

	options = append(options,
		WithMinContains[beerproto.DiastaticPowerUnit]([]string{">", "wk", "min"}),
		WithMinTrim[beerproto.DiastaticPowerUnit]([]string{">", "wk", "min"}),
		WithMaxTrim[beerproto.DiastaticPowerUnit]([]string{"<", "max"}),
		WithMaxContains[beerproto.DiastaticPowerUnit]([]string{"<", "wk", "max"}),
		WithUnit(beerproto.DiastaticPowerUnit_DIASTATIC_POWER_UNIT_WK),
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
