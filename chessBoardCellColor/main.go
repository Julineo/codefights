/*
Given two cells on the standard chess board, determine whether they have the same color or not.

Example

For cell1 = "A1" and cell2 = "C3", the output should be
chessBoardCellColor(cell1, cell2) = true.

For cell1 = "A1" and cell2 = "H3", the output should be
chessBoardCellColor(cell1, cell2) = false.
*/
func chessBoardCellColor(cell1 string, cell2 string) bool {
    return 0 == (cell1[0]-64 + cell1[1]-48 + cell2[0]-64 + cell2[1]-48)%2
}
