package main

import (
	"testing"

	asrt "github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	x, y := 2, 2

	c := Calculate{}

	actual := c.Add(x, y)
	expected := 4

	if actual != expected {
		t.Errorf("Calculate.Add(%d,%d) failed. Expected: %d, Actual:%d", x, y, expected, actual)
	}
}

func TestSubtract(t *testing.T) {
	c := Calculate{}

	tables := []struct {
		x        int
		y        int
		expected int
	}{
		{2, 2, 0},
		{3, 2, 1},
	}

	for _, v := range tables {
		actual := c.Subtract(v.x, v.y)

		if actual != v.expected {
			t.Errorf("Calculate.Subtract(%d,%d) failed. Expected: %d, Actual:%d", v.x, v.y, v.expected, actual)
		}

	}
}

func TestMultiply(t *testing.T) {
	assert := asrt.New(t)

	c := Calculate{}

	tables := []struct {
		x        int
		y        int
		expected int
	}{
		{2, 2, 4},
		{3, 2, 6},
	}

	for _, v := range tables {
		actual := c.Multiply(v.x, v.y)

		assert.Equal(v.expected, actual)
	}
}

func TestDivide(t *testing.T) {
	assert := asrt.New(t)
	c := Calculate{}

	tables := []struct {
		x        float64
		y        float64
		expected float64
	}{
		{2, 2, 1},
		{6, 2, 3},
	}

	for _, v := range tables {
		actual := c.Divide(v.x, v.y)

		assert.Equal(v.expected, actual)
	}
}
