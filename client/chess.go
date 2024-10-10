package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	lipgloss "github.com/charmbracelet/lipgloss"
	"os"
	"rules"
)

type model struct {
	board          [8][8]rune
	availableMoves [8][8]bool
	cursor         [2]int
	selectedPiece  [2]int
	whiteTurn      bool
}

// Foregrounds
var whiteForeground lipgloss.Style
var blackForeground lipgloss.Style

// Backgrounds
var boardBackground [2]lipgloss.Style
var cursorBackground lipgloss.Style
var validCursorBackground lipgloss.Style
var pieceSelectedBackground lipgloss.Style
var availableBackground lipgloss.Style
var invalidBackground lipgloss.Style

func initialModel() model {
	return model{
		board: [8][8]rune{
			{'♖', '♘', '♗', '♔', '♕', '♗', '♘', '♖'},
			{'♙', '♙', '♙', '♙', '♙', '♙', '♙', '♙'},
			{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
			{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
			{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
			{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
			{'♟', '♟', '♟', '♟', '♟', '♟', '♟', '♟'},
			{'♜', '♞', '♝', '♚', '♛', '♝', '♞', '♜'},
		},
		availableMoves: rules.InitializeMovesRepresentation(),
		cursor:         [2]int{0, 0},
		selectedPiece:  [2]int{-1, -1},
		whiteTurn:      true,
	}
}

func (m model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	//oldCursor := m.cursor
	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c", "q":
			return m, tea.Quit

		// The "up" and "k" keys move the cursor up
		case "up", "k":
			m.cursor[1]--
			if m.cursor[1] < 0 {
				m.cursor[1] = 7
			}
			break

		// The "down" and "j" keys move the cursor down
		case "down", "j":
			m.cursor[1] = (m.cursor[1] + 1) % 8
			break

		// The "down" and "j" keys move the cursor down
		case "right", "l":
			m.cursor[0] = (m.cursor[0] + 1) % 8
			break

		// The "down" and "j" keys move the cursor down
		case "left", "h":
			m.cursor[0]--
			if m.cursor[0] < 0 {
				m.cursor[0] = 7
			}
			break

		// The "enter" key and the spacebar (a literal space) toggle
		// the selected state for the item that the cursor is pointing at.
		case "enter", " ":
			piece := m.board[m.cursor[1]][m.cursor[0]]
			if piece != ' ' && rules.IsWhite(piece) == m.whiteTurn {
				m.selectedPiece = m.cursor
				m.availableMoves = rules.GetPossibleMoves(m.selectedPiece, m.board)
			} else if m.selectedPiece[0] >= 0 || m.selectedPiece[1] >= 0 {
				//Move confirmation mode
				if m.availableMoves[m.cursor[1]][m.cursor[0]] {
					m.board[m.cursor[1]][m.cursor[0]] = m.board[m.selectedPiece[1]][m.selectedPiece[0]]
					m.board[m.selectedPiece[1]][m.selectedPiece[0]] = ' '
					m.availableMoves = rules.InitializeMovesRepresentation()
					m.selectedPiece[0] = -1
					m.selectedPiece[1] = -1
					m.whiteTurn = !m.whiteTurn
				}
			}
		}
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}

func (m model) View() string {
	s := ""
	if m.whiteTurn {
		s += "White"
	} else {
		s += "Black"
	}
	s += "'s Turn\n"
	for y, row := range m.board {
		for x, space := range row {
			foreground := lipgloss.NewStyle()
			background := lipgloss.NewStyle()

			var styleOffset int = ((x + y) % 2)
			background = boardBackground[styleOffset]

			if m.cursor[0] == x && m.cursor[1] == y {
				if space == ' ' {
					background = cursorBackground
				} else if rules.IsWhite(space) == m.whiteTurn {
					background = validCursorBackground
				} else {
					background = invalidBackground
				}
			}

			if m.selectedPiece[0] >= 0 && m.selectedPiece[1] >= 0 {
				if m.selectedPiece[0] == x && m.selectedPiece[1] == y {
					background = pieceSelectedBackground
				}
				if m.availableMoves[y][x] {
					if m.cursor[0] == x && m.cursor[1] == y {
						background = validCursorBackground
					} else {
						background = availableBackground
					}
				}
			}

			if rules.IsWhite(space) {
				foreground = whiteForeground
				if !rules.IsPawn(space) {
					space += 6
				}
			} else {
				foreground = blackForeground
			}

			if rules.IsPawn(space) {
				space = '♙'
			}

			style := lipgloss.NewStyle().
				Inherit(background).
				Inherit(foreground)

			s += style.Render(string(space) + " ")
		}
		s += "\n"
	}

	// The footer
	s += "\nPress q to quit.\n"

	// Send the UI for rendering
	return s
}

func main() {
	const padding int = 1
	const dimensions int = 1

	whiteForeground = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FFFFFF"))

	blackForeground = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#000000"))

	boardBackground[0] = lipgloss.NewStyle().
		Bold(true).
		Background(lipgloss.Color("#7D56F4"))

	boardBackground[1] = lipgloss.NewStyle().
		Bold(true).
		Background(lipgloss.Color("#567DF4"))

	cursorBackground = lipgloss.NewStyle().
		Bold(true).
		Background(lipgloss.Color("#888888"))

	validCursorBackground = lipgloss.NewStyle().
		Bold(true).
		Background(lipgloss.Color("#228822"))

	pieceSelectedBackground = lipgloss.NewStyle().
		Bold(true).
		Background(lipgloss.Color("#888888"))

	availableBackground = lipgloss.NewStyle().
		Background(lipgloss.Color("#0000FF"))

	invalidBackground = lipgloss.NewStyle().
		Background(lipgloss.Color("#FF0000"))

	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
