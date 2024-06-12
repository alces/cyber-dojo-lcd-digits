package lcd

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

var digitsTestResults = []struct{
    argument int
    expected string
} {
    {1, "   \n"+
        "  |\n"+
        "  |"},
}

func TestDigits(t *testing.T) {
    for _, r := range digitsTestResults {    
        assert.Equal(t, r.expected, Digits(r.argument), r.argument)
    }    
}