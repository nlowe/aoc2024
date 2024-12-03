package day3

import (
	"fmt"
	"io"
	"regexp"

	"github.com/spf13/cobra"

	"github.com/nlowe/aoc2024/challenge"
	"github.com/nlowe/aoc2024/util"
)

var validMUL = regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

func aCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "a",
		Short: "Day 3, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partA(challenge.InputFile()))
		},
	}
}

func partA(input io.Reader) int {
	var result int

	for _, instr := range validMUL.FindAllStringSubmatch(challenge.Raw(input), -1) {
		result += util.MustAtoI(instr[1]) * util.MustAtoI(instr[2])
	}

	return result
}
