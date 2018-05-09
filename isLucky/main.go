func isLucky(n int) bool {
    s = strconv.Itoa(n)
    var l, r int
    for i, j = 0, len(s)-1; i  j; i, j = i + 1, j - 1 {
        v, _ = strconv.Atoi(string(byte(s[i])))
        l = l + v
        v, _ = strconv.Atoi(string(byte(s[j])))
        r = r + v
    }
    if l == r {
        return true
    }
    return false
}

