package engine

type Cell struct {
	isBlackHole             bool
	isOpen                  bool
	adjacentBlackHolesCount uint8
}

func (c *Cell) IsOpen() bool {
	return c.isOpen
}

func (c *Cell) IsBlackHole() bool {
	return c.isBlackHole
}

func (c *Cell) GetAdjacentBlackHolesCount() uint8 {
	return c.adjacentBlackHolesCount
}

// GameData game setup, statistics, and two-dimensional array of cells represent board,
// indexes represent the location of each cell on the board
type GameData struct {
	size             uint8
	blackHolesCount  uint8
	openedCellsCount uint8
	winCount         int
	loseCount        int
	cells            [][]Cell
}

func (d *GameData) init() {
	d.openedCellsCount = 0
	d.cells = make([][]Cell, d.size)
	for i := uint8(0); i < d.size; i++ {
		d.cells[i] = make([]Cell, d.size)
	}
}

func (d *GameData) GetSize() uint8 {
	return d.size
}

func (d *GameData) GetTotalEmptyCellsCount() uint8 {
	return (d.size * d.size) - d.blackHolesCount
}

func (d *GameData) GetOpenedCellsCount() uint8 {
	return d.openedCellsCount
}

func (d *GameData) isAllCellOpened() bool {
	return d.GetOpenedCellsCount() == d.GetTotalEmptyCellsCount()
}

func (d *GameData) GetCells() [][]Cell {
	return d.cells
}

func NewGameData(size, blackHolesCount uint8) *GameData {
	return &GameData{size: size, blackHolesCount: blackHolesCount}
}
