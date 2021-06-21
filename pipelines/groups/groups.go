package groups

type GroupConfig struct {
	Name string   `yaml:"name"`
	Jobs []string `yaml:"jobs"`
}

func New(name string, jobs []string) *GroupConfig {
	return &GroupConfig{
		Name: name,
		Jobs: jobs,
	}
}
