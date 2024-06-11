package lcd

import (
    "strings"
)

func join(digits []number) (result number) {
    numDigits := len(digits)
    row := make([]string, numDigits)
    
    for i := 0; i < len(result); i++ {
        result[i] = strings.Join(row(digits, i), " ")
    }
    
    return
}

func row(digits []number, index int) []string {
    numDigits := len(digits)
    row := make([]string, numDigits)
    
    for i := 0; i < numDigits; i++ {
        row[i] = digits[i][index]
    }
    
    return row
}