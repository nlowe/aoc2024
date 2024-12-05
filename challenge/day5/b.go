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

func bCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "b",
		Short: "Day 5, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partB(challenge.InputFile()))
		},
	}
}

func partB(input io.Reader) int {
	ruleList, updateList, _ := strings.Cut(challenge.Raw(input), "\n\n")
	rules := parseRules(ruleList)

	var result int
	for _, update := range strings.Split(strings.TrimSpace(updateList), "\n") {
		if _, ok := ordered(rules, update); ok {
			continue
		}

		pages := strings.Split(update, ",")
		slices.SortFunc(pages, func(x, y string) int {
			xx, yy := util.MustAtoI(x), util.MustAtoI(y)

			if _, ok := rules[xx][yy]; ok {
				return -1
			}

			if _, ok := rules[yy][xx]; !ok {
				return 1
			}

			return 0
		})

		result += util.MustAtoI(pages[len(pages)/2])
	}

	return result
}
