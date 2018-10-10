package main

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
	l.Value = 1
	for i := 2; i <= 5; i++ {
		l.AddAtTail(i)
	}

	l.showList()

	fmt.Println(isListPalindrome(l))

	// odd palindrome test
	l = &ListNode{}
	l.Value = 2
	l.AddAtTail(5)
	l.AddAtTail(3)
	l.AddAtTail(5)
	l.AddAtTail(2)
	fmt.Println(isListPalindrome(l))

	// even palindrome test
	l = &ListNode{}
	l.Value = 2
	l.AddAtTail(5)
	l.AddAtTail(3)
	l.AddAtTail(3)
	l.AddAtTail(5)
	l.AddAtTail(2)
	fmt.Println(isListPalindrome(l))


}

// checks if list is palindrome
func isListPalindrome (l *ListNode) bool {
	if l == nil { return true }

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
		cur = next
	}
	next = next.Next

	// check if palindrome
	// cur in head of first part, next in head of second part
	if fast.Next == nil { // odd list
		cur = prev
	} else { // even list
		cur.Next = prev
	}

	for cur != nil {
		if cur.Value != next.Value { return false }
		cur = cur.Next
		next = next.Next
	}
	return true
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

// shows linked list
func (this *ListNode) showList() {
	var cur *ListNode
	cur = this
	for cur != nil {
		fmt.Println(cur, &cur)
		cur = cur.Next
	}
}
