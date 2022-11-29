package unit

import (
	"reflect"
	"testing"

	beerproto "github.com/beerproto/beerproto_go"
	"golang.org/x/text/language"
	"tawesoft.co.uk/go/lxstrconv"
)

func TestPercent(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		options []OptionsFunc[beerproto.PercentUnit]

		wantConcentrationRange *beerproto.PercentRangeType
	}{
		{
			value: "~79.0",
			options: []OptionsFunc[beerproto.PercentUnit]{
				WithFormatter[beerproto.PercentUnit](lxstrconv.NewDecimalFormat(language.BritishEnglish)),
				WithUnit(beerproto.PercentUnit_PERCENT_UNIT_PERCENT_SIGN),
				WithDefault[beerproto.PercentUnit](Max),
			},
			wantConcentrationRange: &beerproto.PercentRangeType{
				Maximum: &beerproto.PercentType{
					Value: 79,
					Unit:  beerproto.PercentUnit_PERCENT_UNIT_PERCENT_SIGN,
				},
			},
		},
		{
			value: "5.5%",
			options: []OptionsFunc[beerproto.PercentUnit]{
				WithFormatter[beerproto.PercentUnit](lxstrconv.NewDecimalFormat(language.BritishEnglish)),
				WithUnit(beerproto.PercentUnit_PERCENT_UNIT_PERCENT_SIGN),
				WithDefault[beerproto.PercentUnit](Max),
			},
			wantConcentrationRange: &beerproto.PercentRangeType{
				Maximum: &beerproto.PercentType{
					Value: 5.5,
					Unit:  beerproto.PercentUnit_PERCENT_UNIT_PERCENT_SIGN,
				},
			},
		},
		{
			value: "<13.0",
			options: []OptionsFunc[beerproto.PercentUnit]{
				WithFormatter[beerproto.PercentUnit](lxstrconv.NewDecimalFormat(language.BritishEnglish)),
				WithUnit(beerproto.PercentUnit_PERCENT_UNIT_PERCENT_SIGN),
				WithDefault[beerproto.PercentUnit](Max),
			},
			wantConcentrationRange: &beerproto.PercentRangeType{
				Maximum: &beerproto.PercentType{
					Value: 13,
					Unit:  beerproto.PercentUnit_PERCENT_UNIT_PERCENT_SIGN,
				},
			},
		},
		{
			value: "4 and 5%",
			options: []OptionsFunc[beerproto.PercentUnit]{
				WithFormatter[beerproto.PercentUnit](lxstrconv.NewDecimalFormat(language.BritishEnglish)),
				WithUnit(beerproto.PercentUnit_PERCENT_UNIT_PERCENT_SIGN),
				WithDefault[beerproto.PercentUnit](Max),
			},
			wantConcentrationRange: &beerproto.PercentRangeType{
				Minimum: &beerproto.PercentType{
					Value: 4,
					Unit:  beerproto.PercentUnit_PERCENT_UNIT_PERCENT_SIGN,
				},
				Maximum: &beerproto.PercentType{
					Value: 5,
					Unit:  beerproto.PercentUnit_PERCENT_UNIT_PERCENT_SIGN,
				},
			},
		},
		{
			value: "%4.9",
			options: []OptionsFunc[beerproto.PercentUnit]{
				WithFormatter[beerproto.PercentUnit](lxstrconv.NewDecimalFormat(language.BritishEnglish)),
				WithUnit(beerproto.PercentUnit_PERCENT_UNIT_PERCENT_SIGN),
				WithDefault[beerproto.PercentUnit](Max),
			},
			wantConcentrationRange: &beerproto.PercentRangeType{
				Maximum: &beerproto.PercentType{
					Value: 4.9,
					Unit:  beerproto.PercentUnit_PERCENT_UNIT_PERCENT_SIGN,
				},
			},
		},
		{
			value: "≥ 85% ",
			options: []OptionsFunc[beerproto.PercentUnit]{
				WithFormatter[beerproto.PercentUnit](lxstrconv.NewDecimalFormat(language.BritishEnglish)),
				WithUnit(beerproto.PercentUnit_PERCENT_UNIT_PERCENT_SIGN),
				WithDefault[beerproto.PercentUnit](Max),
			},
			wantConcentrationRange: &beerproto.PercentRangeType{
				Maximum: &beerproto.PercentType{
					Value: 85,
					Unit:  beerproto.PercentUnit_PERCENT_UNIT_PERCENT_SIGN,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotDiastaticPower := Percent(tt.value, tt.options...); !reflect.DeepEqual(gotDiastaticPower, tt.wantConcentrationRange) {
				t.Errorf("Percent() = %v, want %v", gotDiastaticPower, tt.wantConcentrationRange)
			}
		})
	}
}
