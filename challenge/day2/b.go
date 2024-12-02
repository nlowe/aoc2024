package day2

import (
	"fmt"
	"io"
	"strings"

	"github.com/spf13/cobra"

	"github.com/nlowe/aoc2024/challenge"
)

func bCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "b",
		Short: "Day 2, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partB(challenge.InputFile()))
		},
	}
}

func partB(input io.Reader) int {
	var result int

	for line := range challenge.Lines(input) {
		levels := strings.Fields(line)

		if isSafe(levels) {
			result++
			continue
		}

		for i := 0; i < len(levels); i++ {
			clone := make([]string, len(levels))
			copy(clone, levels)

			if isSafe(append(clone[:i], clone[i+1:]...)) {
				result++
				break
			}
		}
	}

	return result
}
