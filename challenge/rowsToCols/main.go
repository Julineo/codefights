/*
https://app.codesignal.com/challenge/LfWCaKHdmeCQk6jZP

Input:
rows: [13, 1, 1, 10]
Output:
Run the code to see output
Expected Output:
[9, 8, 1, 14]

nput:
rows: [3, 190, 112, 133, 91, 156, 123, 38]
Output:
Run the code to see output
Expected Output:
[84, 42, 99, 110, 78, 85, 203, 154]
*/
import . "strconv"

func rowsToCols(rows []int) []int {

    // making transposed matrix
    trans := make([][]rune, len(rows))
    for i := 0; i < len(rows); i++ {
        trans[i] = make([]rune, len(rows))
    }

    // filling transposed matrix
    for j, v := range rows {
        bits := fmt.Sprintf("%0" + Itoa(len(rows)) + "b", v)
        for i, bit := range bits {
            trans[i][j] = bit
        }
    }

    //generating int from bit strings
    res := []int{}
    for i := 0; i < len(rows); i++ {
        r, _ := ParseInt(string(trans[i]), 2, 0)
        res = append(res, int(r))
    }
    return res
}

/*
func rowsToCols(r []int) (c []int) {
    l := len(r)
    for l > 0 {
        v := 0
        l--
        for _, b := range r {
            v += b >> uint(l) & 1 + v
        }
        c = append(c, v)
    }
    return 
}
*/
