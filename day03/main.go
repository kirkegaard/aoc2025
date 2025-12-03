package main

import (
	"aoc2025/utils"
	"log/slog"
)

// findMaxJoltage finds the maximum joltage for a given bank
func findMaxJoltage(bank string) int {
	maxJoltage := 0
	largestDigitToRight := -1

	// Scan from right to left
	for i := len(bank) - 1; i >= 0; i-- {
		currentDigit := int(bank[i] - '0')

		// If we've seen a digit to the right, try forming a two digit number
		if largestDigitToRight != -1 {
			twoDigitNumber := currentDigit*10 + largestDigitToRight
			// If the two digit number is greater than the max joltage, update it
			if twoDigitNumber > maxJoltage {
				maxJoltage = twoDigitNumber
			}
		}

		// Update the largest digit we've seen to the right
		if currentDigit > largestDigitToRight {
			largestDigitToRight = currentDigit
		}
	}

	return maxJoltage
}

func main() {
	logger := slog.Default()
	input := utils.ReadFile("./day03/input")

	result := 0

	for _, bank := range input {
		maxJoltage := findMaxJoltage(bank)
		result += maxJoltage
		logger.Info("Max joltage", slog.Int("joltage", maxJoltage))
	}

	logger.Info("Result", slog.Int("joltage", result))
}
