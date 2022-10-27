package unit

import (
	"strings"

	"github.com/beerproto/tools/utils"
	"tawesoft.co.uk/go/lxstrconv"
)

type Option[TUnit Unit] struct {
	formatter   lxstrconv.NumberFormat
	unit        TUnit
	splitter    []string
	minContains []string
	minTrim     []string

	maxContains []string
	maxTrim     []string

	defaultRange *Default
}

type Default int

const (
	Min Default = 0
	Max         = 1
)

func WithMinTrim[TUnit Unit](trime []string) func(*Option[TUnit]) {
	return func(s *Option[TUnit]) {
		s.minTrim = trime
	}
}

func WithMinContains[TUnit Unit](contains []string) func(*Option[TUnit]) {
	return func(s *Option[TUnit]) {
		s.minContains = contains
	}
}

func WithFormatter[TUnit Unit](formatter lxstrconv.NumberFormat) func(*Option[TUnit]) {
	return func(s *Option[TUnit]) {
		s.formatter = formatter
	}
}

func WithUnit[TUnit Unit](unit TUnit) func(*Option[TUnit]) {
	return func(s *Option[TUnit]) {
		s.unit = unit
	}
}

func WithSplitter[TUnit Unit](splitter []string) func(*Option[TUnit]) {
	return func(s *Option[TUnit]) {
		s.splitter = splitter
	}
}

func Parse[TUnit Unit](value string, rangeType *RangeType[TUnit], options ...func(*Option[TUnit])) {
	opts := &Option[TUnit]{}
	for _, o := range options {
		o(opts)
	}

	value = strings.ToLower(strings.TrimSpace(value))
	length := len(value)
	if length == 0 {
		return
	}

	arr := []string{}

	if opts.splitter != nil {
		arr = utils.Split(value, opts.splitter)
	}

	if len(arr) == 0 {
		arr = []string{value}
	}

	found := false

	if opts.minContains != nil {
		for _, s := range arr {
			if ok := utils.Contains(s, opts.minContains); ok {
				min := ""
				if opts.minTrim != nil {
					min = utils.TrimAny(s, opts.minTrim)
				}
				if ok, unit := unit(min, opts); ok {
					found = true
					rangeType.Minimum = *unit
					break
				}
			}
		}
	}

	if opts.maxContains != nil {
		for _, s := range arr {
			if ok := utils.Contains(s, opts.maxContains); ok {
				max := ""
				if opts.maxTrim != nil {
					max = utils.TrimAny(s, opts.maxTrim)
				}
				if ok, unit := unit(max, opts); ok {
					found = true
					rangeType.Maximum = *unit
					break
				}
			}
		}
	}

	if opts.defaultRange != nil && !found {
		switch *opts.defaultRange {
		case Min:
			if opts.minTrim != nil {
				min := utils.TrimAny(arr[0], opts.minTrim)
				if ok, unit := unit(min, opts); ok {
					rangeType.Minimum = *unit
					return
				}
			}
		case Max:
			if opts.maxTrim != nil {
				max := utils.TrimAny(arr[0], opts.maxTrim)
				if ok, unit := unit(max, opts); ok {
					rangeType.Maximum = *unit
					return
				}
			}
		}
	}
}

func unit[TUnit Unit](value string, opts *Option[TUnit]) (bool, *UnitType[TUnit]) {
	if value, err := opts.formatter.ParseFloat(value); err == nil {
		t := new(UnitType[TUnit])
		t.Value = value
		t.Unit = opts.unit
		return true, t
	}

	return false, nil
}
