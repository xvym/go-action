/**
 * @author: xv
 * @date: 2021/11/17 23:50
 * @link: https://leetcode-cn.com/problems/first-bad-version/
 */

package leetcode

import "testing"

func TestLc278(t *testing.T) {
	firstBadVersion(9)
}

func firstBadVersion(n int) int {
	start := 1
	end := n
	for start <= end {
		mid := start + (end-start)/2
		if isBadVersion(mid) {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return start
}

func isBadVersion(num int) bool {
	return true
}
