package tui

import (
	"fmt"
	"io"
	"strings"

	picker "github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	pickerTitleStyle        = lipgloss.NewStyle().Bold(true)
	pickerItemStyle         = lipgloss.NewStyle().PaddingLeft(2)
	pickerSelectedItemStyle = lipgloss.NewStyle().Bold(true)
)

type itemDelegate struct{}

func (d itemDelegate) Height() int                               { return 1 }
func (d itemDelegate) Spacing() int                              { return 0 }
func (d itemDelegate) Update(_ tea.Msg, _ *picker.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m picker.Model, index int, listItem picker.Item) {
	item, ok := listItem.(fmt.Stringer)
	if !ok {
		return
	}

	str := item.String()

	fn := pickerItemStyle.Render
	if index == m.Index() {
		fn = func(s ...string) string {
			return pickerSelectedItemStyle.Render("> " + strings.Join(s, " "))
		}
	}

	fmt.Fprint(w, fn(str))
}

type pickerModel struct {
	picker picker.Model
	choice picker.Item
}

func (m pickerModel) Init() tea.Cmd {
	return nil
}

func (m pickerModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.picker.SetWidth(msg.Width)
		return m, nil

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "q", "ctrl+c":
			return m, tea.Quit

		case "enter":
			m.choice = m.picker.SelectedItem()
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.picker, cmd = m.picker.Update(msg)
	return m, cmd
}

func (m pickerModel) View() string {
	return m.picker.View()
}

func pick(prompt string, items []picker.Item) picker.Item {
	defaultWidth := 20
	listHeight := len(items) + 4

	l := picker.New(items, itemDelegate{}, defaultWidth, listHeight)
	l.Title = prompt
	l.SetShowHelp(false)
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.Styles.Title = pickerTitleStyle

	pm := pickerModel{picker: l}

	tm, _ := tea.NewProgram(pm).Run()

	pm = tm.(pickerModel)

	return pm.choice
}
