package day4

import (
	"io"
	"strconv"

	aoc "github.com/gregdel/aoc2024/lib"
)

func init() {
	aoc.Register(&day{}, 4)
}

type day struct{}

func (d *day) Expect(part int, test bool) string {
	return aoc.NewResult("18", "2644", "9", "1952").Expect(part, test)
}

func (d *day) Solve(r io.Reader, part int) (string, error) {
	m := aoc.NewMap2DFromReader(r)

	var total int
	if part == 1 {
		total = solve1(m)
	} else {
		total = solve2(m)
	}

	return strconv.Itoa(total), nil
}

func solve1(m *aoc.Map2D) int {
	total := 0
	m.ForAllPoints(func(p *aoc.Point) {
		if p.C == 'X' {
			total += findXMAS(m, p)
		}
	})
	return total
}

func solve2(m *aoc.Map2D) int {
	total := 0
	m.ForAllPoints(func(p *aoc.Point) {
		if p.C == 'A' && hasCross(m, p) {
			total++
		}
	})
	return total
}

func findXMAS(m *aoc.Map2D, p *aoc.Point) int {
	found := 0
	for _, d := range aoc.AllDirectionWithDiags {
		np := p
		l := 0
		for _, c := range "MAS" {
			np = m.Next(d, np)
			if np == nil || np.C != c {
				break
			}
			l++
		}

		if l == 3 {
			found++
		}
	}
	return found
}

func hasCross(m *aoc.Map2D, p *aoc.Point) bool {
	for _, d := range []aoc.Direction{
		aoc.DirectionUpLeft, aoc.DirectionDownLeft,
	} {
		np := m.Next(d, p)
		if np == nil || (np.C != 'M' && np.C != 'S') {
			return false
		}

		other := 'M'
		if np.C == other {
			other = 'S'
		}

		np = m.Next(aoc.OppositeDirection(d), p)
		if np == nil || np.C != other {
			return false
		}
	}

	return true
}
