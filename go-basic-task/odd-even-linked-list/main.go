package main

import "fmt"

func main() {

	fmt.Printf("328. 奇偶链表：https://leetcode.cn/problems/odd-even-linked-list/description/\n")

	var head *ListNode = new(ListNode)
	head.Val = 1
	head.Next = new(ListNode)
	head.Next.Val = 2
	head.Next.Next = new(ListNode)
	head.Next.Next.Val = 3
	head.Next.Next.Next = new(ListNode)
	head.Next.Next.Next.Val = 4
	head.Next.Next.Next.Next = new(ListNode)
	head.Next.Next.Next.Next.Val = 5
	head = oddEvenList(head)

	for {
		if head == nil {
			break
		}
		// 修改为 fmt.Println
		fmt.Println(head.Val)
		head = head.Next
	}
}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	index := 1
	preNode := head
	runNode := head
	firstEvent := head.Next
	for {
		if index != 0 {
			mergePreNodeAndNextNode(runNode, preNode)
		}
		if runNode.Next == nil {
			if index%2 == 0 {
				preNode.Next = firstEvent
			} else {
				runNode.Next = firstEvent
			}
			break
		}
		preNode = runNode
		runNode = runNode.Next
		index++
	}
	return head
}

func mergePreNodeAndNextNode(curNode, preNode *ListNode) {
	if curNode.Next == nil {
		preNode.Next = nil
	} else {
		preNode.Next = curNode.Next
	}
}
