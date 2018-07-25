/*
https://app.codesignal.com/challenge/8FdeLisamv6cFZPAc

Hi all, I have a tiny challenge for all of you today. I was given some arrays of integers and my mission is to find the smallest product of 3 elements in an array. I will need your help on this problem.

Please note that our integers can be negative and positive.

Examples

For arr = [1, 2, 3], the output should be smallestProduct(arr) = 6

The only option is 1 * 2 * 3 = 6.

For arr = [-1, 0, -2, 3], the output should be smallestProduct(arr) = 0

There are four possible products, and the smallest one is 0:

-1 * 0 * -2 = 0
-1 * 0 * 3 = 0
-1 * -2 * 3 = 6
0 * -2 * 3 = 0
Input/Output

[execution time limit] 4 seconds (go)

[input] array.integer arr

An array of integers.

Guaranteed constraints:
3 ≤ arr.length ≤ 105
 -103 ≤ arr[i] ≤ 103

[output] integer

The smallest product of any 3 elements in the given array.
*/

import . "sort"

import . "sort"
func smallestProduct(a []int) int {
    Ints(a)
    l := len(a)
    x := a[0] * a[1] * a[2]
    y := a[0] * a[l-1] * a[l-2]
    if y < x {
        x = y
    }
    return x
}
