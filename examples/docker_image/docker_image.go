package main

import (
	"fmt"

	"github.com/adamthomason/go-concourse/icons"
	"github.com/adamthomason/go-concourse/job"
	"github.com/adamthomason/go-concourse/pipeline"
	"github.com/adamthomason/go-concourse/plan"
	"github.com/adamthomason/go-concourse/plan/get"
	"github.com/adamthomason/go-concourse/plan/task"
	"github.com/adamthomason/go-concourse/resource"
	"github.com/adamthomason/go-concourse/source"
	"github.com/adamthomason/go-concourse/types"
)

func main() {
	// Initialize Pipeline
	dockerPipeline := pipeline.New()

	// Next we'll add a resource. In this instance, we're adding a git resource which
	// points to a git repository.
	golangMockGit := resource.New("golang-mock-git", types.Git, icons.GitHub, source.Source{
		"uri": "https://github.com/golang/mock.git",
	})
	dockerPipeline.AttachResource(golangMockGit)

	// We'll also add a registry image resource which will make an image available for
	// our task later down the line.
	golang_1_16_x_image := resource.New("golang-1.16.x-image", types.RegistryImage, icons.Docker, source.Source{
		"repository": "golang",
		"tag":        "1.16-stretch",
	})
	dockerPipeline.AttachResource(golang_1_16_x_image)

	// Create a plan to attach to a job
	buildJobPlan := &plan.Plan{
		get.New(golangMockGit.Name, true),
		get.New(golang_1_16_x_image.Name, true),
		task.New("run-tests", golang_1_16_x_image.Name, "example config"),
	}

	// Create the job and attach our plan
	buildJob := job.New("Build magical Docker image", buildJobPlan)

	// Initialize and attach a job to the pipeline
	dockerPipeline.AttachJob(buildJob)

	// Render our pipeline and print it out
	fmt.Println(dockerPipeline.Render())
}
