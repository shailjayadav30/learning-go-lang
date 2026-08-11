package arrays

import "fmt"

func Array() {
	var nums [4]string
	nums[0] = "h"
	nums[1] = "e"
	nums[2] = "l"
	nums[3] = "o"
	var vals [5]bool

	num := [2][2]int{{1, 3}, {2, 4}}
	fmt.Println(num)
	values := [4]int{3, 4, 5, 6}
	fmt.Println(values)
	fmt.Println(vals)

	fmt.Println(len(nums))
	fmt.Println(nums[0])
	fmt.Println(nums)
	// generally we dont use arrays much we only use when we the size is predictable otherwise we use slice
	// memory optimization
	// can access in constant time
}
