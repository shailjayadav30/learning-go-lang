package patternques

import (
	"fmt"
	// "strings"
)

// *****
// *****
// *****
// *****
// *****
func Pattern1(row, col int) {

	for i := 1; i <= row; i++ {
		for j := 1; j <= col; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

}

// *
// **
// ***
// ****
// *****

func Pattern2(row, col int) {
	for i := 1; i <= row; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

}

// 1
// 12
// 123
// 1234
// 12345

func Pattern3(row, col int) {
	for i := 1; i <= row; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(j)
		}
		fmt.Println()
	}

}

// 1
// 22
// 333
// 4444
// 55555

func Pattern4(row, col int) {
	for i := 1; i <= row; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(i)
		}
		fmt.Println()
	}

}

// *****
// ****
// ***
// **
// *
func Pattern5(row, col int) {
	for i := row; i >= 1; i-- {
		for j := i; j >= 1; j-- {
			fmt.Print("*")
		}
		fmt.Println()
	}

}

// 12345
// 1234
// 123
// 12
// 1
func Pattern6(row, col int) {
	for i := 5; i >= 1; i-- {
		for j := 1; j <= i; j++ {
			fmt.Print(j)
		}
		fmt.Println()
	}

}

//	   *
//	  ***
//	 *****
//	*******
//
// *********
func Pattern7(N int) {
	for i := 0; i <= N-1; i++ {
		for l := 0; l < N-i-1; l++ {
			fmt.Print(" ")
		}
		for k := 0; k < 2*i+1; k++ {
			fmt.Print("*")

		}
		for r := 0; r < N-i-1; r++ {
			fmt.Print(" ")
		}

		fmt.Println()

	}

}

// *********
//
//	*******
//	 *****
//	  ***
//	   *
func Pattern8(N int) {
	for i := 0; i <= N-1; i++ {
		for l := 0; l < i; l++ {
			fmt.Print(" ")
		}
		for k := 0; k < 2*N-(2*i+1); k++ {
			fmt.Print("*")

		}
		for r := 0; r < i; r++ {
			fmt.Print(" ")
		}

		fmt.Println()

	}

}

//	   *
//	  ***
//	 *****
//	*******
//
// *********
// *********
//
//	*******
//	 *****
//	  ***
//	   *
func Pattern9(N int) {
	for i := 0; i <= N-1; i++ {
		for l := 0; l < N-i-1; l++ {
			fmt.Print(" ")
		}
		for k := 0; k < 2*i+1; k++ {
			fmt.Print("*")

		}
		for r := 0; r < N-i-1; r++ {
			fmt.Print(" ")
		}

		fmt.Println()

	}

	for i := 0; i <= N-1; i++ {
		for l := 0; l < i; l++ {
			fmt.Print(" ")
		}
		for k := 0; k < 2*N-(2*i+1); k++ {
			fmt.Print("*")

		}
		for r := 0; r < i; r++ {
			fmt.Print(" ")
		}

		fmt.Println()

	}
}

// *
// **
// ***
// ****
// *****
// ****
// ***
// **
// *

func Pattern10(n int) {
	for i := 0; i < n; i++ {

		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()

	}
	for k := n - 1; k >= 1; k-- {
		for l := 1; l <= k; l++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

}

// 1
// 01
// 101
// 0101
// 10101
func Pattern11(n int) {
	var start int
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			start = 1
		} else {
			start = 0
		}
		for j := 0; j <= i; j++ {
			fmt.Print(start)
			if start == 0 {
				start = 1
			} else {
				start = 0
			}
		}

		fmt.Println()
	}
}

// 1        1
// 12      21
// 123    321
// 1234  4321
// 1234554321
func Pattern12(n int) {
	space := 2 * (n - 1)
	for i := 1; i <= n; i++ {
		for l := 1; l <= i; l++ {
			fmt.Print(l)
		}
		for s := 1; s <= space; s++ {
			fmt.Print(" ")
		}
		for r := i; r >= 1; r-- {
			fmt.Print(r)
		}
		fmt.Println()
		space -= 2
	}
}

// 1
// 2 3
// 4 5 6
// 7 8 9 10
// 11 12 13 14 15

func Pattern13(n int) {
	value := 1
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print(value)
			fmt.Print(" ")
			value++
		}
		fmt.Println()
	}
}

// A
// AB
// ABC
// ABCD
// ABCDE
func Pattern14(n int) {
	for i := 65; i < 65+n; i++ {
		for j := 65; j <= i; j++ {
			fmt.Printf("%c", j)
		}
		fmt.Println()
	}
}

// ABCDE
// ABCD
// ABC
// AB
// A

func Pattern15(n int) {
	for i := 5; i > 0; i-- {
		for j := 65; j < 65+i; j++ {
			fmt.Printf("%c", j)
		}
		fmt.Println()
	}
}

// A
// BB
// CCC
// DDDD
// EEEEE

func Pattern16(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Printf("%c", (65 + i))
		}
		fmt.Println()
	}
}

func Pattern17(n int) {

	for i := 0; i < n; i++ {

		for l := 0; l < n-i-1; l++ {
			fmt.Print(" ")
		}

		ch := 'A'

		for k := 0; k <= i; k++ {
			fmt.Printf("%c", ch)
			ch++
		}
		ch -= 2

		for k := 0; k < i; k++ {
			fmt.Printf("%c", ch)
			ch--
		}

		fmt.Println()
	}
}
