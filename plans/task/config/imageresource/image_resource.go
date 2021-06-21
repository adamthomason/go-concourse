package imageresource

import (
	"github.com/adamthomason/go-concourse/source"
)

type ImageResource struct {
	Type   string         `yaml:"type"`
	Source *source.Source `yaml:"source"`
}

func New(typeString string, source *source.Source) *ImageResource {
	return &ImageResource{
		Type:   typeString,
		Source: source,
	}
}
