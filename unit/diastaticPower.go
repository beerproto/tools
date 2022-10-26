package unit

import (
	"strings"

	"github.com/beerproto/tools/utils"
)

type Option[TUnit Unit, TNumber Number] struct {
	formatter   func(string) (TNumber, error)
	unit        TUnit
	splitter    []string
	minContains []string
	minTrim     []string

	maxContains []string
	maxTrim     []string

	defaultRange Default
}

type Default int

const (
	Min Default = 0
	Max         = 1
)

func WithFormatter[TUnit Unit, TNumber Number](formatter func(string) (TNumber, error)) func(*Option[TUnit, TNumber]) {
	return func(s *Option[TUnit, TNumber]) {
		s.formatter = formatter
	}
}

func WithUnit[TUnit Unit, TNumber Number](unit TUnit) func(*Option[TUnit, TNumber]) {
	return func(s *Option[TUnit, TNumber]) {
		s.unit = unit
	}
}

func WithSplitter[TUnit Unit, TNumber Number](splitter []string) func(*Option[TUnit, TNumber]) {
	return func(s *Option[TUnit, TNumber]) {
		s.splitter = splitter
	}
}

func Parse[TUnit Unit, TNumber Number](value string, rangeType RangeType[TUnit, TNumber], options ...func(*Option[TUnit, TNumber])) {
	opts := &Option[TUnit, TNumber]{}
	for _, o := range options {
		o(opts)
	}

	value = strings.ToLower(strings.TrimSpace(value))
	length := len(value)
	if length == 0 {
		return
	}

	arr := utils.Split(value, opts.splitter)

	if len(arr) == 0 {
		arr = []string{value}
	}

	if ok := utils.Contains(arr[0], opts.minContains); ok {
		min := utils.TrimAny(arr[0], opts.minTrim)
		if min, err := opts.formatter(min); err == nil {
			t := new(UnitType[TUnit, TNumber])
			t.Value = min
			t.Unit = opts.unit
			rangeType.Minimum = *t
		}

	} else if ok = utils.Contains(arr[0], opts.maxContains); ok {
		max := utils.TrimAny(arr[0], opts.maxTrim)
		if max, err := opts.formatter(max); err == nil {
			t := new(UnitType[TUnit, TNumber])
			t.Value = max
			t.Unit = opts.unit
			rangeType.Maximum = *t
		}

	} else {
		min := utils.TrimAny(arr[0], opts.minTrim)
		if min, err := opts.formatter(min); err == nil {
			t := new(UnitType[TUnit, TNumber])
			t.Value = min
			t.Unit = opts.unit
			rangeType.Minimum = *t
		}

	}

	if len(arr) == 2 {
		max := utils.TrimAny(arr[0], opts.maxTrim)
		if max, err := opts.formatter(max); err == nil {
			t := new(UnitType[TUnit, TNumber])
			t.Value = max
			t.Unit = opts.unit
			rangeType.Maximum = *t
		}

	}

	return
}
