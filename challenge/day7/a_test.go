package day7

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

const example = `190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20`

func TestA(t *testing.T) {
	input := strings.NewReader(example)

	result := partA(input)

	require.EqualValues(t, 3749, result)
}
