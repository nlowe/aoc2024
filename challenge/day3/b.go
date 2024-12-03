package day3

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
		Short: "Day 3, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partB(challenge.InputFile()))
		},
	}
}

func partB(input io.Reader) int {
	var result int

	// Split by "do()" to find sections where mul is enabled, which works because it is enabled for the first section
	for section := range challenge.SectionsOf(input, "do()") {
		// Split by "don't()" and discard the right half as those instructions are disabled until the next section
		check, _, _ := strings.Cut(section, "don't()")

		// Treat the left half as another input for partA to solve
		result += partA(strings.NewReader(check))
	}

	return result
}
