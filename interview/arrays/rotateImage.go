/*
Note: Try to solve this task in-place (with O(1) additional memory), since this is what you'll be asked to do during an interview.

You are given an n x n 2D matrix that represents an image. Rotate the image by 90 degrees (clockwise).

Example

For

a = [[1, 2, 3],
     [4, 5, 6],
     [7, 8, 9]]
the output should be

rotateImage(a) =
    [[7, 4, 1],
     [8, 5, 2],
     [9, 6, 3]]
*/

func rotateImage(a [][]int) [][]int {
    // transpose upper triangle
    for l := 1; l < len(a); l++ {
        for i, j := 0, l; i < j; i, j = i + 1, j - 1 {
            a[i][j], a[j][i] = a[j][i], a[i][j]
        }
    }

    // transpose lower triangle
    for l := 1; l < len(a) - 1; l++ {
        for i, j := l, len(a) - 1; i < j; i, j = i + 1, j - 1 {
            fmt.Println(a[i][j], a[j][i])
            a[i][j], a[j][i] = a[j][i], a[i][j]
        }
    }

    // reverse horizontally
    for i := 0; i < len(a); i++ {
        for p, q := 0, len(a) - 1; p < q; p, q = p + 1, q - 1 {
            a[i][p], a[i][q] = a[i][q], a[i][p]
        }
    }

    return a
}
