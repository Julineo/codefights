/*
https://app.codesignal.com/challenge/W3iEJPAy39CsM2B6g

You've been chosen as a contestant on a popular TV game show celebrating the excesses of consumerism!

You and several other contestants will be taking turns estimating the price of a fancy item. The bidding is done according to the following rules:

Bids are rounded to the nearest dollar amount
The lowest acceptable bid is $1
All bids must be unique (no repeats)
The contestant whose estimate was closest to the actual price without going over wins
You're the last one to bid, and you'd like to strategize - your plan is to choose the number that would give the broadest range of possible winning values. You have a pretty good idea of what the maximum possible price could be, so you can use that in calculating the best option.

Given an integer maxPrice representing the highest price possible for the fancy item, and an array of integers bids representing the other contestants' estimates, your task is to find the winning bid (ie: the value of the bid that would give you the highest chance of winning, assuming a uniform distribution over all possible prices). In the event of a tie, choose the smallest possible bid among the winners.

NOTE: It's possible that some contestants bid above maxPrice!

Example

For maxPrice = 3250 and bids = [2599, 972, 700, 2013], the output should be 973.

With a bid of 973, you would have 1040 chances of winning (you win if the price is anything within the inclusive range from 973 to 2012).

For maxPrice = 75 and bids = [15, 57, 36, 250], the output should be 16.

With a bid of 16, you have 20 chances of winning (anything between 16 and 35 would work). You would have the same number of chances to win with a bid of 37, but since 16 is smaller, it's the one you should choose.

For maxPrice = 10000 and bids = [5000, 3500, 7250, 8999], the output should be 1

The broadest range would be between 1 and 3499, so you should choose 1.
*/

import "sort"

func comeOnDown(maxPrice int, bids []int) int {
    bids = append([]int{0}, bids...)
    bids = append(bids, maxPrice + 1)
    sort.Ints(bids)
    i, r, d := 0, 0, 0
    for {
        if i + 1 == len(bids) {
            break
        }
        if bids[i] >= maxPrice {
            break
        }
        if bids[i+1] - bids[i] > d {
            d = bids[i+1] - bids[i]
            r = bids[i] + 1
        }
        i++
    }
    return r
}
