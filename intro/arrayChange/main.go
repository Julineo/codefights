/*You are given an array of integers. On each move you are allowed to increase exactly one of its element by one. Find the minimal number of moves required to obtain a strictly increasing sequence from the input.

Example

For inputArray = [1, 1, 1], the output should be
arrayChange(inputArray) = 3.

Input/Output

[execution time limit] 4 seconds (go)

[input] array.integer inputArray

Guaranteed constraints:
3 = inputArray.length = 105,
-105 = inputArray[i] = 105.
*/

func arrayChange(inputArray []int) int {
    count := 0
    for i := 1; i < len(inputArray); i++ {
        if inputArray[i] > inputArray[i-1] {
            continue
        }
        for inputArray[i] <= inputArray[i-1] {
            inputArray[i]++
            count++
        }
    }
    return count
}
