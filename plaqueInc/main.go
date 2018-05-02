package main

import (
	"fmt"
)

func main() {

	people := [][]int{
		{0, 1},
		{1, 0},
	}

	/*people = [][]int{
		{0, 1},
		{1, 0, 3},
		{2, 3},
		{3, 1, 2},
	}*/

	/*people = [][]int{
		{0, 1, 2},
		{1, 0, 2},
		{2, 1, 0},
		{3, 5, 4},
		{4, 3, 5},
		{5, 4, 3},
	}*/

	/*people = [][]int{
		{0, 1},
		{1, 0, 2},
		{2, 1, 3},
		{3, 2, 4},
		{4, 3, 5},
		{5, 4, 6},
		{6, 5, 7},
		{7, 6, 8},
		{8, 7, 9},
		{9, 8, 10},
		{10, 9, 11},
		{11, 10, 12},
		{12, 11, 13},
		{13, 12, 14},
		{14, 13, 15},
		{15, 14, 16},
		{16, 15, 17},
		{17, 16, 18},
		{18, 17, 19},
		{19, 18, 20},
		{20, 19, 21},
		{21, 20, 22},
		{22, 21, 23},
		{23, 22, 24},
		{24, 23, 25},
		{25, 24, 26},
		{26, 25, 27},
		{27, 26, 28},
		{28, 27, 29},
		{29, 28, 30},
		{30, 29, 31},
		{31, 30, 32},
		{32, 31, 33},
		{33, 32, 34},
		{34, 33, 35},
		{35, 34, 36},
		{36, 35, 37},
		{37, 36, 38},
		{38, 37, 39},
		{39, 38, 40},
		{40, 39, 41},
		{41, 40, 42},
		{42, 41, 43},
		{43, 42, 44},
		{44, 43, 45},
		{45, 44, 46},
		{46, 45, 47},
		{47, 46, 48},
		{48, 47, 49},
		{49, 48, 50},
		{50, 49},
	}*/
	//Expected output: 25

	/*people = [][]int{
		{0, 2},
		{1, 2},
		{2, 0, 1},
	}*/

	people = [][]int{
		{0, 1},
		{1, 0, 2},
		{2, 1, 3},
		{3, 2, 4},
		{4, 3, 5},
		{5, 4, 6},
		{6, 5, 7},
		{7, 6},
		{8},
		{9, 10},
		{10, 9},
	}

	people = [][]int{
		{0, 1},
		{1, 0, 2},
		{2, 1, 3},
		{3, 2, 4, 18},
		{4, 3, 5},
		{5, 4, 6},
		{6, 5, 7},
		{7, 6, 8},
		{8, 7, 9},
		{9, 8, 10},
		{10, 9, 11},
		{11, 10, 12},
		{12, 11, 13},
		{13, 12, 14},
		{14, 13, 15},
		{15, 14, 16},
		{16, 15, 17},
		{17, 16, 18},
		{18, 17, 3, 19},
		{19, 18}}
	//expected output: 3

	fmt.Println(people)
	fmt.Println("result: ", plagueInc(people))
}

// BFS-based solution
func plagueInc(p [][]int) int {
	n := len(p)
	// output, best distance
	o, b := -1, n
	for i := range p {
		// enqueue i
		que := []int{i}
		visited := make([]bool,len(p))
		// step, left to visit
		s, v := -2, n

		// iterate until nothing left to visit
		for v > 0 && s < b - 1 {
			// visit every node in the queue that hasn't been visited
			for _, f := range que {
				if !visited[f] {
					visited[f] = true
					v--
					que = append(que, p[f]...)
				}
				que = que[1:]
			}
			s++
		}
		if v < 1 {
			b = s
			o = i
		}
	}
	return o
}
