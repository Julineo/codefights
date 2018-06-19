/*
Consider integer numbers from 0 to n - 1 written down along the circle in such a way that the distance between any two neighbouring numbers is equal (note that 0 and n - 1 are neighbouring, too).

Given n and firstNumber, find the number which is written in the radially opposite position to firstNumber.

Example

For n = 10 and firstNumber = 2, the output should be
circleOfNumbers(n, firstNumber) = 7.
*/

func circleOfNumbers(n int, firstNumber int) int {
	return (firstNumber + n/2) % n
}

/*
func circleOfNumbers(n int, firstNumber int) int {
    if firstNumber >= n/2 {
        return firstNumber - n/2
    }
    return n/2 + firstNumber
}
*/
