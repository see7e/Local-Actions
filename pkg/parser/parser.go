package parser

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type Step struct {
	Name    string   `yaml:"name"`
	Run     string   `yaml:"run"`
	Env     map[string]string `yaml:"env,omitempty"`
	Outputs map[string]string `yaml:"outputs,omitempty"`
}

type Job struct {
	Name  string `yaml:"name"`
	Steps []Step `yaml:"steps"`
}

type Workflow struct {
	Jobs []Job `yaml:"jobs"`
}

func ParseWorkflow(file string) (*Workflow, error) {
	data, err := ioutil.ReadFile(file)
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
