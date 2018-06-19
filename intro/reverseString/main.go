/*
Given some integer, find the maximal number you can obtain by deleting exactly one digit of the given number.

Example

For n = 152, the output should be
deleteDigit(n) = 52;
For n = 1001, the output should be
deleteDigit(n) = 101.
Input/Output

[execution time limit] 4 seconds (go)

[input] integer n

Guaranteed constraints:
10 = n = 106.
*/

func deleteDigit(n int) int {
    max := 0
    for i := 0; i < len(strconv.Itoa(n)); i++ {
        st := strconv.Itoa(n)
        rs := remove(st, i)
        ri, _ := strconv.Atoi(rs)
        if ri > max {
            max = ri
        }
    }
    return max
}

func remove(st string, i int) string  {
    fmt.Println(st[:i], st[i+1:])
    return st[:i] + st[i+1:]
}
