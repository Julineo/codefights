/*
https://app.codesignal.com/arcade/intro/level-10/ppZ9zSufpjyzAsSEx/description

Given a string, find the shortest possible string which can be achieved by adding characters to the end of initial string to make it a palindrome.

Example

For st = "abcdc", the output should be
buildPalindrome(st) = "abcdcba".
*/

func buildPalindrome(st string) string {
    at := []byte(st)
    bt := []byte{}
    if isPalindrome(at) {
        return st
    }
    for i := 0; i < len(at); i++ {
        bt = append([]byte{at[i]}, bt...)
        rt := append(at, bt...)
        if isPalindrome(rt) {
            return string(rt)
        }
    }
    return st
}

func isPalindrome(at []byte) bool {
    for i, j := 0, len(at)-1; i < j; i, j = i + 1, j - 1 {
        if at[i] != at[j] {
            return false
        }
    }
    return true
}
