package prefer

import "encoding/json"

type Serializer interface {
	Serialize(obj interface{}) ([]byte, error)
	Deserialize(data []byte, obj interface{}) error
}

func NewSerializer(identifier string) (serializer Serializer, err error) {
	// TODO: Automatic discovery of which loader type to use
	serializer = JSONSerializer{}
	return serializer, err
}

type JSONSerializer struct{}

func (serializer JSONSerializer) Serialize(obj interface{}) ([]byte, error) {
	return json.Marshal(obj)
}

func (serializer JSONSerializer) Deserialize(data []byte, obj interface{}) error {
	return json.Unmarshal(data, &obj)
}
