/*Given a rectangular matrix of characters, add a border of asterisks(*) to it.

Example

For

picture = ["abc",
           "ded"]
the output should be

addBorder(picture) = ["*****",
                      "*abc*",
                      "*ded*",
                      "*****"]*/

import "strings"

func addBorder(picture []string) []string {
    //left and right borders
    for i := 0; i < len(picture); i++ {
        picture[i] = "*" + picture[i] + "*"
    }

    //lower border and upper border
    s := strings.Repeat("*", len(picture[0]))
    picture = append(picture, s)
    picture = append([]string{s}, picture...)

    return picture
}
