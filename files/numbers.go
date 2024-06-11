package lcd

import (
    "strings"
)

func join(digits []number) (result number) {
    numDigits := len(digits)
    row := make([]string, numDigits)
    
    for i := 0; i < len(result); i++ {
        for j := 0; j < numDigits; j++ {
            row[j] = digits[j][i]
        }
        
        result[i] = strings.Join(row, " ")
    }
    
    return
}

func row(digits []number, index int) []string {
    return []string{}
}