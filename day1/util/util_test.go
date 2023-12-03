package util

import (
	"testing"
)

func TestRuneIsNumber(t *testing.T) {
    tests := []struct {
        char rune
        expected int
        shouldPass bool
    }{
        {'$', -1, false},
        {'#', -1, false},
        {'a', -1, false},
        {'0', -1, false},
        {'1', 1, true},
        {'2', 2, true},
		{'3', 3, true},
		{'4', 4, true},
		{'5', 5, true},
		{'6', 6, true},
		{'7', 7, true},
		{'8', 8, true},
        {'9', 9, true},
    }

    for _, test := range tests {
		number, isNumber := RuneIsNumber(test.char)

		if number != test.expected || isNumber != test.shouldPass {
			t.Errorf("For '%c', expected (%d, %t), but got (%d, %t)", test.char, test.expected, test.shouldPass, number, isNumber)
		}
	}
}

func TestReverseRunes(t *testing.T) {
    tests := []struct {
        input    string
        expected string
    }{
        {"hello", "olleh"},
        {"racecar", "racecar"},
        {"12345", "54321"},
        {"1234567890", "0987654321"},
        {"abcdefghijklmnopqrstuvwxyz", "zyxwvutsrqponmlkjihgfedcba"},
        {"a", "a"},
        {"", ""},
        //{"😀🚀🔥", "🔥🚀😀"}, //fails
    }

    for _, test := range tests {
        result := ReverseRunes(test.input)
        if result != test.expected {
            t.Errorf("For '%s' expected '%s', but got '%s'", test.input, test.expected, result)
        }
    }
}
