package luhn_test

import (
	"fmt"
	"testing"

	"github.com/serpant1ne/starting-go/luhn"
)

type testData struct {
	number int
	result bool
}

var tests = []testData{
	{154, false},
	{1234567890123456, false},
	{17893729974, true},
	{5356173103211482, true},
}

func TestCheckNumber(t *testing.T) {
	for i, data := range tests {
		out, err := luhn.CheckNumber(data.number)
		if err != nil {
			t.Error(
				"Error:", err,
			)
		} else if out != data.result {
			t.Error(
				"For", data.number,
				"expected", data.result,
				"got", out,
			)
		} else {
			fmt.Printf(`test %d passed; `, i+1)
		}
	}
}
