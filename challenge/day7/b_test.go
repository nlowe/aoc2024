package day7

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestB(t *testing.T) {
	input := strings.NewReader(example)

	result := partB(input)

	require.EqualValues(t, 11387, result)
}
