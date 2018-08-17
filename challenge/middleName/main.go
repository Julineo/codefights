/*
https://app.codesignal.com/challenge/DBQa6cjuqHt3b6qTx/solutions/4vDnQxvZTswXoJPdt

RegExp is intended.

You are looking for your long lost cousin! Unfortunately, you don't remember their middle name. Fortunately, your good friend Travis found a bunch of full name references. Unfortunately, what Travis gives you is a mess; it is not a list of names but a paragraph of garbage. Fortunately, you are a master programmer and know just how to extract from this.

Given the first and last name of the cousin, and a paragraph of text, search for middle names and return an array of all possible middle names so that hopefully, someday you could be reunited with your long lost cousin!

NOTES

CaSe SeNsItIvItY mAtTeRs
Ordering should be as they appear, from left to right (repeats are allowed)
All middle names must begin with a capital letter, followed by lowercase letters
Middle names cannot contain spaces, and can be separated from the first and last names by at most one space
Example

For firstName = "Timmy", lastName = "Turner", and para = "Garbagetext#### @%^##RFEH Timmy Tiberius Turner !@#AS", the output should be middleName(firstName, lastName, para) = ["Tiberius"]

For firstName = "Mike", lastName = "Johnston", and para = "Mike Guy Johnston 203-497-3156mike leroy johnston 203-979-134MikeCodyJohnston203-456-789", the output should be middleName(firstName, lastName, para) = ["Guy", "Cody"]

For firstName = "Tom", lastName = "Tucker", and para = "Tom Theodore Tucker; Toom Tuesday Tucker; Tom Toodles Tucker", the output should be middleName(firstName, lastName, para) = ["Theodore", "Toodles"]

For firstName = "Jim", lastName = "Halpert", and para = "Jim jeff HalpertJimJamesHalpert Jim JOhnson Halpert JimJimJimJimJimJimJimJimHalpert", the output should be middleName(firstName, lastName, para) = ["James", "Jim"]
*/
import . "regexp"
func middleName(firstName, lastName, para string) []string {
    r := []string{}
    for {
        v := MustCompile(firstName + " ?([A-Z][a-z]*) ?" + lastName).FindStringSubmatchIndex(para)
        fmt.Println(v)
        if v == nil {
            break
        }
        r = append(r, para[v[2]:v[3]])
        para = para[v[2]:]
    }
    return r
}
