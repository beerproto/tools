package samples_test

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"html"
	"io/ioutil"
	"net/http/httptest"
	"reflect"
	"strconv"
	"testing"

	"github.com/beerproto/beerxml.go/handlers"
	"github.com/beerproto/beerxml.go/reader"
)

func TestSchemas_Generate(t *testing.T) {
	tests := []struct {
		name string
		xml  string
	}{
		{
			name: "recipes",
			xml:  "recipes.xml",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			file, err := ioutil.ReadFile(tt.xml)
			if err != nil {
				t.Error(err)
			}

			reader := bytes.NewBuffer(file)

			req := httptest.NewRequest("POST", "http://example.com", reader)
			w := httptest.NewRecorder()
			recipes := &beerXML.RECIPES{}

			handlers.BeerXML(w, req, recipes)

			buf := bytes.Buffer{}
			encoder := xml.NewEncoder(&buf)
			err = encoder.Encode(recipes)
			if err != nil {
				t.Error(err)
			}

			err = ShouldEqualXMLObject(file, buf.Bytes())
			if err != nil {
				t.Error(err)
			}
		})
	}
}

func ShouldEqualXMLObject(data1, data2 []byte) error {
	dec := xml.NewDecoder(bytes.NewBuffer(data1))
	dec.CharsetReader = reader.MakeCharsetReader

	x := &Node{}
	err := dec.Decode(&x)

	if err != nil {
		return fmt.Errorf("unmarshal of data1 failed: %w", err)
	}

	dec2 := xml.NewDecoder(bytes.NewBuffer(data2))
	dec2.CharsetReader = reader.MakeCharsetReader

	y := &Node{}
	err = dec2.Decode(&y)
	if err != nil {
		return fmt.Errorf("unmarshal of data2 failed: %w", err)
	}

	xNodes := walk([]*Node{x})
	yNodes := walk([]*Node{y})

	jx, err := json.Marshal(xNodes)
	if err != nil {
		return fmt.Errorf("marshal of data1 failed: %w", err)
	}

	jy, err := json.Marshal(yNodes)
	if err != nil {
		return fmt.Errorf("marshal of data2 failed: %w", err)
	}

	if !reflect.DeepEqual(jx, jy) {
		return fmt.Errorf("object not equal \nexpected \n%v \ngot \n%v", string(jx), string(jy))
	}

	return nil
}

type Node struct {
	XMLName xml.Name
	Attrs   []xml.Attr `xml:"-"`
	Content []byte     `xml:",innerxml"`
	Nodes   []*Node    `xml:",any"`
}

func (n *Node) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	n.Attrs = start.Attr
	type node Node

	return d.DecodeElement((*node)(n), &start)
}

func walk(nodes []*Node) map[string]interface{} {
	m := make(map[string]interface{})
	for _, n := range nodes {
		if len(n.Nodes) > 0 {
			m[n.XMLName.Local] = walk(n.Nodes)
		} else {
			content := bytes.TrimSpace(n.Content)
			str := string(content)

			if f, err := strconv.ParseInt(str, 10, 32); err == nil {
				m[n.XMLName.Local] = f
			} else if f, err := strconv.ParseFloat(str, 32); err == nil {
				m[n.XMLName.Local] = f
			} else if f, err := strconv.ParseBool(str); err == nil {
				m[n.XMLName.Local] = f
			} else {
				m[n.XMLName.Local] = html.UnescapeString(str)
			}
		}
	}
	return m
}
