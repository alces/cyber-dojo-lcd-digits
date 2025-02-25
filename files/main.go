package lcd

func Digits(num int) string {
    integers := splitNumber(num)
    size := len(integers)
    digits := make([]number, size)
    
    for i := 0; i < size; i++ {
        digits[i] = digit(integers[i])
    }
    
    return join(digits).String()
}