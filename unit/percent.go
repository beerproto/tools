package unit

import (
	beerproto "github.com/beerproto/beerproto_go"
)

func Percent(value string, options ...OptionsFunc[beerproto.PercentType_PercentUnitType]) *beerproto.PercentRangeType {
	rangeType := &RangeType[beerproto.PercentType_PercentUnitType, float64]{}

	options = append(options,
		WithMinContains[beerproto.PercentType_PercentUnitType]([]string{"min"}),
		WithMinTrim[beerproto.PercentType_PercentUnitType]([]string{"≤", "min", "%"}),
		WithMaxTrim[beerproto.PercentType_PercentUnitType]([]string{"≥", "max", "%"}),
		WithMaxContains[beerproto.PercentType_PercentUnitType]([]string{"max"}),
		WithUnit(beerproto.PercentType_PERCENT_SIGN),
		WithDefault[beerproto.PercentType_PercentUnitType](Max),
	)

	parse(value, rangeType, options...)

	percent := &beerproto.PercentRangeType{}

	if rangeType.Maximum != nil {
		percent.Maximum = &beerproto.PercentType{
			Unit:  rangeType.Maximum.Unit,
			Value: rangeType.Maximum.Value,
		}
	}

	if rangeType.Minimum != nil {
		percent.Minimum = &beerproto.PercentType{
			Unit:  rangeType.Minimum.Unit,
			Value: rangeType.Minimum.Value,
		}
	}

	return percent
}
