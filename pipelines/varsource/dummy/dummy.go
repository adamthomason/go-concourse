package dummy

import (
	"github.com/adamthomason/go-concourse/vars"
)

type DummyConfig struct {
	Vars *vars.Vars `yaml:"vars"`
}

func (dc *DummyConfig) EnableVarSourceConfig() {}
