/*
https://app.codesignal.com/arcade/intro/level-11/pwRLrkrNpnsbgMndb

Given a position of a knight on the standard chessboard, find the number of different moves the knight can perform.

The knight can move to a square that is two squares horizontally and one square vertically, or two squares vertically and one square horizontally away from it. The complete move therefore looks like the letter L. Check out the image below to see all valid moves for a knight piece that is placed on one of the central squares.


Example

For cell = "a1", the output should be
chessKnight(cell) = 2.

For cell = "c2", the output should be
chessKnight(cell) = 6.
*/
func chessKnight(cell string) (r int) {
    x := int(cell[0] - 96)
    y := int(cell[1] - 48)
    sl := [][]int{{1, 2},{2, 1},{2, -1},{1, -2},{-1, -2},{-2, -1},{-2, 1},{-1, 2},}
    for i := 0; i < 8; i++ {
        if x + sl[i][0] > 0 &&
            y + sl[i][1] > 0 &&
            x + sl[i][0] < 9 &&
            y + sl[i][1] < 9 { r++ }
    }
    return
}

/*
func chessKnight(cell string) (r int) {
    x := cell[0] - 96
    y := cell[1] - 48
        if x + 1 > 0 && y + 2 > 0 && x + 1 < 9 && y + 2 < 9 {r++}
        if x + 2 > 0 && y + 1 > 0 && x + 2 < 9 && y + 1 < 9 {r++}
        if x + 2 > 0 && y - 1 > 0 && x + 2 < 9 && y - 1 < 9 {r++}
        if x + 1 > 0 && y - 2 > 0 && x + 1 < 9 && y - 2 < 9 {r++}
        if x - 1 > 0 && y - 2 > 0 && x - 1 < 9 && y - 2 < 9 {r++}
        if x - 2 > 0 && y - 1 > 0 && x - 2 < 9 && y - 1 < 9 {r++}
        if x - 2 > 0 && y + 1 > 0 && x - 2 < 9 && y + 1 < 9 {r++}
        if x - 1 > 0 && y + 2 > 0 && x - 1 < 9 && y + 2 < 9 {r++}
    return
}
*/
