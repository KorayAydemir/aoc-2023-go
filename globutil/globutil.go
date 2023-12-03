package globutil

import (
	"fmt"
	"os"
)

func CreateAndWriteToFile(input string) {
	file, fileCreateErr := os.Create("/Users/kaydemir/code/adventofcode/2023/go/day1/output.txt")
	if fileCreateErr != nil {
		fmt.Println(fileCreateErr)
		return
	}
	defer file.Close()

	_, fileWriteErr := file.WriteString(input)
	if fileWriteErr != nil {
		fmt.Println(fileWriteErr)
		return
	}
}

func ReadInputFile(filePath string) ([]byte) {
	input, err := os.ReadFile(filePath)
	if err != nil {
        panic(err)
	}

	return input 
}
