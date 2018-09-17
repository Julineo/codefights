/*
https://app.codesignal.com/interview-practice/task/mj4qdJqZknbkHNzhK

Implement the missing code, denoted by ellipses. You may not modify the pre-existing code.
A sentence is considered correct if:

it starts with a capital letter;
it ends with a full stop, question mark or exclamation point ('.', '?' or '!');
full stops, question marks and exclamation points don't appear anywhere else in the sentence.
Given a sentence, return true if it is correct and false otherwise.

Example

For sentence = "This is an example of *correct* sentence.",
the output should be
isSentenceCorrect(sentence) = true;

For sentence = "!this sentence is TOTALLY incorrecT",
the output should be
isSentenceCorrect(sentence) = false.
*/

import "regexp"

func isSentenceCorrect(sentence string) bool {
  re := regexp.MustCompile("^[A-Z][^?.!]*[?.!]$")
  return re.MatchString(sentence)
}
