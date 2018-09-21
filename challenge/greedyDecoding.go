/*
https://app.codesignal.com/challenge/JhtBKjkeDwaaERovM

You've intercepted a top secret message that's been encoded using a simple cipher. The original message consisted of uppercase English letters, each of which have been converted to a number, using the following mapping:

'A' -> 1
'B' -> 2
...
'Z' -> 26
Since the message now appears as an uninterrupted string of digits, there are many ways it could be interpreted, but we're going to do it in the greediest way possible - any time the next number could be interpreted as a double-digit number, we'll interpret it that way (ignoring the possibility of a single-digit interpretation).

It is guaranteed that message is a valid encoding, and no number should be interpreted as having a leading zero.

Example

For message = "195318520", the output should be greedyDecoding(message) = "SECRET".

example

There are other possible decodings, but this is the greediest one.

For message = "89", the output should be greedyDecoding(message) = "HI".

There's only one way to interpret this one, since there's no 89th letter in the alphabet.
*/

func greedyDecoding(message string) (res string) {
    //Z-26
    i := 0
    a := 0
    for i < len(message) {
        if i + 2 > len(message) { goto short }
        a, _ = strconv.Atoi(message[i:i+2])
        if a < 27 {
            res = res + string(a + 64)
            i++
            i++
            continue
        }
        short:
        b, _ := strconv.Atoi(string(message[i]))
        res = res + string(b + 64)
        i++
    }
    return
}
