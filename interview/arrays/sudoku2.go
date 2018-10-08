/*
https://app.codesignal.com/interview-practice/task/SKZ45AF99NpbnvgTn

Sudoku is a number-placement puzzle. The objective is to fill a 9 × 9 grid with numbers in such a way that each column, each row, and each of the nine 3 × 3 sub-grids that compose the grid all contain all of the numbers from 1 to 9 one time.

Implement an algorithm that will check whether the given grid of numbers represents a valid Sudoku puzzle according to the layout rules described above. Note that the puzzle represented by grid does not have to be solvable.

Example

For

grid = [['.', '.', '.', '1', '4', '.', '.', '2', '.'],
        ['.', '.', '6', '.', '.', '.', '.', '.', '.'],
        ['.', '.', '.', '.', '.', '.', '.', '.', '.'],
        ['.', '.', '1', '.', '.', '.', '.', '.', '.'],
        ['.', '6', '7', '.', '.', '.', '.', '.', '9'],
        ['.', '.', '.', '.', '.', '.', '8', '1', '.'],
        ['.', '3', '.', '.', '.', '.', '.', '.', '6'],
        ['.', '.', '.', '.', '.', '7', '.', '.', '.'],
        ['.', '.', '.', '5', '.', '.', '.', '7', '.']]
the output should be
sudoku2(grid) = true;

For

grid = [['.', '.', '.', '.', '2', '.', '.', '9', '.'],
        ['.', '.', '.', '.', '6', '.', '.', '.', '.'],
        ['7', '1', '.', '.', '7', '5', '.', '.', '.'],
        ['.', '7', '.', '.', '.', '.', '.', '.', '.'],
        ['.', '.', '.', '.', '8', '3', '.', '.', '.'],
        ['.', '.', '8', '.', '.', '7', '.', '6', '.'],
        ['.', '.', '.', '.', '.', '2', '.', '.', '.'],
        ['.', '1', '.', '2', '.', '.', '.', '.', '.'],
        ['.', '2', '.', '.', '3', '.', '.', '.', '.']]
the output should be
sudoku2(grid) = false.

The given grid is not correct because there are two 1s in the second column. Each column, each row, and each 3 × 3 subgrid can only contain the numbers 1 through 9 one time.

Input/Output

[execution time limit] 4 seconds (go)

[input] array.array.char grid

A 9 × 9 array of characters, in which each character is either a digit from '1' to '9' or a period '.'.

[output] boolean

Return true if grid represents a valid Sudoku puzzle, otherwise return false.
*/

func sudoku2(grid [][]string) bool {
    //rows
    for i := 0; i < 9; i++ {
        m := make(map[string]int)
        for j := 0; j < 9; j++ {
            if grid[i][j] != "." {
                m[grid[i][j]]++
                if m[grid[i][j]] > 1 {
                    return false }
            }
        }
    }
    //columns
    for j := 0; j < 9; j++ {
        m := make(map[string]int)
        for i := 0; i < 9; i++ {
            if grid[i][j] != "." {
                m[grid[i][j]]++
                if m[grid[i][j]] > 1 { return false }
            }
        }
    }

    //groups
    for i := 0; i < 9; i = i + 3 {
        for j := 0; j < 9; j = j + 3 {
            m := make(map[string]int)

            if grid[i][j] != "." {
                m[grid[i][j]]++
            }
            if m[grid[i][j]] > 1 { return false }

            if grid[i+1][j] != "." {
                m[grid[i+1][j]]++
            }
            if m[grid[i+1][j]] > 1 { return false }

            if grid[i+2][j] != "." {
                m[grid[i+2][j]]++
            }
            if m[grid[i+2][j]] > 1 { return false }


            if grid[i][j+1] != "." {
                m[grid[i][j+1]]++
            }
            if m[grid[i][j+1]] > 1 { return false }

            if grid[i+1][j+1] != "." {
                m[grid[i+1][j+1]]++
            }
            if m[grid[i+1][j+1]] > 1 { return false }

            if grid[i+2][j+1] != "." {
                m[grid[i+2][j+1]]++
            }
            if m[grid[i+2][j+1]] > 1 { return false }


            if grid[i][j+2] != "." {
                m[grid[i][j+2]]++
            }
            if m[grid[i][j+2]] > 1 { return false }

            if grid[i+1][j+2] != "." {
                m[grid[i+1][j+2]]++
            }
            if m[grid[i+1][j+2]] > 1 { return false }

            if grid[i+2][j+2] != "." {
                m[grid[i+2][j+2]]++
            }
            if m[grid[i+2][j+2]] > 1 { return false }

        }
    }
    return true
}
