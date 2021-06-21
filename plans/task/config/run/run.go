package run

type Run struct {
	Path string   `yaml:"path"`
	Args []string `yaml:"args"`
}

func New(path string, args []string) *Run {
	return &Run{
		Path: path,
		Args: args,
	}
}
