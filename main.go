package main

import (
	"Proxx/engine"
	"Proxx/ui/console"
	"fmt"
)

func main() {
	runConsoleMiner()
}

func runConsoleMiner() {
	data := engine.NewGameData(console.GetBoardSizeAndBlockHolesCount())
	ui := console.NewUi(data)
	var ask string
	for {
		game := engine.NewGame(data, ui)
		game.Play()

		fmt.Print("New game? [yes/no] ")
		_, _ = fmt.Scan(&ask)
		if ask == "no" {
			break
		}
	}

}
