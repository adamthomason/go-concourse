package main

import (
	"fmt"

	"github.com/adamthomason/go-concourse/job"
	"github.com/adamthomason/go-concourse/pipeline"
	"github.com/adamthomason/go-concourse/plan"
	"github.com/adamthomason/go-concourse/plan/get"
	"github.com/adamthomason/go-concourse/plan/task"
	"github.com/adamthomason/go-concourse/resource"
	"github.com/adamthomason/go-concourse/resourcetype"
)

func main() {
	// Initialize Pipeline
	pl := pipeline.New()

	// Define a resource type. In this example, we're creating a Teams notification
	// resource type. This enables us to use a custom Docker image to send out
	// notifications to a Teams channel.
	pl.AttachResourceType(
		resourcetype.New("teams-notification", "docker-image", resourcetype.ResourceTypeSource{
			"repository": "some-docker-repo",
			"username":   "((dockerhub.username))",
			"password":   "((dockerhub.password))",
		}))

	// Next we'll add a resource. In this instance, we're adding a git resource which
	// points to a git repository.
	golangMockGit := resource.New("golang-mock-git", resource.TYPE_GIT, resource.ICON_GITHUB, resource.ResourceSource{
		"uri": "https://github.com/golang/mock.git",
	})
	pl.AttachResource(golangMockGit)

	// We'll also add a registry image resource which will make an image available for
	// our task later down the line.
	golang_1_16_x_image := resource.New("golang-1.16.x-image", resource.TYPE_REGISTRY_IMAGE, resource.ICON_DOCKER, resource.ResourceSource{
		"repository": "golang",
		"tag":        "1.16-stretch",
	})
	pl.AttachResource(golang_1_16_x_image)

	// Initialize and attach a job to the pipeline
	pl.AttachJob(job.New("Build magical Docker image", &plan.Plan{
		get.New(golangMockGit.Name, true),
		get.New(golang_1_16_x_image.Name, true),
		task.New("run-tests", golang_1_16_x_image.Name, "example config"),
	}))

	// Render our pipeline and print it out
	fmt.Println(pl.Render())
}
