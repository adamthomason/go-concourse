package resourcetype

type ResourceType struct {
	Name   string      `yaml:"name"`
	Source interface{} `yaml:"source"`
	Type   string      `yaml:"type"`
}

func New(name, typeString string) *ResourceType {
	return &ResourceType{
		Name: name,
		Type: typeString,
	}
}

func (rt *ResourceType) AddSource(source interface{}) {
	rt.Source = source
}
