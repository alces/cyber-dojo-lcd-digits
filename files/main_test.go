package lcd

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestDigits(t *testing.T) {
    expected := "   \n"+
                "  |\n"+
                "  |"
    
    assert.Equal(t, expected, Digits(1), "one")
}