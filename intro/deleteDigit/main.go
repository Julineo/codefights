/*
https://app.codesignal.com/arcade/intro/level-11/vpfeqDwGZSzYNm2uX

Given some integer, find the maximal number you can obtain by deleting exactly one digit of the given number.

Example

For n = 152, the output should be
deleteDigit(n) = 52;
For n = 1001, the output should be
deleteDigit(n) = 101.
*/
func deleteDigit(n int) int {
    ns := strconv.Itoa(n)
    max := 0
    for i := 0; i < len(s); i++ {
        s := string(ns[:i]) + string(ns[i+1:])
        r, _ := strconv.Atoi(string(s))
        if r > max {
            max = r
        }
    }
    return max
}

/*
func deleteDigit(n int) int {
    s := []byte(strconv.Itoa(n))
    max := 0
    for i := 0; i < len(s); i++ {
        z := make([]byte, len(s))
        copy(z, s)
        c := append(z[:i], z[i+1:]...)
        r, _ := strconv.Atoi(string(c))
        if r > max {
            max = r
        }
    }
    return max
}
*/
