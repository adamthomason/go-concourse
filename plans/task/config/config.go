package config

import (
	"github.com/adamthomason/go-concourse/plans/task/config/imageresource"
	"github.com/adamthomason/go-concourse/plans/task/config/run"
)

type Config struct {
	Platform      string                       `yaml:"platform"`
	ImageResource *imageresource.ImageResource `yaml:"image_resource"`
	Run           *run.Run                     `yaml:"run"`
}

func New(platform string, imageResource *imageresource.ImageResource, run *run.Run) *Config {
	return &Config{
		Platform:      platform,
		ImageResource: imageResource,
		Run:           run,
	}
}
