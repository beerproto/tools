package unit

import (
	beerproto "github.com/beerproto/beerproto_go"
)

func Acidity(value string, options ...OptionsFunc[beerproto.AcidityUnit]) *beerproto.AcidityRangeType {
	rangeType := &RangeType[beerproto.AcidityUnit, float64]{}

	options = append(options,
		WithMinContains[beerproto.AcidityUnit]([]string{"min"}),
		WithMinTrim[beerproto.AcidityUnit]([]string{"min"}),
		WithMaxTrim[beerproto.AcidityUnit]([]string{"max"}),
		WithMaxContains[beerproto.AcidityUnit]([]string{"max"}),
		WithUnit(beerproto.AcidityUnit_ACIDITY_UNIT_PH),
		WithDefault[beerproto.AcidityUnit](Max),
	)

	parse(value, rangeType, options...)

	acidity := &beerproto.AcidityRangeType{}

	if rangeType.Maximum != nil {
		acidity.Maximum = &beerproto.AcidityType{
			Unit:  rangeType.Maximum.Unit,
			Value: rangeType.Maximum.Value,
		}
	}

	if rangeType.Minimum != nil {
		acidity.Minimum = &beerproto.AcidityType{
			Unit:  rangeType.Minimum.Unit,
			Value: rangeType.Minimum.Value,
		}
	}

	return acidity
}
