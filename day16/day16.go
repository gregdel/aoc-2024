package day16

import (
	"io"
	"slices"
	"strconv"

	aoc "github.com/gregdel/aoc2024/lib"
)

func init() {
	aoc.Register(&day{}, 16)
}

type day struct{}

func (d *day) Expect(part int, test bool) string {
	return aoc.NewResult(
		"7036", "99488",
		"45", "516",
	).Expect(part, test)
}

func (d *day) Solve(r io.Reader, part int) (string, error) {
	m := aoc.NewMap2DFromReader(r)

	start := m.At(1, m.Height()-2)
	end := m.At(m.Width()-2, 1)

	result := 0
	bestPath, score := findPath(m, start, end)
	result = score
	if part == 2 {
		points := aoc.NewSet[*aoc.Point]()
		for _, move := range bestPath {
			worthIt := false
			for _, d := range aoc.AllDirection {
				np := m.Next(d, move.P)
				if np != nil && np.C == '.' && !points.Has(np) {
					worthIt = true
					break
				}
			}

			if !worthIt {
				continue
			}

			or := move.P.C
			move.P.C = '#'
			bp, score := findPath(m, start, end)
			if score == result {
				for _, p := range bp {
					points.Add(p.P)
				}
			}
			move.P.C = or
		}
		result = points.Len()
	}

	return strconv.Itoa(result), nil
}

func findPath(m *aoc.Map2D, start, end *aoc.Point) ([]aoc.Move, int) {
	moves := aoc.NewSet[aoc.Move]()
	distances := map[aoc.Move]int{}
	ancestors := map[aoc.Move]aoc.Move{}

	pq := aoc.NewPriorityQueue[aoc.Move]()
	pq.Push(aoc.NewMove(start, aoc.DirectionRight), 0)

	var lastMove aoc.Move
	bestScore := 0
	for pq.Len() > 0 {
		move, score := pq.Pop()
		if move.P == end {
			lastMove = move
			bestScore = score
			break
		}

		if moves.Has(move) {
			continue
		}
		moves.Add(move)

		dir := move.D
		for i := 0; i < 4; i++ {
			s := score + 1
			dir = aoc.RotateCW(dir)
			switch i {
			case 0:
				dir = move.D
			case 2:
				continue
			case 1, 3:
				s += 1000
			}

			n := m.Next(dir, move.P)
			if n == nil || n.C == '#' {
				continue
			}

			nm := aoc.NewMove(n, dir)
			if moves.Has(nm) {
				continue
			}

			pd, ok := distances[nm]
			if ok && s > pd {
				continue
			}

			ancestors[nm] = move
			distances[nm] = s
			pq.Push(nm, s)
		}
	}

	if bestScore == 0 {
		return nil, 0
	}

	ret := []aoc.Move{}
	move, ok := lastMove, true
	for ok {
		ret = append(ret, move)
		move, ok = ancestors[move]
	}
	slices.Reverse(ret)

	return ret, bestScore
}
