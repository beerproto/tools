package main

import (
	"os"
	"path/filepath"

	"github.com/beerproto/beerproto_go/fermentables"
	"github.com/beerproto/tools/scraper"
	"google.golang.org/protobuf/encoding/protojson"
)

func main() {
	list := &fermentables.GrainsType{
		Grains: []*fermentables.GrainType{},
	}
	list.Grains = append(list.Grains, scraper.Agraria()...)

	mops := &protojson.MarshalOptions{}
	b, _ := mops.Marshal(list)

	absPath, _ := filepath.Abs("data.json")
	err := os.WriteFile(absPath, b, 0644)
	if err != nil {
		panic(err)
	}
}
