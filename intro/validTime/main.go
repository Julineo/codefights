/*
https://app.codesignal.com/arcade/intro/level-12/ywMyCTspqGXPWRZx5

Check if the given string is a correct time representation of the 24-hour clock.

Example

For time = "13:58", the output should be
validTime(time) = true;
For time = "25:51", the output should be
validTime(time) = false;
For time = "02:76", the output should be
validTime(time) = false.
*/
import . "regexp"
func validTime(time string) bool {
    r, _ := MatchString("(^[10][0-9]|^[2][0-3]):[0-5][0-9]", time)
    return r
}
