package api

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	tests := []struct {
		label    string
		int1     int
		int2     int
		limit    int
		str1     string
		str2     string
		expected []string
	}{
		{
			label:    "nominal",
			int1:     3,
			int2:     5,
			limit:    100,
			str1:     "fizz",
			str2:     "buzz",
			expected: []string{"1", "2", "fizz", "4", "buzz", "fizz", "7", "8", "fizz", "buzz", "11", "fizz", "13", "14", "fizzbuzz", "16", "17", "fizz", "19", "buzz", "fizz", "22", "23", "fizz", "buzz", "26", "fizz", "28", "29", "fizzbuzz", "31", "32", "fizz", "34", "buzz", "fizz", "37", "38", "fizz", "buzz", "41", "fizz", "43", "44", "fizzbuzz", "46", "47", "fizz", "49", "buzz", "fizz", "52", "53", "fizz", "buzz", "56", "fizz", "58", "59", "fizzbuzz", "61", "62", "fizz", "64", "buzz", "fizz", "67", "68", "fizz", "buzz", "71", "fizz", "73", "74", "fizzbuzz", "76", "77", "fizz", "79", "buzz", "fizz", "82", "83", "fizz", "buzz", "86", "fizz", "88", "89", "fizzbuzz", "91", "92", "fizz", "94", "buzz", "fizz", "97", "98", "fizz", "buzz"},
		},
		{
			label:    "empty",
			int1:     0,
			int2:     0,
			limit:    0,
			str1:     "",
			str2:     "",
			expected: []string{},
		},
		{
			label:    "divided by zero",
			int1:     0,
			int2:     0,
			limit:    5,
			str1:     "",
			str2:     "",
			expected: []string{"1", "2", "3", "4", "5"},
		},
	}

	for _, test := range tests {
		t.Run(test.label, func(t *testing.T) {
			res := FizzBuzz(test.int1, test.int2, test.limit, test.str1, test.str2)
			assert.Equal(t, test.expected, res)
		})
	}
}
