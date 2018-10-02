/*
https://app.codesignal.com/interview-practice/task/D7n2Cj9JDYbQzWiZE

working algorithm, but exceeds time limit. Needs to be redone
*/

func primesSum2(n int) int {
    if n == 1 { return 0 }
    if n == 2 { return 2 }
    if n == 3 { return 5 }
    res := 5

    primes := []int{2, 3}

    flag := true
    for c := 5; c <= n; c = c + 2 {
        flag = true
        for _, a := range primes {
            if c % a == 0 {
                flag = false
                break
            }
        }
        if flag {
            primes = append(primes, c)
            res += c
        }
    }
    fmt.Println(10000000007)
    return res % (10 * 10 * 10 * 10 * 10 * 10 * 10 * 10 * 10 * 10 + 7)
}
