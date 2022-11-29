package unit

import (
	beerproto "github.com/beerproto/beerproto_go"
)

func Time(value string, options ...OptionsFunc[beerproto.TimeUnit]) *beerproto.TimeRangeType {
	rangeType := &RangeType[beerproto.TimeUnit, int64]{}

	options = append(options,
		WithMinContains[beerproto.TimeUnit]([]string{"min"}),
		WithMinTrim[beerproto.TimeUnit]([]string{"min"}),
		WithMaxTrim[beerproto.TimeUnit]([]string{"max"}),
		WithMaxContains[beerproto.TimeUnit]([]string{"max"}),
		WithUnit(beerproto.TimeUnit_TIME_UNIT_MIN),
		WithDefault[beerproto.TimeUnit](Max),
	)

	parse(value, rangeType, options...)

	time := &beerproto.TimeRangeType{}

	if rangeType.Maximum != nil {
		time.Maximum = &beerproto.TimeType{
			Unit:  rangeType.Maximum.Unit,
			Value: rangeType.Maximum.Value,
		}
	}

	if rangeType.Minimum != nil {
		time.Minimum = &beerproto.TimeType{
			Unit:  rangeType.Minimum.Unit,
			Value: rangeType.Minimum.Value,
		}
	}

	return time
}
