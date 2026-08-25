package basicMath

import "math"

// Count digits in a number
func CountDigit(n int) int {
	if n < 10 {
		return 1
	}
	return 1 + CountDigit(n/10)

}

// Reverse Digits of A Number
func ReverseNO(n int) int {
	s := 0
	// temp := n
	for n > 0 {
		r := n % 10
		s = s*10 + r
		n = n / 10
	}
	return s
}

// Palindrome number

func PalindromeNO(n int) bool {
	s := 0
	temp := n
	for temp > 0 {
		r := temp % 10
		s = s*10 + r
		temp = temp / 10
	}
	if s == n {
		return true
	}
	return false
}

// GCD if two numbers

func GCD(a, b int) int {
	d := a
	gcd := 1
	if b < d {
		d = b
	}

	for i := 1; i <= d; i++ {
		if a%i == 0 && b%i == 0 {
			gcd = i
		}
	}
	return gcd
}

// Check if a number is Armstrong Number or not

func ArmstrongNo(n int) bool {
	temp := n
	count := 0
	ans := 0
	newTemp := n
	for temp > 0 {
		count++
		temp = temp / 10
	}
	for newTemp > 0 {
		r := newTemp % 10
		ans = ans + int(math.Pow(float64(r), float64(count)))
		newTemp = newTemp / 10
	}
	if ans == n {
		return true
	}
	return false
}
