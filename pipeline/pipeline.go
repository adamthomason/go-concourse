package pipeline

import (
	"github.com/adamthomason/go-concourse/job"
	"github.com/adamthomason/go-concourse/resourcetype"
	"gopkg.in/yaml.v2"
)

type Pipeline struct {
	Jobs          []*job.Job                   `yaml:"jobs"`
	ResourceTypes []*resourcetype.ResourceType `yaml:"resource_types"`
}

func New() *Pipeline {
	return &Pipeline{}
}

func (p *Pipeline) AttachJob(job *job.Job) {
	p.Jobs = append(p.Jobs, job)
}

func (p *Pipeline) AttachResourceType(rt *resourcetype.ResourceType) {
	p.ResourceTypes = append(p.ResourceTypes, rt)
}

func (p *Pipeline) Render() string {
	data, err := yaml.Marshal(p)
	if err != nil {
		panic(err)
	}

	return string(data)
}
