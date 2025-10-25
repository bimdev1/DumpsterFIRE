package ui

import "github.com/charmbracelet/lipgloss"

var (
	appStyle = lipgloss.NewStyle().Padding(1, 2)

	titleStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("205"))

	subtitleStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))

	sectionTitleStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("63"))

	errorStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("9")).Bold(true)

	helpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("241")).Italic(true)
)
