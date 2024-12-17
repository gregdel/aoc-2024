package day17

import (
	"bufio"
	"io"
	"slices"
	"strconv"
	"strings"

	aoc "github.com/gregdel/aoc2024/lib"
)

func init() {
	aoc.Register(&day{}, 17)
}

type day struct {
}

func (d *day) Expect(part int, test bool) string {
	return aoc.NewResult(
		"5,7,3,0",
		"5,1,3,4,3,7,2,1,7",
		"117440", "216584205979245",
	).Expect(part, test)
}

func (d *day) Solve(r io.Reader, part int) (string, error) {
	scanner := bufio.NewScanner(r)
	i := 0
	regs := []int{}
	program := []int{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		line = strings.Split(line, ": ")[1]
		line = strings.ReplaceAll(line, ",", " ")
		ints := aoc.IntsFromString(line)
		if i < 3 {
			regs = append(regs, ints[0])
		} else {
			program = ints
		}
		i++
	}

	result := ""
	if part == 1 {
		out := run(uint64(regs[0]), program)
		for i := 0; i < len(out); i++ {
			if i > 0 {
				result += ","
			}
			result += strconv.Itoa(out[i])
		}
	} else {
		r, _ := find(program, len(program)-1, 0)
		result = strconv.FormatUint(r, 10)
	}

	return result, nil
}

func run(a uint64, program []int) []int {
	b, c, fp := uint64(0), uint64(0), 0
	out := []int{}

	combo := func(comb uint64) uint64 {
		switch comb {
		case 0, 1, 2, 3:
			return comb
		case 4:
			return a
		case 5:
			return b
		case 6:
			return c
		default:
			panic("Unexpected combo")
		}
	}

	for fp < len(program) {
		if fp == len(program)-1 {
			break
		}

		op := uint64(program[fp+1])
		move := true
		switch program[fp] {
		case 0:
			a >>= combo(op)
		case 1:
			b ^= op
		case 2:
			b = combo(op) % 8
		case 3:
			if a != 0 {
				fp = int(op)
				move = false
			}
		case 4:
			b ^= c
		case 5:
			op = combo(op) % 8
			out = append(out, int(op))
		case 6:
			b = a >> combo(op)
		case 7:
			c = a >> combo(op)
		}

		if move {
			fp += 2
		}
	}

	return out
}

func find(program []int, i int, result uint64) (uint64, bool) {
	expected := program[i:]

	result <<= 3
	for a := uint64(0); a < 1<<7; a++ {
		out := run(a|result, program)
		if len(out) > 0 {
			v := result | a
			if !slices.Equal(out, expected) {
				continue
			}

			if slices.Equal(out, program) {
				return v, true
			}

			r, ok := find(program, i-1, v)
			if ok {
				return r, true
			}
		}
	}

	return result, false
}
