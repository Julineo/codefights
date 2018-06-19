func fifthBitA2PlusB2(a int, b int) int {
    //fmt.Printf("%032b\n", a)
    //fmt.Printf("%032b\n", b)
    s := fmt.Sprintf("%064b\n", (a * a + b * b) >> 5 )
    fmt.Println(s)
    fmt.Println(s[len(s)-2])
    z, _ := strconv.Atoi(string(s[len(s)-2]))
    return z
}
