package lcd

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

var chopTailResults = []struct{
    argument int
    head     int
    tail     int
} {
    {5, 0, 5},
    {17, 1, 7},
    {590, 59, 0},
}

func TestChopTail(t *testing.T) {   
    for _, r := range chopTailResults {
        actualHead, actualTail := chopTail(r.argument)

        assert.Equal(t, r.head, actualHead, "unexpected head")
        assert.Equal(t, r.tail, actualTail, "unexpected tail")
    }
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