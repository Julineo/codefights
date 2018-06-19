/*
https://codefights.com/arcade/intro/level-7/ZFnQkq9RmMiyE6qtq/description
Given a sorted array of integers a, find an integer x from a such that the value of

abs(a[0] - x) + abs(a[1] - x) + ... + abs(a[a.length - 1] - x)
is the smallest possible (here abs denotes the absolute value).
If there are several possible answers, output the smallest one.

Example

For a = [2, 4, 7], the output should be
absoluteValuesSumMinimization(a) = 4.
*/

func absoluteValuesSumMinimization(a []int) int {
    if len(a)%2 == 0 {
        return a[len(a)/2 - 1]
    }
    return a[len(a)/2]
}
