package day20

import (
	"fmt"
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
		"0", "1365", "x", "x",
	).Expect(part, test)
}

// 1339 too low
func (d *day) Solve(r io.Reader, part int) (string, error) {
	m := aoc.NewMap2DFromReader(r)
	// fmt.Println(m)

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

	startEndPath, distStart := m.FindPathDistances(start, end, '.')
	_, distEnd := m.FindPathDistances(end, start, '.')

	baseScore := len(startEndPath)
	fmt.Println("Base score", baseScore)

	points := []*aoc.Point{}
	m.ForAllPoints(func(p *aoc.Point) bool {
		if p.C != '#' {
			return true
		}

		for _, d := range aoc.AllDirection {
			np := m.Next(d, p)
			if np == nil || np.C != '#' {
				points = append(points, p)
				break
			}
		}

		return true
	})
	fmt.Println(len(points))

	maxCheat := 1

	yo := map[int]int{}
	result := 0
	for _, p := range points {
		ds, sok := distStart[p]
		if !sok {
			continue
		}

		// p.C = '.'
		// realScore := len(m.FindPath(start, end, '.'))
		// p.C = '#'

		// if score != realScore {
		// 	fmt.Println(p, ds, de, score, realScore)
		// 	break
		// }

		explored := aoc.NewSet[*aoc.Point]()
		pq := aoc.NewPriorityQueue[*aoc.Point]()
		pq.Push(p, 1)
		for pq.Len() > 0 {
			cp, prio := pq.Pop()
			if prio > maxCheat {
				continue
			}
			if explored.Has(cp) {
				continue
			}

			de, dok := distEnd[p]
			if dok {
				score := ds + de
				saved := baseScore - score
				if score > baseScore {
					continue
				}

				yo[saved]++
				// fmt.Println(score, "saved:", saved)
				if saved >= 100 {
					result++
				}
			}

			for _, d := range aoc.AllDirection {
				np := m.Next(d, p)
				if np == nil || np.C != '#' {
					continue
				}
				pq.Push(np, prio+1)
			}
		}
	}

	for s, c := range yo {
		fmt.Println(c, "saved", s)
	}

	fmt.Println(yo)

	return strconv.Itoa(result), nil
}
