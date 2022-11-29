package unit

import (
	beerproto "github.com/beerproto/beerproto_go"
)

func Viscosity(value string, options ...OptionsFunc[beerproto.ViscosityUnit]) *beerproto.ViscosityRangeType {
	rangeType := &RangeType[beerproto.ViscosityUnit, float64]{}

	options = append(options,
		WithMinContains[beerproto.ViscosityUnit]([]string{"min"}),
		WithMinTrim[beerproto.ViscosityUnit]([]string{"cp", "min"}),
		WithMaxTrim[beerproto.ViscosityUnit]([]string{"cp", "max"}),
		WithMaxContains[beerproto.ViscosityUnit]([]string{"max"}),
		WithUnit(beerproto.ViscosityUnit_VISCOSITY_UNIT_CP),
		WithDefault[beerproto.ViscosityUnit](Max),
		WithDecimals[beerproto.ViscosityUnit](2),
	)

	parse(value, rangeType, options...)

	viscosity := &beerproto.ViscosityRangeType{}

	if rangeType.Maximum != nil {
		viscosity.Maximum = &beerproto.ViscosityType{
			Unit:  rangeType.Maximum.Unit,
			Value: rangeType.Maximum.Value,
		}
	}

	if rangeType.Minimum != nil {
		viscosity.Minimum = &beerproto.ViscosityType{
			Unit:  rangeType.Minimum.Unit,
			Value: rangeType.Minimum.Value,
		}
	}

	return viscosity
}
