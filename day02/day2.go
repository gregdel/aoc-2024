package day2

import (
	"bufio"
	"io"
	"strconv"

	aoc "github.com/gregdel/aoc2024/lib"
)

func init() {
	aoc.Register(&day{}, 2)
}

type day struct{}

func (d *day) Solve(r io.Reader, part int) (string, error) {
	var check func([]int) bool
	check = isSafe
	if part == 2 {
		check = isSafeWithoutOne
	}

	var safe int
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		ints := aoc.IntsFromString(scanner.Text())
		if check(ints) {
			safe++
		}
	}
	return strconv.Itoa(safe), nil
}

func (d *day) Expect(part int, test bool) string {
	return aoc.NewResult("2", "534", "4", "577").Expect(part, test)
}

func isSafe(ints []int) bool {
	increase := ints[0] < ints[1]
	for i := 0; i < len(ints)-1; i++ {
		cur := ints[i]
		next := ints[i+1]

		if (cur < next) != increase {
			return false
		}

		diff := aoc.Abs(cur - next)
		if diff < 1 || diff > 3 {
			return false
		}
	}

	return true
}

func isSafeWithoutOne(ints []int) bool {
	if isSafe(ints) {
		return true
	}

	for i := 0; i < len(ints); i++ {
		newInts := make([]int, len(ints))
		copy(newInts, ints)
		newInts = append(newInts[:i], newInts[i+1:]...)

		if isSafe(newInts) {
			return true
		}
	}
	return false
}
