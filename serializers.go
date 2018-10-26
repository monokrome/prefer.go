package prefer

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"path"

	"gopkg.in/h2non/filetype.v0"

	"github.com/go-yaml/yaml"
)

// Serializers is a mapping of file extensions to their related serializers
var serializers map[string]Serializer

// Serializer provides an interface for packing and unpacking data
type Serializer interface {
	Serialize(interface{}) ([]byte, error)
	Deserialize([]byte, interface{}) error
}

// SerializerFor gets a serializer matching the given identifier
func SerializerFor(identifier string, content []byte) (serializer Serializer, err error) {
	var extension string

	kind, unknown := filetype.Match(content)

	// Attempt to detect the Formatter when no identifier is provided
	if unknown == nil && kind.Extension != "unknown" {
		extension = kind.Extension[1:]
	} else {
		extension = path.Ext(identifier)

		if extension != "" {
			extension = extension[1:]
		}
	}

	serializer, ok := serializers[extension]

	if ok != true {
		return nil, errors.New("No Serializer was found for the given file")
	}

	return serializer, nil
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

// Set up serializers
func init() {
	serializers = make(map[string]Serializer)

	serializers["json"] = JSONSerializer{}
	serializers["xml"] = XMLSerializer{}
	serializers["yaml"] = YAMLSerializer{}
	serializers["yml"] = YAMLSerializer{}
}
