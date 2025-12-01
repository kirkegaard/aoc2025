package main

import (
	"aoc2025/utils"
	"fmt"
	"strconv"
)

func wrap(value int, min int, max int) int {
	rangeSize := max - min + 1
	v := (value - min) % rangeSize
	if v < 0 {
		v += rangeSize
	}
	return v + min
}

func main() {
	result := 0
	dial := 50
	input := utils.ReadFile("./day01/input")

	for _, line := range input {
		direction := string(line[0])
		steps, _ := strconv.Atoi(line[1:])

		switch direction {
		case "R":
			dial = wrap(dial+steps, 0, 99)
		case "L":
			dial = wrap(dial-steps, 0, 99)
		}

		if dial == 0 {
			result++
		}
	}

	fmt.Println(result)
}
