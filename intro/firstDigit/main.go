/*
https://codefights.com/arcade/intro/level-8/rRGGbTtwZe2mA8Wov
Find the leftmost digit that occurs in a given string.

Example

For inputString = "var_1__Int", the output should be
firstDigit(inputString) = '1';
For inputString = "q2q-q", the output should be
firstDigit(inputString) = '2';
For inputString = "0ss", the output should be
firstDigit(inputString) = '0'.
*/
import . "strconv"

func firstDigit(inputString string) string {
    runes := []rune(inputString)
    for _, v := range runes {
        if _, err := Atoi(string(v)); err == nil {
            return string(v)
        }
    }
    return "0"
}
