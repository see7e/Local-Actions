package main

import (
	"fmt"
	"log"
	"os"

	"gh-actions-simulator/pkg/executor"
	"gh-actions-simulator/pkg/parser"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: gh-actions-simulator <workflow-file>")
	}

	workflowFile := os.Args[1]
	config, err := parser.ParseWorkflow(workflowFile)
	if err != nil {
		log.Fatalf("Failed to parse workflow: %v", err)
	}

	fmt.Println("Starting GitHub Actions Simulation...")
	for jobName, job := range config.Jobs {
		fmt.Printf("\nRunning job: %s\n", jobName)
		for serviceName, service := range job.Services {
			fmt.Printf("Starting service: %s (Image: %s)\n", serviceName, service.Image)
			// Simulate service startup (implement actual logic if needed)
		}
		for _, step := range job.Steps {
			fmt.Printf("Running step: %s\n", step.Name)
			if err := executor.ExecuteStep(step); err != nil {
				log.Fatalf("Step failed: %s, Error: %v", step.Name, err)
			}
		}
	}
	fmt.Println("\nSimulation complete.")
}
