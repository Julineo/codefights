/*
Given an array of integers, find the maximal absolute difference between any two of its adjacent elements.

Example

For inputArray = [2, 4, 1, 0], the output should be
arrayMaximalAdjacentDifference(inputArray) = 3.

Input/Output

[execution time limit] 4 seconds (go)

[input] array.integer inputArray

Guaranteed constraints:
3 = inputArray.length = 10,
-15 = inputArray[i] = 15.
*/

func arrayMaximalAdjacentDifference(inputArray []int) int {
    mxdiff := 0
    for i := 1; i < len(inputArray); i++ {
        tmp := abs(inputArray[i-1] - inputArray[i])
        if tmp > mxdiff {
            mxdiff = tmp
        }
    }
    return mxdiff
}

func abs(n int) int {
    if n < 0 {return -n}
    return n
}
