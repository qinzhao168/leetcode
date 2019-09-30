/*
选择和冒泡的区别：
	选择排序是走一趟找出来一个最小的值和第一个同学交换位置。
	而冒泡排序是相邻同学比较高低，这样走一趟，最高个就沉到末尾了。

*/
package sort

// 1.1冒泡排序，向后移动，从小到大排序,其思想为相邻两个数进行比较，将较大的滞后，时间复杂度O(N^2
func BubbleSort(array []int) {
	/*
		它重复地走访过要排序的元素列，依次比较两个相邻的元素，如果他们的顺序（如从大到小、首字母从A到Z）错误就把他们交换过来。
		走访元素的工作是重复地进行直到没有相邻元素需要交换，也就是说该元素列已经排序完成。
	*/
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array)-1-i; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	return
}

// 1.2
func bubbleSort2(array []int) {
	/*每一趟扫描的时候，对那些有序的部分，设置一个标志，如果为true表示发生了交换，如果为false，
	则没有发生交换，那么我们就可以直接跳过去了。*/
	flag := true
	for flag {
		flag = false
		for j := 1; j < len(array); j++ {
			if array[j-1] > array[j] {
				array[j-1], array[j] = array[j], array[j-1]
				flag = true
			}
		}
	}
	return
}

// 1.3
func bubbleSort3(array []int) {
	/*
		这个改进就是把flag变为具体的位置了，这样我们就可以记录末尾的边界。这个边界是排序与不排序的边界。
	*/
	flag := len(array)
	var k int
	for flag > 0 {
		k = flag
		flag = 0
		for j := 1; j < k; j++ {
			if array[j-1] > array[j] {
				array[j-1], array[j] = array[j], array[j-1]
				flag = j
			}
		}
	}
	return
}

//2.选择排序，从后面找个最小的放在前面的位置，从小到大排序,时间复杂度O(N^2)
func SelectSort(array []int) {
	/*
		选择排序（Selection sort）是一种简单直观的排序算法。它的工作原理是：第一次从待排序的数据元素中选出最小（或最大）的一个元素，
		存放在序列的起始位置，然后再从剩余的未排序元素中寻找到最小（大）元素，然后放到已排序的序列的末尾。以此类推，
		直到全部待排序的数据元素的个数为零。选择排序是不稳定的排序方法。*/
	for i := 0; i < len(array)-1; i++ {
		for j := i + 1; j < len(array); j++ {
			if array[i] > array[j] {
				array[i], array[j] = array[j], array[i]
			}
		}
	}

	return
}

//3.快速排序
func QuickSort(array []int) {

	return
}

//4.堆排序
func HeapSort(array []int) {

	return
}
