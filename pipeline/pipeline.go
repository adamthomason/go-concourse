package pipeline

import (
	"github.com/adamthomason/go-concourse/job"
	"github.com/adamthomason/go-concourse/resource"
	"github.com/adamthomason/go-concourse/resourcetype"
	"gopkg.in/yaml.v2"
)

// Pipeline contains the top-level data which maps
// one-to-one with the transpiled Concourse-native YAML.
type Pipeline struct {
	Jobs          []*job.Job                   `yaml:"jobs"`
	Resources     []*resource.Resource         `yaml:"resources,omitempty"`
	ResourceTypes []*resourcetype.ResourceType `yaml:"resource_types,omitempty"`
}

// New is a helper function to generate a new Pipeline struct and
// returns a pointer.
func New() *Pipeline {
	return &Pipeline{}
}

// AttachJob is a helper function which takes a Job struct pointer
// as an argument and appends it onto the slice of Jobs on the
// Pipline struct.
func (p *Pipeline) AttachJob(job *job.Job) {
	p.Jobs = append(p.Jobs, job)
}

// AttachResource is a helper function which takes a pointer to a
// Resource struct as an argument and appends it to the slice of
// Resources on the Pipeline struct.
func (p *Pipeline) AttachResource(resource *resource.Resource) {
	p.Resources = append(p.Resources, resource)
}

// AttachResourceType is a helper function which takes a pointer to
// a ResourceType struct as an argument and appends it to the slice
// of ResourceTypes on the Pipeline struct.
func (p *Pipeline) AttachResourceType(resourceType *resourcetype.ResourceType) {
	p.ResourceTypes = append(p.ResourceTypes, resourceType)
}

// Render marshals the Pipeline struct into Concourse-native
// YAML data, which is cast to a string.
func (p *Pipeline) Render() string {
	data, err := yaml.Marshal(p)
	if err != nil {
		panic(err)
	}

	return string(data)
}
