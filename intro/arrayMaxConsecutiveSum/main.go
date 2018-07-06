/*
https://codefights.com/arcade/intro/level-8/Rqvw3daffNE7sT7d5/description
Given array of integers, find the maximal possible sum of some of its k consecutive elements.

Example

For inputArray = [2, 3, 5, 1, 6] and k = 2, the output should be
arrayMaxConsecutiveSum(inputArray, k) = 8.
All possible sums of 2 consecutive elements are:

2 + 3 = 5;
3 + 5 = 8;
5 + 1 = 6;
1 + 6 = 7.
Thus, the answer is 8.
*/
func arrayMaxConsecutiveSum(arr []int, k int) (max int) {
    for i := 0; i < len(arr) - k + 1; i++ {
        sum, j := 0, 0
        for j < k {
            if j + i == len(arr) {
                return
            }
            sum += arr[i + j]
            j++
        }
        if max < sum {
            max = sum
        }
    }
    return
}
