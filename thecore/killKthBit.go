/*
https://app.codesignal.com/arcade/code-arcade/corner-of-0s-and-1s/b5z4P2r2CGCtf8HCR

You figured out that some numbers have a modified single digit in their binary representation. More specifically, in the given number n the kth bit from the right was initially set to 0, but its current value might be different. It's now up to you to write a function that will change the kth bit of n back to 0.

Example

For n = 37 and k = 3, the output should be
killKthBit(n, k) = 33.

3710 = 1001012 ~> 1000012 = 3310.

For n = 37 and k = 4, the output should be
killKthBit(n, k) = 37.

The 4th bit is 0 already (looks like the Mad Coder forgot to encrypt this number), so the answer is still 37.
*/
func killKthBit(n int, k int) int {
  return n & ^(1 << uint8(k-1))
}
