package day12

import (
	"io"
	"strconv"

	aoc "github.com/gregdel/aoc2024/lib"
)

func init() {
	aoc.Register(&day{}, 12)
}

type day struct{}

func (d *day) Expect(part int, test bool) string {
	return aoc.NewResult("1930", "1359028", "1206", "839780").Expect(part, test)
}

func (d *day) Solve(r io.Reader, part int) (string, error) {
	m := aoc.NewMap2DFromReader(r)
	total := 0
	explored := aoc.NewSet[*aoc.Point]()
	fences := map[*aoc.Point]aoc.Set[aoc.Direction]{}
	m.ForAllPoints(func(p *aoc.Point) bool {
		if explored.Has(p) {
			return true
		}

		area, perimeter, fc := explore(m, p, explored, fences)
		if part == 1 {
			total += area * perimeter
		} else {
			total += area * fc
		}

		return true
	})

	return strconv.Itoa(total), nil
}

func explore(m *aoc.Map2D, p *aoc.Point, explored aoc.Set[*aoc.Point], fences map[*aoc.Point]aoc.Set[aoc.Direction]) (int, int, int) {
	if explored.Has(p) {
		return 0, 0, 0
	}
	fences[p] = aoc.NewSet[aoc.Direction]()
	explored.Add(p)

	todo := []*aoc.Point{}
	area, perimeter, fencesCount := 1, 4, 0
	for _, d := range aoc.AllDirection {
		next := m.Next(d, p)
		if next == nil || next.C != p.C {
			fences[p].Add(d)
			fencesCount++
			continue
		}

		todo = append(todo, next)
		perimeter--
	}

	for _, next := range todo {
		if !explored.Has(next) {
			continue
		}
		for _, d := range fences[p].Slice() {
			if fences[next].Has(d) {
				fencesCount--
			}
		}
	}

	for _, next := range todo {
		na, np, nf := explore(m, next, explored, fences)
		area += na
		fencesCount += nf
		perimeter += np
	}

	return area, perimeter, fencesCount
}
