package day20

import (
	"io"
	"strconv"

	aoc "github.com/gregdel/aoc2024/lib"
)

func init() {
	aoc.Register(&day{}, 20)
}

type day struct{}

func (d *day) Expect(part int, test bool) string {
	return aoc.NewResult(
		"0", "1365", "0", "986082",
	).Expect(part, test)
}

func (d *day) Solve(r io.Reader, part int) (string, error) {
	m := aoc.NewMap2DFromReader(r)
	var start, end *aoc.Point
	m.ForAllPoints(func(p *aoc.Point) bool {
		switch p.C {
		case 'S':
			start = p
		case 'E':
			end = p
		default:
			if start != nil && end != nil {
				return false
			}
		}
		return true
	})

	points := append([]*aoc.Point{start}, m.FindPath(start, end, '.')...)

	result, maxCheat, toSave := 0, 2, 100
	if part == 2 {
		maxCheat = 20
	}

	for i := 0; i < len(points)-toSave; i++ {
		for j := aoc.Min(len(points), i+toSave); j < len(points); j++ {
			distance := aoc.ManhattanDistance(points[i], points[j])
			if part == 1 && distance != 2 {
				continue
			}

			if part == 2 && distance > maxCheat {
				continue
			}

			distStart, distEnd := i, len(points)-1-j
			saved := (len(points) - 1) - (distStart + distEnd + distance)
			if saved <= 0 {
				continue
			}

			if saved >= toSave {
				result++
			}
		}
	}

	return strconv.Itoa(result), nil
}
