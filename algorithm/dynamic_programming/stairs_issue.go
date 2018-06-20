package main

func stairsRecursion(number int) int {
	if number <= 2 {
		return number
	}
	return stairsRecursion(number-1) + stairsRecursion(number-2)
}

func stairs(number int) int {
	memo := make(map[int]int)
	memo[1] = 1
	memo[2] = 2
	count := 2
	for count < number {
		count++
		memo[count] = memo[count-1] + memo[count-2]
	}
	return memo[number]
}

func main() {
	print(stairs(50))
}
