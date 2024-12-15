package day7

import (
	"fmt"
	"io"
	"slices"
	"strings"

	"github.com/spf13/cobra"

	"github.com/nlowe/aoc2024/challenge"
	"github.com/nlowe/aoc2024/util"
)

func aCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "a",
		Short: "Day 7, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partA(challenge.InputFile()))
		},
	}
}

func partA(input io.Reader) uint64 {
	var result uint64

	for eqn := range challenge.Lines(input) {
		rawWant, rawValues, _ := strings.Cut(eqn, ":")
		want := uint64(util.MustAtoI(rawWant))
		values := slices.Collect(challenge.Fields(strings.TrimSpace(rawValues), util.MustAtoUI))

		if check(want, values[0], values[1:]) {
			result += want
		}
	}

	return result
}

func check(want, total uint64, remaining []uint64) bool {
	if len(remaining) == 1 {
		return (want == total+remaining[0]) || (want == total*remaining[0])
	}

	return check(want, total+remaining[0], remaining[1:]) || check(want, total*remaining[0], remaining[1:])
}
