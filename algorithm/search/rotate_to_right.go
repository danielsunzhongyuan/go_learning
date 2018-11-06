package search

/*
一个整数数组本是有序的（升序），但被右移了X位置（对给定的整数X (0 <= X <= (Length- 1))，
数组的每个元素向右移动X位置，也就是array[I]移动到array[(i+X)%Length]。
请实现函数int findX(into[] array), 找出X。要求时间复杂度越小越好。

用二分法查找比array[0]小的第一个元素，时间复杂度为0(logn)
*/
func RotateToRight(arr []int, length int) int {
	if length < 2 {
		return 0
	}

	if arr[0] > arr[1] {
		return 1
	}

	left, right := 1, length-1
	for left < right {
		mid := (left + right) / 2
		if arr[0] > arr[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	if arr[0] > arr[left] {
		return left % length
	}
	return (left + 1) % length // just in case the array does not rotate
}
