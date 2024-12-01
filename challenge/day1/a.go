package day1

import (
	"fmt"
	"io"
	"slices"
	"strings"

	"github.com/spf13/cobra"

	"github.com/nlowe/aoc2024/challenge"
	"github.com/nlowe/aoc2024/util"
	"github.com/nlowe/aoc2024/util/gmath"
)

func aCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "a",
		Short: "Day 1, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partA(challenge.InputFile()))
		},
	}
}

func partA(input io.Reader) int {
	var left []int
	var right []int

	for line := range challenge.Lines(input) {
		parts := strings.Fields(line)

		left = append(left, util.MustAtoI(parts[0]))
		right = append(right, util.MustAtoI(parts[1]))
	}

	slices.Sort(left)
	slices.Sort(right)

	var result int
	for i, l := range left {
		result += gmath.Abs(l - right[i])
	}

	return result
}
