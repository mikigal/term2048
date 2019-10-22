package main

import (
	"encoding/json"
	ui "github.com/gizak/termui/v3"
	"github.com/mitchellh/go-homedir"
	"io/ioutil"
	"strings"
)

type Save struct {
	BestScore    int     `json:"BestScore"`
	CurrentScore int     `json:"CurrentScore"`
	Cells        [][]int `json:"Cells"`
}

func loadGame() bool {
	bytes, err := ioutil.ReadFile(getHomeDir() + ".term2048_data.json")
	if err != nil {
		return false
	}

	save := Save{}
	err = json.Unmarshal(bytes, &save)
	if err != nil {
		return false
	}

	bestScore = save.BestScore
	currentScore = save.CurrentScore
	cells = save.Cells
	return true
}

func saveGame() {
	output, _ := json.Marshal(Save{
		BestScore:    bestScore,
		CurrentScore: currentScore,
		Cells:        cells,
	})

	err := ioutil.WriteFile(getHomeDir()+".term2048_data.json", output, 0644)
	check(err)
}

func resetGame() {
	gameOverWaiting = false
	currentScore = 0
	cells = [][]int{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}

	addRandomCell()
	addRandomCell()

	ui.Clear()
	renderGame()

	go saveGame()
}

func getHomeDir() string {
	home, err := homedir.Dir()
	check(err)

	if !strings.HasSuffix(home, "/") {
		home += "/"
	}

	return home
}
