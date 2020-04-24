package samples_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"
	"testing"

	mapping "github.com/beerproto/tools/mapping/beerJSON"
)

func TestSchemas_Generate(t *testing.T) {

	tests := []struct {
		name    string
		json    string
		wantErr bool
	}{
		{
			name: "Boil Whirlpool Chill",
			json: "boil_whirlpool_chill.json",
		},
		{
			name: "BrettDosedKegsSaison",
			json: "BrettDosedKegsSaison.json",
		},
		{
			name: "CheaterHops",
			json: "CheaterHops.json",
		},
		{
			name: "CorianderSpice",
			json: "CorianderSpice.json",
		},
		{
			name: "CrystalMaltSpecialtyGrain",
			json: "CrystalMaltSpecialtyGrain.json",
		},
		{
			name: "EquipmentSet",
			json: "EquipmentSet.json",
		},
		{
			name: "FermentableRecord",
			json: "FermentableRecord.json",
		},
		{
			name: "HoppedExtract",
			json: "HoppedExtract.json",
		},
		{
			name: "HopRecordWithAllFields",
			json: "HopRecordWithAllFields.json",
		},
		{
			name: "HopWithRequiredFieldsOnly",
			json: "HopWithRequiredFieldsOnly.json",
		},
		{
			name: "IrishMoss",
			json: "IrishMoss.json",
		},
		{
			name: "MashSingleStepInfusion",
			json: "MashSingleStepInfusion.json",
		},
		{
			name: "MashTwoStepTemperature",
			json: "MashTwoStepTemperature.json",
		},
		{
			name: "MedievalAle",
			json: "MedievalAle.json",
		},
		{
			name: "SampleWaterProfile",
			json: "SampleWaterProfile.json",
		},
		{
			name: "StyleBohemianPilsner",
			json: "StyleBohemianPilsner.json",
		},
		{
			name: "StyleDryIrishStoutWithAllFields",
			json: "StyleDryIrishStoutWithAllFields.json",
		},
		{
			name: "YeastWithMorePopularFields",
			json: "YeastWithMorePopularFields.json",
		},
		{
			name: "YeastWithRequiredFieldsOnly",
			json: "YeastWithRequiredFieldsOnly.json",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file, err := ioutil.ReadFile(tt.json)
			if err != nil {
				t.Error(err)
			}

			beerJSON, _ := mapping.MapToProto(file)

			data, err := json.Marshal(beerJSON)
			if err != nil {
				t.Error(err)
			}

			err = ShouldEqualJSONObject(file, data)
			if err != nil {
				t.Error(err)
			}
		})
	}
}

func ShouldEqualJSONObject(data1, data2 []byte) error {
	x := make(map[string]interface{})
	err := json.Unmarshal(data1, &x)
	if err != nil {
		return fmt.Errorf("unmarshal of data1 failed: %w", err)
	}
	y := make(map[string]interface{})
	err = json.Unmarshal(data2, &y)
	if err != nil {
		return fmt.Errorf("unmarshal of data2 failed: %w", err)
	}

	if !reflect.DeepEqual(x, y) {
		jx, err := json.Marshal(x)
		if err != nil {
			return fmt.Errorf("marshal of data1 failed: %w", err)
		}

		jy, err := json.Marshal(y)
		if err != nil {
			return fmt.Errorf("marshal of data2 failed: %w", err)
		}
		return fmt.Errorf("object not equal \nexpected \n%v \ngot \n%v", string(jx), string(jy))
	}

	return nil
}


