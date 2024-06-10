package lcd

type number [3]string

func digit(num int) (result number) {
    switch(num) {
    case 0:
        result = number{
            " _ ",
            "| |",
            "|_|",
        }
    }
    
    return result
}
