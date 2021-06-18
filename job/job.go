package job

import (
	"github.com/adamthomason/go-concourse/plan"
)

// Job holds Concourse job specific data, to be attached
// to a Pipeline struct.
type Job struct {
	Name string     `yaml:"name"`
	Plan *plan.Plan `yaml:"plan"`
}

// New is a helper function which takes a pointer to a Plan
// struct as an argument and returns a pointer to a Job struct.
func New(name string, plan *plan.Plan) *Job {
	return &Job{
		Name: name,
		Plan: plan,
	}
}
