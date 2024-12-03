package aoc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSet(t *testing.T) {
	input := []int{
		1, 2, 3, 4, 5, 6,
		1, 2, 3, 4, 5, 6,
		1, 2, 3, 4, 5, 6,
	}

	set := NewSet[int]()
	for _, i := range input {
		set.Add(i)
	}

	assert.Equal(t, set.Len(), 6)
	assert.True(t, set.Has(6))

	set.Remove(6)
	assert.Equal(t, set.Len(), 5)
	assert.False(t, set.Has(6))
	assert.ElementsMatch(t, set.Slice(), []int{1, 2, 3, 4, 5})

	set.Reset()
	assert.Equal(t, set.Len(), 0)
}
