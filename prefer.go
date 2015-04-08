package prefer

type filterable func(identifier string) bool

type Configuration struct {
	identifier  string
	loaders     map[Loader]filterable
	serializers map[Serializer]filterable
}

func NewConfiguration(identifier string) *Configuration {
	configuration := Configuration{
		identifier: identifier,
	}

	return &configuration
}

func (configuration *Configuration) Load(out interface{}) error {
	loader, err := NewLoader(configuration.identifier)
	checkError(err)

	serializer, err := NewSerializer(configuration.identifier)
	checkError(err)

	content, err := loader.Load(configuration.identifier)
	checkError(err)

	err = serializer.Deserialize(content, out)
	return err
}
