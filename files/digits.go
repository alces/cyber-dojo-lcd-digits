package lcd

func digit(num int) (result number) {
    switch(num) {
    case 0:
        result = number{
            " _ ",
            "| |",
            "|_|",
        }
    case 1:
        result = number{
            "   ",
            "  |",
            "  |",
        }
    case 2:
        result = number{
            " _ ",
            " _|",
            "|_ ",
        }
    case 3:
        result = number{
            " _ ",
            " _|",
            " _|",
        }
    case 4:
        result = number{
            "   ",
            "|_|",
            "  |",
        }
    case 5:
        result = number{
            " _ ",
            "|_ ",
            " _|",
        }
    case 6:
        result = number{
            " _ ",
            "|_ ",
            "|_|",
        }
    case 7:
        result = number{
            " _ ",
            "  |",
            "  |",
        }
    case 8:
        result = number{
            " _ ",
            "|_|",
            "|_|",
        }
    case 9:
        result = number{
            " _ ",
            "|_|",
            "  |",
        }
    }
    
    return result
}
