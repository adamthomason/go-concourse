package main

import (
	"fmt"

	"github.com/adamthomason/go-concourse/jobs"
	"github.com/adamthomason/go-concourse/pipelines"
	"github.com/adamthomason/go-concourse/plans"
	"github.com/adamthomason/go-concourse/plans/task"
)

func main() {
	pipeline := pipelines.New()

	helloWorldPlan := plans.Plan{
		task.New("simple-task", "", ""),
	}

	helloWorldJob := jobs.New("job", &helloWorldPlan)

	pipeline.AttachJob(helloWorldJob)

	fmt.Println(pipeline.Render())
}
