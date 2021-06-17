package job

import (
	"github.com/adamthomason/go-concourse/plan"
)

type Job struct {
	Name string `yaml:"name"`
	Plan *plan.Plan
}

func New(name string) *Job {
	return &Job{
		Name: name,
	}
}

func (job *Job) SetPlan(plan *plan.Plan) {
	job.Plan = plan
}
