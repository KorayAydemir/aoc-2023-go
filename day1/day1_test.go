package day1

import (
	"testing"
)


func TestFindFirstWordTarget(t *testing.T) {
    wordTargets := [9]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

    testCases := []struct {
        input          string
        expectedNumber int
        expectedFound  bool
    }{
        {"one hundred", 1, true},
        {"This is two", 2, true},
        {"No numbers", 0, false}, 
        {"Four", 0, false},
        {"five", 5, true},
        {"evif", 0, false},
    }

    for _, tc := range testCases {
        num, found := findFirstWordTarget(tc.input, wordTargets)
        if num != tc.expectedNumber || found != tc.expectedFound {
            t.Errorf("For input '%s', expected (%d, %t), but got (%d, %t)", tc.input, tc.expectedNumber, tc.expectedFound, num, found)
        }
    }
}
