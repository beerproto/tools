package unit

import (
	"reflect"
	"testing"

	beerproto "github.com/beerproto/beerproto_go"
	"golang.org/x/text/language"
	"tawesoft.co.uk/go/lxstrconv"
)

func TestViscosity(t *testing.T) {
	tests := []struct {
		name              string
		value             string
		options           []OptionsFunc[beerproto.ViscosityUnitType]
		wantViscositRange *beerproto.ViscosityRangeType
	}{
		{
			value: "1.65cP Max",
			options: []OptionsFunc[beerproto.ViscosityUnitType]{
				WithFormatter[beerproto.ViscosityUnitType](lxstrconv.NewDecimalFormat(language.BritishEnglish)),
				WithUnit(beerproto.ViscosityUnitType_CP),
			},
			wantViscositRange: &beerproto.ViscosityRangeType{
				Maximum: &beerproto.ViscosityType{
					Value: 1.65,
					Unit:  beerproto.ViscosityUnitType_CP,
				},
			},
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotDiastaticPower := Viscosity(tt.value, tt.options...); !reflect.DeepEqual(gotDiastaticPower, tt.wantViscositRange) {
				t.Errorf("Viscosity() = %v, want %v", gotDiastaticPower, tt.wantViscositRange)
			}
		})
	}
}
