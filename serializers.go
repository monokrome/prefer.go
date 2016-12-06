package prefer

import (
	"encoding/json"
	"encoding/xml"
	"path"

	"gopkg.in/h2non/filetype.v0"
)

// NOTE: It may make more sense to use a map to these instead of creating
// potentially unnecessray structs for implementing interfaces on.
type Serializer interface {
	Serialize(interface{}) ([]byte, error)
	Deserialize([]byte, interface{}) error
}

func NewSerializer(identifier string, content []byte) (serializer Serializer, err error) {
	var extension string

	if kind, unknown := filetype.Match(content); err == nil && unknown == nil {
		extension = kind.Extension
	} else {
		extension = path.Ext(identifier)
	}

	switch extension {
	default:
		return JSONSerializer{}, nil
	}
}

type JSONSerializer struct{}

func (this JSONSerializer) Serialize(input interface{}) ([]byte, error) {
	return json.Marshal(input)
}

func (this JSONSerializer) Deserialize(input []byte, obj interface{}) error {
	return json.Unmarshal(input, &obj)
}

type XMLSerializer struct{}

func (this XMLSerializer) Serialize(input interface{}) ([]byte, error) {
	return xml.Marshal(input)
}

func (this XMLSerializer) Deserialize(input []byte, obj interface{}) error {
	return xml.Unmarshal(input, &obj)
}
