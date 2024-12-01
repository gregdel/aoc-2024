package day1

import (
	"bufio"
	"io"
	"sort"
	"strconv"

	aoc "github.com/gregdel/aoc2024/lib"
)

func init() {
	aoc.Register(&day{}, 1)
}

type day struct{}

func (d *day) Solve(r io.Reader, part int) (string, error) {
	left, right, err := parseInput(r)
	if err != nil {
		return "", err
	}

	if part == 1 {
		return solve1(left, right)
	}
	return solve2(left, right)
}

func (d *day) Expect(part int, test bool) string {
	return aoc.NewResult("11", "2264607", "31", "19457120").Expect(part, test)
}

func parseInput(r io.Reader) ([]int, []int, error) {
	left, right := []int{}, []int{}

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		values := aoc.IntsFromString(scanner.Text())
		left = append(left, values[0])
		right = append(right, values[1])
	}

	return left, right, scanner.Err()
}

func solve1(left, right []int) (string, error) {
	sort.Ints(left)
	sort.Ints(right)

	total := 0
	for i := 0; i < len(left); i++ {
		total += aoc.Abs(left[i] - right[i])
	}

	return strconv.Itoa(total), nil
}

func solve2(left, right []int) (string, error) {
	counts := map[int]int{}
	for _, n := range right {
		counts[n]++
	}

	total := 0
	for _, n := range left {
		total += n * counts[n]
	}

	return strconv.Itoa(total), nil
}
