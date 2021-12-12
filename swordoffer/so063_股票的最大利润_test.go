/**
 * @author: xv
 * @date: 2021/12/12 17:38
 * @link: https://leetcode-cn.com/problems/gu-piao-de-zui-da-li-run-lcof/
 * @tag: 动态规划
 */

package swordoffer

import (
	"fmt"
	"testing"
)

func TestSo063(t *testing.T) {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	res := 0
	min := prices[0]
	for i := range prices {
		temp := prices[i] - min
		if temp > res {
			res = temp
		}
		if prices[i] < min {
			min = prices[i]
		}
	}
	return res
}
