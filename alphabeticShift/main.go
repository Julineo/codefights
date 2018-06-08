/*
Given a string, replace each its character by the next one in the English alphabet (z would be replaced by a).

Example

For inputString = "crazy", the output should be
alphabeticShift(inputString) = "dsbaz"
*/
func alphabeticShift(input string) (output string) {
    for _, v := range input {
        if v == 'z' {
            output += "a"
            continue
        }
        output += string(v+1)
    }
    return
}
