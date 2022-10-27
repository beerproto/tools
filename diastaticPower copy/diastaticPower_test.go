package diastaticPower

import (
	"reflect"
	"testing"

	beerproto "github.com/beerproto/beerproto_go"
	"golang.org/x/text/language"
	"tawesoft.co.uk/go/lxstrconv"
)

func Test_diastaticPower(t *testing.T) {

	tests := []struct {
		value       string
		options     []func(*DiastaticPowerOption)
		wantPercent *beerproto.DiastaticPowerRangeType
	}{
		{
			value: "245 wk min",
			options: []func(*DiastaticPowerOption){
				WithUnit(beerproto.DiastaticPowerUnitType_WK),
				WithFormatter(lxstrconv.NewDecimalFormat(language.BritishEnglish)),
			},
			wantPercent: &beerproto.DiastaticPowerRangeType{
				Minimum: &beerproto.DiastaticPowerType{
					Unit:  beerproto.DiastaticPowerUnitType_WK,
					Value: 245,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.value, func(t *testing.T) {
			if gotPercent := Parse(tt.value, tt.options...); !reflect.DeepEqual(gotPercent, tt.wantPercent) {
				t.Errorf("DiastaticPower() = %v, want %v", gotPercent, tt.wantPercent)
			}
		})
	}
}
