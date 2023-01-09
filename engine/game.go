package engine

import (
	"fmt"
	"math/rand"
)

type GameUi interface {
	GetCoordinations() (row, column uint8, err error)
	DisplayErrorMessage(msg string)
	DisplaySuccessMessage(msg string)
	Render(isHideNotOpened bool)
}

type Game struct {
	Data *GameData
	Ui   GameUi
}

func NewGame(data *GameData, ui GameUi) *Game {
	data.init()
	return &Game{data, ui}
}

const gameWinMessage = "You're win!"
const gameFailMessage = "You're failed, try again."

func (g *Game) getRandomCoordination() (row, column uint8) {
	for {
		max := int32(g.Data.size)
		row, column = uint8(rand.Int31n(max)), uint8(rand.Int31n(max))
		cell := &g.Data.cells[row][column]
		if cell.isOpen == false && cell.isBlackHole == false {
			return
		}
	}
}

/**
add black holes by random locations and increment adjacent cells with count of black holes near
*/
func (g *Game) addBlackHoles() {
	for count := uint8(0); count < g.Data.blackHolesCount; count++ {
		row, column := g.getRandomCoordination()
		g.Data.cells[row][column].isBlackHole = true
		g.updateAdjacentCells(row, column, func(rowItem, columnItem uint8) {
			g.Data.cells[rowItem][columnItem].adjacentBlackHolesCount++
		})
	}
}

func (g *Game) updateAdjacentCells(row, column uint8, cellCallback func(rowItem, columnItem uint8)) {
	//adjacent cells locations
	adjacents := map[uint8][]uint8{
		row:     {column - 1, column + 1},
		row - 1: {column, column - 1, column + 1},
		row + 1: {column, column - 1, column + 1},
	}

	for rowItem, columns := range adjacents {
		for _, columnItem := range columns {
			if rowItem < g.Data.size && columnItem < g.Data.size {
				cellCallback(rowItem, columnItem)
			}
		}
	}
}

func (g *Game) openCell(row, column uint8) error {
	cell := &g.Data.cells[row][column]
	if cell.isOpen == true {
		return nil
	}

	if cell.isBlackHole {
		return fmt.Errorf("black hole")
	}

	cell.isOpen = true
	g.Data.openedCellsCount++

	if g.Data.GetOpenedCellsCount() == 1 {
		g.addBlackHoles()
	}

	//if the opened cell does not occur near any black holes, recursively open all adjacent cells with same rule
	if cell.GetAdjacentBlackHolesCount() == 0 {
		g.updateAdjacentCells(row, column, func(rowItem, columnItem uint8) {
			_ = g.openCell(rowItem, columnItem)
		})
	}

	return nil
}

func (g *Game) Play() {
	for {
		g.Ui.Render(true)

		row, column, err := g.Ui.GetCoordinations()
		if err != nil {
			g.Ui.DisplayErrorMessage("Error: " + err.Error())
			continue
		}

		err = g.openCell(row, column)

		if err != nil {
			g.Data.loseCount++
			g.Ui.Render(false)
			g.Ui.DisplayErrorMessage(gameFailMessage)
			return
		}

		if g.Data.isAllCellOpened() {
			g.Data.winCount++
			g.Ui.Render(true)
			g.Ui.DisplaySuccessMessage(gameWinMessage)
			return
		}
	}
}
