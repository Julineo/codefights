/*
https://app.codesignal.com/challenge/NZLcmPoRjWAMfCbx2

Reverse the order of all characters in a string that occur an odd amount of times (spaces included). All other characters should remain in the same position; only odd-frequency characters are eligible to swap with each other.

Case-sensitivity is important, so for example "a" is considered different than "A" when counting character frequencies.

Example

For str = "hello world", the output should be reverseOddCount(str) = "dlrwo loleh".

example

Occurrences:
h: 1 (reverse because it occurs an odd # of times)
e: 1 (reverse because it occurs an odd # of times)
l: 3 (reverse because it occurs an odd # of times)
o: 2 (leave it in place because it occurs an even # of times)
w: 1 (reverse because it occurs an odd # of times)
r: 1 (reverse because it occurs an odd # of times)
d: 1 (reverse because it occurs an odd # of times)
*/
func reverseOddCount(str string) string {
    b := []byte(str)
    m := make(map[byte]int)

    for _, v := range b {
        m[v]++
    }

    i, j := 0, len(str)-1

    //two pointers technique
    for i < j {
        if m[str[i]] % 2 != 0 && m[str[j]] % 2 != 0 {
            b[i], b[j] = b[j], b[i]
            j--
            i++
            continue
        }
        if m[str[j]] % 2 == 0 {
            j--
        }
        if m[str[i]] % 2 == 0 {
            i++
        }
    }

    return string(b)
}
