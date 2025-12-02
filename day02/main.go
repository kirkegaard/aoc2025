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
		parts := strings.Split(id, "-")

		start, startErr := strconv.Atoi(parts[0])
		end, endErr := strconv.Atoi(parts[1])

		logger.Info("Iterating over", "start", start, "end", end)

		if startErr != nil || endErr != nil {
			logger.Error("Error parsing input", "error", startErr, "error", endErr)
		}

		if start > end {
			logger.Error("Start is greater than end", "start", start, "end", end)
		}

		for i := start; i <= end; i++ {
			s := strconv.Itoa(i)

			if len(s)%2 != 0 {
				continue
			}

			half := len(s) / 2
			if s[:half] == s[half:] {
				result += i
			}
		}
	}

	logger.Info("Result", "result", result)
}
