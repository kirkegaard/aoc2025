package main

import (
	"aoc2025/utils"
	"log/slog"
	"strconv"
	"strings"
)

func main() {
	logger := slog.Default()

	// The input
	input := utils.ReadFile("./day02/input")
	line := input[0]

	result := 0

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

			// Skip if the number has an odd number of digits
			if len(s)%2 != 0 {
				continue
			}

			// Calculate the halfway point
			half := len(s) / 2

			// Check if the first half of the string is equal to the second half
			if s[:half] == s[half:] {
				result += i
			}
		}
	}

	logger.Info("Result", "result", result)
}
