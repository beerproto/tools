package unit

import (
	beerproto "github.com/beerproto/beerproto_go"
)

func Acidity(value string, options ...OptionsFunc[beerproto.AcidityUnitType]) *beerproto.AcidityRangeType {
	rangeType := &RangeType[beerproto.AcidityUnitType, float64]{}

	options = append(options,
		WithMinContains[beerproto.AcidityUnitType]([]string{"min"}),
		WithMinTrim[beerproto.AcidityUnitType]([]string{"min"}),
		WithMaxTrim[beerproto.AcidityUnitType]([]string{"max"}),
		WithMaxContains[beerproto.AcidityUnitType]([]string{"max"}),
		WithUnit(beerproto.AcidityUnitType_PH),
		WithDefault[beerproto.AcidityUnitType](Max),
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
