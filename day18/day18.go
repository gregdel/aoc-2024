package day18

import (
	"bufio"
	"io"
	"strconv"
	"strings"

	aoc "github.com/gregdel/aoc2024/lib"
)

func init() {
	aoc.Register(&day{}, 18)
}

type day struct {
}

func (d *day) Expect(part int, test bool) string {
	return aoc.NewResult(
		"22", "308", "6,1", "46,28",
	).Expect(part, test)
}

func (d *day) Solve(r io.Reader, part int) (string, error) {
	scanner := bufio.NewScanner(r)
	points := []*aoc.Point{}
	for scanner.Scan() {
		coords := aoc.IntsFromString(strings.ReplaceAll(scanner.Text(), ",", " "))
		points = append(points, aoc.NewPoint(coords[0], coords[1], '#'))
	}

	mx, my, bytes := 71, 71, 1024
	if len(points) < 100 {
		mx, my, bytes = 7, 7, 12
	}

	m := aoc.NewEmptyMap2D(mx, my, '.')
	start, end := m.At(0, 0), m.At(mx-1, my-1)

	result := ""
	for i, p := range points {
		if part == 1 && i == bytes {
			result = strconv.Itoa(m.FindPathDistance(start, end, '.'))
			break
		}
		m.At(p.X, p.Y).C = p.C

		if i < bytes {
			continue
		}

		d := m.FindPathDistance(start, end, '.')
		if d == 0 {
			result = strconv.Itoa(p.X) + "," + strconv.Itoa(p.Y)
			break
		}
	}

	return result, nil
}
