package d1p2

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
    fmt.Println("result of day1 part2: ", sum)
}

// numberTargets are number variants: "1", "2", "3", "4", "5", "6", "7", "8", "9"
var wordTargets = [9]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
var reversedWordTargets = [9]string{"eno", "owt", "eerht", "ruof", "evif", "xis", "neves", "thgie", "enin"} 

func findFirstTarget(slice string) int {
    firstNumber, firstNumberIdx := findFirstNumberTarget(slice)
    firstWordTarget, firstWordIdx := findFirstWordTarget(slice, wordTargets)

    noFirstWord := firstWordIdx == -1
    noFirstNumber := firstNumberIdx == -1

    if noFirstWord && noFirstNumber {
        panic("findFirstTarget: did not find number or word in slice: " + slice)
    }

    if firstWordIdx < firstNumberIdx {
        return firstWordTarget
    } else {
        return firstNumber
    }
}

func findLastTarget(slice string) int {
    reversedSlice := util.ReverseRunes(slice)

    firstNumber, firstNumberIdx := findFirstNumberTarget(reversedSlice)
    firstWordTarget, firstWordIdx := findFirstWordTarget(reversedSlice, reversedWordTargets)

    noFirstWord := firstWordIdx == -1
    noFirstNumber := firstNumberIdx == -1

    if noFirstWord && noFirstNumber {
        panic("findFirstTarget: did not find number or word in slice: " + slice)
    }

    if firstWordIdx < firstNumberIdx {
        return firstWordTarget
    } else {
        return firstNumber
    }

}

func findFirstNumberTarget(slice string) (firstNumber int, idx int) {
    idx = -1
    for i, char := range slice {
        if num, isNumber := util.RuneIsNumber(char); isNumber {
            firstNumber = num
            idx = i
            break
        }
    }
    return
}

func findFirstWordTarget(slice string, wordTargets [9]string) (firstWordNumber int, smallestIdx int) {
    var foundWordTargetsIndexes [9]int
    var idxOfWordTarget int

    for i, numberWord := range wordTargets {
        idxOfWordTarget = strings.Index(slice, numberWord)
        foundWordTargetsIndexes[i] = idxOfWordTarget
    }

    smallestIdx = len(slice)
    for _, idx := range foundWordTargetsIndexes {
        if idx != -1 && idx < smallestIdx {
            smallestIdx = idx
        }
    }

    firstWordNumber = util.IndexOf(foundWordTargetsIndexes[:], smallestIdx) + 1
    return
}

