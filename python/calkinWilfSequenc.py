"""
https://app.codesignal.com/arcade/python-arcade/yin-and-yang/ynSRuyh93ZffkPjtv

Given a rational number, your task is to return its 0-based index in the Calkin-Wilf sequence.

Example

For number = [1, 3], the output should be
calkinWilfSequence(number) = 3.

As you can see in the image above, 1 / 3 is the 3rd 0-based number in the sequence.
"""
def calkinWilfSequence(number):
    def fractions():
        a = b = 1
        while True:
            yield [a, b]
            a, b = b, 2 * (a - a % b) + b - a

    gen = fractions()
    res = 0
    while next(gen) != number:
        res += 1
    return res
