/*
https://app.codesignal.com/interview-practice/task/tNajEfDckCgSvrjQp
*/
func modInverse(a int64, m int64) int64 {
    // extended euclid algorithm. iterative to avoid slow recursion
    m0 := m
    var y int64 = 0
    var x int64 = 1

    if a == 0 {
        return -1
    }

    if m == 1 {
        return 0
    }

    for a > 1 {
        if m == 0 {
            return -1
        }
        // quotient
        q := a / m
        t := m

        // m is reminder now
        m = a % m
        a = t
        t = y

        // update x, y
        y = x - q * y
        x = t
    }

    // make x positive
    if x < 0 {
        x = x + m0
    }
    return x
}
