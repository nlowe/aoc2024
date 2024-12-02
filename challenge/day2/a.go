package day2

import (
	"fmt"
	"io"
	"strings"

	"github.com/spf13/cobra"

	"github.com/nlowe/aoc2024/challenge"
	"github.com/nlowe/aoc2024/util"
	"github.com/nlowe/aoc2024/util/gmath"
)

func aCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "a",
		Short: "Day 2, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partA(challenge.InputFile()))
		},
	}
}

func partA(input io.Reader) int {
	var result int

	for line := range challenge.Lines(input) {
		levels := strings.Fields(line)

		if isSafe(levels) {
			result++
		}
	}

	return result
}

func isSafe(levels []string) bool {
	last := util.MustAtoI(levels[0])
	dir := util.MustAtoI(levels[1]) - last

	if dir == 0 || gmath.Abs(dir) > 3 {
		return false
	}

	last = util.MustAtoI(levels[1])

	for i := 2; i < len(levels); i++ {
		v := util.MustAtoI(levels[i])

		delta := v - last

		if delta == 0 {
			return false
		}

		if gmath.Sign(delta) != gmath.Sign(dir) {
			return false
		}

		if gmath.Abs(delta) > 3 {
			return false
		}

		last = v
	}

	return true
}
