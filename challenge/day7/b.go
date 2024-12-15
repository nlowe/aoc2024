package day7

import (
	"fmt"
	"io"

	"github.com/spf13/cobra"

	"github.com/nlowe/aoc2024/challenge"
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

func partB(input io.Reader) int {
	panic("Not implemented!")
}
