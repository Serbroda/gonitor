package tui

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"gonitor/monitors"
	"os"
	"time"
)

type Notes struct {
	current []MonHandler
	sub     chan []MonHandler
}

type MonHandler struct {
	monitor monitors.Monitor
	res     []int
}

func (n Notes) awaitNext() Notes {
	return Notes{current: <-n.sub, sub: n.sub}
}

type model struct {
	currentNotes []MonHandler
}

func (m model) Init() tea.Cmd {
	channelOut := make(chan []MonHandler)

	go func() {
		for {
			time.Sleep(time.Second * 5)
			for _, m := range m.currentNotes {
				ok, _ := m.monitor.Monitor()
				for i := len(m.res) - 1; i > 0; i-- {
					m.res[i] = m.res[i-1]
				}
				if ok {
					m.res[0] = 1
				} else {
					m.res[0] = 0
				}
			}
			channelOut <- m.currentNotes
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
		for _, v2 := range v.res {
			switch v2 {
			case 0:
				str += "▯"
				break
			case 1:
				str += "▮"
				break
			default:
				str += " "
				break
			}
			str += " "
		}
		str += "\n"
	}
	return str
}

func Start(monitors []monitors.Monitor) {
	ms := make([]MonHandler, len(monitors))
	for i, m := range monitors {
		ms[i] = MonHandler{
			monitor: m,
			res:     []int{-1, -1, -1, -1, -1},
		}
	}
	p := tea.NewProgram(model{
		currentNotes: ms,
	}, tea.WithAltScreen())
	if err := p.Start(); err != nil {
		fmt.Println("Error starting app")
		os.Exit(2)
	}
}
