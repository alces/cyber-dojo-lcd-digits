package lcd

import (
    "testing"
    "github.com/stretchr/testify/assert"
)


func TestChopTail(t *testing.T) {
    head, tail := chopTail(5)

    assert.Equal(t, 0, head, "unexpected head")
    assert.Equal(t, 5, tail, "unexpected tail")
}

var digits423 = []number{
    {
        "   ",
        "|_|",
        "  |",
    },
    {
        " _ ",
        " _|",
        "|_ ",
    },
    {
        " _ ",
        " _|",
        " _|",
    },
}

func TestJoin423(t *testing.T) {
    expected := number{
        "     _   _ ",
        "|_|  _|  _|",
        "  | |_   _|",
    }
    
    assert.Equal(t, expected, join(digits423), "423")
}

func TestJoinEmpty(t *testing.T) {
    assert.Equal(t, number{}, join([]number{}), "empty number")
}

func TestRow(t *testing.T) {
    expected := []string{
        "|_|", " _|", " _|",
    }
    
    assert.Equal(t, expected, row(digits423, 1), "middle row of 423")
}