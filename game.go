package main

import (
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"math/rand"
	"strconv"
	"time"
)

const (
	Left  int = 0
	Right int = 1
	Up    int = 2
	Down  int = 3
)

var gameOverWaiting = false
var bestScore = 0
var currentScore = 0
var cells = [][]int{
	{1024, 2048, 512, 1024},
	{2, 1024, 32, 16},
	{8, 4, 2, 64},
	{1024, 1024, 8, 32},
}

var colors = make(map[int]ui.Style)

func initGame() {
	colors[2] = ui.NewStyle(ui.ColorYellow)
	colors[4] = ui.NewStyle(ui.ColorYellow)
	colors[8] = ui.NewStyle(ui.ColorYellow)
	colors[16] = ui.NewStyle(ui.ColorCyan)
	colors[32] = ui.NewStyle(ui.ColorCyan)
	colors[64] = ui.NewStyle(ui.ColorCyan)
	colors[128] = ui.NewStyle(ui.ColorBlue)
	colors[256] = ui.NewStyle(ui.ColorBlue)
	colors[512] = ui.NewStyle(ui.ColorBlue)
	colors[1024] = ui.NewStyle(ui.ColorGreen)
	colors[2048] = ui.NewStyle(ui.ColorGreen)
	colors[-1] = ui.NewStyle(ui.ColorRed)

	rand.Seed(time.Now().UnixNano())

	if !loadGame() {
		addRandomCell()
		addRandomCell()
	}

	renderGame()
}

func move(directory int) {
	if gameOverWaiting {
		return
	}

	if recalculate(directory) {
		addRandomCell()
	}

	if currentScore > bestScore {
		bestScore = currentScore
	}

	go saveGame()
	renderGame()

	if !isMoveAvailable() {
		gameOverWaiting = true
		renderGame()
	}
}

func addRandomCell() {
	for {
		row := rand.Intn(4)
		col := rand.Intn(4)

		if cells[row][col] == 0 {
			if rand.Intn(11) == 10 {
				cells[row][col] = 4
			} else {
				cells[row][col] = 2
			}
			break
		}
	}
}

func renderGame() {
	table := NewColorfulTable()
	table.TextAlignment = ui.AlignCenter
	tableContent := make([][]string, 4)

	for row := 0; row < 4; row++ {
		rowContent := make([]string, 4)
		for col := 0; col < 4; col++ {
			color, exists := colors[cells[row][col]]
			if !exists {
				color = colors[-1]
			}
			table.StyleCell(row, col, color)

			value := strconv.Itoa(cells[row][col])
			if value == "0" {
				value = ""
			}
			rowContent[col] = value
		}
		tableContent[row] = rowContent
	}

	table.Rows = tableContent
	table.SetRect(0, 0, 21, 9)

	scores := widgets.NewParagraph()
	scores.Text = "Best Score: " + strconv.Itoa(bestScore) + "\nCurrent Score: " + strconv.Itoa(currentScore)
	scores.SetRect(27, 0, 50, 4)

	help := widgets.NewParagraph()
	help.Text = "Press ESC to exit\nPress Backspace to undo"
	help.SetRect(27, 4, 54, 8)

	if !gameOverWaiting {
		ui.Render(table, scores, help)
	} else {
		gameOver := widgets.NewParagraph()
		gameOver.Text = "Game Over!\nPress Enter to continue"
		gameOver.SetRect(0, 9, 26, 13)

		ui.Render(table, scores, help, gameOver)
	}
}
