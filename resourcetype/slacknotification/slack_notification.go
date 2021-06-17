package slacknotification

type SlackNotificationSource struct {
	InsecureRegistries []string `yaml:"insecure_registries,omitempty"`
	Password           string   `yaml:"password,omitempty"`
	Repository         string   `yaml:"repository,omitempty"`
	Username           string   `yaml:"username,omitempty"`
}

func New(insecureRegistries []string, password, repositroy, username string) *SlackNotificationSource {
	return &SlackNotificationSource{
		InsecureRegistries: insecureRegistries,
		Password:           password,
		Repository:         repositroy,
		Username:           username,
	}
}
