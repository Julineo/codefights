/*
https://app.codesignal.com/challenge/DfQ3jYW9sKJj6r4Pv

You're practicing writing numbers in word form, according to the following rules:

The first letter is capitalized
Hyphenated words are used for numbers under 100 (eg: "Thirty-two")
Given an integer number, return a string representing the number in word form.

Examples

For number = 1, the output should be numberInWords(number) = "One".

For number = 12, the output should be numberInWords(number) = "Twelve".

For number = 14, the output should be numberInWords(number) = "Fourteen".

For number = 21, the output should be numberInWords(number) = "Twenty-one".

For number = 299, the output should be numberInWords(number) = "Two hundred ninety-nine".

For number = 123456, the output should be numberInWords(number) = "One hundred twenty-three thousand four hundred fifty-six".

Input / Output

[execution time limit] 4 seconds (go)

[input] integer number

An integer representing the number you need to write out.

Guaranteed constraints:
1 ≤ number ≤ 106
*/
import (
    . "strconv"
    . "strings"
)

func numberInWords(num int) string {
    if num == 0 {
        return "Zero"
    }
    if num == 1000000 {
        return "One million"
    }

    sl := []string{"","one","two","three","four","five","six","seven","eight","nine",
"ten","eleven","twelve","thirteen","fourteen","fifteen","sixteen","seventeen","eighteen","nineteen",
"twenty","twenty-one","twenty-two","twenty-three","twenty-four","twenty-five","twenty-six","twenty-seven","twenty-eight","twenty-nine",
"thirty","thirty-one","thirty-two","thirty-three","thirty-four","thirty-five","thirty-six","thirty-seven","thirty-eight","thirty-nine",
"forty","forty-one","forty-two","forty-three","forty-four","forty-five","forty-six","forty-seven","forty-eight","forty-nine",
"fifty","fifty-one","fifty-two","fifty-three","fifty-four","fifty-five","fifty-six","fifty-seven","fifty-eight","fifty-nine",
"sixty","sixty-one","sixty-two","sixty-three","sixty-four","sixty-five","sixty-six","sixty-seven","sixty-eight","sixty-nine",
"seventy","seventy-one","seventy-two","seventy-three","seventy-four","seventy-five","seventy-six","seventy-seven","seventy-eight","seventy-nine",
"eighty","eighty-one","eighty-two","eighty-three","eighty-four","eighty-five","eighty-six","eighty-seven","eighty-eight","eighty-nine",
"ninety","ninety-one","ninety-two","ninety-three","ninety-four","ninety-five","ninety-six","ninety-seven","ninety-eight","ninety-nine"}
    snum := Itoa(num)
    if len(snum) < 3 {
        return fl(sl[num])
    }
    if len(snum) == 3 {
        t := Ai(snum[0:1])
        o := Ai(snum[1:])
        return fl(sl[t] + " hundred " + sl[o])
    }
    if len(snum) == 4 {
        f := Ai(snum[0:1])
        t := Ai(snum[1:2])
        o := Ai(snum[2:])
        if t == 0 {
            return fl(sl[f] + " thousand " + sl[o])
        }
        return fl(sl[f] + " thousand " + sl[t] + " hundred " + sl[o])
    }
    if len(snum) == 5 {
        f := Ai(snum[0:2])
        t := Ai(snum[2:3])
        o := Ai(snum[3:])

        if o == 0 && t == 0 {
            return fl(sl[f] + " thousand ")
        }
        if t == 0 {
            return fl(sl[f] + " thousand " + sl[o])
        }
        return fl(sl[f] + " thousand " + sl[t] + " hundred " + sl[o])
    }
    if len(snum) == 6 {
        v := Ai(snum[0:1])
        f := Ai(snum[1:3])
        t := Ai(snum[3:4])
        o := Ai(snum[4:])

        if o == 0 && t == 0 {
            return fl(sl[v] + " hundred " + sl[f] + " thousand ")
        }
        if t == 0 {
            return fl(sl[v] + " hundred " + sl[f] + " thousand " + sl[o])
        }
        return fl(sl[v] + " hundred " + sl[f] + " thousand " + sl[t] + " hundred " + sl[o])
    }
    return ""
}

func fl(s string) string {
    return Replace(Trim(Title(s[0:1]) + s[1:], " "), "  ", " ", -1)
}

func Ai(s string) int {
    t, _ := Atoi(s)
    return t
}
