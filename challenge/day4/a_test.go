package day4

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

const example = `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`

func TestA(t *testing.T) {
	input := strings.NewReader(example)

	result := partA(input)

	require.Equal(t, 18, result)
}
