/**
 * @author: xv
 * @date: 2021/12/16 14:31
 * @link: https://leetcode-cn.com/problems/kth-largest-element-in-an-array/
 * @tag: 排序
 */

package leetcode

import (
	"fmt"
	"testing"
)

func TestLc215(t *testing.T) {
	fmt.Println(findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4))
}

func findKthLargest(nums []int, k int) int {
	heapSort(nums)
	return nums[k-1]
}

func heapSort(nums []int) {
	// 第一步，构建堆，将无序序列构建成一个小顶堆
	// 叶子节点不用调整（因为每个叶子节点就已经构成了堆），对每个叶子节点都构建小顶堆
	// 从最后一个非叶子节点开始,按从右到左的顺序开始构建堆，根据完全二叉树性质可知最后一个非叶子节点的索引为 len/2-1
	length := len(nums)
	for i := length / 2; i >= 0; i-- {
		// 每个非叶子节点的左右子节点索引分别为 i * 2 - 1和i * 2 + 1
		buildMinHeap(nums, i, length)
	}
	// 此时我们已经构建了一个堆顶为数组中最小值的堆，但数组还不是有序的。这时我们要将最小值放和数组未排序部分的末尾作交换，然后再重新构建小顶堆
	for j := length - 1; j > 0; j-- {
		nums[0], nums[j] = nums[j], nums[0]
		length--
		buildMinHeap(nums, 0, length)
	}
}

func buildMinHeap(nums []int, i int, length int) {
	smallestIndex := i
	leftChildIndex := i*2 + 1
	rightChildIndex := i*2 + 2
	if leftChildIndex < length && nums[leftChildIndex] < nums[smallestIndex] {
		smallestIndex = leftChildIndex
	}
	if rightChildIndex < length && nums[rightChildIndex] < nums[smallestIndex] {
		smallestIndex = rightChildIndex
	}
	if smallestIndex != i {
		nums[i], nums[smallestIndex] = nums[smallestIndex], nums[i]
		// 因为交换后子节点的值发生了变化（变大了），所以要对以子节点为顶的堆进行最小堆化
		buildMinHeap(nums, smallestIndex, length)
	}
}