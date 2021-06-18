package resources

import (
	"github.com/adamthomason/go-concourse/sources"
)

// Resource is a struct to house Concourse resource data.
type Resource struct {
	Icon   string         `yaml:"icon,omitempty"`
	Name   string         `yaml:"name"`
	Source sources.Source `yaml:"source"`
	Type   string         `yaml:"type"`
}

// New is a helper function which returns a pointer to a Resource.
func New(name, typeString, icon string, source sources.Source) *Resource {
	return &Resource{
		Icon:   icon,
		Name:   name,
		Source: source,
		Type:   typeString,
	}
}
