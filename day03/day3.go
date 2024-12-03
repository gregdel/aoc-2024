package day3

import (
	"bufio"
	"io"
	"regexp"
	"strconv"
	"strings"

	aoc "github.com/gregdel/aoc2024/lib"
)

func init() {
	aoc.Register(&day{}, 3)
}

var (
	opPattern  = regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	insPattern = regexp.MustCompile(`do(?:n't)?\(\)`)
)

type day struct{}

func (d *day) Expect(part int, test bool) string {
	return aoc.NewResult("161", "170807108", "48", "74838033").Expect(part, test)
}

func (d *day) Solve(r io.Reader, part int) (string, error) {
	var b strings.Builder
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		b.Write(scanner.Bytes())
	}

	if part == 1 {
		return d.solve1(b.String())
	}

	return d.solve2(b.String())
}

func compute(input string) int {
	total := 0
	matches := opPattern.FindAllStringSubmatch(input, -1)
	for _, m := range matches {
		total += aoc.MustGet(strconv.Atoi(m[1])) * aoc.MustGet(strconv.Atoi(m[2]))
	}
	return total
}

func (d *day) solve1(input string) (string, error) {
	return strconv.Itoa(compute(input)), nil
}

func (d *day) solve2(input string) (string, error) {
	enabled := true
	idx := []int{0}

	instructions := insPattern.FindAllStringIndex(input, -1)
	for i := 0; i < len(instructions); i++ {
		isDo := (instructions[i][1] - instructions[i][0]) == 4
		if isDo && !enabled {
			idx = append(idx, instructions[i][1])
			enabled = true
		}

		if !isDo && enabled {
			idx = append(idx, instructions[i][0])
			enabled = false
		}
	}

	if len(idx)%2 != 0 {
		idx = append(idx, len(input))
	}

	total := 0
	for i := 0; i < len(idx); i += 2 {
		total += compute(input[idx[i]:idx[i+1]])
	}

	return strconv.Itoa(total), nil
}
