package day13

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"

	aoc "github.com/gregdel/aoc2024/lib"
)

func init() {
	aoc.Register(&day{}, 14)
}

type day struct{}

func (d *day) Expect(part int, test bool) string {
	return aoc.NewResult(
		"12", "217328832",
		"0", "7412",
	).Expect(part, test)
}

type data struct {
	p aoc.Point
	v aoc.Vec
}

func (d *day) Solve(r io.Reader, part int) (string, error) {
	input := []data{}
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")
		coords := strings.Split(parts[0][2:], ",")
		vec := strings.Split(parts[1][2:], ",")

		input = append(input, data{
			p: aoc.Point{
				X: aoc.MustGet(strconv.Atoi(coords[0])),
				Y: aoc.MustGet(strconv.Atoi(coords[1])),
			},
			v: aoc.NewVec(
				aoc.MustGet(strconv.Atoi(vec[0])),
				aoc.MustGet(strconv.Atoi(vec[1])),
			),
		})
	}

	maxX, maxY := 101, 103
	if len(input) < 20 {
		maxX, maxY = 11, 7
	}

	var err error
	result := 0
	if part == 1 {
		result, err = solve1(input, 100, maxX, maxY)
	} else {
		result, err = solve2(input, maxX, maxY)
	}
	return strconv.Itoa(result), err
}

func solve1(input []data, seconds, maxX, maxY int) (int, error) {
	points := map[aoc.Point]int{}
	for _, data := range input {
		points[data.p.Translate(data.v.Times(seconds)).Wrap(maxX, maxY)]++
	}

	type quad struct {
		minX, minY, maxX, maxY int
	}

	halfX, halfY := maxX/2, maxY/2
	counts := map[quad]int{}
	for _, q := range []quad{
		{0, 0, halfX, halfY},
		{halfX + 1, 0, maxX, halfY},
		{0, halfY + 1, halfX, maxY},
		{halfX + 1, halfY + 1, maxX, maxY},
	} {
		for p, c := range points {
			if !aoc.IsWithinMap(p.X, p.Y, q.minX, q.minY, q.maxX, q.maxY) {
				continue
			}
			counts[q] += c
		}
	}

	total := 1
	for _, c := range counts {
		total *= c
	}

	i := 0
	p := make([]aoc.Point, len(points))
	for k := range points {
		p[i] = k
		i++
	}
	m := mapFromData(p, maxX, maxY)
	fmt.Println(m)

	return total, nil
}

func mapFromData(points []aoc.Point, maxX, maxY int) *aoc.Map2D {
	m := aoc.NewEmptyMap2D(maxX, maxY, ' ')
	for _, p := range points {
		mp := m.At(p.X, p.Y)
		mp.C = 'X'
	}
	return m
}

func solve2(input []data, maxX, maxY int) (int, error) {
	if len(input) < 20 {
		return 0, nil
	}

	for i := 1; ; i++ {
		ps := aoc.NewSet[aoc.Point]()
		for _, data := range input {
			ps.Add(data.p.Translate(data.v.Times(i)).Wrap(maxX, maxY))
		}
		points := ps.Slice()

		for _, p := range points {
			count := 1
			x := p.X
			for ps.Has(aoc.Point{X: x, Y: p.Y, C: p.C}) {
				x++
				count++
			}

			if count >= 10 {
				return i, nil
			}
		}
	}
}
