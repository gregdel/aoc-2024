package main

import (
	"flag"
	"fmt"
	"os"

	aoc "github.com/gregdel/aoc2024/lib"

	_ "github.com/gregdel/aoc2024/day01"
	_ "github.com/gregdel/aoc2024/day02"
	_ "github.com/gregdel/aoc2024/day03"
	_ "github.com/gregdel/aoc2024/day04"
	_ "github.com/gregdel/aoc2024/day05"
	_ "github.com/gregdel/aoc2024/day06"
)

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run() error {
	var day, part int
	var all bool
	var test bool

	flag.IntVar(&day, "day", 0, "day to run")
	flag.IntVar(&part, "part", 0, "part to run")
	flag.BoolVar(&all, "all", false, "run test and non test")
	flag.BoolVar(&test, "test", false, "run on test input only")
	flag.Parse()

	tests := []bool{true, false}
	if test {
		tests = []bool{true}
	}

	days := []int{}
	if all {
		days = aoc.AllDays()
	} else if day != 0 {
		days = []int{day}
	}

	if len(days) == 0 {
		return fmt.Errorf("Missing day")
	}

	parts := []int{1, 2}
	if part != 0 {
		parts = []int{part}
	}

	for _, day := range days {
		if err := aoc.FetchInput(day); err != nil {
			return err
		}

		for _, part := range parts {
			for _, test := range tests {
				result, err := aoc.Run(day, part, test)
				if err != nil {
					return err
				}

				result.Show()
			}
		}
	}

	return nil
}
