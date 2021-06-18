package task

type Task struct {
	Task   string `yaml:"task"`
	Image  string `yaml:"image"`
	Config string `yaml:"config"`
}

func (t *Task) EnableItem() {}

func New(task, image, config string) *Task {
	return &Task{
		Task:   task,
		Image:  image,
		Config: config,
	}
}
