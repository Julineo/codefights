/*
https://app.codesignal.com/interview-practice/task/PLCrGrJmBxQdj8QKX

You are supposed to label a bunch of boxes with numbers from 0 to n, and all the labels are stored in the array arr. Unfortunately one of the labels is missing. Your task is to find it.

Example

For arr = [3, 1, 0], the output should be
missingNumber(arr) = 2.
*/

import . "sort"

func missingNumber(arr []int) int {
    Ints(arr)
    for i := 0; i < len(arr); i++ {
        if arr[i] != i { return i}
    }
    return len(arr)
}
