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
	"sync"

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

    var wg sync.WaitGroup

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

    wg.Add(3)

    go func() {
        inOrderTr(root)
        wg.Done()
    }()
    go func() {
        preOrderTr(root)
        wg.Done()
    }()
    go func() {
        postOrderTr(root)
        wg.Done()
    }()

    wg.Wait()
    return min(a1, min(a2, a3))
}

func min (a, b []int) []int {
    for i := range a {
        if a[i] < b[i] { return a }
        if b[i] < a[i] { return b }
    }
    return a
}
