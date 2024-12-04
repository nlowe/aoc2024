package day4

import (
	"fmt"
	"io"

	"github.com/spf13/cobra"

	"github.com/nlowe/aoc2024/challenge"
	"github.com/nlowe/aoc2024/util/tilemap"
)

func bCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "b",
		Short: "Day 4, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partB(challenge.InputFile()))
		},
	}
}

func partB(input io.Reader) int {
	crossword := tilemap.FromInput(input)

	var result int
	for r, p := range crossword.Values() {
		// Look for the center of an X
		if r != 'A' {
			continue
		}

		// Get the corners
		tl, tlOK := crossword.TileAt(p.X-1, p.Y-1)
		tr, trOK := crossword.TileAt(p.X+1, p.Y-1)
		bl, blOK := crossword.TileAt(p.X-1, p.Y+1)
		br, brOK := crossword.TileAt(p.X+1, p.Y+1)

		if !(tlOK && trOK && blOK && brOK) {
			// on an edge
			continue
		}

		// X marks the spot, check for M.S and S.M on both diagonals
		if ((tl == 'M' && br == 'S') || (tl == 'S' && br == 'M')) &&
			((tr == 'M' && bl == 'S') || (tr == 'S' && bl == 'M')) {
			result++
		}
	}

	return result
}
