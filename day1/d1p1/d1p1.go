package d1p1


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
    fmt.Println("result of day1 part1: ", sum)
}

// targets are: "1", "2", "3", "4", "5", "6", "7", "8", "9"

func findFirstTarget(slice string) int {
    firstNumber := findFirstNumberTarget(slice)

    return firstNumber
}

func findLastTarget(slice string) int {
    reversedSlice := util.ReverseRunes(slice)

    firstNumber := findFirstNumberTarget(reversedSlice)

    return firstNumber
}

func findFirstNumberTarget(slice string) (firstNumber int) {
    for _, char := range slice {
        if num, isNumber := util.RuneIsNumber(char); isNumber {
            firstNumber = num
            break
        }
    }
    return
}
