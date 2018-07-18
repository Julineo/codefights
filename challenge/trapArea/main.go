/*
https://app.codesignal.com/challenge/sGDJsXcFYhkoejcrq

For a regular polygon with n sides and perimeter p, we can find a new similar polygon by extending the radius from the center to each vertex. Given the length of the new radius r, we'd like to find the difference in area between the two shapes.

trapArea

To avoid precision issues, we'll round the final answer to the nearest integer.

Note: in the case that the given radius is smaller than that of the original polygon, the difference in area should be negative.

Example

For n = 5, p = 40, and r = 16, the output should be trapArea(n, p, r) = 499.

The difference between the two areas is about 498.565617, which is rounded to 499.
*/
import . "math"

func trapArea(n int, p float32, r float32) int {
    n1 := float64(n)
    p1 := float64(p)
    r2 := float64(r)
    s1 := p1/n1
    A1 := (n1/4) * Pow(s1, 2) * Cos(Pi/n1)/Sin(Pi/n1)
    A2 := 0.5 * n1 * Pow(r2, 2)  * Sin(2*Pi/n1)
    return int(Round(A2 - A1))
}
