/*
https://app.codesignal.com/challenge/6nrk2rLGZRkH9gDZ5

Jumping Jimmy is trying to reach the top of a tall tower!

jumping jimmy

Given an integer jumpHeight representing the maximum vertical distance he can ascend in one jump, and an array tower representing the vertical distances between each floor of the tower, your task is to find the total vertical distance Jimmy is able to travel (until he reaches the top of the tower or isn't able to complete the next jump).

Example

For tower = [3, 1, 2] and jumpHeight = 3, the output should be jumpingJimmy(tower, jumpHeight) = 6.

example 1

After reaching the top of the tower, Jimmy has ascended a total distance of 3 + 1 + 2 = 6.

For tower = [1, 2, 3, 4] and jumpHeight = 2, the output should be jumpingJimmy(tower, jumpHeight) = 3.

example 2

Jimmy is able to complete the first two jumps, but he can't clear a gap of size 3, so the total vertical distance is 1 + 2 = 3.

For tower = [5, 1, 8, 2, 4, 3, 1, 9, 8, 5, 1] and jumpHeight = 1, the output should be jumpingJimmy(tower, jumpHeight) = 0.

example 3

Jimmy can't even clear the first gap, so he can't ascend the tower at all 
*/
func jumpingJimmy(t []int, j int) (r int) {
    for _, v := range t {
        if j < v { break }
        r += v
    }
    return
}
