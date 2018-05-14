func matrixElementsSum(matrix [][]int) int {
    var cost int
    for j := 0; j < len(matrix[0]); j++ {
        uh := false
        for i := 0; i < len(matrix); i++ {
            if matrix[i][j] > 0 && !uh {
                cost = cost + matrix[i][j]
                uh = false
            }
            if matrix[i][j] == 0{
                uh = true
            }
        }
    }
    return cost
}
