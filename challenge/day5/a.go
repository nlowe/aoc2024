package day5

import (
	"fmt"
	"io"
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
		if middle, ok := ordered(rules, update); ok {
			result += middle
		}
	}

	return result
}

func ordered(rules map[int]map[int]struct{}, update string) (int, bool) {
	pages := strings.Split(update, ",")

	for i, page := range pages {
		v := util.MustAtoI(page)

		for j := i + 1; j < len(pages); j++ {
			vv := util.MustAtoI(pages[j])

			// Found a later page that should have come before it
			if _, ok := rules[vv][v]; ok {
				return 0, false
			}
		}
	}

	return util.MustAtoI(pages[len(pages)/2]), true
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
