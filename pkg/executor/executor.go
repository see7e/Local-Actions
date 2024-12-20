package executor

import (
	"fmt"
	"os"
	"os/exec"
)

type Step struct {
	Name string
	Run  string
	Env  map[string]string
}

func ExecuteStep(step Step) error {
	cmd := exec.Command("bash", "-c", step.Run)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Set environment variables
	env := os.Environ()
	for k, v := range step.Env {
		env = append(env, fmt.Sprintf("%s=%s", k, v))
	}
	cmd.Env = env

	return cmd.Run()
}
