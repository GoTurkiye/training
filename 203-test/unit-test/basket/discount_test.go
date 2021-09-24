package main

import (
	"fmt"
	"testing"

	asrt "github.com/stretchr/testify/assert"
)

type TestCalculator struct {
	add      map[string]int
	subtract map[string]int
	multiply map[string]int
	divide   map[string]float64
}

// TODO: https://github.com/vektra/mockery ile auto gen mocklar uretebilirsiniz
func (c TestCalculator) On(f string, x, y int, rt float64) {
	switch f {
	case "Add":
		c.add[fmt.Sprintf("%d%d", x, y)] = int(rt)
	case "Subtract":
		c.subtract[fmt.Sprintf("%d%d", x, y)] = int(rt)
	case "Multiply":
		c.multiply[fmt.Sprintf("%d%d", x, y)] = int(rt)
	case "Divide":
		c.divide[fmt.Sprintf("%d%d", x, y)] = rt
	}
}

func (c TestCalculator) Add(x, y int) int {
	return c.add[fmt.Sprintf("%d%d", x, y)]
}

func (c TestCalculator) Subtract(x, y int) int {
	return c.subtract[fmt.Sprintf("%d%d", x, y)]
}

func (c TestCalculator) Multiply(x, y int) int {
	return c.multiply[fmt.Sprintf("%d%d", x, y)]
}

func (c TestCalculator) Divide(x, y int) float64 {
	return c.divide[fmt.Sprintf("%d%d", x, y)]
}

func TestAmount(t *testing.T) {
	assert := asrt.New(t)

	testCalculator := TestCalculator{
		add:      map[string]int{},
		subtract: map[string]int{},
		multiply: map[string]int{},
		divide:   map[string]float64{},
	}

	d := NewMinPriceDiscount(11, testCalculator)

	tables := []struct {
		amount     float64
		percentage float64
		expected   float64
	}{
		{100, 20, 80},
		{10, 20, 10},
		{100, 120, 100},
	}

	testCalculator.On("Subtract", 100, 20, 80)

	for _, v := range tables {
		actual := d.Amount(v.amount, v.percentage)

		assert.Equal(v.expected, actual)
	}
}

func TestPercentage(t *testing.T) {
	assert := asrt.New(t)

	testCalculator := TestCalculator{
		add:      map[string]int{},
		subtract: map[string]int{},
		multiply: map[string]int{},
		divide:   map[string]float64{},
	}

	d := NewMinPriceDiscount(11, testCalculator)

	tables := []struct {
		amount     float64
		percentage float64
		expected   float64
	}{
		{100, 20, 80},
		{10, 20, 10},
		{100, 101, 100},
	}

	testCalculator.On("Multiply", 100, 20, 2000)
	testCalculator.On("Multiply", 100, 101, 10100)
	testCalculator.On("Divide", 2000, 100, 20)
	testCalculator.On("Divide", 10100, 100, 101)
	testCalculator.On("Subtract", 100, 20, 80)

	for _, v := range tables {
		actual := d.Percentage(v.amount, v.percentage)

		assert.Equal(v.expected, actual)
	}
}
