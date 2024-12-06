package day6

import (
	"io"
	"strconv"

	aoc "github.com/gregdel/aoc2024/lib"
)

func init() {
	aoc.Register(&day{}, 6)
}

type day struct{}

func (d *day) Expect(part int, test bool) string {
	return aoc.NewResult("41", "4374", "6", "1705").Expect(part, test)
}

func (d *day) Solve(r io.Reader, part int) (string, error) {
	m := aoc.NewMap2DFromReader(r)
	p := start(m)
	path := explore(m, p)
	if part == 1 {
		return strconv.Itoa(path.Len()), nil
	}

	loops := 0
	path.Remove(p)
	for _, pp := range path.Slice() {
		pp.C = '#'
		if looping(m, p) {
			loops++
		}
		pp.C = '.'
	}

	return strconv.Itoa(loops), nil
}

func start(m *aoc.Map2D) *aoc.Point {
	var ret *aoc.Point
	m.ForAllPoints(func(p *aoc.Point) {
		if p.C == '^' {
			ret = p
		}
	})
	return ret
}

func explore(m *aoc.Map2D, p *aoc.Point) aoc.Set[*aoc.Point] {
	direction := aoc.DirectionUp
	explored := aoc.NewSet[*aoc.Point]()
	next := p
	for {
		explored.Add(next)
		n := m.Next(direction, next)
		if n == nil {
			break
		}

		if n.C == '#' {
			direction = aoc.RotateCW(direction)
			continue
		}

		next = n
	}
	return explored
}

func looping(m *aoc.Map2D, p *aoc.Point) bool {
	type pd struct {
		p *aoc.Point
		d aoc.Direction
	}
	explored := aoc.NewSet[pd]()
	direction := aoc.DirectionUp
	next := p
	explored.Add(pd{p: next, d: direction})
	for {
		n := m.Next(direction, next)
		if n == nil {
			return false
		}

		if n.C == '#' {
			direction = aoc.RotateCW(direction)
			continue
		}

		if explored.Has(pd{p: n, d: direction}) {
			return true
		}
		explored.Add(pd{p: next, d: direction})

		next = n
	}
}
