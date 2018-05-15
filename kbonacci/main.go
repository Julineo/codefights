func kbonacci(k, n int) int {
	if n <= k - 1 {
		return 1
	}
	total := 0
	i := n - k
	for i < n {
		total = total + kbonacci(k, i)
		i = i + 1
	}
return total
}
