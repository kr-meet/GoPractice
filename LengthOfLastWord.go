func lengthOfLastWord(s string) int {
    
    split := strings.Fields(s)
    fmt.Println(split)
    return len(split[len(split) - 1])
}