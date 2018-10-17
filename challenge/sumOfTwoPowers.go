/*
https://app.codesignal.com/challenge/TnnSSXh97pX5ypbaJ

For a given number n, your task is to find how it can be expressed as the sum of two integers each raised to the power of k.

You are to return the list of possibilities as an array of 2-element arrays of summands, sorted lexicographically.

Example

For n = 17 and k = 2, the output should be sumOfTwoPowers(n, k) = [[1, 16]].

Since 17 can be expressed as the sum of two squares, 1 ** 2 + 4 ** 2, we return [[1, 4]].

For n = 27 and k = 3, the output should be sumOfTwoPowers(n, k) = [[0, 27]].

Since 27 can be expressed as the sum of two cubes, 0 ** 3 + 3 ** 3, we return [[0, 27]].

For n = 16 and k = 5, the output should be sumOfTwoPowers(n, k) = [].

There are no ways to express 16 as a sum of two integers raised to the 5th power, so we return an empty array, [].

For n = 65 and k = 2, the output should be sumOfTwoPowers(n, k) = [[1, 64], [16, 49]].

There are two ways to express 65 as a sum of two squares: 1 ** 2 + 8 ** 2 or 4 ** 2 + 7 ** 2, so we return [[1, 64], [16, 49]].
*/
import . "math"

func sumOfTwoPowers(n int, k int) [][]int {
    r := [][]int{}
    p := int(Pow(float64(n), 1/float64(k))) + 1
    for i := 0; i < p;  i++ {
        for j := i; j <= p; j++ {
            a, b := i, j
            for p := 1; p < k; p++ {
                a *= i
                b *= j
            }
            if a + b == n {
                r = append(r, []int{a, b})
            }
        }
    }
    return r
}
