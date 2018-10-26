/*
https://app.codesignal.com/challenge/fBMqpka3pvk9jSHXp

You're spending the night at your aunt's 300-year-old mansion which you're convinced is haunted. While in bed, the sibilant and shrill noises from around the house keep you awake. Are they really ghosts or just the wind?

You recorded the sounds you had heard as a sequence of low-pitched noises (represented by o) and high-pitched ones (represented by u).

The wind blows wildly and randomly, but ghosts' boos and wails follow a predictable pattern:

They begin with a non-zero length of low-pitched os, then a non-zero length of high-pitched us, and finally another length of low-pitched os that are of the same length as the initial sequence of os.

For example:

"ouo" = ghost
"oouuuuuoo" = ghost
"ouuooo" = ghost at first, then wind
"uo" = wind
Given an uninterrupted sequence of sounds, your task is to determine whether it can be divided into non-overlapping, contiguous subsequences that all follow the pattern of ghosts' wails. If so, return true; otherwise, return false.

Example

For sounds = "oouuuoo", the output should be wailingGhosts(sounds) = true.

This is a typical ghost's wail, and takes up the entire input.

For sounds = "uuouoo", the output should be wailingGhosts(sounds) = false.

A wail must start with an o, and there is no o at the beginning of the input.

For sounds = "ouoouo", the output should be wailingGhosts(sounds) = true.

Here we have two ouo ghost wails, one after the other.

For sounds = "ouououo", the output should be wailingGhosts(sounds) = false.

Wails do not overlap, so the middle u cannot be placed in either wail surrounding it.
*/
import "strings"

func wailingGhosts(s string) (r bool) {
    // Two pointers technique
    for len(s) > 0 {
        i := strings.Index(s, "u")
        j := strings.Index(s[i + 1:], "o") + i + 1
        fmt.Println("b")
        k := i + j
        if i < 1 || len(s) < k { return false }
        s = s[k:]
    }
    return !r
}
