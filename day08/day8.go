package day8

import (
	"io"
	"strconv"

	aoc "github.com/gregdel/aoc2024/lib"
)

func init() {
	aoc.Register(&day{}, 8)
}

type day struct{}

func (d *day) Expect(part int, test bool) string {
	return aoc.NewResult("14", "222", "34", "884").Expect(part, test)
}

func (d *day) Solve(r io.Reader, part int) (string, error) {
	m := aoc.NewMap2DFromReader(r)
	points := map[rune][]*aoc.Point{}
	m.ForAllPoints(func(p *aoc.Point) bool {
		if p.C == '.' {
			return true
		}
		points[p.C] = append(points[p.C], p)
		return true
	})

	antinodes := aoc.NewSet[*aoc.Point]()
	for _, antennas := range points {
		for i := 0; i < len(antennas)-1; i++ {
			for j := i + 1; j < len(antennas); j++ {
				ai := antennas[i]
				aj := antennas[j]
				v := aoc.NewVecFromPoints(*ai, *aj)
				if part == 2 {
					antinodes.Add(ai)
					antinodes.Add(aj)
				}
				for {
					ai = m.At(ai.X-v.U, ai.Y-v.V)
					if ai != nil {
						ai.C = '#'
					}
					antinodes.Add(ai)
					if ai == nil || part == 1 {
						break
					}
				}
				for {
					aj = m.At(aj.X+v.U, aj.Y+v.V)
					if aj != nil {
						aj.C = '#'
					}
					antinodes.Add(aj)
					if aj == nil || part == 1 {
						break
					}
				}
			}
		}
	}

	antinodes.Remove(nil)
	return strconv.Itoa(antinodes.Len()), nil
}
