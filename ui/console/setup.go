package console

import (
	"fmt"
)

const minBoardSize = 5
const maxBoardSize = 30

func GetBoardSizeAndBlockHolesCount() (boardSize, blackHolesCount uint8) {
	boardSize = getBoardSize()
	blackHolesCount = getBlackHolesCount(getMaxBlackHolesCount(boardSize))
	return
}

func getBoardSize() uint8 {
	var sizeStr string

	for {
		fmt.Printf("Enter board size [%d-%d]: ", minBoardSize, maxBoardSize)
		_, _ = fmt.Scanln(&sizeStr)
		size, err := getUint8FromString(sizeStr, minBoardSize, maxBoardSize)
		if err != nil {
			displayErrorMessage(err.Error())
			continue
		}

		return size
	}
}

func getMaxBlackHolesCount(boardSize uint8) uint8 {
	return (boardSize * boardSize) / 2
}

func getBlackHolesCount(maxBlackHolesCount uint8) uint8 {
	var blackHolesCountStr string

	for {
		fmt.Printf("Enter black holes count [1-%d]: ", maxBlackHolesCount)
		_, _ = fmt.Scanln(&blackHolesCountStr)

		blackHolesCount, err := getUint8FromString(blackHolesCountStr, 1, maxBlackHolesCount)
		if err != nil {
			displayErrorMessage(err.Error())
			continue
		}

		return blackHolesCount
	}
}
