package unit

import (
	beerproto "github.com/beerproto/beerproto_go"
)

func Percent(value string, options ...OptionsFunc[beerproto.PercentUnit]) *beerproto.PercentRangeType {
	rangeType := &RangeType[beerproto.PercentUnit, float64]{}

	options = append(options,
		WithMinContains[beerproto.PercentUnit]([]string{">", "min"}),
		WithMinTrim[beerproto.PercentUnit]([]string{">", "≤", "min", "%"}),
		WithMaxTrim[beerproto.PercentUnit]([]string{"<", "≥", "max", "%"}),
		WithMaxContains[beerproto.PercentUnit]([]string{"<", "max"}),
		WithUnit(beerproto.PercentUnit_PERCENT_UNIT_PERCENT_SIGN),
		WithDefault[beerproto.PercentUnit](Max),
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
