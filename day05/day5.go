package day5

import (
	"bufio"
	"io"
	"sort"
	"strconv"
	"strings"

	aoc "github.com/gregdel/aoc2024/lib"
)

func init() {
	aoc.Register(&day{}, 5)
}

type day struct{}

func (d *day) Expect(part int, test bool) string {
	return aoc.NewResult("143", "7024", "123", "4151").Expect(part, test)
}

func (d *day) Solve(r io.Reader, part int) (string, error) {
	rules := map[int]aoc.Set[int]{}
	updates := [][]int{}
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		if strings.Contains(line, "|") {
			ints := aoc.IntsFromString(strings.ReplaceAll(line, "|", " "))
			k, v := ints[0], ints[1]
			if _, ok := rules[k]; !ok {
				rules[k] = aoc.NewSet[int]()
			}
			rules[k].Add(v)
			continue
		}

		updates = append(updates, aoc.IntsFromString(strings.ReplaceAll(line, ",", " ")))
	}

	total := 0
	for _, u := range updates {
		if isValid(rules, u) {
			if part == 2 {
				continue
			}
		} else {
			if part == 1 {
				continue
			}
			sort.Slice(u, func(i, j int) bool {
				return rules[u[j]].Has(u[i])
			})
		}

		total += u[(len(u) / 2)]
	}

	return strconv.Itoa(total), nil
}

func isValid(rules map[int]aoc.Set[int], elems []int) bool {
	for i := 1; i < len(elems); i++ {
		for j := 0; j < i; j++ {
			if !rules[elems[j]].Has(elems[i]) {
				return false
			}
		}
	}

	return true
}
