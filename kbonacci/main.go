/*
Define the k-bonacci sequence as follows. The first k terms are all 1. The nth term, for n >= k, is the sum of k previous terms in the sequence. For example, the first five terms of the 3-bonacci sequence are 1, 1, 1, 3, 5.

Given n and k, write a function which returns the nth term (0-indexed) of the k-bonacci sequence. Since the answer may be quite large, return it as a string.

Example
For k = 3 and n = 4, the output should be kbonacci(k, n) = "5".
*/

package main

import (
	"fmt"
	"time"
	"math/big"
	"strconv"
)

var (
	m          map[int]int
	nCacheHits int
	k	= 40
	n       = 25000
)

func do1() {
	timeStart := time.Now()
	res := kbonacci(k, n)
	fmt.Printf("res: %d, time: %s\n", res, time.Since(timeStart))
}

func do2() {
	m = map[int]int{}
	nCacheHits = 0
	timeStart := time.Now()
	res := kbonaccim(k, n)
	fmt.Printf("res: %d, time: %s, cache hits: %d\n", res, time.Since(timeStart), nCacheHits)
}

func main() {
	fmt.Printf("Running for n=%d\n", n)
	//do1()
	do2()
}

func kbonaccim(k int, n int) string {
	m := make(map[string]*big.Int)
	
	key := big.NewInt(int64(k))
	en := big.NewInt(int64(n))
	c := big.NewInt(0)
	start := big.NewInt(0)
	one := big.NewInt(1)

	var helper func(k, n *big.Int) *big.Int
	helper = func(k, n *big.Int) *big.Int {
		c.Sub(k, one)
		if n.Cmp(c) <= 0 {
			return one
		}
		if v, ok := m[n.String()]; ok {
			nCacheHits++
			return v
		}
		total := big.NewInt(0)
		start.Sub(n, k)
		for i := new(big.Int).Set(start); i.Cmp(n) == -1; i.Add(i, one) {
			total.Add(total, helper(k, i))
		}
		m[n.String()] = total
		return m[n.String()]
	}
	return helper(key, en).String()
}

func kbonacci(k int, n int) string {
	m := make(map[int]int)
	
	var helper func(k, n int) int
	helper = func(k, n int) int {
		if n <= k-1 {
			return 1
		}
		if v, ok := m[n]; ok {
			return v
		}
		total := 0
		for i := n - k; i < n; i++ {
			total += helper(k, i)
		}
		m[n] = total
		return m[n]
	}
	return strconv.Itoa(helper(k, n))
}
