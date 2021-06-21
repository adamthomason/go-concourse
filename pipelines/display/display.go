package display

type DisplayConfig struct {
	BackgroundImage string `yaml:"background_image"`
}

func New(backgroundImage string) *DisplayConfig {
	return &DisplayConfig{
		BackgroundImage: backgroundImage,
	}
}
