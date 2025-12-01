package main

import (
	"aoc2025/utils"
	"fmt"
	"strconv"
)

func wrap(position int, signedSteps int) (int, int) {
	newPosition := (position + signedSteps) % 100
	if newPosition < 0 {
		newPosition += 100
	}

	zeroCrossings := 0
	newPos := position + signedSteps

	if newPos > 0 {
		zeroCrossings = newPos / 100
	} else if newPos == 0 {
		zeroCrossings = 1
	} else {
		zeroCrossings = ((100-position)%100 + (-signedSteps)) / 100
	}

	return newPosition, zeroCrossings
}

func main() {
	input := utils.ReadFile("./day01/input")

	result1 := 0
	result2 := 0
	dial := 50

	for _, line := range input {
		direction := string(line[0])
		steps, _ := strconv.Atoi(line[1:])

		signedSteps := steps
		if direction == "L" {
			signedSteps = -steps
		}

		zeroes := 0
		dial, zeroes = wrap(dial, signedSteps)

		if dial == 0 {
			result1++
		}
		result2 += zeroes
	}

	fmt.Println("Part 1:", result1)
	fmt.Println("Part 2:", result2)
}
