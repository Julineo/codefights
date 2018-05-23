/*
Two strings are called anagrams, if they contain the same characters, but the order of the characters may be different.

Given a string consisting of lowercase letters and question marks, s1, and another string consisting of lowercase letters, s2, determine whether these two strings could become anagrams by replacing each ? character in s1 with a letter.

Examples

For s1 = listen and s2 = silent, the output should be couldBeAnagram(s1, s2) = true. The letters of s1 could be rearranged to form s2.

For s1 = cat and s2 = dog, the output should be couldBeAnagram(s1, s2) = false. There's no way s2 could be formed using the letters of s1.

For s1 = n?ce and s2 = nice, the output should be couldBeAnagram(s1, s2) = true. By replacing the ? with i in s1, the two strings will have the same characters.
*/
func couldBeAnagram(s1 string, s2 string) bool {
    m := make(map[rune]int)
    //add line 2 to map
    for _, r := range s2 {
        m[r]++
    }
    //substract line 2 from map and chek if something there which not suppose to be
    for _, r := range s1 {
        m[r]--
        if r != '?' && m[r] < 0 {
            return false
        }
    }
    return len(s1) == len(s2)
}
