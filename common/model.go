/**
 * @author: xv
 * @date: 2021/11/24 14:59
 */

package common

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Node struct {
	Val    int
	Left   *Node
	Right  *Node
	Next   *Node
	Random *Node
}

func BuildDefaultListNode(count int) *ListNode {
	head := &ListNode{}
	cur := &ListNode{}
	next := &ListNode{}
	for i := 1; i <= count; i++ {
		cur.Val = i
		if i == 1 {
			head = cur
		}
		if i != count {
			cur.Next = next
			cur = cur.Next
			next = &ListNode{}
		}
	}
	return head
}

func BuildListNode(nums []int) *ListNode {
	head := &ListNode{}
	cur := &ListNode{}
	next := &ListNode{}
	for i := 0; i < len(nums); i++ {
		cur.Val = nums[i]
		if i == 0 {
			head = cur
		}
		if i != len(nums)-1 {
			cur.Next = next
			cur = cur.Next
			next = &ListNode{}
		}
	}
	return head
}

func BuildDefaultNode(count int) *Node {
	head := &Node{}
	cur := &Node{}
	next := &Node{}
	for i := 1; i <= count; i++ {
		cur.Val = i
		if i == 1 {
			head = cur
		}
		if i != count {
			cur.Next = next
			cur = cur.Next
			next = &Node{}
		}
	}
	return head
}