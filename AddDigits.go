func addDigits(num int) int {
    
    for num > 9 {
        temp := num
        num = 0
        for temp > 0 {
            rem := temp % 10
            temp /= 10
            num += rem
        }
    }

    return num
}