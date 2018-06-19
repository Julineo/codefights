/*
Given a string, find out if its characters can be rearranged to form a palindrome.

Example

For inputString = "aabb", the output should be
palindromeRearranging(inputString) = true.

We can rearrange "aabb" to make "abba", which is a palindrome.
*/

func palindromeRearranging(inputString string) bool {
	m := make(map[rune]int)
	for _, v := range inputString {
		m[v]++
	}
	fmt.Println(m)
	c := 0
	for _, v:= range m {
		if v % 2 > 0 {
			c++
		}
	}
	if c > 1 {
		return false
	}
	return true
}
