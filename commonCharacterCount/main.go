func commonCharacterCount(s1 string, s2 string) int {
    count := 0
    r1 := []rune(s1)
    r2 := []rune(s2)
    for _, v1 := range(r1) {
        for i, v2 := range(r2) {
            if v1 == v2 {
                count++
                r2[i] = '"'
                break
            }
        }
    }
    return count
}
