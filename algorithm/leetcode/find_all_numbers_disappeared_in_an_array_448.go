package main

import "fmt"

/*
Given an array of integers where 1 ≤ a[i] ≤ n (n = size of array), some elements appear twice and others appear once.

Find all the elements of [1, n] inclusive that do not appear in this array.

Could you do it without extra space and in O(n) runtime? You may assume the returned list does not count as extra space.

Example:

Input:
[4,3,2,7,8,2,3,1]

Output:
[5,6]
*/

func findDisappearedNumbers(nums []int) []int {
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
            ret = append(ret, i+1)
        }
    }
    return ret
}

func main() {
    fmt.Println(findDisappearedNumbers([]int{4, 3, 2, 7, 8, 2, 3, 1})) // [5,6]
}
