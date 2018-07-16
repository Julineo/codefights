/*
https://app.codesignal.com/arcade/intro/level-9/hoLtYWbjdrD2PF6yo/description

Let's define digit degree of some positive integer as the number of times we need to replace this number with the sum of its digits until we get to a one digit number.

Given an integer, find its digit degree.

Example

For n = 5, the output should be
digitDegree(n) = 0;
For n = 100, the output should be
digitDegree(n) = 1.
1 + 0 + 0 = 1.
For n = 91, the output should be
digitDegree(n) = 2.
9 + 1 = 10 -> 1 + 0 = 1.
*/
func digitDegree(n int) (degree int) {
    for n/10 > 0 {
        n = digitSum(n)
        degree++
    }
    return
}

// returns sum of digits 100 -> 1 + 0 + 1 = 2
func digitSum(n int) (sum int) {
    for n > 0 {
        rem := n%10
        n = n/10
        fmt.Println(rem)
        sum += rem
    }
    return
}
