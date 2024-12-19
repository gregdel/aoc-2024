package day19

import (
	"bufio"
	"io"
	"strconv"
	"strings"

	aoc "github.com/gregdel/aoc2024/lib"
)

func init() {
	aoc.Register(&day{}, 19)
}

type day struct{}

func (d *day) Expect(part int, test bool) string {
	return aoc.NewResult(
		"6", "360", "16", "577474410989846",
	).Expect(part, test)
}

type cacheValue struct {
	value int
	ok    bool
}

var cache map[string]cacheValue

func (d *day) Solve(r io.Reader, part int) (string, error) {
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	patterns := strings.Split(scanner.Text(), ", ")

	result := 0
	i := 0
	scanner.Scan()
	for scanner.Scan() {
		input := scanner.Text()

		filterd := []string{}
		for _, p := range patterns {
			if strings.Contains(input, p) {
				filterd = append(filterd, p)
			}
		}

		cache = map[string]cacheValue{}
		v, ok := match(filterd, input, 0, part)
		if ok {
			if part == 1 {
				result++
			} else {
				result += v
			}
		}
		i++
	}

	return strconv.Itoa(result), nil
}

func match(patterns []string, input string, depth, part int) (int, bool) {
	if len(input) == 0 {
		return 1, true
	}

	total := 0
	for _, p := range patterns {
		if strings.HasPrefix(input, p) {
			nextInput := input[len(p):]

			var c int
			var ok bool
			v, ok := cache[nextInput]
			if ok {
				c, ok = v.value, v.ok
			} else {
				c, ok = match(patterns, nextInput, depth+1, part)
				cache[nextInput] = cacheValue{value: c, ok: ok}
			}
			total += c

			if part == 1 {
				if ok {
					return 0, ok
				}
				continue
			}
		}
	}

	return total, total > 0
}
