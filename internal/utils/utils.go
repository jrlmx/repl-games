package utils

import (
	"fmt"
	"strconv"
)

func ParseXYInput(xname string, yname string, args ...string) (int, int, error) {
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

func ParseXInput(xname string, args ...string) (int, error) {
	if len(args) < 1 {
		return 0, fmt.Errorf("missing %s", xname)
	}

	x, err := strconv.Atoi(args[0])

	if err != nil {
		return 0, fmt.Errorf("invalid x")
	}

	return x, nil
}