package day2

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestB(t *testing.T) {
	input := strings.NewReader(example)

	result := partB(input)

	require.Equal(t, 4, result)
}
