package prefer

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"path"

	"gopkg.in/h2non/filetype.v0"

	"github.com/go-yaml/yaml"
)

// NOTE: It may make more sense to use a map to these instead of creating
// potentially unnecessray structs for implementing interfaces on.
type Serializer interface {
	Serialize(interface{}) ([]byte, error)
	Deserialize([]byte, interface{}) error
}

func NewSerializer(identifier string, content []byte) (serializer Serializer, err error) {
	var extension string

	if kind, unknown := filetype.Match(content); err == nil && unknown == nil && kind.Extension != "unknown" {
		extension = kind.Extension
	} else {
		extension = path.Ext(identifier)
	}

	switch extension {
	case ".xml":
		return XMLSerializer{}, nil
	case ".json":
		return JSONSerializer{}, nil
	case ".yaml":
		return YAMLSerializer{}, nil
	case ".yml":
		return YAMLSerializer{}, nil
	default:
		return nil, errors.New("No matching serializer for " + identifier)
	}
}

// JSONSerializer serializers to/from JSON format
type JSONSerializer struct{}

// Serialize receives an object and returns a serialized version of it
func (serializer JSONSerializer) Serialize(input interface{}) ([]byte, error) {
	return json.Marshal(input)
}

// Deserialize receives []byte and fills obj with it's deserialized values
func (serializer JSONSerializer) Deserialize(input []byte, obj interface{}) error {
	return json.Unmarshal(input, &obj)
}

// XMLSerializer serializers to/from XML format
type XMLSerializer struct{}

// Serialize receives an object and returns a serialized version of it
func (serializer XMLSerializer) Serialize(input interface{}) ([]byte, error) {
	return xml.Marshal(input)
}

// Deserialize receives []byte and fills obj with it's deserialized values
func (serializer XMLSerializer) Deserialize(input []byte, obj interface{}) error {
	return xml.Unmarshal(input, &obj)
}

// YAMLSerializer serializers to/from YAML format
type YAMLSerializer struct{}

// Serialize receives an object and returns a serialized version of it
func (serializer YAMLSerializer) Serialize(input interface{}) ([]byte, error) {
	return yaml.Marshal(input)
}

// Deserialize receives []byte and fills obj with it's deserialized values
func (serializer YAMLSerializer) Deserialize(input []byte, obj interface{}) error {
	return yaml.Unmarshal(input, &obj)
}
