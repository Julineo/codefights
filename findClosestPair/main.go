/*
Given an array of integers numbers, we'd like to find the closest pair of elements that add up to sum. Return the distance between the closest pair (absolute difference between the two indices). If there isn't a pair that adds up to sum, return -1.

Example

For numbers = [1, 0, 2, 4, 3, 0] and sum = 5 the output should be findClosestPair(numbers, sum) = 2. 1 and 4 have a sum of 5, but 2 and 3 are closer.

for numbers = [2, 3, 7] and sum = 8 the output should be findClosestPair(numbers, sum) = -1. There are no pairs that have a sum of 8.
*/
func findClosestPair(numbers []int, sum int) int {
    min := len(numbers) + 1
    got := false
    m := make(map[int]int)
    for i := 0; i < len(numbers); i++ {
        j, ok := m[sum - numbers[i]]
            if ok {
                if i - j < min {
                    min = i - j
                    got = true

            }
        }
        m[numbers[i]] = i
    } 
    if got {
        return min
    }
    return -1
}
