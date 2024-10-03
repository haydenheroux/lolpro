package tui

import (
	"fmt"

	input "github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

var (
	inputTitleStyle = pickerTitleStyle.PaddingLeft(2)
)

type inputModel struct {
	input  input.Model
	prompt string
}

func (m inputModel) Init() tea.Cmd {
	return input.Blink
}

func (m inputModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter, tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		}
	}

	m.input, cmd = m.input.Update(msg)
	return m, cmd
}

func (m inputModel) View() string {
	return fmt.Sprintf(
		"%s\n\n%s\n",
		inputTitleStyle.Render(m.prompt),
		m.input.View(),
	) + "\n"
}

func ask(prompt, placeholder string) string {
	input := input.New()
	input.Placeholder = placeholder
	input.Focus()
	input.CharLimit = 156
	input.Width = 20

	im := inputModel{input, prompt}

	tm, _ := tea.NewProgram(im).Run()

	im = tm.(inputModel)

	return im.input.Value()
}
