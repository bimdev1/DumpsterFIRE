package data

import (
	"context"
	"time"
)

// Workflow represents a DumpsterFIRE guided workflow definition.
type Workflow struct {
	Title       string
	Description string
	Steps       []Step
}

// Step captures a single guided action a responder should take.
type Step struct {
	Title       string
	Prompt      string
	Placeholder string
}

// Repository exposes methods for retrieving workflows from storage.
type Repository struct{}

// NewRepository creates a repository instance. In a real application this could
// perform IO such as opening a database or API client. The constructor exists so
// the UI can depend on an interface instead of concrete data.
func NewRepository() *Repository {
	return &Repository{}
}

// LoadWorkflows simulates retrieving workflows from persistent storage.
func (r *Repository) LoadWorkflows(ctx context.Context) ([]Workflow, error) {
	// Simulate latency so the UI can display a loading spinner.
	select {
	case <-time.After(600 * time.Millisecond):
	case <-ctx.Done():
		return nil, ctx.Err()
	}

	return []Workflow{
		{
			Title:       "Ransomware Triage",
			Description: "Identify scope, contain systems, and gather intel.",
			Steps: []Step{
				{Title: "Identify", Prompt: "Impacted hostnames", Placeholder: "host1, host2"},
				{Title: "Contain", Prompt: "Containment actions taken", Placeholder: "Isolated VLAN"},
				{Title: "Collect", Prompt: "Artifacts collected", Placeholder: "Memory dump, logs"},
			},
		},
		{
			Title:       "Phishing Investigation",
			Description: "Validate, scope, and respond to reported phishing messages.",
			Steps: []Step{
				{Title: "Validate", Prompt: "Reporter and channel", Placeholder: "Jane Doe via email"},
				{Title: "Scope", Prompt: "Recipients", Placeholder: "Distribution list"},
				{Title: "Respond", Prompt: "Response actions", Placeholder: "Blocked sender"},
			},
		},
		{
			Title:       "Credential Theft",
			Description: "Reset credentials and evaluate blast radius.",
			Steps: []Step{
				{Title: "Reset", Prompt: "Accounts reset", Placeholder: "user@example.com"},
				{Title: "Investigate", Prompt: "Suspicious activity", Placeholder: "Unusual logins"},
				{Title: "Notify", Prompt: "Stakeholders notified", Placeholder: "Security leadership"},
			},
		},
	}, nil
}
