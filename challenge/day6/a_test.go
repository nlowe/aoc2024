package day6

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

const example = `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`

func TestA(t *testing.T) {
	input := strings.NewReader(example)

	result := partA(input)

	require.Equal(t, 41, result)
}
