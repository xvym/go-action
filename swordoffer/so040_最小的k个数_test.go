/**
 * @author: xv
 * @date: 2021/12/16 10:55
 * @link: https://leetcode-cn.com/problems/zui-xiao-de-kge-shu-lcof
 * @tag: 排序
 */

package swordoffer

import (
	"fmt"
	"testing"
)

func TestSo040(t *testing.T) {
	fmt.Println(getLeastNumbers([]int{0, 1, 2, 1}, 1))
}

func getLeastNumbers(arr []int, k int) []int {
	// 快排超时了
	heapSortInSo040(arr, k)
	arr = append(arr[:k])
	return arr
}

func heapSortInSo040(arr []int, k int) {
	length := len(arr)
	for i := length / 2; i >= 0; i-- {
		buildMaxHeapInSo040(arr, i, length)
	}
	round := length - k
	for j := length - 1; j >= 0 && round >= 0; j-- {
		arr[0], arr[length-1] = arr[length-1], arr[0]
		length--
		round--
		buildMaxHeapInSo040(arr, 0, length)
	}
}

func buildMaxHeapInSo040(arr []int, i int, length int) {
	biggestIndex := i
	leftChildIndex := i*2 + 1
	rightChildIndex := i*2 + 2
	if leftChildIndex < length && arr[biggestIndex] < arr[leftChildIndex] {
		biggestIndex = leftChildIndex
	}
	if rightChildIndex < length && arr[biggestIndex] < arr[rightChildIndex] {
		biggestIndex = rightChildIndex
	}
	if biggestIndex != i {
		arr[biggestIndex], arr[i] = arr[i], arr[biggestIndex]
		buildMaxHeapInSo040(arr, biggestIndex, length)
	}
}