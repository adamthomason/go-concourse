package task

import (
	"github.com/adamthomason/go-concourse/plans/task/config"
)

type Task struct {
	Task   string         `yaml:"task"`
	Config *config.Config `yaml:"config"`
}

func (t *Task) EnableItem() {}

func New(task string, taskConfig *config.Config) *Task {
	return &Task{
		Task:   task,
		Config: taskConfig,
	}
}
