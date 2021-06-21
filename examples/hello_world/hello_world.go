package main

import (
	"fmt"

	"github.com/adamthomason/go-concourse/constants/platforms"
	"github.com/adamthomason/go-concourse/constants/types"
	"github.com/adamthomason/go-concourse/jobs"
	"github.com/adamthomason/go-concourse/pipelines"
	"github.com/adamthomason/go-concourse/plans"
	"github.com/adamthomason/go-concourse/plans/task"
	"github.com/adamthomason/go-concourse/plans/task/config"
	"github.com/adamthomason/go-concourse/plans/task/config/imageresource"
	"github.com/adamthomason/go-concourse/plans/task/config/run"
	"github.com/adamthomason/go-concourse/source"
)

func main() {
	pipeline := pipelines.New()

	simpleTask := task.New("simple-task", config.New(
		platforms.Linux,
		imageresource.New(types.RegistryImage, &source.Source{"repository": "busybox"}),
		run.New("echo", []string{"Hello World!"}),
	))

	helloWorldPlan := &plans.Plan{
		simpleTask,
	}

	pipeline.AttachJob(jobs.New("job", helloWorldPlan))

	fmt.Println(pipeline.Render())
}
