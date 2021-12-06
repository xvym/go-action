/**
 * @author: xv
 * @date: 2021/12/6 10:58
 * @link: https://leetcode-cn.com/problems/cong-wei-dao-tou-da-yin-lian-biao-lcof/
 * @tag: 链表
 */

package swordoffer

import (
	"fmt"
	. "go-action/common"
	"testing"
)

func TestSo006(t *testing.T) {
	head := BuildDefaultListNode(5)
	fmt.Println(reversePrint(head))
}

func reversePrint(head *ListNode) []int {
	length := 0
	cur := head
	for cur != nil {
		cur = cur.Next
		length++
	}
	res := make([]int, length)

	cur = head
	for i := 1; i <= length; i++ {
		res[length-i] = cur.Val
		cur = cur.Next
	}
	return res
}