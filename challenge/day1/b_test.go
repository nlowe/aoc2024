package day1

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestB(t *testing.T) {
	input := strings.NewReader(exampleInput)

	result := partB(input)

	require.Equal(t, 31, result)
}
