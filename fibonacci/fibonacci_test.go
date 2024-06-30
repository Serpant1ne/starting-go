package fibonacci_test

import (
	"testing"

	"github.com/serpant1ne/starting-go/fibonacci"
)

type testData struct {
	name  string
	index int
	value int
}

var tests = []testData{
	{"2nd element", 2, 1},
	{"3rd element", 3, 2},
	{"10th element", 10, 54},
	{"50th element", 50, 12586269025},
}

func TestFibonacci(t *testing.T) {
	for _, data := range tests {
		t.Run(data.name, func(t *testing.T) {
			v := fibonacci.Fibonacci(data.index)
			if v != data.value {
				t.Error(
					"For", data.index,
					"expected", data.value,
					"got", v,
				)
			}
		})
	}
}

func TestFibonacciAsync(t *testing.T) {
	c1 := make(chan int)
	c2 := make(chan int)

	for _, data := range tests {
		t.Run(data.name, func(t *testing.T) {
			go fibonacci.FibonacciAsync(data.index-1, c1)
			go fibonacci.FibonacciAsync(data.index-2, c2)
			firstHalf := <-c1
			secondHalf := <-c2
			v := firstHalf + secondHalf
			if v != data.value {
				t.Error(
					"For", data.index,
					"expected", data.value,
					"got", v,
				)
			}
		})

	}
}
