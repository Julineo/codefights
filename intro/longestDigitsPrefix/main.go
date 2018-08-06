/*
https://app.codesignal.com/arcade/intro/level-9/AACpNbZANCkhHWNs3/description

Given a string, output its longest prefix which contains only digits.

Example

For inputString="123aa1", the output should be
longestDigitsPrefix(inputString) = "123".
*/

import "unicode"

func longestDigitsPrefix(inputString string) string {
    max := 0
    for _, v := range inputString {
        if unicode.IsDigit(v) {
            max++
        } else {
            break
        }
    }
    return inputString[:max]
}
