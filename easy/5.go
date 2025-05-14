package easy

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// took hard time to understand linkedlist | 1ms Beats 4.09%
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	}

	head := list1

	for {
		// exit := false
		curr := head
		prev := head
		temp := *list2
		for {
			if temp.Val <= curr.Val {
				temp.Next = curr
				if curr == head {
					head = &temp
				} else {
					prev.Next = &temp
				}
				break
			}
			prev = curr
			if curr.Next != nil {
				curr = curr.Next
			} else {
				curr.Next = &temp
				list2.Next = nil
				break
			}
		}

		if list2.Next == nil {
			break
		}
		list2 = list2.Next
	}

	return head
}

func createList(l []int) *ListNode {
	if len(l) == 0 {
		return nil
	}
	prev := &ListNode{
		Val: l[0],
	}
	head := prev
	for i := 1; i < len(l); i++ {
		curr := &ListNode{
			Val: l[i],
		}

		prev.Next = curr
		if i == 1 {
			head = prev
		}
		prev = curr
	}
	return head
}

func loopLL(l *ListNode) {
	if l == nil {
		fmt.Println("")
		return
	}
	for {
		fmt.Printf("%d ", l.Val)
		if l.Next == nil {
			break
		}
		l = l.Next
	}
	fmt.Println()
}

func RunMrg2List() {
	list1 := createList([]int{1, 2})
	list2 := createList([]int{0})
	fmt.Print("Input: list1 = ")
	loopLL(list1)
	fmt.Print("Input: list2 = ")
	loopLL(list2)
	l := mergeTwoLists(list1, list2)
	fmt.Println("After merge")
	loopLL(l)

}

// You are given the heads of two sorted linked lists list1 and list2.

// Merge the two lists into one sorted list. The list should be made by splicing together the nodes of the first two lists.

// Return the head of the merged linked list.

// Example 1:
// Input: list1 = [1,2,4], list2 = [1,3,4]
// Output: [1,1,2,3,4,4]

// Example 2:
// Input: list1 = [], list2 = []
// Output: []

// Example 3:
// Input: list1 = [], list2 = [0]
// Output: [0]
