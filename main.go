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
	for _, job := range config.Jobs {
		fmt.Printf("\nRunning job: %s\n", job.Name)

		// Start services
		for serviceName, service := range job.Services {
			err := executor.StartService(serviceName, service)
			if err != nil {
				log.Fatalf("Failed to start service %s: %v", serviceName, err)
			}
		}

		// Execute steps
		for _, step := range job.Steps {
			if err := executor.ExecuteStep(step); err != nil {
				log.Printf("> Step failed: %s, Error: %v", step.Name, err)
				executor.CleanData()
				log.Fatal("Exiting due to failure.")
			}
		}

		// Stop services
		for serviceName := range job.Services {
			_ = executor.StopService(serviceName)
		}
	}
	
	exitStatus := executor.CleanData()
	if exitStatus != nil {
		fmt.Printf("\nSimulation complete (Error: %s).\n", exitStatus.Error())
	} else {
		fmt.Println("\nSimulation complete (Success).")
	}
}
