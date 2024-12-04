package day4

import (
	"fmt"
	"io"
	"strings"

	"github.com/spf13/cobra"

	"github.com/nlowe/aoc2024/challenge"
	"github.com/nlowe/aoc2024/util/tilemap"
)

const needle = "XMAS"

func aCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "a",
		Short: "Day 4, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partA(challenge.InputFile()))
		},
	}
}

func partA(input io.Reader) int {
	crossword := tilemap.FromInput(input)

	var result int
	for _, p := range crossword.Values() {
		for _, candidate := range wordSlicesFrom(crossword, p.X, p.Y) {
			if candidate == needle {
				result++
			}
		}
	}

	return result
}

func wordSlicesFrom(crossword *tilemap.Map[rune], x, y int) []string {
	result := make([]string, 0, 8)

	for _, delta := range []struct {
		x int
		y int
	}{
		{1, 1},   // down right
		{1, 0},   //      right
		{1, -1},  // up   right
		{0, 1},   // down
		{0, -1},  // up
		{-1, 1},  // down left
		{-1, 0},  //      left
		{-1, -1}, // up   left
	} {
		if word := wordSliceAlong(crossword, x, y, delta.x, delta.y); word != "" {
			result = append(result, word)
		}
	}

	return result
}

func wordSliceAlong(crossword *tilemap.Map[rune], x, y, dx, dy int) string {
	var word strings.Builder

	var r rune
	var ok bool

	for i := range len(needle) {
		r, ok = crossword.TileAt(x+i*dx, y+i*dy)
		if !ok {
			return ""
		}

		word.WriteRune(r)
	}

	return word.String()
}
