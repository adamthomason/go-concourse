package resource

var ICON_DOCKER string = "docker"
var ICON_GITHUB string = "github"

var TYPE_GIT string = "git"
var TYPE_REGISTRY_IMAGE string = "registry-image"

// ResourceSource is a helper type to be used for
// the source field of a Resource struct.
type ResourceSource map[string]string

// Resource is a struct to house Concourse resource data.
type Resource struct {
	Icon   string         `yaml:"icon,omitempty"`
	Name   string         `yaml:"name"`
	Source ResourceSource `yaml:"source"`
	Type   string         `yaml:"type"`
}

// New is a helper function which returns a pointer to a Resource.
func New(name, typeString, icon string, source ResourceSource) *Resource {
	return &Resource{
		Icon:   icon,
		Name:   name,
		Source: source,
		Type:   typeString,
	}
}
