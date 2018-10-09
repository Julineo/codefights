package "main"

import (
	"fmt"
)

type ListNode struct {
   Value interface{}
   Next *ListNode
}

func main() {
	var l *ListNode
	l = &ListNode{}
	l.Val = 1
	for i := 2; i <= 5; i++ {
		l.AddAtTail(i)
	}
	fmt.Println(isListPalindrome(l))

	// odd palindrome test
	l = &ListNode
	l.Val = 2
	l.AddAtTail(5)
	l.AddAtTail(3)
	l.AddAtTail(5)
	l.AddAtTail(2)
	fmt.Println(isListPalindrome(l))

	// even palindrome test
	l = &ListNode
	l.Val = 2
	l.AddAtTail(5)
	l.AddAtTail(3)
	l.AddAtTail(3)
	l.AddAtTail(5)
	l.AddAtTail(2)
	fmt.Println(isListPalindrome(l))


}

// checks if list is palindrome
func isListPalindrome (l *ListNode) bool {
	// for tracking the end and the middle of list
	var fast *ListNode
	fast = l

	// for reversing the first half of list
	var prev, cur, next *ListNode
	cur, next = l, l

	// cur goes to the middle, reversing first part of list
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next

		next = next.Next
		cur.Next = prev
		prev = cur
		cur = Next
	}
	next = next.Next
	cur.Next = prev

	// check if palindrome
	// cur in head of first part, next in head of second part

	next = next.Next
	if fast.Next = nil { // odd list
		cur = prev
	} else { // even list
		cur.next = prev
	}

}

// adds nodes at tail
func (this *ListNode) AddAtTail(val int)  {
	var newTail, cur *ListNode
	newTail = &ListNode{val, nil}
	cur = this
	for {
		if cur.Next == nil {
			cur.Next = newTail
			return
		}
		cur = cur.Next
	}
}
