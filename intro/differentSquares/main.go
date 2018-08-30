/*
https://app.codesignal.com/arcade/intro/level-12/fQpfgxiY6aGiGHLtv

Given a rectangular matrix containing only digits, calculate the number of different 2 × 2 squares in it.

Example

For

matrix = [[1, 2, 1],
          [2, 2, 2],
          [2, 2, 2],
          [1, 2, 3],
          [2, 2, 1]]
the output should be
differentSquares(matrix) = 6.

Here are all 6 different 2 × 2 squares:

1 2
2 2
2 1
2 2
2 2
2 2
2 2
1 2
2 2
2 3
2 3
2 1
*/
func differentSquares(matrix [][]int) int {
    m := len(matrix)
    n := len(matrix[0])
    if m < 2 || n < 2 {
        return 0
    }
    mp := make(map[int]int)
    for i := 0; i < m - 1; i++ {
        for j := 0; j < n - 1; j++ {
            s := matrix[i][j] + matrix[i+1][j]*10 + matrix[i][j+1]*100 + matrix[i+1][j+1]*1000
            mp[s]++
        }
    }
    return len(mp)
}
