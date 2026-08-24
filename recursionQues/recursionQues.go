package recursionques

import "fmt"

// Print Name N times using Recursion

func PrintName(name string, n int) {
	if n > 0 {
		fmt.Println("times", n, name)
		PrintName(name, n-1)
	}

}

// Print 1 to N using Recursion

func Print1ToN(n int) {
	if n >= 1 {
		Print1ToN(n - 1)
		fmt.Println(n)
	}

}

// Print N to 1 using Recursion

func PrintNTo1(n int) {
	if n >= 1 {
		fmt.Println(n)
		PrintNTo1(n - 1)
	}

}

// Sum of first N Natural Numbers

func SumOfN(n int) int {
	if n < 1 {
		return 0
	}
	return n + SumOfN(n-1)
}

// Factorial of a Number : Iterative and Recursive

func Factorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return n * Factorial(n-1)
}

// Reverse a given Array
func ReverseArray(arr []int, n int) []int {
	if n < 0 || n == len(arr)/2 {
		return arr
	}
	l := len(arr)
	temp := arr[l-n]
	arr[l-n] = arr[n-1]
	arr[n-1] = temp

	return ReverseArray(arr, n-1)

}

// Check if the given String is Palindrome or not

func PalindromeStr(str string, start, end int) bool {
	if start >= end {
		return true
	}
	if str[start] != str[end] {
		return false
	}

	return PalindromeStr(str, start+1, end-1)

}

func PalindromeStr2(str string) bool {
	if len(str) <= 1 {
		return true
	}
	if str[0] != str[len(str)-1] {
		return false
	}
	return PalindromeStr2(str[1 : len(str)-1])

}

func FibonacciSeries(n int) int {
	if n <= 1 {
		return n
	}

	return FibonacciSeries(n-1) + FibonacciSeries(n-2)

}
