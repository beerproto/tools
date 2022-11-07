package unit

import (
	beerproto "github.com/beerproto/beerproto_go"
)

func Viscosity(value string, options ...OptionsFunc[beerproto.ViscosityUnitType]) *beerproto.ViscosityRangeType {
	rangeType := &RangeType[beerproto.ViscosityUnitType, float64]{}

	options = append(options,
		WithMinContains[beerproto.ViscosityUnitType]([]string{"min"}),
		WithMinTrim[beerproto.ViscosityUnitType]([]string{"cp", "min"}),
		WithMaxTrim[beerproto.ViscosityUnitType]([]string{"cp", "max"}),
		WithMaxContains[beerproto.ViscosityUnitType]([]string{"max"}),
		WithUnit(beerproto.ViscosityUnitType_CP),
		WithDefault[beerproto.ViscosityUnitType](Max),
		WithDecimals[beerproto.ViscosityUnitType](2),
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
