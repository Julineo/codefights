/*
https://app.codesignal.com/challenge/cEpnffZAHbeFQjqvC

Little Panda likes candies very much, so his father gives him coins[i] coins on the ith day to buy some. There is only one candy store in Panda's home town, and on the ith day they sell only one candy with price candies[i][0] and tastiness candies[i][1].

Little Panda buys the candy on the ith day if he can afford it and that makes him happier by candies[i][1] points. He saves all the coins left after the purchase to buy more candies later. If he can't afford the candy, it doesn't necessarily make him less happy (obviously it doesn't make him happier anyway). It only decreases his happiness if he hasn't bought any candy tastier than the current one before. In that case, his happiness decreases by candies[i][1] points.

Your task is to find how happy Panda will be after all the days. Panda has 0 coins before the day number 0.
*/
func happyPanda(coins []int, candies [][]int) (happiness int) {
    wallet := 0
    max := -1
    for i, coin := range coins {
        wallet += coin
        if wallet >= candies[i][0] {
            wallet -= candies[i][0]
            if max < candies[i][1] {
                max = candies[i][1]
            }
            happiness += candies[i][1]
            continue
        }
        if max <= candies[i][1]{
            happiness -= candies[i][1]
        }
    }
    return happiness
}
