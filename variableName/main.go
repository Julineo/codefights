/*
Correct variable names consist only of English letters, digits and underscores and they can't start with a digit.

Check if the given string is a correct variable name.

Example

For name = &quot;var_1__Int&quot;, the output should be
variableName(name) = true;
For name = &quot;qq-q&quot;, the output should be
variableName(name) = false;
For name = &quot;2w2&quot;, the output should be
variableName(name) = false.
*/

import . "unicode"

func variableName(name string) bool {
    runes := []rune(name)
    if IsDigit(runes[0]) {return false}
    for _,v := range runes {
        if !(IsDigit(v) || IsLetter(v) || v == '_') {
            return false
        }
    }
    return true
}
