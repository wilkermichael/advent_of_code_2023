package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseSum(t *testing.T) {
	testCases := []struct {
		name     string
		input    []string
		expected int
	}{
		{
			name: "first and last",
			input: []string{
				"1abc2",
				"pqr3stu8vwx",
				"a1b2c3d4e5f",
				"treb7uchet",
			},
			expected: 142,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			out, _ := ParseSum(tc.input)
			assert.Equal(t, tc.expected, out)
		})
	}
}

func TestParseSumWithStrings(t *testing.T) {
	testCases := []struct {
		name     string
		input    []string
		expected int
	}{
		{
			name: "first and last",
			input: []string{
				"two1nine",
				"eightwothree",
				"abcone2threexyz",
				"xtwone3four",
				"4nineeightseven2",
				"zoneight234",
				"7pqrstsixteen",
			},
			expected: 281,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			out, _ := ParseSumWithStrings(tc.input)
			assert.Equal(t, tc.expected, out)
		})
	}
}

func TestParseStringsToInt(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "first and last",
			input:    "1asdfad4",
			expected: 14,
		},
		{
			name:     "embedded",
			input:    "asdf3dvjk4lkh",
			expected: 34,
		},
		{
			name:     "embedded and last",
			input:    "asdf3dvjklkh4",
			expected: 34,
		},
		{
			name:     "first and embedded",
			input:    "3asdfdvjklkh4",
			expected: 34,
		},
		{
			name:     "multi",
			input:    "3adsfasd4asdfasdf6",
			expected: 36,
		},
		{
			name:     "only digits",
			input:    "24601",
			expected: 21,
		},
		{
			name:     "single digits",
			input:    "abcd8sgjk",
			expected: 88,
		},
		{
			name:     "single digit",
			input:    "8",
			expected: 88,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			out, _ := parseStringToInt(tc.input)
			assert.Equal(t, tc.expected, out)
		})
	}
}

func TestParseStringstoIntsWithStrings(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "first and last",
			input:    "1asdfad4",
			expected: 14,
		},
		{
			name:     "embedded",
			input:    "asdf3dvjk4lkh",
			expected: 34,
		},
		{
			name:     "embedded and last",
			input:    "asdf3dvjklkh4",
			expected: 34,
		},
		{
			name:     "first and embedded",
			input:    "3asdfdvjklkh4",
			expected: 34,
		},
		{
			name:     "multi",
			input:    "3adsfasd4asdfasdf6",
			expected: 36,
		},
		{
			name:     "only digits",
			input:    "24601",
			expected: 21,
		},
		{
			name:     "single digits",
			input:    "abcd8sgjk",
			expected: 88,
		},
		{
			name:     "single digit",
			input:    "8",
			expected: 88,
		},
		{
			name:     "text digit",
			input:    "eight",
			expected: 88,
		},
		{
			name:     "mix text digit",
			input:    "eight1",
			expected: 81,
		},
		{
			name:     "two text digit",
			input:    "eightnine",
			expected: 89,
		},
		{
			name:     "big mix text digit",
			input:    "asdfasveightninesix3four",
			expected: 84,
		},
		{
			name:     "example 1",
			input:    "two1nine",
			expected: 29,
		},
		{
			name:     "example 2",
			input:    "eightwothree",
			expected: 83,
		},
		{
			name:     "example 3",
			input:    "abcone2threexyz",
			expected: 13,
		},
		{
			name:     "example 4",
			input:    "xtwone3four",
			expected: 24,
		},
		{
			name:     "example 5",
			input:    "4nineeightseven2",
			expected: 42,
		},
		{
			name:     "example 6",
			input:    "zoneight234",
			expected: 14,
		},
		{
			name:     "example 7",
			input:    "7pqrstsixteen",
			expected: 76,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			out, _ := parseStringstoIntsWithStrings(tc.input)
			assert.Equal(t, tc.expected, out)
		})
	}
}
