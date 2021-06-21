package varsource

var Vault string = "vault"
var Dummy string = "dummy"

type VarSourceConfig interface {
	EnableVarSourceConfig()
}

type VarSource struct {
	Name   string           `yaml:"name"`
	Type   string           `yaml:"type"`
	Config *VarSourceConfig `yaml:"config"`
}

func New(name, sourceType string, config *VarSourceConfig) *VarSource {
	return &VarSource{
		Name:   name,
		Type:   sourceType,
		Config: config,
	}
}
