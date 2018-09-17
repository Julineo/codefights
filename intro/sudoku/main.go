/*
https://app.codesignal.com/arcade/intro/level-12/tQgasP8b62JBeirMS

Sudoku is a number-placement puzzle. The objective is to fill a 9 × 9 grid with digits so that each column, each row, and each of the nine 3 × 3 sub-grids that compose the grid contains all of the digits from 1 to 9.

This algorithm should check if the given grid of numbers represents a correct solution to Sudoku.

Example

For the first example below, the output should be true. For the other grid, the output should be false: each of the nine 3 × 3 sub-grids should contain all of the digits from 1 to 9.



Input/Output

[execution time limit] 4 seconds (go)

[input] array.array.integer grid

A matrix representing 9 × 9 grid already filled with numbers from 1 to 9.

[output] boolean

true if the given grid represents a correct solution to Sudoku, false otherwise.
*/
// using map
func sudoku(grid [][]int) bool {
    for i := 0; i < 9; i = i + 3 {
        m := make(map[int]int, 9)
        for j := 0; j < 9; j = j + 3 {

            val := grid[i][j]
            m[val]++
            if m[val] > 1 { return false }

            val = grid[i+1][j]
            m[val]++
            if m[val] > 1 { return false }

            val = grid[i+2][j]
            m[val]++
            if m[val] > 1 { return false }

            val = grid[i][j+1]
            m[val]++
            if m[val] > 1 { return false }

            val = grid[i+1][j+1]
            m[val]++
            if m[val] > 1 { return false }

            val = grid[i+2][j+1]
            m[val]++
            if m[val] > 1 { return false }

            val = grid[i][j+2]
            m[val]++
            if m[val] > 1 { return false }

            val = grid[i+1][j+2]
            m[val]++
            if m[val] > 1 { return false }

            val = grid[i+2][j+2]
            m[val]++
            if m[val] > 1 { return false }
        }
    }
    return true
}
