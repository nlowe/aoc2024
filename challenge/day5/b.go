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
		pages := strings.Split(update, ",")
		sortedPages := strings.Split(update, ",")
		slices.SortStableFunc(sortedPages, sortFunc(rules))

		// Skip anything that didn't get re-ordered
		if slices.Equal(pages, sortedPages) {
			continue
		}

		result += util.MustAtoI(sortedPages[len(sortedPages)/2])
	}

	return result
}
