package get

type Get struct {
	Get     string `yaml:"get"`
	Trigger bool   `yaml:"trigger"`
}

func (g *Get) EnableItem() {}

func New(get string, trigger bool) *Get {
	return &Get{
		Get:     get,
		Trigger: trigger,
	}
}
