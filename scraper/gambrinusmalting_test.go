package scraper

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/beerproto/beerproto_go/fermentables"
	"google.golang.org/protobuf/encoding/protojson"
)

func TestGambrinusMalting(t *testing.T) {

	list := &fermentables.GrainsType{
		Grains: []*fermentables.GrainType{},
	}
	list.Grains = append(list.Grains, NewGambrinusMalting().Parse()...)

	mops := &protojson.MarshalOptions{}
	b, _ := mops.Marshal(list)

	absPath, _ := filepath.Abs("data.json")
	err := os.WriteFile(absPath, b, 0644)
	if err != nil {
		panic(err)
	}
}
