package ui

import (
	"context"
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/DumpsterFIRE/dumpsterfire/internal/data"
)

type stage int

const (
	stageLoading stage = iota
	stageList
	stageWorkflow
	stageSummary
)

type workflowsLoadedMsg struct {
	workflows []data.Workflow
	err       error
}

type Model struct {
	ctx       context.Context
	repo      *data.Repository
	stage     stage
	spinner   spinner.Model
	list      list.Model
	input     textinput.Model
	workflows []data.Workflow

	selectedWorkflow data.Workflow
	responses        map[int]string
	currentStep      int

	width  int
	height int

	err error
}

type workflowItem struct {
	workflow data.Workflow
}

func (w workflowItem) Title() string       { return w.workflow.Title }
func (w workflowItem) Description() string { return w.workflow.Description }
func (w workflowItem) FilterValue() string { return strings.ToLower(w.workflow.Title) }

// NewModel constructs the Bubble Tea model for the DumpsterFIRE CLI UI.
func NewModel(ctx context.Context) Model {
	if ctx == nil {
		ctx = context.Background()
	}

	repo := data.NewRepository()
	sp := spinner.New()
	sp.Spinner = spinner.Dot
	sp.Style = subtitleStyle

	delegate := list.NewDefaultDelegate()
	delegate.ShowDescription = true
	lst := list.New([]list.Item{}, delegate, 0, 0)
	lst.Title = "Guided Workflows"
	lst.SetShowFilter(false)
	lst.SetShowHelp(false)
	lst.SetStatusBarItemName("workflow", "workflows")

	input := textinput.New()
	input.Prompt = "> "
	input.Placeholder = "Document your response"
	input.Focus()

	return Model{
		ctx:       ctx,
		repo:      repo,
		stage:     stageLoading,
		spinner:   sp,
		list:      lst,
		input:     input,
		responses: make(map[int]string),
	}
}

func loadWorkflowsCmd(ctx context.Context, repo *data.Repository) tea.Cmd {
	return func() tea.Msg {
		workflows, err := repo.LoadWorkflows(ctx)
		return workflowsLoadedMsg{workflows: workflows, err: err}
	}
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(m.spinner.Tick, loadWorkflowsCmd(m.ctx, m.repo))
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.list.SetSize(m.width-4, m.height-8)

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "esc":
			if m.stage == stageWorkflow || m.stage == stageSummary {
				m.stage = stageList
				m.selectedWorkflow = data.Workflow{}
				m.responses = make(map[int]string)
				m.currentStep = 0
			}
		case "enter":
			switch m.stage {
			case stageList:
				if item, ok := m.list.SelectedItem().(workflowItem); ok {
					m.selectedWorkflow = item.workflow
					m.stage = stageWorkflow
					m.currentStep = 0
					m.responses = make(map[int]string)
					if len(item.workflow.Steps) > 0 {
						step := item.workflow.Steps[0]
						m.input.Placeholder = step.Placeholder
						m.input.SetValue("")
						m.input.Focus()
					}
				}
			case stageWorkflow:
				if m.selectedWorkflow.Steps != nil {
					m.responses[m.currentStep] = strings.TrimSpace(m.input.Value())
					m.currentStep++
					if m.currentStep >= len(m.selectedWorkflow.Steps) {
						m.stage = stageSummary
						m.input.Blur()
					} else {
						next := m.selectedWorkflow.Steps[m.currentStep]
						m.input.SetValue("")
						m.input.Placeholder = next.Placeholder
						m.input.Focus()
					}
				}
			case stageSummary:
				m.stage = stageList
				m.selectedWorkflow = data.Workflow{}
				m.responses = make(map[int]string)
				m.currentStep = 0
				m.input.SetValue("")
				m.input.Placeholder = "Document your response"
				m.input.Blur()
			}
		}

	case workflowsLoadedMsg:
		if msg.err != nil {
			m.err = msg.err
			m.stage = stageSummary
			break
		}
		m.workflows = msg.workflows
		items := make([]list.Item, 0, len(msg.workflows))
		for _, wf := range msg.workflows {
			items = append(items, workflowItem{workflow: wf})
		}
		m.list.SetItems(items)
		if len(items) > 0 {
			m.stage = stageList
		} else {
			m.stage = stageSummary
		}
	}

	switch m.stage {
	case stageLoading:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		cmds = append(cmds, cmd)
	case stageList:
		var cmd tea.Cmd
		m.list, cmd = m.list.Update(msg)
		cmds = append(cmds, cmd)
	case stageWorkflow:
		var cmd tea.Cmd
		m.input, cmd = m.input.Update(msg)
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	switch m.stage {
	case stageLoading:
		return appStyle.Render(fmt.Sprintf("%s Loading DumpsterFIRE workflows...", m.spinner.View()))
	case stageList:
		if len(m.workflows) == 0 {
			return appStyle.Render("No workflows available. Press q to quit.")
		}
		return appStyle.Render(strings.Join([]string{
			titleStyle.Render("DumpsterFIRE"),
			subtitleStyle.Render("Select a workflow to begin guided response."),
			m.list.View(),
			helpStyle.Render("↑/↓ to navigate • enter to open • esc to cancel • q to quit"),
		}, "\n\n"))
	case stageWorkflow:
		if len(m.selectedWorkflow.Steps) == 0 {
			return appStyle.Render("Workflow has no steps. Press esc to return.")
		}
		step := m.selectedWorkflow.Steps[m.currentStep]
		progress := fmt.Sprintf("Step %d of %d", m.currentStep+1, len(m.selectedWorkflow.Steps))
		body := lipgloss.JoinVertical(lipgloss.Left,
			titleStyle.Render(m.selectedWorkflow.Title),
			subtitleStyle.Render(m.selectedWorkflow.Description),
			sectionTitleStyle.Render(progress+": "+step.Title),
			subtitleStyle.Render(step.Prompt),
			m.input.View(),
			helpStyle.Render("enter to continue • esc to cancel • q to quit"),
		)
		return appStyle.Render(body)
	case stageSummary:
		if m.err != nil {
			return appStyle.Render(errorStyle.Render(fmt.Sprintf("Error loading workflows: %v", m.err)))
		}
		if m.selectedWorkflow.Title == "" {
			return appStyle.Render(strings.Join([]string{
				titleStyle.Render("DumpsterFIRE"),
				subtitleStyle.Render("Press enter to return to the workflows list or q to quit."),
			}, "\n\n"))
		}
		var b strings.Builder
		b.WriteString(titleStyle.Render(m.selectedWorkflow.Title))
		b.WriteString("\n")
		b.WriteString(subtitleStyle.Render("Captured responses:"))
		b.WriteString("\n\n")
		for i, step := range m.selectedWorkflow.Steps {
			response := strings.TrimSpace(m.responses[i])
			if response == "" {
				response = "(no response recorded)"
			}
			b.WriteString(sectionTitleStyle.Render(step.Title))
			b.WriteString("\n")
			b.WriteString(response)
			b.WriteString("\n\n")
		}
		b.WriteString(helpStyle.Render("enter to return • esc to cancel • q to quit"))
		return appStyle.Render(b.String())
	default:
		return appStyle.Render("Press q to quit.")
	}
}
