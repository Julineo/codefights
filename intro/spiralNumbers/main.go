func spiralNumbers(n int) [][]int {
    rows := make([][]int, n)
    for i := range rows {
        rows[i] = make([]int, n)
    }
    k, d := 0, 0
    i, j := 0, -1
    for {
        //right
        for c := 0; c < n - d; c++ {
            j++
            k++
            rows[i][j] = k
        }
        if k >= n * n { break }
        d++

        //down
        for c := 0; c < n - d; c++ {
            i++
            k++
            rows[i][j] = k
        }
        if k >= n * n { break }

        //left
        for c := 0; c < n - d; c++ {
            j--
            k++
            rows[i][j] = k
        }
        if k >= n * n { break }
        d++

        //up
        for c := 0; c < n - d; c++ {
            i--
            k++
            rows[i][j] = k
        }
        if k >= n * n { break }
    }
    return rows
}
