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

func TestThree(t *testing.T) {
    expected := number{
        " _ ",
        " _|",
        " _|",
    }
    
    assert.Equal(t, expected, digit(3))
}

func TestFour(t *testing.T) {
    expected := number{
        "   ",
        "|_|",
        "  |",
    }
    
    assert.Equal(t, expected, digit(4))
}

func TestFive(t *testing.T) {
    expected := number{
        " _ ",
        "|_ ",
        " _|",
    }
    
    assert.Equal(t, expected, digit(5))
}

func TestSix(t *testing.T) {
    expected := number{
        " _ ",
        "|_ ",
        "|_|",
    }
    
    assert.Equal(t, expected, digit(6))
}

func TestSeven(t *testing.T) {
    expected := number{
        " _ ",
        "  |",
        "  |",
    }
    
    assert.Equal(t, expected, digit(7))
}

func TestEight(t *testing.T) {
    expected := number{
        " _ ",
        "|_|",
        "|_|",
    }
    
    assert.Equal(t, expected, digit(8))
}

func TestNine(t *testing.T) {
    expected := number{
        " _ ",
        "|_|",
        "  |",
    }
    
    assert.Equal(t, expected, digit(9))
}
