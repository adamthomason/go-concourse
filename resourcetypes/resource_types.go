package resourcetypes

import (
	"github.com/adamthomason/go-concourse/source"
)

// ResourceType holds the data which represents a Concourse
// resource type.
type ResourceType struct {
	Name   string        `yaml:"name"`
	Source source.Source `yaml:"source"`
	Type   string        `yaml:"type"`
}

// New is a helper method which returns a pointer to a ResourceType struct.
func New(name, typeString string, source source.Source) *ResourceType {
	return &ResourceType{
		Name:   name,
		Source: source,
		Type:   typeString,
	}
}
