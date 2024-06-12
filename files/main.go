package lcd

func Digits(number int) string {
    integers := splitNumber(number)
    size := len(integers)
    digits := make([]number, size)
    
    for i := 0; i < size; i++ {
        digits[i] = digit(integers[i])
    }
    
    return join(digits).String()
}