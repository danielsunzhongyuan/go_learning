package search

func InterplotionSearch(arr []int, x int) int {
	length := len(arr)
	i, j := 0, length-1
	for i <= j {
		mid := i + (x-arr[i])/(arr[j]-arr[i])*(j-i)
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
