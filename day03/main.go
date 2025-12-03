package main

import (
	"aoc2025/utils"
	"log/slog"
	"strconv"
)

// findMaxJoltage finds the maximum joltage for a given bank using a greedy
// sliding window.
func findMaxJoltage(bank string, digitsToSelect int) int {
	result := ""
	digitsRemaining := digitsToSelect
	windowSize := len(bank) - digitsToSelect
	startPos := 0

	for digitsRemaining > 0 {
		// Find the largest digit in the valid window
		largestDigit := 0
		largestPos := startPos

		// The window from startPos to (startPos + skipBudget)
		for i := startPos; i <= startPos+windowSize && i < len(bank); i++ {
			digit := int(bank[i] - '0')
			if digit > largestDigit {
				largestDigit = digit
				largestPos = i
			}
		}

		// Select digit
		result += strconv.Itoa(largestDigit)
		digitsRemaining--

		// Update window size and start position
		windowSize -= (largestPos - startPos)
		startPos = largestPos + 1
	}

	// Convert result to integer
	joltage, _ := strconv.Atoi(result)
	return joltage
}

func main() {
	logger := slog.Default()
	input := utils.ReadFile("./day03/input")

	result1 := 0
	result2 := 0

	for _, bank := range input {
		twoDigitJoltage := findMaxJoltage(bank, 2)
		twelveDigitJoltage := findMaxJoltage(bank, 12)

		result1 += twoDigitJoltage
		result2 += twelveDigitJoltage

		logger.Info("Bank joltage", slog.Int("twoDigit", twoDigitJoltage), slog.Int("twelveDigit", twelveDigitJoltage))
	}

	logger.Info("Result 1", slog.Int("total", result1))
	logger.Info("Result 2", slog.Int("total", result2))
}
