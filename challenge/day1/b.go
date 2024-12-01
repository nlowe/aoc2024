package day1

import (
	"fmt"
	"io"
	"strings"

	"github.com/spf13/cobra"

	"github.com/nlowe/aoc2024/challenge"
	"github.com/nlowe/aoc2024/util"
)

func bCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "b",
		Short: "Day 1, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partB(challenge.InputFile()))
		},
	}
}

func partB(input io.Reader) int {
	var left []int
	right := map[int]int{}

	for line := range challenge.Lines(input) {
		parts := strings.Fields(line)

		left = append(left, util.MustAtoI(parts[0]))
		right[util.MustAtoI(parts[1])]++
	}

	var result int
	for _, i := range left {
		result += i * right[i]
	}

	return result
}
