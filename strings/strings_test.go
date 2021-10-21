package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Test struct {
	in string
	outOdds int
	outEvens int
}

var tests = []Test{
	{
		in:       "abc",
		outOdds:  2,
		outEvens: 1,
	},
	{
		in:       "123",
		outOdds:  2,
		outEvens: 1,
	},
}

func TestCountOddEven(t *testing.T) {
	for _, test := range tests {
		odds, evens := CountOddEven(test.in)
		assert.EqualValues(t, test.outOdds, odds)
		assert.EqualValues(t, test.outEvens, evens)
	}
}
