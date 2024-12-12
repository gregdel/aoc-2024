package day10

import (
	"io"
	"strconv"

	aoc "github.com/gregdel/aoc2024/lib"
)

func init() {
	aoc.Register(&day{}, 10)
}

type day struct{}

func (d *day) Expect(part int, test bool) string {
	return aoc.NewResult("36", "514", "81", "1162").Expect(part, test)
}

func (d *day) Solve(r io.Reader, part int) (string, error) {
	m := aoc.NewMap2DFromReader(r)
	total := 0
	m.ForAllPoints(func(p *aoc.Point) bool {
		if p.C == '0' {
			ends := aoc.NewSet[*aoc.Point]()
			v := explore(m, p, ends, part)
			total += v
		}
		return true
	})

	return strconv.Itoa(total), nil
}

func explore(m *aoc.Map2D, from *aoc.Point, ends aoc.Set[*aoc.Point], part int) int {
	if from.C == '9' {
		if ends.Has(from) {
			return 0
		}

		if part == 1 {
			ends.Add(from)
		}
		return 1
	}

	total := 0
	for _, d := range aoc.AllDirection {
		next := m.Next(d, from)
		if next == nil || next.C != from.C+1 {
			continue
		}

		total += explore(m, next, ends, part)
	}

	return total
}
