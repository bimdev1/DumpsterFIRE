# DumpsterFIRE

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE) [![Go Reference](https://pkg.go.dev/badge/github.com/charmbracelet/bubbletea.svg)](https://pkg.go.dev/github.com/charmbracelet/bubbletea)

DumpsterFIRE is a terminal user interface (TUI) application that helps incident response teams triage, document, and coordinate their work from the command line. Built on the Charmbracelet ecosystem—Bubble Tea for the Elm-inspired architecture, Bubbles for reusable components, and Lip Gloss for styling—it delivers a responsive and accessible workflow that mirrors the design review guidance for the project.

## Table of Contents
- [Project Goals](#project-goals)
- [Architecture Overview](#architecture-overview)
  - [Bubble Tea Model / Update / View](#bubble-tea-model--update--view)
  - [Bubbles Components](#bubbles-components)
  - [Lip Gloss Styling](#lip-gloss-styling)
- [Setup Requirements](#setup-requirements)
- [Usage](#usage)
  - [Keyboard Navigation](#keyboard-navigation)
  - [Workflow Modes](#workflow-modes)
  - [Cobra Integration](#cobra-integration)
- [Resources](#resources)

## Project Goals
- Provide a streamlined TUI for incident responders to log findings, switch between tasks, and coordinate evidence collection without leaving the keyboard.
- Offer a predictable architecture so contributors can extend existing workflows or add new modules quickly.
- Maintain parity with the design review documentation so the README doubles as a quick-start guide for new team members.

## Architecture Overview
DumpsterFIRE follows the typical Bubble Tea architecture with clear separation between state, events, and rendering.

### Bubble Tea Model / Update / View
- **Model**: Centralizes application state including active workflow mode, selected incident items, filters, and command palette visibility.
- **Update**: Handles incoming `tea.Msg` values from user input, timers, or asynchronous tasks. Each message transitions the model deterministically so state changes remain explicit and testable.
- **View**: Produces the rendered TUI based on the current model. Views compose Bubbles components and apply Lip Gloss styles to ensure consistent layout, padding, and color usage.

### Bubbles Components
DumpsterFIRE leverages Charmbracelet Bubbles to deliver ergonomic widgets:
- **List** bubbles present incident queues, evidence collections, and timeline views.
- **Text input** components capture command palette entries, search queries, and incident notes.
- **Viewport** and **table** bubbles render detailed artifacts or structured timelines.
Each Bubble is wrapped with thin adapters so we can reuse them across workflow modes while keeping update logic minimal.

### Lip Gloss Styling
Lip Gloss controls theming across the application. Styles define:
- Color palettes for severity levels (e.g., critical red, investigation blue, remediation green).
- Spacing and padding to keep lists and detail panes legible even on small terminals.
- Typography choices such as bold headers and subtle borders to separate panes.
Styles live alongside Bubble components so design tweaks remain localized.

## Setup Requirements
1. Install [Go 1.21+](https://go.dev/dl/).
2. Fetch project dependencies:
   ```bash
   go mod download
   ```
3. (Optional) Install [direnv](https://direnv.net/) or your preferred environment manager to preload credentials or API keys when running DumpsterFIRE.

## Usage
Run DumpsterFIRE from the project root:
```bash
go run ./cmd/dumpsterfire
```
If you installed the CLI via `go install`, invoke it as `dumpsterfire` from anywhere in your `$PATH`.

### Keyboard Navigation
- `Tab` / `Shift+Tab`: cycle focus through interactive panes (queues, detail views, palette).
- Arrow keys or `j`/`k`: move selection within lists.
- `Enter`: open the selected item or confirm palette commands.
- `Esc`: close dialogs or return to the parent mode.
- `Ctrl+C` or `q`: quit the application gracefully.
Shortcuts align with the design review so responders can operate the UI without touching the mouse.

### Workflow Modes
DumpsterFIRE exposes multiple modes tailored to incident lifecycle stages:
- **Triage Mode**: prioritize inbound alerts, assign responders, and tag follow-up actions.
- **Investigation Mode**: inspect evidence timelines, annotate findings, and link related incidents.
- **Remediation Mode**: track runbooks, update status fields, and broadcast communications.
Mode switching occurs via the command palette (`:`) or keyboard shortcuts (e.g., `g t` for triage), and each mode configures the layout, filters, and component bindings relevant to its tasks.

### Cobra Integration
The CLI wraps the Bubble Tea program with [Cobra](https://github.com/spf13/cobra) to deliver structured commands and flags:
- `dumpsterfire` launches the interactive TUI.
- `dumpsterfire export --format=json` emits incident data for external tooling.
- Subcommands like `workflow` or `config` manage profiles, API tokens, and automation hooks.
Cobra's command tree mirrors the TUI's modes, enabling automation scripts and CI pipelines to trigger the same workflows available in the interface.

## Resources
- [Bubble Tea Documentation](https://github.com/charmbracelet/bubbletea)
- [Bubbles Components](https://github.com/charmbracelet/bubbles)
- [Lip Gloss Styling](https://github.com/charmbracelet/lipgloss)
- [Cobra CLI Framework](https://github.com/spf13/cobra)
- [Charmbracelet Community Discord](https://charm.sh/chat)

Contributions, feedback, and feature ideas are welcome! Open an issue or submit a pull request to help keep DumpsterFIRE blazing ahead.
