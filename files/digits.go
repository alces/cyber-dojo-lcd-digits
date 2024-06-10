package lcd

func digit(num int) [3]string {
    var result [3]string

    switch(num) {
    case 0:
        result = [3]string{
            " _ ",
            "| |",
            "|_|",
        }
    }
    
    return result
}
