/*
We'd like to construct a diverse array of numbers. At each step, we'll be given two choices for the next number we can add, and we'd like to select the number that appears least frequently in our array so far. If both numbers appear with equal frequency, we'll choose the smaller one.

Our choices will be given in the form of an array, choices, consisting of 2-element arrays of integers.

Example

For choices = [[1, 2], [3, 4], [1, 2]], the output should be leastAppearance(choices) = [1, 3, 2].

Initially, our array is empty, so given the choice between 1 and 2, we'll pick 1 since it's smaller.

On the next step, our array looks like [1], which doesn't contain 3 or 4, so we'll pick 3 (again, because it's smaller than 4).

On the final step, our array looks like [1, 3], so we'll pick 2 since the array already contains a 1.
*/

package main

import "fmt"

func main() {
	fmt.Println(leastAppearance(
		[][]int{
			{1, 2},
			{3, 4},
			{1, 2}}))

	fmt.Println(leastAppearance(
		[][]int{
			{1, 2}}))

	fmt.Println(leastAppearance(
		[][]int{
			{}, {1, 2}}))

	fmt.Println(leastAppearance(
		[][]int{
			{}, {}}))
}

func leastAppearance(choices [][]int) []int {
	counts := make(map[int]int)
	res := []int{}
	if len(choices) == 0 {
		return res
	}
	for _, c := range choices {
		if len(c) == 0 {
			continue
		}
		if counts[c[0]] > counts[c[1]] {
			res = append(res, c[1])
			counts[c[1]]++
			continue
		}
		if counts[c[0]] < counts[c[1]] {
			res = append(res, c[0])
			counts[c[0]]++
			continue
		}
		if counts[c[0]] == counts[c[1]] {
			m := min(c[0], c[1])
			res = append(res, m)
			counts[m]++
			continue
		}
	}
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

