/*
https://codefights.com/arcade/intro/level-8/8N7p3MqzGQg5vFJfZ/description

Given a string, find the number of different characters in it.

Example

For s = "cabca", the output should be
differentSymbolsNaive(s) = 3.

There are 3 different characters a, b and c.
*/

func differentSymbolsNaive(ss string) int {
    m := make(map[rune]int)
    for _, s := range ss {
        m[s]++
    }
    return len(m)
}
