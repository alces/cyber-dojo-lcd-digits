package lcd

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

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

var splitNumberResults = []struct{
    argument int
    expected []int
} {
    {5, []int{5}},
    {17, []int{1, 7}},
    {130, []int{1, 3, 0}},
}

func TestSplitNumber(t *testing.T) {
    for _, r := range splitNumberResults {
        assert.Equal(t, r.expected, splitNumber(r.argument), r.argument)
    }    
}