/*
https://app.codesignal.com/arcade/code-arcade/at-the-crossroads/7jaup9HprdJno2diw/

In tennis, a set is finished when one of the players wins 6 games and the other one wins less than 5, or, if both players win at least 5 games, until one of the players wins 7 games.

Determine if it is possible for a tennis set to be finished with the score score1 : score2.

Example

For score1 = 3 and score2 = 6, the output should be
tennisSet(score1, score2) = true;

For score1 = 8 and score2 = 5, the output should be
tennisSet(score1, score2) = false.

Since both players won at least 5 games, the set would've ended once one of them won the 7th one.

For score1 = 6 and score2 = 5, the output should be
tennisSet(score1, score2) = false.
*/

func tennisSet(score1 int, score2 int) bool {
    if score1 < score2 {
        score1, score2 = score2, score1
    }

    if score1 == 6 && score2 < 5 { return true }
    if score1 == 7 && (score2 == 5 || score2 == 6) { return true }
    return false
}
