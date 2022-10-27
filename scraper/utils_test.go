package scraper

import (
	"reflect"
	"testing"

	beerproto "github.com/beerproto/beerproto_go"
	"golang.org/x/text/language"
	"tawesoft.co.uk/go/lxstrconv"
)

func Test_concentration(t *testing.T) {
	formater := lxstrconv.NewDecimalFormat(language.BritishEnglish)

	tests := []struct {
		value string
		want  *beerproto.ConcentrationRangeType
	}{
		{
			value: "630 -730",
			want: &beerproto.ConcentrationRangeType{
				Minimum: &beerproto.ConcentrationType{
					Unit:  beerproto.ConcentrationUnitType_MGL,
					Value: 630,
				},
				Maximum: &beerproto.ConcentrationType{
					Unit:  beerproto.ConcentrationUnitType_MGL,
					Value: 730,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.value, func(t *testing.T) {
			if gotColor := concentration(tt.value, formater, beerproto.ConcentrationUnitType_MGL); !reflect.DeepEqual(gotColor, tt.want) {
				t.Errorf("concentration() = %v, want %v", gotColor, tt.want)
			}
		})
	}
}

func Test_percent(t *testing.T) {
	formater := lxstrconv.NewDecimalFormat(language.BritishEnglish)

	tests := []struct {
		value       string
		wantPercent *beerproto.PercentRangeType
	}{
		{
			value: "%4.9",
			wantPercent: &beerproto.PercentRangeType{
				Maximum: &beerproto.PercentType{
					Unit:  beerproto.PercentType_PERCENT_SIGN,
					Value: 4.9,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.value, func(t *testing.T) {
			if gotPercent := percent(tt.value, formater, beerproto.PercentType_PERCENT_SIGN); !reflect.DeepEqual(gotPercent, tt.wantPercent) {
				t.Errorf("percent() = %v, want %v", gotPercent, tt.wantPercent)
			}
		})
	}
}

func Test_diastaticPower(t *testing.T) {
	formater := lxstrconv.NewDecimalFormat(language.BritishEnglish)

	tests := []struct {
		value       string
		wantPercent *beerproto.DiastaticPowerRangeType
		unit        beerproto.DiastaticPowerUnitType
	}{
		{
			value: "245 wk min",
			wantPercent: &beerproto.DiastaticPowerRangeType{
				Minimum: &beerproto.DiastaticPowerType{
					Unit:  beerproto.DiastaticPowerUnitType_WK,
					Value: 245,
				},
			},
			unit: beerproto.DiastaticPowerUnitType_WK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.value, func(t *testing.T) {
			if gotPercent := diastaticPower(tt.value, formater, tt.unit); !reflect.DeepEqual(gotPercent, tt.wantPercent) {
				t.Errorf("diastaticPower() = %v, want %v", gotPercent, tt.wantPercent)
			}
		})
	}
}
