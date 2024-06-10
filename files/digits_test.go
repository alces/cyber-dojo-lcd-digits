package lcd

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestZero(t *testing.T) {
    expected := number{
        " _ ",
        "| |",
        "|_|",
    }

    assert.Equal(t, expected, digit(0))
}

func TestOne(t *testing.T) {
    expected := number{
        "   ",
        "  |",
        "  |",
    }
    
    assert.Equal(t, expected, digit(1))
}

func TestTwo(t *testing.T) {
    expected := number{
        " _ ",
        " _|",
        "|_ ",
    }
    
    assert.Equal(t, expected, digit(2))
}
