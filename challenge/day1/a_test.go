package day1

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

const exampleInput = `3   4
4   3
2   5
1   3
3   9
3   3`

func TestA(t *testing.T) {

	input := strings.NewReader(exampleInput)

	result := partA(input)

	require.Equal(t, 11, result)
}
