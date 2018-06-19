/*
Check if the given matrix is an identity matrix.

Example

For

matrix = [[1, 0, 0], 
          [0, 1, 0], 
          [0, 0, 1]]
the output should be
isIdentityMatrix(matrix) = true;

For

matrix = [[1, 0, 0], 
          [0, 0, 0], 
          [0, 0, 1]]
the output should be
isIdentityMatrix(matrix) = false.
*/

func isIdentityMatrix(matrix [][]int) bool {
    m := len(matrix)
    n := len(matrix[0])
    for i := 0; i < m ; i++ {
        for j := 0; j < n; j++ {
            if matrix[i][j] != 1 && i == j {
                return false
            }
            if  matrix[i][j] != 0 && i != j {
                return false
            } 
        }
    }
     return true
}
