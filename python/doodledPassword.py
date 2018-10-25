"""
https://app.codesignal.com/arcade/python-arcade/caravan-of-collections/aarR4B273h5D2x8ry

Implement the missing code, denoted by ellipses. You may not modify the pre-existing code.
Your friend has been doodling during the lecture and wrote down several digits in a circle. You're wondering if these digits might form the password to your friend's computer. You're planning to prank him some time in the future, and hacking into his computer will definitely help. If the digits written in the clockwise order indeed form a password, all you need to do is figure out which digit comes in it first.

Given a list of digits as they are written in the clockwise order, find all other combinations the password could possibly have.

Example

For digits = [1, 2, 3, 4, 5], the output should be

doodledPassword(digits) = [[1, 2, 3, 4, 5], [2, 3, 4, 5, 1], [3, 4, 5, 1, 2],
                           [4, 5, 1, 2, 3], [5, 1, 2, 3, 4]]
"""
from collections import deque

def doodledPassword(digits):
    n = len(digits)
    res = [deque(digits) for _ in range(n)]
    map(lambda x,y: x.rotate(y), res, range(n,0,-1))
    return [list(d) for d in res]
