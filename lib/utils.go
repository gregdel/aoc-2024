package aoc

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// MustGet returns v as is. It panics if err is non-nil.
func MustGet[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}

// IntsFromString returns a slice of ints from a string.
func IntsFromString(input string) []int {
	fields := strings.Fields(input)
	output := make([]int, len(fields))
	for i := 0; i < len(fields); i++ {
		output[i] = MustGet(strconv.Atoi(fields[i]))
	}
	return output
}

func BinStr(i int) string {
	return strconv.FormatInt(int64(i), 2)
}

func BinUStr(i uint64) string {
	return strconv.FormatUint(i, 2)
}

func WaitInput() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
}
