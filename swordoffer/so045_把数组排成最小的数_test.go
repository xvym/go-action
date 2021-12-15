/**
 * @author: xv
 * @date: 2021/12/15 12:00
 * @link: https://leetcode-cn.com/problems/ba-shu-zu-pai-cheng-zui-xiao-de-shu-lcof/
 * @tag: 排序
 */

package swordoffer

import (
	"fmt"
	"sort"
	"strconv"
	"testing"
)

func TestSo045(t *testing.T) {
	fmt.Println(minNumber([]int{3, 30, 34, 5, 9}))
}

func minNumber(nums []int) string {
	strs := make([]string, len(nums))
	for i := range nums {
		strs[i] = strconv.Itoa(nums[i])
	}
	sort.Sort(byCustom(strs))
	res := string("")
	for i := range strs {
		res += strs[i]
	}
	return res
}

type byCustom []string

func (p byCustom) Len() int {
	return len(p)
}

func (p byCustom) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p byCustom) Less(i, j int) bool {
	return p[i]+p[j] < p[j]+p[i]
}