/*
https://app.codesignal.com/challenge/fRE4enG7w5fgZAcEy

Ms. Jackson is a talented and inspiring math teacher. She always finds a way to keep even the most disinterested students engaged in her lessons. Today's lesson was on fractions, which tends to be an exhausting subject for most students, so to maintain the class's interest and motivation, she decided to offer extra credit to anyone brave enough to solve her extra-challenging problem on repeating decimal numbers.

You've learned that some fractions produce repeating decimals (for example 2 / 3 = 0.666...). Ms. Jackson has also taught you that each repeating decimal number can be expressed as a fraction. So with these facts in mind, you'll be given a repeating decimal number, as well as the denominator of the fraction that produced it, and you'll have to find the numerator!

Your friend Bob tends to struggle with math, so he could really use the extra credit, but he's not really sure how to solve this type of problem. You figure Bob might get it after seeing several different examples, so you'd like to write an algorithm for solving this type of problem efficiently, so that Bob can check his answers against your algorithm's results.

Notes:

Parentheses will be used to represent the repeating part of the decimal (eg: 3.6(81) means 3.6818181...)
If the decimal has a repeating part, it'll always begin after the decimal point.
It's guaranteed that for each decimal and denominator, there is a valid integer value for the numerator.
Example

For decimal = "0.(3)" and denominator = 3, the output should be 1.

example 1

For decimal = "0.0(25)" and denominator = 594, the output should be 15.

example 2

For decimal = "4.6" and denominator = 5, the output should be 23.

example 3

For decimal = "3.6(81)" and denominator = 264, the output should be 972.

example 4

Input / Output

[execution time limit] 4 seconds (go)

[input] string decimal

A string representing the (potentially repeating) decimal number. If there's a repeating part, it'll be enclosed in parentheses.

Guaranteed constraints:
1 <= decimal.length <= 104
decimal[i] in {"0" - "9", ".", "(", ")"}
*/

// solution 3
import . "strings"

func extraCreditAssignment(d string, n int) int {
    s := Replace(d, "(", "", 1)
    f := .0
    fmt.Sscan(s, &f)
    return int(f * float64(n) + 0.99)
}



// solution 2
/*
func extraCreditAssignment(decimal string, denominator int) int {
    t := ""
    for _, v := range decimal {
        if v > 45 {
            t += string(v)
        }
    }
    fmt.Println(t)
    f, _ := strconv.ParseFloat(t, 64)
    return int(f * float64(denominator) + 0.99)
}
*/

// solution 1
/*
import (
    . "strings"
)

func extraCreditAssignment(decimal string, denominator int) int {
    r := NewReplacer("(", "", ")", "")
    v := r.Replace(decimal)
    f, _ := strconv.ParseFloat(v, 64)
    return int(f * float64(denominator) + 0.99)
}
*/

// solution 0
/*
import (
    . "strings"
        "math"
        "regexp")

func extraCreditAssignment(decimal string, denominator int) int {
    var rgx = regexp.MustCompile(`\((.*?)\)`)
	rs := rgx.FindStringSubmatch(decimal)
    add := ""
    if len(rs) > 1 {
        add = rs[1]
    } 
    decimal = Replace(decimal, "(", "", -1)
    decimal = Replace(decimal, ")", "", -1)
    decimal = decimal + add
    f, _ := strconv.ParseFloat(decimal, 64)
    return int(math.Round(f * float64(denominator)))
}
*/
