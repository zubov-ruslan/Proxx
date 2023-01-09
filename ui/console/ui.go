package console

import (
	"Proxx/engine"
	"errors"
	"fmt"
	gocolor "github.com/TwiN/go-color"
	"os"
	"os/exec"
)

func NewUi(data *engine.GameData) *Ui {
	return &Ui{data}
}

type Ui struct {
	data *engine.GameData
}

func (ui *Ui) Render(isHideNotOpened bool) {
	ui.clear()
	fmt.Printf("To open: %d / %d", ui.data.GetOpenedCellsCount(), ui.data.GetTotalEmptyCellsCount())
	fmt.Println("")
	for _, row := range ui.data.GetCells() {
		for _, cell := range row {
			ui.renderCell(cell, isHideNotOpened)
		}
		fmt.Println()
	}
}

func (ui *Ui) renderCell(c engine.Cell, isHideNotOpened bool) {
	if isHideNotOpened && c.IsOpen() == false {
		fmt.Print(" x ")
		return
	}

	if c.IsBlackHole() {
		fmt.Print(gocolor.Ize(gocolor.Red, " * "))
	} else {
		fmt.Printf(gocolor.Ize(gocolor.Green, " %d "), c.GetAdjacentBlackHolesCount())
	}
}

func (ui *Ui) clear() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	_ = c.Run()
}

func (ui *Ui) GetCoordinations() (row, column uint8, err error) {
	fmt.Printf("\nEnter row and column through a space [1-%d] [1-%d]: ", ui.data.GetSize(), ui.data.GetSize())

	var rawRow, rawColumn string
	_, err = fmt.Scanln(&rawRow, &rawColumn)
	if err != nil {
		err = errors.New("incorrect amount of arguments, should be exactly 2")
		return
	}

	row, err = getUint8FromString(rawRow, 1, ui.data.GetSize())
	if err != nil {
		return
	}

	column, err = getUint8FromString(rawColumn, 1, ui.data.GetSize())
	if err != nil {
		return
	}

	return row - 1, column - 1, err
}

func (ui *Ui) DisplayErrorMessage(msg string) {
	displayErrorMessage(msg)
}

func (ui *Ui) DisplaySuccessMessage(msg string) {
	displayMessage(msg, gocolor.Green)
}
