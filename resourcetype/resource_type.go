package resourcetype

// ResourceTypeSource is a helper type to be used for
// the Source field of a ResourceType.
type ResourceTypeSource map[string]string

// ResourceType holds the data which represents a Concourse
// resource type.
type ResourceType struct {
	Name   string             `yaml:"name"`
	Source ResourceTypeSource `yaml:"source"`
	Type   string             `yaml:"type"`
}

// New is a helper method which returns a pointer to a ResourceType struct.
func New(name, typeString string, source ResourceTypeSource) *ResourceType {
	return &ResourceType{
		Name:   name,
		Source: source,
		Type:   typeString,
	}
}
