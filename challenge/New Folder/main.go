/*
https://app.codesignal.com/challenge/fHMuKwFnuqMZip7Ff

Before the World Cup starts, all teams are split into 8 groups, each containing 4 teams. After all the teams within each group have played each other, the two highest-performing teams move on to the next round - playoff.



You're given a list of group results in the following format:

[[A1, A2, A3, A4],
 [B1, B2, B3, B4],
 [C1, C2, C3, C4],
 ...
 [H1, H2, H3, H4]]
Here each row contains the results of one group (first element of the row is the command got 1st place, second element is for command on the 2nd place, etc.)

So for example, if the first row looks like ["Uruguay", "Russia", "Egypt", "Saudi Arabia"], it means that within group A, Uruguay is first, Russia is second, Egypt placed third, and Saudi Arabia placed fourth. Therefore, only Russia and Uruguay will move on to playoffs.

The matchups for the next round are the following: (A1, B2), (B1, A2), (C1, D2), (D1, C2), (E1, F2), (F1, E2), (G1, H2), (H1, G2) (the first place team from group A will face the second place team from group B, the first place team from group B will face the second place team from group A, the first place team from group C will face the second place team from group D, etc).

Given a matrix of results, your task is to return a table of matchups for the next round. The columns should be vertically aligned, with the team names aligned to the left.

Example

For

results = [
    ["Uruguay", "Russia", "Egypt", "Saudi Arabia"],
    ["Spain", "Portugal", "Iran", "Morocco"],
    ["France", "Denmark", "Peru", "Australia"],
    ["Croatia", "Argentina", "Nigeria", "Iceland"],
    ["Brazil", "Switzerland", "Serbia", "Costa Rica"],
    ["Sweden", "Mexico", "South Korea", "Germany"],
    ["Belgium", "England", "Tunisia", "Panama"],
    ["Colombia", "Japan", "Senegal", "Poland"]
]
the output should be

worldCupStages(results) = ["|Uruguay |Spain |France   |Croatia|Brazil|Sweden     |Belgium|Colombia|",
                           "|Portugal|Russia|Argentina|Denmark|Mexico|Switzerland|Japan  |England |"]
*/
func worldCupStages(res [][]string) (s []string) {
    var a1, a2 []string
    for i := 1; i < len(res); i += 2 {
        a1 = append(a1, res[i-1][0], res[i][0])
        a2 = append(a2, res[i][1], res[i-1][1])
    }
    var s1, s2 string
    for i := 0 ; i < len(a1); i++ {
        len := max(len(a1[i]), len(a2[i]))
        s1 = s1 + "|" + pad(a1[i], len)
        s2 = s2 + "|" + pad(a2[i], len)
    }
    return []string{s1 + "|", s2 + "|"}
}

//Pads right with spaces
func pad(str string, lenght int) string {
	for {
		str += " "
		if len(str) > lenght {
			return str[0:lenght]
		}
	}
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
