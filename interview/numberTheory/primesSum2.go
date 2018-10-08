/*
https://app.codesignal.com/interview-practice/task/D7n2Cj9JDYbQzWiZE

Given a positive integer n, calculate the sum of all the prime numbers from 1 to n, inclusive. Because this number may be very big, return it modulo 109 + 7 in your output.

Example

For n = 6, the output should be
primesSum2(n) = 10.

The sum of the prime numbers from 1 to 6, inclusive, (2 + 3 + 5) is 10.

For n = 11, the output should be
primesSum2(n) = 28.

The sum of the prime numbers from 1 to 11, inclusive, (2 + 3 + 5 + 7 + 11) is 28.*/

func primesSum2(n int) int {
    isNotPrime := make([]bool, n + 1)
    sum := 0
    mod := 1000000007
    for i := 2; i <= n; i++ {
        if !isNotPrime[i] {
            if i * i <= n {
                for j := 2 * i; j <= n; j += i {
                    isNotPrime[j] = true
                }
            }
            sum = (sum + i) % mod
        }
    }
    return sum
}

//time limit error, not working with slices/arrays?
/*func primesSum2(n int) int {
    if n == 1 { return 0 }
    if n == 2 { return 2 }
    if n == 3 { return 5 }
    res := 5
    
    primes := []int{2, 3}
   
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

/*
//for some reason linked list even slower than slice
// linked list
type listNode struct {
    Val int
    Next *listNode
}

// adds element at tail to linked list
func (this *listNode) AddAtTail(val int) {
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
func (this *listNode) checkInList(val int) bool {AddAtTail
	var cur *istNode
	cur = this
	for cur != nil {
        if cur.Val == val { return true }
		cur = cur.Next
	}
    return false
}
*/

/*
func primesSum2(n int) int {
    if n == 1 { return 0 }
    if n == 2 { return 2 }
    if n == 3 { return 5 }
    var res int
    res = 5

    // linked list should be faster than slice/array?
    var primes *listNode
    primes = &listNode{}

    // fill list with initial 2, 3, 5 values
    primes.Val = 2
    primes.AddAtTail(3)

    primes.showList()


    flag := true
    var c int
    for c = 5; c <= n; c = c + 2 {//count only odd numbers
        flag = true

        var cur *listNode
	    cur = primes
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
    fmt.Println(10000000007)
              //1099511627776
    return res % (10 * 10 * 10 * 10 * 10 * 10 * 10 * 10 * 10 * 10 + 7)
}
*/

//Maps are slowere than slices x2 times
/*
func primesSum2(n int) int {
    if n == 1 { return 0 }
    if n == 2 { return 2 }
    if n == 3 { return 5 }
    res := 5
    
    primes = make(map[int]struct{})
    
    primes[2] = struct{}{}
    primes[3] = struct{}{}
   
    flag := true
    for c := 5; c <= n; c = c + 2 {//count only odd numbers
        flag = true
        for a, _ := range primes {
            if c % a == 0 {
                flag = false
                break
            }
        }
        if flag {
            primes[c] = struct{}{}
            res += c
        }
    }
    fmt.Println(10000000007)
              //1099511627776
    return res % (10000000007)
}
*/
