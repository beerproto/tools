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
		options []OptionsFunc[beerproto.PercentType_PercentUnitType]

		wantConcentrationRange *beerproto.PercentRangeType
	}{
		{
			value: "%4.9",
			options: []OptionsFunc[beerproto.PercentType_PercentUnitType]{
				WithFormatter[beerproto.PercentType_PercentUnitType](lxstrconv.NewDecimalFormat(language.BritishEnglish)),
				WithUnit(beerproto.PercentType_PERCENT_SIGN),
				WithDefault[beerproto.PercentType_PercentUnitType](Max),
			},
			wantConcentrationRange: &beerproto.PercentRangeType{
				Maximum: &beerproto.PercentType{
					Value: 4.9,
					Unit:  beerproto.PercentType_PERCENT_SIGN,
				},
			},
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotDiastaticPower := Percent(tt.value, tt.options...); !reflect.DeepEqual(gotDiastaticPower, tt.wantConcentrationRange) {
				t.Errorf("Percent() = %v, want %v", gotDiastaticPower, tt.wantConcentrationRange)
			}
		})
	}
}
