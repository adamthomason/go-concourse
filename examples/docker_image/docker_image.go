package main

import (
	"fmt"

	"github.com/adamthomason/go-concourse/job"
	"github.com/adamthomason/go-concourse/pipeline"
	"github.com/adamthomason/go-concourse/plan"
	"github.com/adamthomason/go-concourse/resourcetype"
	"github.com/adamthomason/go-concourse/resourcetype/slacknotification"
)

func main() {
	// Initialize Pipeline
	pl := pipeline.New()

	// Define a resource type
	slack := resourcetype.New("slack-notification", "docker-image")
	slack.AddSource(slacknotification.New(nil, "password123", "github.com/example", "test123"))
	pl.AttachResourceType(slack)

	// Initialize a plan
	plan := plan.New()

	// Initialize job
	buildJob := job.New("Build magical Docker image")
	buildJob.SetPlan(plan)

	// Attach the job to the pipeline
	pl.AttachJob(buildJob)

	fmt.Println(pl.Render())
}
