/*
https://app.codesignal.com/challenge/sEQJmYPDMryEsXKhQ/solutions/fHREAtfz4dafo9dm2
*/
func countNoFillZones(W string) (r int) {
    m := make(map[rune]int)
    for _, v := range "AabDdegOoPpQqR0469" {
        m[v] = 1
    }
    for _, v := range "B8" {
        m[v] = 2
    }

    for _, w := range W {
        r += m[w]
    }
    return
}
