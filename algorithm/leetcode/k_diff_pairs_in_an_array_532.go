/*
Given an array of integers and an integer k, you need to find the number of unique k-diff pairs in the array. Here a k-diff pair is defined as an integer pair (i, j), where i and j are both numbers in the array and their absolute difference is k.

Example 1:
Input: [3, 1, 4, 1, 5], k = 2
Output: 2
Explanation: There are two 2-diff pairs in the array, (1, 3) and (3, 5).
Although we have two 1s in the input, we should only return the number of unique pairs.
Example 2:
Input:[1, 2, 3, 4, 5], k = 1
Output: 4
Explanation: There are four 1-diff pairs in the array, (1, 2), (2, 3), (3, 4) and (4, 5).
Example 3:
Input: [1, 3, 1, 5, 4], k = 0
Output: 1
Explanation: There is one 0-diff pair in the array, (1, 1).
*/
package main

import "sort"
import "fmt"

func findPairs(nums []int, k int) int {
    sort.Ints(nums)
    length := len(nums)
    count := 0
    if 0 == k {
        for i:=1;i<length;i++ {
            if nums[i] == nums[i-1] {
                count++
            }
        }
    } else {
        for i:=0; i < length - 1; i++ {
            j:=i+1
            for j<length && nums[j]<nums[i]+k {
                j++
            }
            if j < length && nums[j] == nums[i] + k {
                count++
            }
        }
    }
    return count
}

func main() {
    fmt.Println(findPairs([]int{3,1,4,1,5},2))
}
