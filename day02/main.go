package main

import (
	"aoc2025/utils"
	"log/slog"
	"strconv"
	"strings"
)

func isDoubleRepeat(s string) bool {
	// Return early if the number has an odd number of digits
	if len(s)%2 != 0 {
		return false
	}

	// Calculate the halfway point
	half := len(s) / 2

	// Check if the first half of the string is equal to the second half
	return s[:half] == s[half:]
}

func isRepeatedSequence(s string) bool {
	// Get the length of the string
	l := len(s)

	// Try each possible pattern size from 1 to half the string length
	for size := 1; size <= l/2; size++ {
		// Skip sizes that don't divide evenly into the string length
		if l%size != 0 {
			continue
		}

		// Extract the potential repeating unit
		unit := s[:size]

		// Calculate how many times this unit should repeat
		repeats := l / size

		// Check if the unit repeats throughout the entire string
		exp := strings.Repeat(unit, repeats)
		if exp == s {
			return true
		}
	}

	return false
}

func main() {
	logger := slog.Default()

	// The input
	input := utils.ReadFile("./day02/input")
	line := input[0]

	result1 := 0
	result2 := 0

	for id := range strings.SplitSeq(line, ",") {
		// Split the range into start and end
		parts := strings.Split(id, "-")

		// Parse the numbers
		start, startErr := strconv.Atoi(parts[0])
		end, endErr := strconv.Atoi(parts[1])

		logger.Info("Iterating over", "start", start, "end", end)

		// Check if startErr or endErr are not nil
		if startErr != nil || endErr != nil {
			logger.Error("Error parsing input", "error", startErr, "error", endErr)
		}

		// Make sure start is smaller than end
		if start > end {
			logger.Error("Start is greater than end", "start", start, "end", end)
		}

		for i := start; i <= end; i++ {
			// Convert the number to a string
			s := strconv.Itoa(i)

			if isDoubleRepeat(s) {
				logger.Info("Found double repeat match", "number", i)
				result1 += i
			}

			if isRepeatedSequence(s) {
				logger.Info("Found repeated sequence match", "number", i)
				result2 += i
			}
		}
	}

	logger.Info("Result", "result 1", result1, "result 2", result2)
}
