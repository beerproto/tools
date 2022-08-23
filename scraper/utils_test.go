package scraper

import (
	"reflect"
	"testing"

	beerproto "github.com/beerproto/beerproto_go"
	"golang.org/x/text/language"
	"tawesoft.co.uk/go/lxstrconv"
)

func Test_color(t *testing.T) {
	formater := lxstrconv.NewDecimalFormat(language.BritishEnglish)

	tests := []struct {
		value     string
		wantColor *beerproto.ColorRangeType
	}{
		{
			value: "3.0 - 3.5",
			wantColor: &beerproto.ColorRangeType{
				Minimum: &beerproto.ColorType{
					Unit:  beerproto.ColorUnitType_EBC,
					Value: 3,
				},
				Maximum: &beerproto.ColorType{
					Unit:  beerproto.ColorUnitType_EBC,
					Value: 3.5,
				},
			},
		},
		{
			value: "3.0",
			wantColor: &beerproto.ColorRangeType{
				Maximum: &beerproto.ColorType{
					Unit:  beerproto.ColorUnitType_EBC,
					Value: 3,
				},
			},
		},
		{
			value: "max 3.5",
			wantColor: &beerproto.ColorRangeType{
				Maximum: &beerproto.ColorType{
					Unit:  beerproto.ColorUnitType_EBC,
					Value: 3.5,
				},
			},
		},
		{
			value: "min 3.5",
			wantColor: &beerproto.ColorRangeType{
				Minimum: &beerproto.ColorType{
					Unit:  beerproto.ColorUnitType_EBC,
					Value: 3.5,
				},
			},
		},
		{
			value: "min 3.0 - max 3.5",
			wantColor: &beerproto.ColorRangeType{
				Minimum: &beerproto.ColorType{
					Unit:  beerproto.ColorUnitType_EBC,
					Value: 3,
				},
				Maximum: &beerproto.ColorType{
					Unit:  beerproto.ColorUnitType_EBC,
					Value: 3.5,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.value, func(t *testing.T) {
			if gotColor := color(tt.value, formater, beerproto.ColorUnitType_EBC); !reflect.DeepEqual(gotColor, tt.wantColor) {
				t.Errorf("color() = %v, want %v", gotColor, tt.wantColor)
			}
		})
	}
}

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
			if gotPercent := percent(tt.value, formater); !reflect.DeepEqual(gotPercent, tt.wantPercent) {
				t.Errorf("percent() = %v, want %v", gotPercent, tt.wantPercent)
			}
		})
	}
}
