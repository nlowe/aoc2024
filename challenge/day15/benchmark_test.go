package day15

import (
	"testing"

	"github.com/nlowe/aoc2024/challenge"
)

func BenchmarkA(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = partA(challenge.InputFile())
	}
}

func BenchmarkB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = partB(challenge.InputFile())
	}
}