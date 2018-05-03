func almostIncreasingSequence(s []int) bool {
    count := 0
    maxOne := -1000010
    maxTwo := -1000010
    for i := 0; i < len(s); i++ {
        fmt.Println("MaxOne: ", maxOne, "MaxTwo: ", maxTwo)
        fmt.Printf("s[%v]: %v\n", i, s[i])
        if s[i] > maxOne {
            fmt.Println("one")
            maxTwo = maxOne
            maxOne = s[i]
            continue
        }
        if s[i] > maxTwo {
            fmt.Println("two")
            maxOne = s[i]
            count++
            continue
        }
        count++
        fmt.Println("count:", count)
    }
    fmt.Println("count:", count)
    fmt.Println("len(s) - count:", len(s) - count)
    return count < 2
}

