func alternatingSums(a []int) []int {
    r := []int{0, 0}
    for i, v := range a {
        r[i%2] += v
    }
    return r
}

/*func alternatingSums(a []int) []int {
    s1, s2 := 0, 0
    for i, v := range a {
        if i%2 == 0 {
            s1 += v
            continue
        }
        s2 +=v
    }
    return []int{s1, s2}
}*/
