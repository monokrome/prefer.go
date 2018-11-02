package prefer

import "fmt"

type filterable func(identifier string) bool

// Configuration represents a specific configuration loaded by Prefer
type Configuration struct {
	identifier  string
	loaders     map[Loader]filterable
	serializers map[Serializer]filterable
}

// Load creates and loads a Configuration object
func Load(identifier string, options *interface{}, out interface{}) (*Configuration, error) {
	configuration := &Configuration{}

	loader, err := NewLoader(identifier)
	if err != nil {
		return nil, err
	}

	// When inexact identifiers are provided, this updates identifier with a
	// best-guess existing match for the given identifier
	if configuration.identifier, err = loader.Discover(identifier); err != nil {
		return nil, fmt.Errorf("can not find matching configuration for %v", identifier)
	}

	content, identifier, err := loader.Load(configuration.identifier)
	if err != nil {
		return nil, err
	}

	serializer, err := SerializerFor(configuration.identifier, content)
	if err != nil {
		return nil, err
	}

	if err = serializer.Deserialize(content, out); err != nil {
		return nil, err
	}

	return configuration, nil
}

// Identifier returns the identifier for the loaded configuration
func (configuration *Configuration) Identifier() string {
	return configuration.identifier
}
