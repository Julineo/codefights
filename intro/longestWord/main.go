/*
https://app.codesignal.com/arcade/intro/level-12/s9qvXv4yTaWg8g4ma

Define a word as a sequence of consecutive English letters. Find the longest word from the given string.

Example

For text = "Ready, steady, go!", the output should be
longestWord(text) = "steady".
*/
import . "regexp"

func longestWord(text string) string {
    re := MustCompile("[a-zA-Z]+")
    words := re.FindAllString(text, -1)
    max, maxIdx := 0, 0
    for i, word := range words {
        if len(word) > max {
            maxIdx = i
            max = len(word)
        }
    }
    return words[maxIdx]
}
