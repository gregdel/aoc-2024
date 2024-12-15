package day15

import (
	"bufio"
	"io"
	"strconv"

	aoc "github.com/gregdel/aoc2024/lib"
)

func init() {
	aoc.Register(&day{}, 15)
}

type day struct{}

func (d *day) Expect(part int, test bool) string {
	return aoc.NewResult(
		"10092", "1509863",
		"9021", "1548815",
	).Expect(part, test)
}

func (d *day) Solve(r io.Reader, part int) (string, error) {
	scanner := bufio.NewScanner(r)
	m := aoc.NewMap2DFromScanner(scanner)
	if part == 2 {
		m = transformMap(m)
	}

	var start *aoc.Point
	m.ForAllPoints(func(p *aoc.Point) bool {
		if p.C == '@' {
			start = p
			return false
		}
		return true
	})

	for scanner.Scan() {
		line := scanner.Text()
		for _, r := range line {
			if part == 1 {
				start = handleDirection1(m, start, r)
			} else {
				next := handleDirection2(m, start, r)
				if next != start {
					start.C = '.'
					next.C = '@'
					start = next
				}
			}

			if start == nil {
				return "", nil
			}
		}
	}

	result := 0
	m.ForAllPoints(func(p *aoc.Point) bool {
		if p.C == 'O' || p.C == '[' {
			result += p.Y*100 + p.X
			return true
		}
		return true
	})

	return strconv.Itoa(result), nil
}

func handleDirection2(m *aoc.Map2D, start *aoc.Point, r rune) *aoc.Point {
	d := aoc.DirFromRune(r)
	next := m.Next(d, start)
	if next == nil {
		return nil
	}

	if explore(m, next, d, false, false) {
		explore(m, next, d, false, true)
		return next
	}

	return start
}

func explore(m *aoc.Map2D, p *aoc.Point, d aoc.Direction, pair, move bool) bool {
	var other aoc.Direction
	switch p.C {
	case '.':
		return true
	case '#':
		return false
	case '[':
		other = aoc.DirectionRight
	case ']':
		other = aoc.DirectionLeft
	default:
		panic("explore")
	}

	if !pair && (d == aoc.DirectionUp || d == aoc.DirectionDown) {
		if !explore(m, m.Next(other, p), d, true, move) {
			return false
		}
	}

	next := m.Next(d, p)
	if !explore(m, next, d, false, move) {
		return false
	}

	if move {
		next.C, p.C = p.C, next.C
	}

	return true
}

func handleDirection1(m *aoc.Map2D, start *aoc.Point, r rune) *aoc.Point {
	dir := aoc.DirFromRune(r)
	next := m.Next(dir, start)
	if next == nil {
		return start
	}

	switch next.C {
	case '#':
		return start
	case '.':
		start.C, next.C = next.C, start.C
		return next
	case 'O':
		free := m.Next(dir, next)
		for free != nil {
			if free.C == '#' {
				free = nil
				break
			}

			if free.C == 'O' {
				free = m.Next(dir, free)
				continue
			}

			break
		}

		if free == nil {
			return start
		}

		start.C, next.C = next.C, start.C
		start.C, free.C = free.C, start.C
		return next
	default:
		panic("Fuck")
	}
}

func transformMap(m *aoc.Map2D) *aoc.Map2D {
	nm := aoc.NewEmptyMap2D(m.Width()*2, m.Height(), '.')
	for y := 0; y < m.Height(); y++ {
		for x := 0; x < m.Width(); x++ {
			switch m.At(x, y).C {
			case '#':
				nm.At(x*2, y).C = '#'
				nm.At(x*2+1, y).C = '#'
			case '.':
				nm.At(x*2, y).C = '.'
				nm.At(x*2+1, y).C = '.'
			case 'O':
				nm.At(x*2, y).C = '['
				nm.At(x*2+1, y).C = ']'
			case '@':
				nm.At(x*2, y).C = '@'
				nm.At(x*2+1, y).C = '.'
			default:
				panic("Fuck")
			}
		}
	}
	return nm
}
