package main

import (
	"fmt"

	"github.com/adamthomason/go-concourse/icons"
	"github.com/adamthomason/go-concourse/jobs"
	"github.com/adamthomason/go-concourse/pipelines"
	"github.com/adamthomason/go-concourse/plans"
	"github.com/adamthomason/go-concourse/plans/get"
	"github.com/adamthomason/go-concourse/plans/task"
	"github.com/adamthomason/go-concourse/resources"
	"github.com/adamthomason/go-concourse/sources"
	"github.com/adamthomason/go-concourse/types"
)

func main() {
	// Initialize Pipeline
	pipeline := pipelines.New()

	// Next we'll add a resource. In this instance, we're adding a git resource which
	// points to a git repository.
	golangMockGit := resources.New("golang-mock-git", types.Git, icons.GitHub, sources.Source{
		"uri": "https://github.com/golang/mock.git",
	})
	pipeline.AttachResource(golangMockGit)

	// We'll also add a registry image resource which will make an image available for
	// our task later down the line.
	golang_1_16_x_image := resources.New("golang-1.16.x-image", types.RegistryImage, icons.Docker, sources.Source{
		"repository": "golang",
		"tag":        "1.16-stretch",
	})
	pipeline.AttachResource(golang_1_16_x_image)

	// Create a plan to attach to a job
	buildJobPlan := &plans.Plan{
		get.New(golangMockGit.Name, true),
		get.New(golang_1_16_x_image.Name, true),
		task.New("run-tests", golang_1_16_x_image.Name, "example config"),
	}

	// Create the job and attach our plan
	buildJob := jobs.New("Build magical Docker image", buildJobPlan)

	// Initialize and attach a job to the pipeline
	pipeline.AttachJob(buildJob)

	// Render our pipeline and print it out
	fmt.Println(pipeline.Render())
}
