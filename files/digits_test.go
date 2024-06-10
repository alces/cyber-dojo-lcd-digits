package lcd

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestDigit(t *testing.T) {
    zero := number{
        " _ ",
        "| |",
        "|_|",
    }

    assert.Equal(t, zero, digit(0))
}
