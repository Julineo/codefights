/*
https://app.codesignal.com/arcade/code-arcade/intro-gates/SZB5XypsMokGusDhX

Given an integer n, return the largest number that contains exactly n digits.

Example

For n = 2, the output should be
largestNumber(n) = 99.
*/

func largestNumber(n int) int {
    r := 1
    for i := 0; i < n; i++ {
        r *= 10
    }
    return r - 1
}
