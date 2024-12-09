package day7

import (
	"bufio"
	"io"
	"strconv"
	"strings"

	aoc "github.com/gregdel/aoc2024/lib"
)

func init() {
	aoc.Register(&day{}, 7)
}

type day struct{}

func (d *day) Expect(part int, test bool) string {
	return aoc.NewResult("3749", "1620690235709", "11387", "145397611075341").Expect(part, test)
}

func (d *day) Solve(r io.Reader, part int) (string, error) {
	total := 0
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		a, b, _ := strings.Cut(scanner.Text(), ":")
		r := aoc.MustGet(strconv.Atoi(a))
		vs := aoc.IntsFromString(b)
		if eval(r, vs, part) {
			total += r
		}
	}
	return strconv.Itoa(total), nil
}

func eval(expected int, values []int, part int) bool {
	l := len(values)
	if l < 2 {
		return false
	}

	v0 := values[0]
	v1 := values[1]
	v0v1 := v0 * v1
	v0pv1 := v0 + v1

	toTest := []int{v0v1, v0pv1}

	if part == 2 {
		v0cv1 := aoc.MustGet(strconv.Atoi(strconv.Itoa(v0) + strconv.Itoa(v1)))
		toTest = append(toTest, v0cv1)
	}

	skip := false
	if l == 2 {
		for _, n := range toTest {
			if n > expected {
				skip = true
			}

			if n == expected {
				return true
			}
		}
	}

	if skip {
		return false
	}

	reduced := make([]int, l-1)
	copy(reduced[1:], values[2:])
	for _, n := range toTest {
		if n > expected {
			continue
		}

		reduced[0] = n
		if eval(expected, reduced, part) {
			return true
		}
	}

	return false
}
