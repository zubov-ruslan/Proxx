package engine

import (
	"Proxx/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func getCountBlackHolesCells(cells [][]Cell) int {
	countOfBlackHoles := 0
	for _, rows := range cells {
		for _, cell := range rows {
			if cell.isBlackHole == true {
				countOfBlackHoles++
			}
		}
	}
	return countOfBlackHoles
}

func getAdjacentBlackHolesCount(cells [][]Cell, row, column uint8) uint8 {
	var blackHolesCount uint8
	adjacents := map[uint8][]uint8{
		row:     {column - 1, column + 1},
		row - 1: {column, column - 1, column + 1},
		row + 1: {column, column - 1, column + 1},
	}
	boardSize := uint8(len(cells))

	for adjRow, adjColumns := range adjacents {
		for _, adjColumn := range adjColumns {
			if adjRow < boardSize && adjColumn < boardSize && cells[adjRow][adjColumn].isBlackHole {
				blackHolesCount++
			}
		}
	}

	return blackHolesCount
}

func TestGame_Play(t *testing.T) {

	t.Run("full case with game win", func(t *testing.T) {
		data := NewGameData(8, 10)
		data.init()
		ui := mocks.NewGameUi(t)

		ui.On("Render", true)
		ui.On("DisplaySuccessMessage", gameWinMessage).Once()

		g := NewGame(data, ui)
		//no black holes before first cell has opened
		assert.Equal(t, 0, getCountBlackHolesCells(g.Data.cells))

		//open the first cell in order to fill in cells with blacks holes
		assert.Nil(t, g.openCell(uint8(0), uint8(0)))
		assert.Equal(t, 10, getCountBlackHolesCells(g.Data.cells))

		for row, rows := range g.Data.cells {
			for column, cell := range rows {
				//mock GetCoordinations in order to win the game
				if cell.isBlackHole == false && cell.isOpen == false {
					ui.On("GetCoordinations").Return(uint8(row), uint8(column), nil).Once()
				}
				//check cells adjacent black holes calculation
				expectedAdjacentBlackHolesCount := getAdjacentBlackHolesCount(g.Data.cells, uint8(row), uint8(column))
				assert.Equal(t, expectedAdjacentBlackHolesCount, cell.adjacentBlackHolesCount)
			}
		}

		g.Play()
		assert.Equal(t, 1, g.Data.winCount)
		assert.Equal(t, 0, g.Data.loseCount)
		assert.Equal(t, uint8(8), g.Data.size)
		assert.Equal(t, uint8(10), g.Data.blackHolesCount)
		assert.Equal(t, uint8(54), data.openedCellsCount)
	})

	t.Run("game fail case", func(t *testing.T) {
		data := NewGameData(5, 4)
		data.init()
		ui := &mocks.GameUi{}

		ui.On("Render", true)
		ui.On("Render", false).Once()
		ui.On("DisplayErrorMessage", gameFailMessage).Once()

		g := NewGame(data, ui)
		//open the first cell in order to fill in cells with blacks holes
		assert.Nil(t, g.openCell(uint8(0), uint8(0)))

	Loop:
		for row, rows := range data.cells {
			for column, cell := range rows {
				//mock GetCoordinations in order to lose the game
				if cell.isBlackHole == true {
					ui.On("GetCoordinations").Return(uint8(row), uint8(column), nil).Once()
					break Loop
				}
			}
		}

		g.Play()
		assert.Equal(t, 0, data.winCount)
		assert.Equal(t, 1, data.loseCount)
		assert.Equal(t, uint8(4), data.blackHolesCount)
		assert.Equal(t, uint8(1), data.openedCellsCount)
	})

}
