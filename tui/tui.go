package tui

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"gonitor/common"
	"gonitor/monitors"
	"math/rand"
	"os"
	"strings"
	"time"
)

var (
	docStyle = lipgloss.NewStyle().Padding(1, 2, 1, 2)
	line     = lipgloss.NewStyle().Background(lipgloss.AdaptiveColor{Light: "#58D384", Dark: "#58D384"})
)

type res struct {
	mon     monitors.Monitor
	name    string
	monType string
	ok      bool
}

type model struct {
	sub   chan any
	items []res
}

func listen(items []res, sub chan any) tea.Cmd {
	return func() tea.Msg {
		for {
			time.Sleep(time.Millisecond * time.Duration(rand.Int63n(900)+100))

			sub <- struct{}{}
		}
	}
}

func initialModel(conf common.Config) model {
	ress := make([]res, len(conf.Monitors))
	for i, v := range conf.Monitors {
		m := monitors.NewMonitor(v.Type, v.Properties)

		ress[i] = res{
			mon:     m,
			name:    v.Name,
			monType: string(v.Type),
			ok:      false,
		}
	}
	return model{
		sub:   make(chan any),
		items: ress,
	}
}

func (m model) Init() tea.Cmd {
	return tea.Batch()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m model) View() string {
	doc := strings.Builder{}

	for _, v := range m.items {
		doc.WriteString(line.Render(fmt.Sprintf("%s: %v", v.name, v.ok)) + "\n")
	}
	return docStyle.Render(doc.String())
}

func Start(conf common.Config) {
	p := tea.NewProgram(initialModel(conf), tea.WithAltScreen())
	if err := p.Start(); err != nil {
		fmt.Println("Error starting app")
		os.Exit(2)
	}
}
