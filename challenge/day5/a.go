package day5

import (
	"fmt"
	"io"
	"slices"
	"strings"

	"github.com/spf13/cobra"

	"github.com/nlowe/aoc2024/challenge"
	"github.com/nlowe/aoc2024/util"
)

func aCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "a",
		Short: "Day 5, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partA(challenge.InputFile()))
		},
	}
}

func partA(input io.Reader) int {
	ruleList, updateList, _ := strings.Cut(challenge.Raw(input), "\n\n")
	rules := parseRules(ruleList)

	var result int
	for _, update := range strings.Split(strings.TrimSpace(updateList), "\n") {
		pages := strings.Split(update, ",")
		if slices.IsSortedFunc(pages, sortFunc(rules)) {
			result += util.MustAtoI(pages[len(pages)/2])
		}
	}

	return result
}

func sortFunc(rules map[int]map[int]struct{}) func(x, y string) int {
	return func(x, y string) int {
		xx, yy := util.MustAtoI(x), util.MustAtoI(y)

		if _, ok := rules[xx][yy]; ok {
			return -1
		}

		if _, ok := rules[yy][xx]; !ok {
			return 1
		}

		return 0
	}
}

func parseRules(section string) map[int]map[int]struct{} {
	result := map[int]map[int]struct{}{}

	for _, rule := range strings.Split(section, "\n") {
		x, y, _ := strings.Cut(rule, "|")
		xx, yy := util.MustAtoI(x), util.MustAtoI(y)

		if _, ok := result[xx]; !ok {
			result[xx] = map[int]struct{}{}
		}

		result[xx][yy] = struct{}{}
	}

	return result
}
