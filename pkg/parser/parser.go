package parser

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Step struct {
	Name    string            `yaml:"name"`
	Run     string            `yaml:"run"`
	Env     map[string]string `yaml:"env,omitempty"`
	Outputs map[string]string `yaml:"outputs,omitempty"`
}

type Service struct {
	Image string   `yaml:"image"`
	Ports []string `yaml:"ports,omitempty"`
}

type Job struct {
	Name     string             `yaml:"name"`
	Steps    []Step             `yaml:"steps"`
	Services map[string]Service `yaml:"services,omitempty"`
}

type Workflow struct {
	Jobs []Job `yaml:"jobs"`
}

func ParseWorkflow(file string) (*Workflow, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	var workflow Workflow
	err = yaml.Unmarshal(data, &workflow)
	if err != nil {
		return nil, err
	}

	return &workflow, nil
}
