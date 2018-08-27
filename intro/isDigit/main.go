/*
https://app.codesignal.com/arcade/intro/level-11/Zr2XXEpkBPtYWqPJu

Determine if the given character is a digit or not.

Example

For symbol = '0', the output should be
isDigit(symbol) = true;
For symbol = '-', the output should be
isDigit(symbol) = false.
*/

func isDigit(symbol string) bool {
    if _, err := strconv.Atoi(symbol); err != nil {
        return false
    }
    return true
}
