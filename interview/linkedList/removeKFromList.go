/*
https://app.codesignal.com/interview-practice/task/gX7NXPBrYThXZuanm

Note: Try to solve this task in O(n) time using O(1) additional space, where n is the number of elements in the list, since this is what you'll be asked to do during an interview.

Given a singly linked list of integers l and an integer k, remove all elements from list l that have a value equal to k.

Example

For l = [3, 1, 2, 3, 4, 5] and k = 3, the output should be
removeKFromList(l, k) = [1, 2, 4, 5];
For l = [1, 2, 3, 4, 5, 6, 7] and k = 10, the output should be
removeKFromList(l, k) = [1, 2, 3, 4, 5, 6, 7].

*/
// Definition for singly-linked list:
// type ListNode struct {
//   Value interface{}
//   Next *ListNode
// }
//
func removeKFromList(l *ListNode, k int) *ListNode {
    var cur *ListNode
    cur = l
    for cur != nil && cur.Next != nil {
        if cur.Next.Value == k {
            cur.Next = cur.Next.Next
        } else {
            cur = cur.Next
        }
    }
    if l != nil && l.Value == k {
            l = l.Next
    }
    return l
}

