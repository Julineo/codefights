package main

import (
	"fmt"
)
// linked list
type listNode struct {
    Val int64
    Next *listNode
}

// adds element at tail to linked list
func (this *listNode) AddAtTail(val int64) {
    var newTail, cur * listNode
    newTail = &listNode{val, nil}
    cur = this
    for {
        if cur.Next == nil {
            cur.Next = newTail
            return
        }
        cur = cur.Next
    }
}

// prints linked list
func (this *listNode) showList() {
	var cur *listNode
	cur = this
	for cur != nil {
		fmt.Println(cur, &cur)
		cur = cur.Next
	}
}


// checks if value in linked list
/*
func (this *listNode) checkInList(val int64) bool {AddAtTail
	var cur *istNode
	cur = this
	for cur != nil {
        if cur.Val == val { return true }
		cur = cur.Next
	}
    return false
}
*/

func main() {
	fmt.Println(primesSum2(6)) // 10
}

func primesSum2(n int64) int64 {
    if n == 1 { return 0 }
    if n == 2 { return 2 }
    if n == 3 { return 5 }
    var res int64
    res = 5

    // linked list should be faster than slice/array?
    var primes *listNode
    primes = &listNode{}

    // fill list with initial 2, 3, 5 values
    primes.Val = 2
    primes.AddAtTail(3)

    primes.showList()


    flag := true
    var c int64
    for c = 5; c <= n; c = c + 2 {//count only odd numbers
	fmt.Println(c)
        flag = true

        var cur *listNode
	    cur = primes
	    fmt.Println(cur, &cur)
	    for cur != nil {
                if c % cur.Val == 0 {
                    flag = false
                    break
                }
		    cur = cur.Next
	    }

        if flag {
            primes.AddAtTail(c)
            res += c
        }
    }
    fmt.Println(int64(10000000007))
              //1099511627776
    return res % (10 * 10 * 10 * 10 * 10 * 10 * 10 * 10 * 10 * 10 + 7)
}

/* time limit error, not working with slices/arrays?
func primesSum2(n int64) int64 {
    if n == 1 { return 0 }
    if n == 2 { return 2 }
    if n == 3 { return 5 }
    res := 5
    
    primes := []int64{2, 3}
   
    flag := true
    for c := 5; c <= n; c = c + 2 {//count only odd numbers
        flag = true
        for _, a := range primes {
            if c % a == 0 {
                flag = false
                break
            }
        }
        if flag {
            primes = append(primes, c)
            res += c
        }
    }
    fmt.Println(10000000007)
              //1099511627776
    return res % (10 * 10 * 10 * 10 * 10 * 10 * 10 * 10 * 10 * 10 + 7)
}
*/
