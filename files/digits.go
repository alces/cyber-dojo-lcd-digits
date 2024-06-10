package lcd

func digit(num int) [3]string {
    switch(num) {
    case 0:
        return [3]string{
            " _ ",
            "| |",
            "|_|",
        }
    default:
        return [3]string{}
    }
}
