/*
https://app.codesignal.com/arcade/intro/level-11/o2uq6eTuvk7atGadR

Given a string, return its encoding defined as follows:

First, the string is divided into the least possible number of disjoint substrings consisting of identical characters
for example, "aabbbc" is divided into ["aa", "bbb", "c"]
Next, each substring with length greater than one is replaced with a concatenation of its length and the repeating character
for example, substring "bbb" is replaced by "3b"
Finally, all the new strings are concatenated together in the same order and a new string is returned.
Example

For s = "aabbbc", the output should be
lineEncoding(s) = "2a3bc".
*/

func lineEncoding(s string) (res string) {
    count := 1
    s += "#"
    for i := 1; i < len(s); i++ {
        if s[i-1] == s[i] {
            count++
        } else {
            if count > 1 {
                res = res + strconv.Itoa(count)
            }
                res = res + string(s[i-1])
                count = 1
            }
    }
    return
}
