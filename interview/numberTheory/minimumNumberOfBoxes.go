/*
https://app.codesignal.com/interview-practice/task/kKPFt5uegeQSMvR4e

The store has an infinite number of apple boxes and banana boxes. There are a apples in each apple box, and b bananas in each banana box. We are interested in whether we can get exactly f pieces of fruit by ordering boxes of apples or bananas.

Given a, b, and f, your task is to find the minimum n such that it is possible to order n boxes of fruit (either apple or bananas) to get f pieces of fruit in total. Because we can only order whole boxes (and we cannot order a negative number of boxes), it may not be possible to complete the order. If it is not possible to order f pieces of fruit, return -1.

Example

For a = 5, b = 4, and f = 19, the output should be
minimumNumberOfBoxes(a, b, f) = 4.

If we get 3 boxes of apples and 1 box of bananas, then we have 3 · 5 + 1 · 4 = 19 pieces of fruit. The minimum number of boxes we need to order is 4.

For a = 2, b = 8, and f = 1001, the output should be
minimumNumberOfBoxes(a, b, f) = -1.

We can only order an even amount of fruit (either in boxes of 2 or boxes of 8), but we are required to make an odd amount, which cannot be done.

For a = 5, b = 7, and f = 84, the output should be
minimumNumberOfBoxes(a, b, f) = 12.

We can order 12 boxes of 7 bananas to get 84 pieces of fruit.

For a = 5, b = 7, and f = 35, the output should be
minimumNumberOfBoxes(a, b, f) = 5.

We can order 5 boxes of bananas, or 7 boxes of apples. The option with fewer boxes is 5 boxes of bananas.

For a = 5, b = 7, and f = 8, the output should be
minimumNumberOfBoxes(a, b, f) = -1.

Having one box of either apples or bananas is not enough, but any combination of two boxes gives too many pieces of fruit, so we cannot make 8 pieces.

Input/Output

[execution time limit] 4 seconds (go)

[input] integer a

The number of apples in each apple box.

Guaranteed constraints:
0 ≤ a ≤ 109.

[input] integer b

The number of bananas in each bananas box.

Guaranteed constraints:
0 ≤ b ≤ 109.

[input] integer f

The number of fruits we need to order.

Guaranteed constraints:
1 ≤ f ≤ 109.

[output] integer

Return the minimal number of boxes we need in order to get f pieces of fruit, or -1 if it isn't possible.
*/

import "math"

func minimumNumberOfBoxes(a int, b int, f int) int {
    if a + b == 0 {
        return -1
    }
    d := gcd(a, b)
    if a > b {
        a, b = b, a
    }

    if f % d != 0 || d == 0 {
        return -1
    }

    if a == 0 || b == 0 {
        return f / d
    }

    xp, yp := egcd(a / d, b / d)
    x0, y0 := f * xp / d, f * yp / d
    t := int(math.Ceil(-float64(x0) * float64(d) / float64(b)))
    x, y := x0 + b * t / d, y0 - a * t / d
    if x < 0 || y < 0 {
        return -1
    }
    return x + y
}

func egcd(a, b int) (int, int) {
    prevx, x := 1, 0
    prevy, y := 0, 1

    for b > 0 {
        q := a / b
        x, prevx = prevx - q * x, x
        y, prevy = prevy - q * y, y
        a, b = b, a % b
    }
    return prevx, prevy
}

func gcd(a, b int) int {
    if a * b == 0 {
        return a + b
    }
    return gcd(b, a % b)
}
