/*
https://app.codesignal.com/interview-practice/task/uX5iLwhc6L5ckSyNC

Note: Write a solution that only iterates over the string once and uses O(1) additional memory, since this is what you would be asked to do during a real interview.

Given a string s, find and return the first instance of a non-repeating character in it. If there is no such character, return '_'.

Example

For s = "abacabad", the output should be
firstNotRepeatingCharacter(s) = 'c'.

There are 2 non-repeating characters in the string: 'c' and 'd'. Return c since it appears in the string first.

For s = "abacabaabacaba", the output should be
firstNotRepeatingCharacter(s) = '_'.

There are no characters in this string that do not repeat.

Input/Output

[execution time limit] 4 seconds (go)

[input] string s

A string that contains only lowercase English letters.

Guaranteed constraints:
1 ≤ s.length ≤ 105.

[output] char

The first non-repeating character in s, or '_' if there are no characters that do not repeat.
*/

// returns first not repeating character with O(1) additional memory
func firstNotRepeatingCharacter(s string) string {
    for _, v := range s {
        if check(v, s) {
            return string(v)
        }
    }
    return "_"
}

// helper functon checks if the character appears only once in the string
func check(c rune, s string) bool {
    n := 0
    for _, s := range s {
        if c == s {
            n++
        }
        if n > 1 {
            return false
        }
    }
    return true
}
