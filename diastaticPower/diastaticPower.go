package diastaticPower

import (
	"strings"

	beerproto "github.com/beerproto/beerproto_go"
	"github.com/beerproto/tools/utils"
	"tawesoft.co.uk/go/lxstrconv"
)

type DiastaticPowerOption struct {
	formatter lxstrconv.NumberFormat
	unit      beerproto.DiastaticPowerUnitType
}

func WithFormatter(formatter lxstrconv.NumberFormat) func(*DiastaticPowerOption) {
	return func(s *DiastaticPowerOption) {
		s.formatter = formatter
	}
}

func WithUnit(unit beerproto.DiastaticPowerUnitType) func(*DiastaticPowerOption) {
	return func(s *DiastaticPowerOption) {
		s.unit = unit
	}
}

func Parse(value string, options ...func(*DiastaticPowerOption)) (diastaticPower *beerproto.DiastaticPowerRangeType) {
	opts := &DiastaticPowerOption{}
	for _, o := range options {
		o(opts)
	}

	diastaticPower = &beerproto.DiastaticPowerRangeType{}
	value = strings.ToLower(strings.TrimSpace(value))
	length := len(value)
	if length == 0 {
		return
	}

	arr := utils.Split(value, []string{"–", "-"})

	if len(arr) == 0 {
		arr = []string{value}
	}

	if ok := utils.Contains(arr[0], []string{"min"}); ok {
		min := utils.TrimAny(arr[0], []string{"min", "wk"})
		if min, err := opts.formatter.ParseFloat(min); err == nil {
			diastaticPower.Minimum = &beerproto.DiastaticPowerType{
				Value: min,
				Unit:  opts.unit,
			}
		}

	} else if ok = utils.Contains(arr[0], []string{"max"}); ok {
		max := utils.TrimAny(arr[0], []string{"max", "wk"})
		if max, err := opts.formatter.ParseFloat(max); err == nil {
			diastaticPower.Maximum = &beerproto.DiastaticPowerType{
				Value: max,
				Unit:  opts.unit,
			}
		}

	} else {
		min := utils.TrimAny(arr[0], []string{"min", "wk"})
		if min, err := opts.formatter.ParseFloat(min); err == nil {
			diastaticPower.Minimum = &beerproto.DiastaticPowerType{
				Value: min,
				Unit:  opts.unit,
			}
		}

	}

	if len(arr) == 2 {
		max := utils.TrimAny(arr[0], []string{"max", "wk"})
		if max, err := opts.formatter.ParseFloat(max); err == nil {
			diastaticPower.Maximum = &beerproto.DiastaticPowerType{
				Value: max,
				Unit:  opts.unit,
			}
		}

	}

	return
}
