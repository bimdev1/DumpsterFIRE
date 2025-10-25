# DumpsterFIRE

DumpsterFIRE provides guided incident response workflows in a terminal user interface built with [Cobra](https://github.com/spf13/cobra) and [Bubble Tea](https://github.com/charmbracelet/bubbletea).

## Prerequisites

- Go 1.21 or later

## Build

Compile the CLI binary:

```bash
go build ./cmd/dumpsterfire
```

## Run

Execute the application directly without building:

```bash
go run ./cmd/dumpsterfire
```

When launched you will see a list of curated workflows. Use the arrow keys to navigate, <kbd>Enter</kbd> to open a workflow, and follow the guided prompts to document your response. Press <kbd>Esc</kbd> to return to the list or <kbd>q</kbd>/<kbd>Ctrl+C</kbd> to exit.

