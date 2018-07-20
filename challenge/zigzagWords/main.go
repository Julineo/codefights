/*
https://app.codesignal.com/challenge/tu9KcxX573X9qFxfN

Here is a little fun in a blank page and a colorful mind. Your 6 year old daughter is having fun writing words she sees around her. Sometimes from the newspaper, sometimes from TV, sometimes from books lying around the house.

The problem is, she is not really great with vertical alignment. Or maybe it was on purpose, who knows.

She wrote all the words in a zigzag pattern (in l lines, starting at the bottom line), disregarding all whitespace characters.

For example, the word HOSPITAL in 3 lines is now:

    S       A 
  O   P   T   L
H       I  
You decide it is a good idea to teach her how people would try to read it if she does not write in a single line.

Since this would normally be read from left to right, one line at a time, it would look like "SA", "OPTL", and "HI", so you show her the output "SAOPTLHI".

Examples

For t = "HOSPITAL" and l = 3, the output should be zigzagWords(t, l) = "SAOPTLHI"

For t = "my sweet daughter" and l = 3, the output should be zigzagWords(t, l) = "stgryweduhemeat"

Writing in a zigzag pattern and ignoring whitespaces, the text would look like this:

    s       t       g       r
  y   w   e   d   u   h   e
m       e       a       t
Reading from left to right, the lines concatenate to form the string "stgryweduhemeat".

Input / Output

[execution time limit] 4 seconds (go)

[input] string t

A string of text containing English letters, numbers, punctuation, and whitespace.

Guaranteed constraints:
0 ≤ t.length ≤ 2000
t[i] ∈ {"A - Z", "a - z", "0 - 9", "()[]:;.,-=#`'"", space, tab}

[input] integer l

A number representing the number of lines your daughter has used.

Guaranteed constraints:
0 ≤ l ≤ t.length

[output] string

A string representing the concatenation of all l lines in the zigzag version of the text, disregarding whitespaces.
*/
import . "strings"

func zigzagWords(t string, l int) (res string) {

    //removing spaces
    t = Replace(t, string(32), "", -1)
    //t = strings.Replace(t, string(9), "", -1)
    //t = strings.Replace(t, string(10), "", -1)
    //t = strings.Replace(t, string(11), "", -1)
    //t = strings.Replace(t, string(12), "", -1)
    //t = strings.Replace(t, string(13), "", -1)

    // if straight, no need to calculate
    if l == 1 {
        return t
    }

    // making "waves"
    m := make([]string, l)
    o := 0
    s := -1
    for i := 0; i < len(t); i++ {
        if o == l-1 || o == 0 {
            s = s * -1
        }
        m[o] = m[o] + string(t[i])
        o = o + s
    }

    // making output
    for i := len(m) - 1; i >= 0; i-- {
        res += m[i]
    }
    return res
}

