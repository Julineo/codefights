"""
https://app.codesignal.com/arcade/python-arcade/fumbling-in-functional/cYG6vtfv6TJKPsvSY

For the given list of denominators, find the least common denominator by finding their LCM.

Example

For denominators = [2, 3, 4, 5, 6], the output should be
leastCommonDenominator(denominators) = 60
"""
from fractions import gcd

def leastCommonDenominator(denominators):
    return reduce(lambda x,y: x * y/ gcd(x,y), denominators)
