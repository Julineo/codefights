/*
https://codefights.com/arcade/intro/level-7/PTWhv2oWqd6p4AHB9/description
Given an array of equal-length strings, check if it is possible to rearrange the strings in such a way that after the rearrangement the strings at consecutive positions would differ by exactly one character.

Example

For inputArray = ["aba", "bbb", "bab"], the output should be
stringsRearrangement(inputArray) = false;

All rearrangements don't satisfy the description condition.

For inputArray = ["ab", "bb", "aa"], the output should be
stringsRearrangement(inputArray) = true.

Strings can be rearranged in the following way: "aa", "ab", "bb".
*/
func stringsRearrangement(inputArray []string) bool {
    // getting permutations
    inp := []int{}
    for i := 0; i < len(inputArray); i++ {
        inp = append(inp, i)
    }
    pmt := permutations(inp)

    // checking for difference
    for _, v := range pmt {
        flag := true
        for i := 1; i < len(v); i++ {
            if ! diffByOneCharcter(inputArray[v[i-1]] , inputArray[v[i]]) {
                flag = false
                break
            }
        }
        if flag {
            return true
        }
    }
    return false
}

// checks if there is one char difference
func diffByOneCharcter(a, b string) bool {
    if len(a) != len(b) {
        return false
    }
    if a == b {
        return false
    }
    count := 0
    for i, v := range a {
        if v != rune(b[i]) {
            count += 1
        }
    }
    if count == 1 {
        return true
    }
    return false
}

// returns permutations
func permutations(arr []int) [][]int {
    var helper func([]int, int)
    res := [][]int{}

    helper = func(arr []int, n int) {
        if n == 1{
            tmp := make([]int, len(arr))
            copy(tmp, arr)
            res = append(res, tmp)
        } else {
            for i := 0; i < n; i++ {
                helper(arr, n - 1)
                if n % 2 == 1{
                    tmp := arr[i]
                    arr[i] = arr[n - 1]
                    arr[n - 1] = tmp
                } else {
                    tmp := arr[0]
                    arr[0] = arr[n - 1]
                    arr[n - 1] = tmp
                }
            }
        }
    }
    helper(arr, len(arr))
    return res
}
