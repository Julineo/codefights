"""
https://app.codesignal.com/arcade/python-arcade/fumbling-in-functional/TjCNTwysvW6za5Qh4

Given result, return an array of the same length, where the ith element is equal to the ith element of result with the last digit dropped.

Example

For result = [42, 239, 365, 50], the output should be
fixResult(result) = [4, 23, 36, 5]
"""
def fixResult(result):
    def fix(x):
        return x / 10

    return map(fix, result)
