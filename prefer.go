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

func (configuration *Configuration) Reload(out interface{}) error {
	loader, err := NewLoader(configuration.identifier)
	check(err)

	serializer, err := NewSerializer(configuration.identifier)
	check(err)

	content, err := loader.Load(configuration.identifier)
	check(err)

	err = serializer.Deserialize(content, out)
	return err
}
