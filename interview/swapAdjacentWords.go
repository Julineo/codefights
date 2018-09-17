/*
https://app.codesignal.com/interview-practice/task/AMx2WqCoq2PEcJTEp

Implement the missing code, denoted by ellipses. You may not modify the pre-existing code.
You are given a string consisting of words separated by whitespace characters, where the words consist only of English letters. Your task is to swap pairs of consecutive words and return the result.

Example

For s = "CodeFight On", the output should be
swapAdjacentWords(s) = "On CodeFight";
For s = "How are you today guys", the output should be
swapAdjacentWords(s) = "are How today you guys".
*/

import "regexp"

func swapAdjacentWords(s string) string {
  return regexp.MustCompile("([a-zA-Z]+) ([a-zA-Z]+)").ReplaceAllString(s, "$2 $1")
}
