

func perfectSquareOrCube(number int) int {
    ss := permutations(strconv.Itoa(number))
    var res int
    for i := 0; i < len(ss); i++ {
        fmt.Println(ss[i])
        n, _ := strconv.Atoi(string(ss[i]))
        fmt.Println("n: ", n)
        if n <= 900 {
            fmt.Println("n <= 900 ", n)
            for j := 1; j <= 30; j++ {
                if (j * j) == n || (j * j * j) == n {
                    fmt.Println("j: ", j)
                    res++
                    break
                }
            }
        }
        if res == 2 {
            return res
        }
    }
    return res
}

 func join(ins []rune, c rune) (result []string) {
    for i := 0; i <= len(ins); i++ {
        result = append(result, string(ins[:i])+string(c)+string(ins[i:]))
    }
    return
}
  
func permutations(testStr string) []string {
    var n func(testStr []rune, p []string) []string
    n = func(testStr []rune, p []string) []string{
        if len(testStr) == 0 {
            return p
        }else {
            result := []string{}
            for _, e := range p {
                result = append(result, join([]rune(e), testStr[0])...)
            }
            return n(testStr[1:], result)
        }
    }
  
    output := []rune(testStr)
    return n(output[1:], []string{string(output[0])})
}
