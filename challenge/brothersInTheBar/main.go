/*
https://app.codesignal.com/challenge/dHjGPiEYFFcrwKMip

Three brothers walk into a bar. All the beverages are placed in one line at the long bar table. The size of each glass is represented in an array of integers, glasses.

The brothers will drink a round if they can find 3 consecutive glasses of the same size. The barman removes the empty glasses from the table immediately after each round.

Find the maximum number of rounds the three brothers can drink.

Example

For glasses = [1, 1, 2, 3, 3, 3, 2, 2, 1, 1], the output should be brothersInTheBar(glasses) = 3.

The brothers can start with a round of size 3, then after the glasses are cleared, a round of size 2 can be formed, followed by a round of size 1. One glass will be left at the table.



For glasses = [1, 1, 2, 1, 2, 2, 1, 1], the output should be brothersInTheBar(glasses) = 0.

There are no 3 consecutive glasses of the same size.

Input/Output

[execution time limit] 4 seconds (go)

[input] array.integer glasses

The sizes of glasses in the row.

Guaranteed constraints:
1 ≤ glasses.length ≤ 105,
1 ≤ glasses[i] ≤ 106.

[output] integer

The maximum number of rounds the brothers can drink.
*/
func brothersInTheBar(glasses []int) (rounds int) {
    i := 0
    for {
        if i < 0 {
            i = 0
        }
        if i > len(glasses) - 3 {
            break
        }
        if glasses[i] == glasses[i+1] && glasses[i+1] == glasses[i+2] {
            rounds++
            glasses = append(glasses[:i], glasses[i+3:]...)
            i = i - 2
            continue
        }
        i++
    }
    return
}

// Use the original slice as a stack, without using append.
/*
func brothersInTheBar(g []int) (c int) {
    i := 0
    for _, v := range g {
        g[i]=v
        if i > 1 && v == g[i-1] && v == g[i-2] {
            i-=3
            c++
        }
        i++

    }
    return
}
*/
