package unit

import (
	beerproto "github.com/beerproto/beerproto_go"
)

func Time(value string, options ...OptionsFunc[beerproto.TimeType_TimeUnitType]) *beerproto.TimeRangeType {
	rangeType := &RangeType[beerproto.TimeType_TimeUnitType, int64]{}

	options = append(options,
		WithMinContains[beerproto.TimeType_TimeUnitType]([]string{"min"}),
		WithMinTrim[beerproto.TimeType_TimeUnitType]([]string{"min"}),
		WithMaxTrim[beerproto.TimeType_TimeUnitType]([]string{"max"}),
		WithMaxContains[beerproto.TimeType_TimeUnitType]([]string{"max"}),
		WithUnit(beerproto.TimeType_MIN),
		WithDefault[beerproto.TimeType_TimeUnitType](Max),
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
