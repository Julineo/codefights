/*
https://app.codesignal.com/challenge/ZnZu5e2uRSAHiT2q6

Summer road trip time! You and your friends are driving all the way to Banff National Park for a camping trip in the woods.

While staring out of the car window, your friend started reading some road signs in a strange way to break the silence. The messages were all weird, they were nothing like what the road sign wrote when he read it out loud!

Given a sign in its more readable form with sign[i] being one section of the sign, guess what your friend will read out.

Yes, this is a reverse challenge, enjoy!

Example

For sign = ["You matter.", "Don't give up."], the output should be roadSigns(sign) = "You don't matter. Give up.". What a nice message.

https://imgur.com/a/VQC15kH

Input / Output

[execution time limit] 4 seconds (go)

[input] array.string sign

The road sign itself. It is given in its proper form, and is up to you to scramble the words into its garbled message.

Guaranteed constraints:
0 < sign.length < 50
0 < sign[i].length < 237
*/
import . "strings"

func roadSigns(sign []string) (res string) {
    max := 0
    for _, v := range sign {
        if len(Fields(v)) > max {
            max = len(Fields(v))
        }
    }
    for i := 0; i < max; i++ {
        for j := 0; j < len(sign); j++ {
            if i >= len(Fields(sign[j])) {
                continue
            }
            if len(res) == 0 || f(res) == 46 || f(res) == 33 || f(res) == 63 {
                res = res + " " + cap(ToLower(Fields(sign[j])[i]))
                continue
            }
            if Fields(sign[j])[i] == "!" {
                res = res + ToLower(Fields(sign[j])[i])
                continue
            }
            res = res + " " + ToLower(Fields(sign[j])[i])
        }
    }
    res = Trim(res, " ")
    return
}

func cap(s string) string {
    return ToUpper(s[:1]) + ToLower(s[1:])
}

func f(res string) {
    return res[len(res) - 1]
}
