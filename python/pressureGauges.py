"""
https://app.codesignal.com/arcade/python-arcade/drilling-the-lists/SkTfc263CQbGNMtoj

Given the pressures Harry wrote down for each pipe, return two lists: the first one containing the minimum, and the second one containing the maximum pressure of each pipe during the day.

Example

For morning = [3, 5, 2, 6] and evening = [1, 6, 6, 6],
the output should be
pressureGauges(morning, evening) = [[1, 5, 2, 6], [3, 6, 6, 6]]
"""
def pressureGauges(morning, evening):
    return [[min(a, b) for a, b in zip(morning,evening)], [max(a, b) for a, b in zip(morning,evening)]]
