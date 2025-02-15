package executor

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"gh-actions-simulator/pkg/parser" // Import parser package
)

// Start a service as a Docker container
func StartService(name string, service parser.Service) error {
	fmt.Println("--------------------------------------------------")
	fmt.Printf("Starting service: %s (Image: %s)\n", name, service.Image)

	args := []string{"run", "-d", "--name", name}
	for _, port := range service.Ports {
		args = append(args, "-p", port)
	}
	args = append(args, service.Image)

	cmd := exec.Command("docker", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// Stop a service container
func StopService(name string) error {
	fmt.Printf("Stopping service: %s\n", name)
	cmd := exec.Command("docker", "stop", name)
	return cmd.Run()
}

// Execute a single step
func ExecuteStep(step parser.Step) error {
	fmt.Println("--------------------------------------------------")
	fmt.Printf("Executing step: %s\n", step.Name)

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

// CleanData removes all stopped containers, unused images, and volumes
func CleanData() error {
	log.Println("Cleaning up Docker data...")

	// Remove stopped containers
	if err := exec.Command("docker", "container", "prune", "-f").Run(); err != nil {
		log.Printf("Failed to clean containers: %v", err)
		return fmt.Errorf("error cleaning containers: %w", err)
	}

	// Remove unused images
	if err := exec.Command("docker", "image", "prune", "-a", "-f").Run(); err != nil {
		log.Printf("Failed to clean images: %v", err)
		return fmt.Errorf("error cleaning images: %w", err)
	}

	// Remove unused volumes
	if err := exec.Command("docker", "volume", "prune", "-f").Run(); err != nil {
		log.Printf("Failed to clean volumes: %v", err)
		return fmt.Errorf("error cleaning volumes: %w", err)
	}

	log.Println("Docker cleanup complete.")
	return nil // Success
}
