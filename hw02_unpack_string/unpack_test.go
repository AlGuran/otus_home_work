package hw02unpackstring

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnpack(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{input: "a4bc2d5e", expected: "aaaabccddddde"},
		{input: "abccd", expected: "abccd"},
		{input: "", expected: ""},
		{input: "aaa0b", expected: "aab"},
		// uncomment if task with asterisk completed
		{input: `qwe\4\5`, expected: `qwe45`},
		{input: `qwe\45`, expected: `qwe44444`},
		{input: `qwe\\5`, expected: `qwe\\\\\`},
		{input: `qwe\\\3`, expected: `qwe\3`},
		// user cases
		{input: "z1", expected: "z"},
		{input: "", expected: ""},
		{input: "a", expected: "a"},
		{input: "я", expected: "я"},
		{input: "AbC3", expected: "AbCCC"},
		{input: "nn3", expected: "nnnn"},
		{input: "\nn\t3", expected: "\nn\t\t\t"},
		{input: "ё", expected: "ё"},
		{input: "'2^3&1", expected: "''^^^&"},
		{input: "-1", expected: "-"},
		{input: " 3 2 1", expected: "      "},
		{input: `\10\20\30`, expected: ""},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.input, func(t *testing.T) {
			result, err := Unpack(tc.input)
			require.NoError(t, err)
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestUnpackInvalidString(t *testing.T) {
	invalidStrings := []string{"3abc", "45", "aaa10b", "0", "9", "0a", "0.00000", "-0.00000"}
	for _, tc := range invalidStrings {
		tc := tc
		t.Run(tc, func(t *testing.T) {
			_, err := Unpack(tc)
			require.Truef(t, errors.Is(err, ErrInvalidString), "actual error %q", err)
		})
	}
}
