package day6

import (
	"fmt"
	"io"

	"github.com/spf13/cobra"

	"github.com/nlowe/aoc2024/challenge"
	"github.com/nlowe/aoc2024/util/tilemap"
)

func aCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "a",
		Short: "Day 6, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partA(challenge.InputFile()))
		},
	}
}

func partA(input io.Reader) int {
	m := tilemap.FromInput(input)

	start, _ := m.FirstContainerWith('^')
	dx, dy := 0, -1
	x, y := start.Location()

	for p, ok := m.TileAt(x+dx, y+dy); ok; p, ok = m.TileAt(x+dx, y+dy) {
		if p == '#' {
			dx, dy = rotate(dx, dy)
			continue
		}

		m.SetTile(x, y, 'X')
		x, y = x+dx, y+dy
	}

	return len(m.AllContainersWith('X')) + 1
}

func rotate(dx, dy int) (int, int) {
	switch {
	case dx == 0 && dy == -1:
		return 1, 0
	case dx == 1 && dy == 0:
		return 0, 1
	case dx == 0 && dy == 1:
		return -1, 0
	case dx == -1 && dy == 0:
		return 0, -1
	}

	panic(fmt.Errorf("invalid vector: %d, %d", dx, dy))
}
