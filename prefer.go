package prefer

import "log"

type filterable func(identifier string) bool

// Configuration represents a specific configuration loaded by Prefer
type Configuration struct {
	identifier  string
	loaders     map[Loader]filterable
	serializers map[Serializer]filterable
}

// Load creates and loads a Configuration object
func Load(identifier string, out interface{}) (*Configuration, error) {
	configuration := &Configuration{identifier: identifier}

	loader, err := NewLoader(configuration.identifier)
	if err != nil {
		return nil, err
	}

	content, err := loader.Load(configuration.identifier)
	if err != nil {
		return nil, err
	}

	serializer, err := NewSerializer(configuration.identifier, content)
	if err != nil {
		log.Println("4")
		return nil, err
	}

	err = serializer.Deserialize(content, out)
	return configuration, err
}
