package topkbased

import (
	"fmt"
)

type Node struct {
	Frequency int
	Data      int
	Next      *Node
}

func NodeInsert(head *Node, value, frequency int) *Node {
	newNode := &Node{Frequency: frequency, Data: value}

	// If the list is empty or the new node should be the new head
	if head == nil || head.Frequency < frequency {
		newNode.Next = head
		return newNode
	}

	// Traverse the list to find the correct insertion point
	current := head
	for current.Next != nil && current.Next.Frequency >= frequency {
		current = current.Next
	}

	// Insert the new node
	newNode.Next = current.Next
	current.Next = newNode

	return head
}

func TraverseLink(head *Node, array []int, counter int) []int {
	if head == nil || counter == 0 {
		return array
	}
	array = append(array, head.Data)
	counter -= 1
	array = TraverseLink(head.Next, array, counter)
	return array
}

func topKFrequent(nums []int, k int) []int {
	mC := make(map[int]int)
	for _, value := range nums {
		mC[value] += 1

	}
	var head *Node
	for indexValue, frequecy := range mC {
		head = NodeInsert(head, indexValue, frequecy)
	}
	t := TraverseLink(head, []int{}, k)
	return t
}

func TesttopKFrequest() {
	nums := []int{1, 1, 1, 2, 2, 3}
	k := 2
	fmt.Println(topKFrequent(nums, k))
}
