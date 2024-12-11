package day9

import (
	"bufio"
	"io"
	"strconv"

	aoc "github.com/gregdel/aoc2024/lib"
)

func init() {
	aoc.Register(&day{}, 9)
}

type day struct{}

func (d *day) Expect(part int, test bool) string {
	return aoc.NewResult(
		"1928", "6291146824486",
		"2858", "6307279963620",
	).Expect(part, test)
}

func (d *day) Solve(r io.Reader, part int) (string, error) {
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	line := scanner.Text()

	blocks := aoc.NewList[int]()
	id := 0
	for i, d := range line {
		idx := id
		count := int(d - '0')
		if i%2 == 0 {
			id++
		} else {
			idx = -1
		}

		for j := 0; j < count; j++ {
			blocks.Push(aoc.NewListElement(idx))
		}
	}

	var total uint64
	if part == 1 {
		total = solve1(blocks)

	} else {
		total = solve2(blocks)
	}

	return strconv.FormatUint(total, 10), nil
}

func solve1(blocks *aoc.List[int]) uint64 {
	head, tail := blocks.Head, blocks.Tail
	for head != tail {
		if head.Value >= 0 {
			head = head.Next
			continue
		}

		if tail.Value < 0 {
			tail = tail.Prev
			continue
		}

		head.Value, tail.Value = tail.Value, head.Value
		head, tail = head.Next, tail.Prev
	}

	return checksum(blocks)
}

func findFree(blocks *aoc.List[int], size int, limit *aoc.ListElement[int]) (*aoc.ListElement[int], int) {
	current := blocks.Head
	for current != limit && limit != nil {
		if current.Value >= 0 {
			current = current.Next
			continue
		}

		count := 0
		s := current
		for current != nil && current != limit && current.Value == s.Value {
			current = current.Next
			count++
		}

		if count >= size {
			return s, count
		}
	}
	return nil, 0
}

func solve2(blocks *aoc.List[int]) uint64 {
	tail := blocks.Tail
	for tail != nil {
		if tail.Value < 0 {
			tail = tail.Prev
			continue
		}

		count := 0
		e := tail
		for tail != nil && tail.Value == e.Value {
			tail = tail.Prev
			count++
		}

		found, size := findFree(blocks, count, e)
		if found == nil || size == 0 {
			continue
		}

		for i := 0; i < count; i++ {
			e.Value, found.Value = found.Value, e.Value
			e, found = e.Prev, found.Next
		}
	}

	return checksum(blocks)
}

func checksum(blocks *aoc.List[int]) uint64 {
	total := uint64(0)
	idx := 0
	blocks.ForAll(func(e *aoc.ListElement[int]) {
		if e.Value > 0 {
			total += uint64(idx * e.Value)
		}
		idx++
	})
	return total
}
