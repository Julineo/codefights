/*
https://codefights.com/challenge/4PLW8uTndBKB7FghP/solutions/hkX3RufyLP3uydThh

World Cup is here! For most people it is one of the greatest events in their lives, but for commentators it is a huge trouble! It is a real challenge for them to pronounce such complicated names that they are dealing with.

Given the surname of a player, determine how hard it is to pronounce. We assume that the difficulty of the surname is the maximum number of consecutive consonants in it.

Example

For surname = "Blaszczykowski", the output should be
hardSurname(surname) = 6;
For surname = "Papastathopoulos", the output should be
hardSurname(surname) = 2.
*/
func hardSurname(surname string) int {
    tmp, max := 0, 0
    for _, v := range surname {
        if isConsonat(v) {
            tmp += 1
        } else {
            tmp = 0
        }
        if tmp > max {
                max = tmp
        }
    }
    return max
}

func isConsonat(l rune) bool {
    if l == 'a' || l == 'e' || l == 'i' || l == 'o' || l == 'u' || l == 'A' || l == 'E' || l == 'I' || l == 'O' || l == 'U' {
        return false
    }
    return true
}
