package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var carry int = 0

	var root, current *ListNode

	current = new(ListNode)
	for {
		val := add(l1, l2, carry)
		if val == nil {
			current = nil
			break
		}

		*val = *val % 10
		carry = *val / 10

		current.Val = *val
		current.Next = new(ListNode)
		next := current.Next
		current = next

		if root == nil {
			root = &ListNode{Val: *val, Next: next}
		}

		l1 = l1.Next
		l2 = l2.Next
	}
	return root
}

func add(l1, l2 *ListNode, carry int) *int {
	if l1 == nil && l2 == nil && carry == 0 {
		return nil
	}

	if l1 != nil {
		carry += l1.Val
	}

	if l2 != nil {
		carry += l2.Val
	}
	return &carry
}
