package jobs

import (
	"github.com/adamthomason/go-concourse/plans"
)

// Job holds Concourse job specific data, to be attached
// to a Pipeline struct.
type Job struct {
	Name string      `yaml:"name"`
	Plan *plans.Plan `yaml:"plan"`
}

// New is a helper function which takes a pointer to a Plan
// struct as an argument and returns a pointer to a Job struct.
func New(name string, plan *plans.Plan) *Job {
	return &Job{
		Name: name,
		Plan: plan,
	}
}
