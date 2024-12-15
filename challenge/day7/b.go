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

func bCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "b",
		Short: "Day 7, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partB(challenge.InputFile()))
		},
	}
}

func partB(input io.Reader) uint64 {
	var result uint64

	for eqn := range challenge.Lines(input) {
		rawWant, rawValues, _ := strings.Cut(eqn, ":")
		want := uint64(util.MustAtoI(rawWant))
		values := slices.Collect(challenge.Fields(strings.TrimSpace(rawValues), util.MustAtoUI))

		if checkConcat(want, values[0], values[1:]) {
			result += want
		}
	}

	return result
}

func checkConcat(want, total uint64, remaining []uint64) bool {
	if len(remaining) == 1 {
		return (want == total+remaining[0]) || (want == total*remaining[0]) || (want == concat(total, remaining[0]))
	}

	return checkConcat(want, total+remaining[0], remaining[1:]) ||
		checkConcat(want, total*remaining[0], remaining[1:]) ||
		checkConcat(want, concat(total, remaining[0]), remaining[1:])
}

func concat(left, right uint64) uint64 {
	left *= 10
	for tmp := right / 10; tmp > 0; tmp /= 10 {
		left *= 10
	}

	return left + right
}
