package day11

import (
	"io"
	"strconv"

	aoc "github.com/gregdel/aoc2024/lib"
)

func init() {
	aoc.Register(&day{}, 11)
}

type day struct{}

func (d *day) Expect(part int, test bool) string {
	return aoc.NewResult(
		"55312", "203457",
		"65601038650482", "241394363462435",
	).Expect(part, test)
}

func (d *day) Solve(r io.Reader, part int) (string, error) {
	line, err := io.ReadAll(r)
	if err != nil {
		return "", err
	}

	input := map[uint64]int{}
	for _, i := range aoc.IntsFromString(string(line)) {
		input[uint64(i)]++
	}

	count := 25
	if part == 2 {
		count = 75
	}

	for i := 0; i < count; i++ {
		input = blink(input)
	}

	total := 0
	for _, count := range input {
		total += count
	}

	return strconv.Itoa(total), nil
}

func blink(input map[uint64]int) map[uint64]int {
	output := make(map[uint64]int, len(input))
	for k, count := range input {
		c, a, b := transform(k)
		output[a] += count
		if c == 2 {
			output[b] += count
		}
	}
	return output
}

func transform(input uint64) (uint8, uint64, uint64) {
	if input == 0 {
		return 1, 1, 0
	}

	digits := aoc.Digits(input)
	if digits%2 == 0 {
		p := uint64(aoc.Pow10(digits / 2))
		a := input / p
		return 2, a, input - (a * p)
	}

	return 1, input * 2024, 0
}
