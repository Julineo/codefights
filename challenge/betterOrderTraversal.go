/*
There are several ways to traverse a binary tree. Let's talk about some of the most popular ways!

In-order traversal involves looking at the left subtree, then the current node, then the right subtree. This will visit the values of a binary search tree in ascending order.

Pre-order traversal involves looking at the current node, then the left subtree, then the right subtree. This type of traversal is often used for copying a tree.

Post-order traversal involves looking at the left subtree, then the right subtree, then the current node. This type of traversal is often used for deleting a tree.

Given a binary tree, your task is to test each of these traversals - store the values in an array (in the order in which they're visited), and return the one that's lexicographically smallest.

Example

For

{
    "value": -29,
    "left": {
        "value": 95,
        "left": null,
        "right": null
    },
    "right": {
        "value": 2,
        "left": null,
        "right": {
            "value": 100,
            "left": null,
            "right": null
        }
    }
}
the output should be betterOrderTraversal(root) = [-29, 95, 2, 100]

example

The different traversal types would visit the nodes in the following orders:

In-order: [95, -29, 2, 100]
Pre-order: [-29, 95, 2, 100]
Post-order: [95, 100, 2, -29]
So in this case, pre-order traversal produces the lexicographically smallest result.
*/

//
// Definition for binary tree:
// type Tree struct {
//   Value interface{}
//   Left *Tree
//   Right *Tree
// }
package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

func main () {
	t := tree.New(1)
	fmt.Println(betterOrderTraversal(t))
}

func betterOrderTraversal(root *tree.Tree) []int {
    a1 := []int{}
    a2 := []int{}
    a3 := []int{}

    var inOrderTr func(root *tree.Tree)
    inOrderTr = func(root *tree.Tree) {
        if root == nil {
            return
        }

        if root.Left != nil {
            inOrderTr(root.Left)
        }

        a1 = append(a1, root.Value)

        if root.Right != nil {
            inOrderTr(root.Right)
        }
    }

    var preOrderTr func(root *tree.Tree)
    preOrderTr = func(root *tree.Tree) {
        if root == nil {
            return
        }

        a2 = append(a2, root.Value)

        if root.Left != nil {
            preOrderTr(root.Left)
        }

        if root.Right != nil {
            preOrderTr(root.Right)
        }
    }

    var postOrderTr func(root *tree.Tree)
    postOrderTr = func(root *tree.Tree) {
        if root == nil {
            return
        }

        if root.Left != nil {
            postOrderTr(root.Left)
        }

        if root.Right != nil {
            postOrderTr(root.Right)
        }

        a3 = append(a3, root.Value)
    }

    inOrderTr(root)
    preOrderTr(root)
    postOrderTr(root)
	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)

    f1, f2, f3 := true, true, true
    for i := range a1 {
        if f1 && f2 && f3 {
            if a2[i] < a1[i] || a3[i] < a1[i] { f1 = false }
            if a3[i] < a2[i] || a1[i] < a2[i] { f2 = false }
            if a1[i] < a3[i] || a2[i] < a3[i] { f3 = false }
        }

        if f2 && f3 {
            if a3[i] < a2[i] { f2 = false }
            if a2[i] < a3[i] { f3 = false }
        }
        if f1 && f3 {
            if a3[i] < a1[i] { f1 = false }
            if a1[i] < a3[i] { f3 = false }
        }
        if f1 && f2 {
            if a2[i] < a1[i] { f1 = false }
            if a1[i] < a2[i] { f2 = false }
        }

        if f3 && !f1 && !f2 { return a3 }
        if f1 && !f2 && !f3 { return a1 }
        if f2 && !f1 && !f3 { return a2 }
    }

    return a1
}

