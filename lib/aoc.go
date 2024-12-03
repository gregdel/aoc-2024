package aoc

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"time"
)

var challenges = map[int]Challenge{}

// Register registers a challenge.
func Register(c Challenge, day int) {
	challenges[day] = c
}

// AllDays returns a list of all the days regitered.
func AllDays() []int {
	values := []int{}
	for d := range challenges {
		values = append(values, d)
	}
	sort.Ints(values)
	return values
}

// Challenge represents a challenge.
type Challenge interface {
	Solve(r io.Reader, part int) (string, error)
	Expect(part int, test bool) string
}

// Open opens the input for a given day.
func Open(day int, test bool) (io.ReadCloser, error) {
	dir := fmt.Sprintf("day%02d", day)
	filename := "input"
	if test {
		filename += "-test"
	}

	path := filepath.Join(dir, filename)
	return os.Open(path)
}

// Run run a the challenge.
func Run(day, part int, test bool) (*RunResult, error) {
	challenge, ok := challenges[day]
	if !ok {
		return nil, errors.New("missing day")
	}

	input, err := Open(day, test)
	if err != nil {
		return nil, err
	}
	defer input.Close()

	result := newRunResult(day, part, test)
	result.expected = challenge.Expect(part, test)
	result.start = time.Now()
	result.output, err = challenge.Solve(input, part)
	result.stop = time.Now()
	if err != nil {
		return nil, err
	}

	return result, nil
}
