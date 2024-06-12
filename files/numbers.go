package lcd

import (
    "strings"
)

type number [3]string

func (n number) String() string {
    return strings.Join(n, "\n")
}

func join(digits []number) (result number) {
    for i := 0; i < len(result); i++ {
        result[i] = strings.Join(row(digits, i), " ")
    }
    
    return
}

func row(digits []number, index int) []string {
    numDigits := len(digits)
    result := make([]string, numDigits)
    
    for i := 0; i < numDigits; i++ {
        result[i] = digits[i][index]
    }
    
    return result
}

func splitNumber(number int) (result []int) {
    for head := number; head > 0; head /= 10 {
        result = append([]int{head%10}, result...)
    }
    
    return
}