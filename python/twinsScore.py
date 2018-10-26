"""
https://app.codesignal.com/arcade/python-arcade/drilling-the-lists/ESdegETWZmXLJYirj

Given two lists of equal size representing the scores Billy and Mandy got for each exam (b and m, respectively), return a single list containing their combined score.

Example

For b = [22, 13, 45, 32] and m = [28, 41, 13, 32],
the output should be
twinsScore(b, m) = [50, 54, 58, 64]
"""
def twinsScore(b, m):
    return [b + m for b, m in zip(b,m)]
