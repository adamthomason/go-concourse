package resourcetypes

import (
	"github.com/adamthomason/go-concourse/sources"
)

// ResourceType holds the data which represents a Concourse
// resource type.
type ResourceType struct {
	Name   string         `yaml:"name"`
	Source sources.Source `yaml:"source"`
	Type   string         `yaml:"type"`
}

// New is a helper method which returns a pointer to a ResourceType struct.
func New(name, typeString string, source sources.Source) *ResourceType {
	return &ResourceType{
		Name:   name,
		Source: source,
		Type:   typeString,
	}
}
