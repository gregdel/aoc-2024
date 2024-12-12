package aoc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGreaterCommonDivisor(t *testing.T) {
	assert.Equal(t, GreatestCommonDivisor(24, 18), 6)
	assert.Equal(t, GreatestCommonDivisor(15, 10), 5)
	assert.Equal(t, GreatestCommonDivisor(128, 96), 32)
}

func TestLeastCommonMultiple(t *testing.T) {
	assert.Equal(t, LeastCommonMultiple(21, 6), 42)
	assert.Equal(t, LeastCommonMultiple(8, 9, 21), 504)
}

func TestAbs(t *testing.T) {
	assert.Equal(t, Abs(-3), 3)
	assert.Equal(t, Abs(3), 3)
	assert.Equal(t, Abs(-0), 0)
}

func TestMin(t *testing.T) {
	assert.Equal(t, Min(1, 3), 1)
	assert.Equal(t, Min(3, 1), 1)
	assert.Equal(t, Min(1, 1), 1)
}

func TestMax(t *testing.T) {
	assert.Equal(t, Max(1, 3), 3)
	assert.Equal(t, Max(3, 1), 3)
	assert.Equal(t, Max(1, 1), 1)
}

func TestDigits(t *testing.T) {
	assert.Equal(t, Digits(0), 1)
	assert.Equal(t, Digits(10), 2)
	assert.Equal(t, Digits(100), 3)
	assert.Equal(t, Digits(1000), 4)
	assert.Equal(t, Digits(12345), 5)
}

func TestPow10(t *testing.T) {
	assert.Equal(t, Pow10(0), 1)
	assert.Equal(t, Pow10(1), 10)
	assert.Equal(t, Pow10(2), 100)
	assert.Equal(t, Pow10(3), 1000)
	assert.Equal(t, Pow10(4), 10000)
}
