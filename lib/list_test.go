package aoc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	list := NewList[int]()
	assert.Nil(t, list.Pop())

	e1 := NewListElement(1)
	e2 := NewListElement(2)
	e3 := NewListElement(3)

	list.Push(e1)
	list.Push(e2)
	list.Push(e3)

	assert.Equal(t, e1, list.Head)
	assert.Equal(t, e3, list.Tail)
	assert.Equal(t, e3, list.Pop())
	assert.Equal(t, e2, list.Tail)
}
