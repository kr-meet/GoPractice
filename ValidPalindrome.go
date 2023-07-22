func isPalindrome(s string) bool { 

    s = regexp.MustCompile(`[^a-zA-Z0-9 ]+`).ReplaceAllString(s, "")
    s = strings.Replace(s, " ", "", -1)
    s = strings.ToLower(s)

    i:= 0
    j:= len(s) - 1
    
    for i < j {
        if s[i] != s[j] {
            return false
        }
        i++
        j--
    }

    return true
}