package main

import "fmt"
/*
On the first row, we write a 0. Now in every subsequent row, we look at the previous row and replace each occurrence of 0 with 01, and each occurrence of 1 with 10.

Given row N and index K, return the K-th indexed symbol in row N. (The values of K are 1-indexed.) (1 indexed).

Examples:
Input: N = 1, K = 1
Output: 0

Input: N = 2, K = 1
Output: 0

Input: N = 2, K = 2
Output: 1

Input: N = 4, K = 5
Output: 1

Explanation:
row 1: 0
row 2: 01
row 3: 0110
row 4: 01101001
Note:

N will be an integer in the range [1, 30].
K will be an integer in the range [1, 2^(N-1)].
 */

//var level = make([]int, 30, 30)
//
//func init() {
//    for i := uint(0); i < 30; i++ {
//        level[i] = 1 << i
//    }
//}
//func kthGrammar(N int, K int) int {
//    if K == 1 {
//        return 0
//    }
//
//    var i = 0
//    for i = 0; i < 30; i++ {
//        if level[i] >= K {
//            break
//        }
//    }
//    return 1 ^ kthGrammar(N, K-level[i-1])
//
//}

import "math"

func kthGrammar(N int, K int) int {
    i := 0
    for K > 1 {
        K -= int(math.Pow(2, float64(int(math.Log2(float64(K-1))))))
        i += 1
    }
    if i%2 == 0 {
        return 0
    }
    return 1
}

func main() {
    fmt.Println(kthGrammar(1, 1))
    fmt.Println(kthGrammar(2, 1))
    fmt.Println(kthGrammar(2, 2))
    fmt.Println(kthGrammar(4, 5))
}
