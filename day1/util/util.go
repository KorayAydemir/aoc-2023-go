package util

import "strconv"

func RuneIsNumber(char rune) (int, bool) {
    if char < '1' || char > '9' {
        return -1, false
    }

    return int(char - '0'), true
}

func ReverseRunes(s string) string {
    r := []rune(s)
    for i := 0; i < len(s)/2; i++ {
        j := len(s) - i - 1
        r[i], r[j] = r[j], r[i]
    }
    return string(r)
}

func TransformStringToInt (str string) int {
    result, err := strconv.Atoi(str)
    if err != nil {
        panic(err)
    }
    return result
}
