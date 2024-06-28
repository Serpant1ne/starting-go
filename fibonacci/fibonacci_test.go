package fibonacci_test

import (
	"fmt"
	"testing"

	"github.com/serpant1ne/starting-go/fibonacci"
)

type testPair struct {
	index int
	value int
}

var tests = []testPair{
	{2, 1},
	{3, 2},
	{10, 55},
	{50, 12586269025},
}

func TestFibonacci(t *testing.T) {
	for i, pair := range tests {
		v := fibonacci.Fibonacci(pair.index)
		if v != pair.value {
			t.Error(
				"For", pair.index,
				"expected", pair.value,
				"got", v,
			)
		} else {
			fmt.Printf(`test %d passed `, i+1)
		}
	}
}

func TestFibonacciAsync(t *testing.T) {
	c1 := make(chan int)
	c2 := make(chan int)

	for i, pair := range tests {
		go fibonacci.FibonacciAsync(pair.index-1, c1)
		go fibonacci.FibonacciAsync(pair.index-2, c2)
		firstHalf := <-c1
		secondHalf := <-c2
		v := firstHalf + secondHalf
		if v != pair.value {
			t.Error(
				"For", pair.index,
				"expected", pair.value,
				"got", v,
			)
		} else {
			fmt.Printf(`test %d passed`, i+1)
		}
	}
}
