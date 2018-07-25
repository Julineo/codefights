/*
https://app.codesignal.com/challenge/QMBXSpgJvmaAuqPbn

Your garage is completely full of old junk you never use any more, so to help get it cleaned up, you've decided to have a garage sale. You've picked out all the items you plan on selling, and you have a bunch of number stickers for marking the prices.

You really need to get rid of this stuff and you're hoping to sell it all at once, so your plan is to price the items in such a way that the total cost of all the items is as low as possible. You're also trying to get rid of all your stickers, so make sure you use them all!

Given an array of integers stickers, representing the digits on each of the price stickers, and items, an integer representing the total number of items for sale, find the minimum total sum that the items can be sold for (using all the stickers).

Example

For stickers = [9, 4, 7, 1, 6, 3, 8, 4, 5, 1] and items = 3, the output should be garageSaleMystery(stickers, items) = 1974

One way to get this sum would be to price the items at 359, 468, and 1147.

Example

Input / Output

[execution time limit] 4 seconds (go)

[input] array.integer stickers

An array of integers representing the digit values of each of the price stickers.

Guaranteed constraints:
1 ≤ stickers.length ≤ 1000
0 ≤ stickers[i] ≤ 9

[input] integer items

An integer representing the total number of items at the garage sale.

Guaranteed constraints:
1 ≤ items ≤ stickers.length

[output] integer

The minimum total sum of prices that you can get by using all the stickers. It's guaranteed that the output is less than 109.
*/

import . "sort"

func garageSaleMystery(stickers []int, items int) (sum int) {
    Ints(stickers)
    a := make([][]int, items)
    for i, v := range stickers {
        a[i%items] = append(a[i%items], v)
    }
    for i := 0; i < items; i++ {
        tmpSum := a[i][0]
        for j := 1; j < len(a[i]); j++ {
            tmpSum = tmpSum * 10 + a[i][j]
        }
        sum += tmpSum
    }
    return sum
}
