package tui

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"os"
	"time"
)

type Notes struct {
	current []string
	sub     chan []string
}

func (n Notes) awaitNext() Notes {
	return Notes{current: <-n.sub, sub: n.sub}
}

type model struct {
	currentNotes []string
}

func (m model) Init() tea.Cmd {
	channelOut := make(chan []string)

	go func() {
		strs := make([]string, 10)
		for {
			time.Sleep(time.Second * 2)
			for i := 0; i < 10; i++ {
				strs[i] = fmt.Sprintf("Hi from %d", i)
			}
			channelOut <- strs
		}
	}()

	return func() tea.Msg {
		return Notes{<-channelOut, channelOut}
	}
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}

	case Notes:
		m.currentNotes = msg.current
		return m, func() tea.Msg {
			return msg.awaitNext()
		}
	}

	return m, nil
}

func (m model) View() string {
	str := ""
	for _, v := range m.currentNotes {
		str = str + v + time.Now().UTC().String() + "\n"
	}
	return str
}

func Start() {
	p := tea.NewProgram(model{}, tea.WithAltScreen())
	if err := p.Start(); err != nil {
		fmt.Println("Error starting app")
		os.Exit(2)
	}
}
