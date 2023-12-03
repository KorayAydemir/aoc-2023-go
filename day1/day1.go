package day1

import (
	"fmt"
	"strings"
    "github.com/KorayAydemir/aoc-2023-go/globutil"
    "github.com/KorayAydemir/aoc-2023-go/day1/util"
)

func Day1() {
	byteInput := globutil.ReadInputFile("/Users/kaydemir/code/adventofcode/2023/go/day1/input.txt")
    strInput := string(byteInput)

    linesOfInput := strings.Split(string(strInput), "\n")
    linesOfInput = linesOfInput[:len(linesOfInput)-1]

    sum := 0
    // merge first and last target number found on each line, then add to sum
    for _, line := range linesOfInput { 
        firstNum := findFirstTarget(line)
        lastNum := findLastTarget(line)

        numString := fmt.Sprintf("%d%d", firstNum, lastNum)
        numInt := util.TransformStringToInt(numString)

        sum += numInt
    }
    fmt.Println("result of day1: ", sum)
}

// numberTargets are number variants: "1", "2", "3", "4", "5", "6", "7", "8", "9"
var wordTargets = [9]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
//var reversedWordTargets = [9]string{"eno", "owt", "eerht", "ruof", "evif", "xis", "neves", "thgie", "enin"} // how come this isnt required ?

func findFirstTarget(slice string) int {
    firstNumber, didFindNumber := findFirstNumberTarget(slice)
    firstWordTarget, didFindWord := findFirstWordTarget(slice, wordTargets)

    if didFindNumber {
        return firstNumber
    } else if didFindWord {
        return firstWordTarget
    }
    panic("findFirstTarget: did not find number or word in slice: " + slice)
}

func findLastTarget(slice string) int {
    reversedSlice := util.ReverseRunes(slice)

    firstNumber, didFindNumber := findFirstNumberTarget(reversedSlice)
    firstWordTarget, didFindWord := findFirstWordTarget(reversedSlice, wordTargets)

    if didFindNumber {
        return firstNumber
    } else if didFindWord {
        return firstWordTarget
    }
    panic("findLastTarget: did not find number or word in slice: " + slice)
}

func findFirstNumberTarget(slice string) (firstNumber int, didFindNumber bool) {
    for _, char := range slice {
        if num, isNumber := util.RuneIsNumber(char); isNumber {
            firstNumber = num
            didFindNumber = true
            break
        }
    }
    return
}

func findFirstWordTarget(slice string, wordTargets [9]string) (firstWordNumber int, didFindWord bool) {
    for i, numberWord := range wordTargets {
        idxOfNumberOrNumberWord := strings.Index(slice, numberWord)

        if idxOfNumberOrNumberWord != -1 {
            firstWordNumber = i + 1
            didFindWord = true
            break
        }
    }
    return
}
