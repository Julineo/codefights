func checkPalindrome(inputString string) bool {
    for i, j := 0, len(inputString) - 1; i < j; i, j = i + 1, j - 1 {
        if inputString[i] != inputString[j] {
            return false
        }
    }
    return true
}
