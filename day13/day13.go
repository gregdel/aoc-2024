package day13

import (
	"bufio"
	"io"
	"math"
	"strconv"
	"strings"

	aoc "github.com/gregdel/aoc2024/lib"
	"gonum.org/v1/gonum/mat"
)

func init() {
	aoc.Register(&day{}, 13)
}

type day struct{}

func (d *day) Expect(part int, test bool) string {
	return aoc.NewResult(
		"480", "29598",
		"875318608908", "93217456941970",
	).Expect(part, test)
}

const offset = 10000000000000

func (d *day) Solve(r io.Reader, part int) (string, error) {
	total := 0
	scanner := bufio.NewScanner(r)
	matrix := make([]float64, 4)
	i := -1
	for scanner.Scan() {
		i = (i + 1) % 4
		line := scanner.Text()
		if line == "" {
			continue
		}

		line = strings.Split(line, ":")[1]
		parts := strings.Split(line, ",")
		x := float64(aoc.MustGet(strconv.Atoi(parts[0][3:])))
		y := float64(aoc.MustGet(strconv.Atoi(parts[1][3:])))
		if i != 2 {
			matrix[i], matrix[i+2] = x, y
			continue
		}

		if part == 2 {
			x, y = x+offset, y+offset
		}

		m := mat.NewDense(2, 2, matrix)
		r := mat.NewVecDense(2, []float64{x, y})
		var v mat.VecDense
		if err := v.SolveVec(m, r); err != nil {
			continue
		}

		a, b := int(math.Round(v.At(0, 0))), int(math.Round(v.At(1, 0)))
		if (a*int(matrix[0])+b*int(matrix[1])) != int(x) ||
			(a*int(matrix[2])+b*int(matrix[3])) != int(y) ||
			a < 0 || b < 0 {
			continue
		}

		if part == 1 && (a > 100 || b > 100) {
			continue
		}

		total += a*3 + b
	}
	return strconv.Itoa(total), nil
}
