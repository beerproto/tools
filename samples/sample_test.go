package samples_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"
	"testing"

	"github.com/RossMerr/beerjson.go"
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
		// {
		// 	name: "CorianderSpice",
		// 	json: "CorianderSpice.json",
		// },
		// {
		// 	name: "CrystalMaltSpecialtyGrain",
		// 	json: "CrystalMaltSpecialtyGrain.json",
		// },
		// {
		// 	name: "EquipmentSet",
		// 	json: "EquipmentSet.json",
		// },
		// {
		// 	name: "FermentableRecord",
		// 	json: "FermentableRecord.json",
		// },
		// {
		// 	name: "HoppedExtract",
		// 	json: "HoppedExtract.json",
		// },
		// {
		// 	name: "HopRecordWithAllFields",
		// 	json: "HopRecordWithAllFields.json",
		// },
		// {
		// 	name: "HopWithRequiredFieldsOnly",
		// 	json: "HopWithRequiredFieldsOnly.json",
		// },
		// {
		// 	name: "IrishMoss",
		// 	json: "IrishMoss.json",
		// },
		// {
		// 	name: "MashSingleStepInfusion",
		// 	json: "MashSingleStepInfusion.json",
		// },
		// {
		// 	name: "MashTwoStepTemperature",
		// 	json: "MashTwoStepTemperature.json",
		// },
		// {
		// 	name: "MedievalAle",
		// 	json: "MedievalAle.json",
		// },
		// {
		// 	name: "SampleWaterProfile",
		// 	json: "SampleWaterProfile.json",
		// },
		// {
		// 	name: "StyleBohemianPilsner",
		// 	json: "StyleBohemianPilsner.json",
		// },
		// {
		// 	name: "StyleDryIrishStoutWithAllFields",
		// 	json: "StyleDryIrishStoutWithAllFields.json",
		// },
		// {
		// 	name: "YeastWithMorePopularFields",
		// 	json: "YeastWithMorePopularFields.json",
		// },
		// {
		// 	name: "YeastWithRequiredFieldsOnly",
		// 	json: "YeastWithRequiredFieldsOnly.json",
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := ioutil.ReadFile(tt.json)
			if err != nil {
				t.Error(err)
			}

			beer := &beerjson.Beerjson{}
			str :=  &struct {
				Beer *beerjson.Beerjson `json:"beerjson"`
			}{
				Beer: beer,
			}


			err = json.Unmarshal(data, &str)
			if err != nil {
				t.Error(err)
			}

			recipe, err := mapping.MapToProto(beer)
			if err != nil {
				t.Error(err)
			}

			j, err := mapping.MapToJSON(recipe)
			if err != nil {
				t.Error(err)
			}

			bytes, err := json.Marshal(j)
			if err != nil {
				t.Error(err)
			}

			trimedJSON, err := TrimJson(bytes)
			if err != nil {
				t.Error(err)
			}

			innerData, err := json.Marshal(beer)
			if err != nil {
				t.Error(err)
			}

			expectedJSON, err := TrimJson(innerData)
			if err != nil {
				t.Error(err)
			}

			err = ShouldEqualJSONObject(expectedJSON, trimedJSON)
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

func TrimJson(data []byte) ([]byte, error) {
	x := make(map[string]interface{})
	err := json.Unmarshal(data, &x)
	if err != nil {
		return nil, err
	}

	result := trimJson(x)

	results, err := json.Marshal(result)
	return results, err
}



func trimJson(x map[string]interface{}) (map[string]interface{}) {
	for key, value := range x {
		switch v := value.(type) {
		case string:
			if v == "" {
				delete(x, key)
			}
		case float64:
			if v == 0 {
				delete(x, key)
			}
		case bool:
			if v == false {
				delete(x, key)
			}
		case []interface{}:
			for p, iterator := range v {
				if m, ok := iterator.(map[string]interface{}); ok {
					m2 := trimJson(m)
					if len(m) > 0 {
						v[p] = m2
					} else {
						v = append(v[:p], v[p+1:]...)
					}
				}
			}
		case map[string]interface{}:
			 m := trimJson(v)
			 if len(m) > 0 {
				 x[key] = m
			 } else {
				 delete(x, key)
			 }
		}
	}
	return x
}