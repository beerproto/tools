package color

import (
	"strings"

	beerproto "github.com/beerproto/beerproto_go"
	"github.com/beerproto/tools/utils"
	"tawesoft.co.uk/go/lxstrconv"
)

type ColorOption struct {
	formatter lxstrconv.NumberFormat
	unit      beerproto.ColorUnitType
}

func WithFormatter(formatter lxstrconv.NumberFormat) func(*ColorOption) {
	return func(s *ColorOption) {
		s.formatter = formatter
	}
}

func WithUnit(unit beerproto.ColorUnitType) func(*ColorOption) {
	return func(s *ColorOption) {
		s.unit = unit
	}
}

func Parse(value string, options ...func(*ColorOption)) (color *beerproto.ColorRangeType) {
	opts := &ColorOption{}
	for _, o := range options {
		o(opts)
	}

	color = &beerproto.ColorRangeType{}

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
		min := utils.TrimAny(arr[0], []string{"min", "°ebc"})
		if min, err := opts.formatter.ParseFloat(min); err == nil {
			color.Minimum = &beerproto.ColorType{
				Value: min,
				Unit:  opts.unit,
			}
		}

	} else if ok = utils.Contains(arr[0], []string{"max"}); ok {
		max := utils.TrimAny(arr[0], []string{"max", "°ebc"})
		if max, err := opts.formatter.ParseFloat(max); err == nil {
			color.Maximum = &beerproto.ColorType{
				Value: max,
				Unit:  opts.unit,
			}
		}
	}
	if len(arr) == 1 {
		max := utils.TrimAny(arr[0], []string{"max", "°ebc"})
		if max, err := opts.formatter.ParseFloat(max); err == nil {
			color.Maximum = &beerproto.ColorType{
				Value: max,
				Unit:  opts.unit,
			}
		}
	}

	if len(arr) == 2 {
		min := utils.TrimAny(arr[0], []string{"min", "°ebc"})
		if min, err := opts.formatter.ParseFloat(min); err == nil {
			color.Minimum = &beerproto.ColorType{
				Value: min,
				Unit:  opts.unit,
			}
		}
		max := utils.TrimAny(arr[1], []string{"max", "°ebc"})
		if max, err := opts.formatter.ParseFloat(max); err == nil {
			color.Maximum = &beerproto.ColorType{
				Value: max,
				Unit:  opts.unit,
			}
		}

	}

	return
}
