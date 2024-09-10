package main

import (
	"fmt"
	"log"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	SPACER = iota
	VACANT
	ACTIVE
)

type board [10][6]int
type coord struct {
	x int
	y int
}

type model struct {
	board  board
	width  int
	height int
	active coord
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width, m.height = msg.Width, msg.Height
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "up", "k":
			if m.board[m.active.x-1][m.active.y] == VACANT {
				m.board[m.active.x][m.active.y] = VACANT
				m.active.x = m.active.x - 1
				m.board[m.active.x][m.active.y] = ACTIVE
			}
		case "down", "j":
			if m.board[m.active.x+1][m.active.y] == VACANT {
				m.board[m.active.x][m.active.y] = VACANT
				m.active.x = m.active.x + 1
				m.board[m.active.x][m.active.y] = ACTIVE
			}
		case "left", "h":
			if m.board[m.active.x][m.active.y-1] == VACANT {
				m.board[m.active.x][m.active.y] = VACANT
				m.active.y = m.active.y - 1
				m.board[m.active.x][m.active.y] = ACTIVE
			}
		case "right", "l":
			if m.board[m.active.x][m.active.y+1] == VACANT {
				m.board[m.active.x][m.active.y] = VACANT
				m.active.y = m.active.y + 1
				m.board[m.active.x][m.active.y] = ACTIVE
			}
		}
	}
	return m, nil
}

func (m model) View() string {
	return m.renderBoard()
}

func (m model) renderBoard() string {
	rows := make([]string, 0)
	for _, row := range m.board {
		cells := make([]string, 0)
		for _, cell := range row {
			cells = append(cells, renderCell(cell))
		}
		rows = append(rows, lipgloss.JoinHorizontal(lipgloss.Center, cells...))
	}

	board := lipgloss.JoinVertical(lipgloss.Center, rows...)
	board = lipgloss.NewStyle().BorderStyle(lipgloss.DoubleBorder()).Render(board)
	board = lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, board)
	return board
}

func renderCell(cell int) string {
	style := lipgloss.NewStyle()

	switch cell {
	case SPACER:
		style = style.BorderStyle(lipgloss.HiddenBorder())
	case VACANT:
		style = style.BorderStyle(lipgloss.NormalBorder())
	case ACTIVE:
		style = style.BorderStyle(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color("208"))
	}

	row := strings.Repeat(" ", 8)
	square := fmt.Sprintf("%s\n%s\n%s", row, row, row)
	return style.Render(square)
}

func newBoard() [10][6]int {
	return [10][6]int{
		{SPACER, SPACER, SPACER, SPACER, SPACER, SPACER},
		{SPACER, SPACER, VACANT, VACANT, SPACER, SPACER},
		{SPACER, VACANT, VACANT, VACANT, SPACER, SPACER},
		{SPACER, VACANT, VACANT, VACANT, VACANT, SPACER},
		{SPACER, VACANT, VACANT, VACANT, VACANT, SPACER},
		{SPACER, VACANT, VACANT, VACANT, VACANT, SPACER},
		{SPACER, VACANT, VACANT, VACANT, VACANT, SPACER},
		{SPACER, SPACER, VACANT, VACANT, VACANT, SPACER},
		{SPACER, SPACER, ACTIVE, VACANT, SPACER, SPACER},
		{SPACER, SPACER, SPACER, SPACER, SPACER, SPACER},
	}
}

func main() {
	// Initialize our program
	p := tea.NewProgram(model{board: newBoard(), active: coord{x: 8, y: 2}}, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
