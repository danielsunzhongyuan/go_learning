package search

func BinarySearch(arr []int, x int) int {
	length := len(arr)
	i, j := 0, length-1
	for i <= j {
		mid := (i + j) / 2
		if x < arr[mid] {
			j = mid - 1
		} else if x > arr[mid] {
			i = mid + 1
		} else {
			return mid
		}
	}
	return -1
}
