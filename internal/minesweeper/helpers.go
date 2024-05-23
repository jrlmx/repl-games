package minesweeper

import (
	"fmt"
	"strconv"
)

func parseXYInput(xname string, yname string, args ...string) (int, int, error) {
	if len(args) < 2 {
		return 0, 0, fmt.Errorf("missing %s and/or %s", xname, yname)
	}

	x, err := strconv.Atoi(args[0])

	if err != nil {
		return 0, 0, fmt.Errorf("invalid x")
	}

	y, err := strconv.Atoi(args[1])

	if err != nil {
		return 0, 0, fmt.Errorf("invalid y")
	}

	return x, y, nil
}

func displayLogo() {
	fmt.Print(`
          __                                                         
.--------|__.-----.-----.-----.--.--.--.-----.-----.-----.-----.----.
|        |  |     |  -__|__ --|  |  |  |  -__|  -__|  _  |  -__|   _|
|__|__|__|__|__|__|_____|_____|________|_____|_____|   __|_____|__|  
                                                   |__|                 

`)
}
