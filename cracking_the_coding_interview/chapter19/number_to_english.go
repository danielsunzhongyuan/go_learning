package main

import (
	"fmt"
)

/*
Given an integer between 0 and 999,999 print an English phrase that describes the integer.
For example:
    One Thousand, Two Hundred and Thirty Four
*/

func NumberToEnglish(num int) string {
	if num == 0 {
		return "Zero"
	}

	result := ""

	if num < 1000 {
		result += NumberLessThousand(num)
	} else if num == 1000 {
		result += "One Thousand"
	} else if num > 1000 && num%1000 > 0 {
		result += NumberLessThousand(num/1000) + " Thousand, " + NumberLessThousand(num%1000)
	} else if num > 1000 && num%1000 == 0 {
		result += NumberLessThousand(num/1000) + " Thousand"
	}
	return result
}

func NumberLessThousand(x int) string {

	wordArr1 := []string{"", "One", "Two", "Three", "Four",
		"Five", "Six", "Seven", "Eight", "Nine"}

	result := ""

	if x < 100 {
		result += NumberLessHundred(x)
	} else if x == 100 {
		result += "One Hundred"
	} else if x > 100 && x%100 > 0 {
		result += wordArr1[x/100] + " Hundred and " + NumberLessHundred(x%100)
	} else if x > 100 && x%100 == 0 {
		result += wordArr1[x/100] + " Hundred"
	}
	return result
}

func NumberLessHundred(x int) string {
	wordArr1 := []string{"", "One", "Two", "Three", "Four",
		"Five", "Six", "Seven", "Eight", "Nine"}
	wordArr11 := []string{"", "Eleven", "Twelve", "Thirteen", "Fourteen", "Fifteen",
		"Sixteen", "Seventeen", "Eighteen", "NineTeen"}
	wordArr10 := []string{"", "Ten", "Twenty", "Thirty", "Forty", "Fifty", "Sixty",
		"Seventy", "Eighty", "Ninety"}

	result := ""
	if x < 10 {
		result += wordArr1[x]
	} else if x == 10 {
		result += wordArr10[1]
	} else if x > 10 && x < 20 {
		result += wordArr11[x-10]
	} else if x == 20 {
		result += wordArr10[2]
	} else if x > 20 && x%10 == 0 {
		result += wordArr10[x/10]
	} else if x > 20 && x%10 > 0 {
		result += wordArr10[x/10] + " " + wordArr1[x%10]
	}
	return result
}

func main() {
	/*
		0~199 Tested
		900~1199 Tested
	*/
	for i := 1100; i < 1200; i++ {
		fmt.Println(i, NumberToEnglish(i))
	}

	fmt.Println(999999, NumberToEnglish(999999))

}
