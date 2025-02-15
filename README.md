# Project Overview

The **Local Actions** is a Go-based application that simulates running YAML workflows locally. The tool is just a concept proof and is not intended to be a full-fledged replacement for GitHub Actions.  

It allows you to define jobs and steps in a YAML file and execute them, along with managing Docker-based services that might be needed during the execution. The simulator mimics the process of starting services, running commands, and cleaning up resources when finished or if any errors occur.

> [!NOTE]
> # TODO
> - [x] Create simple test workflow
> - [x] Run job (`main.go`)
> - [x] Call Docker (compose)
> - [ ] Test / Create tests
> - [ ] Store the results of the running jobs
> - [ ] Create the report with the job information

## Project Structure

- **`main.go`**: The main entry point for the simulator. This file reads a workflow file, starts services, executes steps, and cleans up resources.
- **`parser.go`**: Contains logic for parsing YAML workflow files.
- **`executor.go`**: Handles starting and stopping services, executing steps, and cleaning up Docker resources.

## Requirements

- Docker must be installed on the machine to run the services and perform cleanup.
- Go 1.16 or newer.
- A valid YAML workflow file with jobs, services, and steps.

## Installation

To run the simulator, follow these steps:

1. Clone the repository:
    ```bash
    git clone https://github.com/your-username/gh-actions-simulator.git
    cd gh-actions-simulator
    ```

2. Build the Go application:
    ```bash
    go build -o gh-actions-simulator
    ```

3. Run the simulator with a workflow file:
    ```bash
    ./gh-actions-simulator <workflow-file>
    ```

---

## Components

> I've kept only the function signatures here to avoid polluting the README. The actual implementation can be found in the respective files.

### `main.go`

The **main.go** file orchestrates the simulation process:

1. **Parsing the Workflow**: Reads the workflow YAML file and converts it into a `Workflow` struct.
2. **Starting Services**: For each job in the workflow, it starts the necessary services (Docker containers).
3. **Executing Steps**: Each step in the job is executed, and if any errors occur, the cleanup function is triggered.
4. **Stopping Services**: After the steps are completed (or if any errors occur), the services (Docker containers) are stopped.
5. **Cleanup**: At the end, it ensures that unused Docker containers, images, and volumes are cleaned up.

### `parser.go`

The **parser.go** file contains the `ParseWorkflow` function that reads the YAML workflow file and unmarshals it into Go structs:

```go
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

func ParseWorkflow(file string) (*Workflow, error) {}
```

### `executor.go`

The **executor.go** file contains the logic for starting services (Docker containers), executing steps, and cleaning up resources:

1. **`StartService`**: Starts a Docker container based on the provided service configuration.
2. **`StopService`**: Stops the specified Docker container.
3. **`ExecuteStep`**: Executes a command (step) in a Docker container or locally.
4. **`CleanData`**: Cleans up all unused Docker containers, images, and volumes.

```go
func StartService(name string, service parser.Service) error {}

func StopService(name string) error {}

func ExecuteStep(step parser.Step) error {}

func CleanData() error {}
```

---

## Example Workflow YAML

Here’s an example of a simple workflow YAML file that could be used with the simulator:

```yaml
jobs:
  - name: Test Job
    services:
      postgres:
        image: postgres:latest
        ports:
          - "5432:5432"
    steps:
      - name: Checkout code
        run: echo "Checking out code..."
      - name: Run Tests
        run: go test ./...
        env:
          DB_HOST: localhost
          DB_PORT: 5432
```
