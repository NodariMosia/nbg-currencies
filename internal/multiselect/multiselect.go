package multiselect

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"

	"nbg-currencies/internal/set"
)

type viewSettings struct {
	header   string
	showHelp bool
}

func newViewSettings(header string, showHelp bool) viewSettings {
	return viewSettings{header, showHelp}
}

type appState struct {
	cursor            int
	options           []string
	selectedOptionIds set.Set[int]
	viewSettings      viewSettings
	submittedOnQuit   bool
}

func newAppStateWithOptions(
	options []string, viewSettings viewSettings,
) appState {
	return appState{
		cursor:            0,
		options:           options,
		selectedOptionIds: set.NewSet[int](),
		viewSettings:      viewSettings,
		submittedOnQuit:   false,
	}
}

func (state appState) Init() tea.Cmd {
	return nil
}

func (state appState) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc", "q":
			state.cursor = -1
			state.viewSettings.showHelp = false
			state.selectedOptionIds.Clear()

			return state, tea.Quit

		case "enter":
			state.cursor = -1
			state.viewSettings.showHelp = false
			state.submittedOnQuit = true

			return state, tea.Quit

		case " ":
			if state.selectedOptionIds.Contains(state.cursor) {
				state.selectedOptionIds.Delete(state.cursor)
			} else {
				state.selectedOptionIds.Add(state.cursor)
			}

		case "up":
			if state.cursor > 0 {
				state.cursor--
			}

		case "down":
			if state.cursor < len(state.options)-1 {
				state.cursor++
			}
		}
	}

	return state, nil
}

func (state appState) View() string {
	s := strings.Builder{}

	s.WriteString(state.viewSettings.header)
	s.WriteString("\n")

	for i, option := range state.options {
		if state.cursor == i {
			s.WriteString("> ")
		} else {
			s.WriteString("  ")
		}

		if state.selectedOptionIds.Contains(i) {
			s.WriteString("[•] ")
		} else {
			s.WriteString("[ ] ")
		}

		s.WriteString(option)
		s.WriteString("\n")
	}

	s.WriteString("\n")

	if state.viewSettings.showHelp {
		s.WriteString("up/down (arrow keys): navigate;\n")
		s.WriteString("space: select/unselect;\n")
		s.WriteString("enter: confirm selection;\n")
		s.WriteString("q or esc: quit.\n\n")
	}

	return s.String()
}

func PromptMultiselect(
	heading string, options []string,
) (set.Set[int], bool, error) {
	viewSettings := newViewSettings(heading, true)
	initialState := newAppStateWithOptions(options, viewSettings)
	teaApp := tea.NewProgram(initialState)

	teaModel, err := teaApp.Run()
	if err != nil {
		return nil, false, err
	}

	state, ok := teaModel.(appState)
	if !ok {
		return nil, false, fmt.Errorf("unexpected model type: %T", state)
	}

	return state.selectedOptionIds, state.submittedOnQuit, nil
}
