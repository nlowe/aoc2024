package day6

import (
	"fmt"
	"io"
	"sync"
	"sync/atomic"

	"github.com/spf13/cobra"

	"github.com/nlowe/aoc2024/challenge"
	"github.com/nlowe/aoc2024/util/tilemap"
)

func bCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "b",
		Short: "Day 6, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partB(challenge.InputFile()))
		},
	}
}

func partB(input io.Reader) int {
	m := tilemap.FromInput(input)

	var result atomic.Int32
	var wg sync.WaitGroup
	for _, candidate := range m.AllContainersWith('.') {
		wg.Add(1)

		// CPU go brrrr, this is stupid inefficient but only takes a few seconds
		// TODO: Speed this up by only testing the places the guard _could_ have hit from part A which reduces the
		//       search space from ~17k to ~4k
		go func() {
			defer wg.Done()

			x, y := candidate.Location()
			if causesLoop(x, y, m) {
				result.Add(1)
			}
		}()
	}

	wg.Wait()
	return int(result.Load())
}

func causesLoop(cx, cy int, m *tilemap.Map[rune]) bool {
	start, _ := m.FirstContainerWith('^')
	dx, dy := 0, -1
	x, y := start.Location()

	seen := tilemap.Of[rune](m.Size())
	// TODO: There's gotta be a better way to test for this
	stamina := 1000
	for p, ok := m.TileAt(x+dx, y+dy); ok && stamina > 0; p, ok = m.TileAt(x+dx, y+dy) {
		if (x+dx == cx && y+dy == cy) || p == '#' {
			dx, dy = rotate(dx, dy)
			continue
		}

		if k, _ := seen.TileAt(x, y); k == 'X' {
			stamina--
		}

		seen.SetTile(x, y, 'X')
		x, y = x+dx, y+dy
	}

	return stamina == 0
}
