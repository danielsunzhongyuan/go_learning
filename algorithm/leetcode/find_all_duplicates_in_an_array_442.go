package main

import "fmt"

/*
Given an array of integers, 1 ≤ a[i] ≤ n (n = size of array), some elements appear twice and others appear once.

Find all the elements that appear twice in this array.

Could you do it without extra space and in O(n) runtime?

Example:
Input:
[4,3,2,7,8,2,3,1]

Output:
[2,3]
 */
func findDuplicates(nums []int) []int {
    ret := make([]int, 0)
    length := len(nums)
    if length == 0 {
        return ret
    }
    for i := 0; i < length; i++ {
        for nums[i] <= length && nums[i] > 0 && nums[i] != i+1 && nums[nums[i]-1] != nums[i] {
            nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
        }
    }
    for i := 0; i < length; i++ {
        if nums[i] != i+1 {
            ret = append(ret, nums[i])
        }
    }
    return ret
}

func main() {
    fmt.Println(findDuplicates([]int{4, 3, 2, 7, 8, 2, 3, 1})) // [2,3]
}
