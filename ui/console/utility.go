package console

import (
	"fmt"
	gocolor "github.com/TwiN/go-color"
	"strconv"
)

func getUint8FromString(numberStr string, min, max uint8) (coordinate uint8, err error) {
	baseErrMsg := `entered data "%s" incorrect, value should be number >= %d and <= %d`
	numberInt, err := strconv.Atoi(numberStr)
	if err != nil {
		err = fmt.Errorf(baseErrMsg, numberStr, min, max)
		return
	}
	if numberInt < int(min) || numberInt > int(max) {
		err = fmt.Errorf(baseErrMsg, numberStr, min, max)
		return
	}

	return uint8(numberInt), nil
}

func displayMessage(msg, colour string) {
	fmt.Println()
	fmt.Println(gocolor.Ize(colour, msg+". Press enter to continue."))
	_, _ = fmt.Scanln()
}

func displayErrorMessage(msg string) {
	displayMessage(msg, gocolor.Red)
}
