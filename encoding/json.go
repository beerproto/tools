package encoding

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/beerproto/tools/beerproto"
	"github.com/beerproto/tools/mapping"
	"google.golang.org/protobuf/encoding/protojson"
)

func JSON(r io.Reader, beer *beerproto.Recipe) error {

	m := make(map[string]json.RawMessage)
	err := json.NewDecoder(r).Decode(&m)
	if err != nil {
		return fmt.Errorf("json: %w", err)
	}

	m = mapping.Enum(m)

	data, err := json.Marshal(&m)
	err = protojson.Unmarshal(data, beer)
	if err != nil {
		return fmt.Errorf("json: %w", err)
	}

	return nil
}
