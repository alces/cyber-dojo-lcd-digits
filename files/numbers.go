package lcd

import (
    "strings"
)

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

func chopTail(number int) (head, tail int) {
    return number/10, number%10
}

func splitNumber(number int) (result []int) {
    head := number
    
    for head > 0 {
        head, tail := chopTail(head)
        result = append(result, tail)
    }
    
    return
}