package lcd

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestJoin423(t *testing.T) {
    arguments := []number{
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
    
    expected := number{
        "     _   _ ",
        "|_|  _|  _|",
        "  | |_   _|",
    }
    
    assert.Equal(t, expected, join(arguments), "423")
}
    