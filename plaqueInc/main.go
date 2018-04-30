package main

import (
	"fmt"
)

func main() {

	people := [][]int{
		{0, 1},
		{1, 0},
	}

	people = [][]int{
		{0, 1},
		{1, 0, 3},
		{2, 3},
		{3, 1, 2},
	}

	fmt.Println(people)
	fmt.Println("result: ", plagueInc(people))
}

func plagueInc(people [][]int) int {
	paths := make(map[int]int)
	var infected map[int]bool

	var touch func(p int) bool
	touch = func(p int) bool {

		infected[p] = true
		fmt.Println(infected)
		paths[p]++
		for i := 1; i < len(people[p]); i++ {
			fmt.Println(p, people[p][i])
			if !infected[people[p][i]] {
				touch(people[p][i])
			}
		}
		return false
	}

	for i := 0; i < len(people); i++ {
		infected = make(map[int]bool)
		fmt.Println(infected)
		touch(i)
	}
	fmt.Println(paths)

	min := paths[0]
	idx := 0
	
	
	//return minimum
	for i := 1; i < len(paths); i++ {
		if paths[i] < min {
			min = paths[i]
			idx = i
		}
	}
	return idx
}
