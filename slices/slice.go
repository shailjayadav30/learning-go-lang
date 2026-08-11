package slices

// import "fmt"
//  search element in slice
func Search(arr []int, target int) (int, int) {
	for i := 0; i < len(arr); i++ {
		if arr[i] == target {
			return i, arr[i]
		}
	}
	return -1, -1
}

// Write a function FilterEvens(nums []int) []int that takes a slice of integers and returns a new slice containing only the even numbers.

func FilterEvenNums(arr []int) []int {
	newArr := make([]int, 0)
	for i := 0; i < len(arr); i++ {
		if arr[i]%2 == 0 {
			newArr = append(newArr, arr[i])
		}

	}
	return newArr
}

// i did this

// func Rotate(arr []int, k int) []int {

// 	l := len(arr)
// 	temp := 0
// 	for i := 0; i < k; i++ {
// 		temp = arr[i]
// 		arr[i] = arr[l-i-1]
// 		arr[l-i-1] = temp
// 	}
// 	// fmt.Println("arr after loop",arr)

// 	arr1 := arr[0:k]
// 	// fmt.Println("arr1",arr1)
// 	arr2 := arr[k:]
// 	// fmt.Println("arr2",arr2)

// 	ans1 := reverse(arr1)
// 	// fmt.Println("ans1",ans1)

// 	ans2 := reverse(arr2)
// 	// fmt.Println("ans2",ans2)

// 	combined := append(ans1, ans2...)
// 	return combined
// }

// func reverse(arr []int) []int {
// 	// fmt.Println("arr inside",arr)

// 	l := len(arr)
// 	temp := 0
// 	for i := 0; i < l-1; i++ {
// 		// fmt.Println("arr loop",arr)

// 		temp = arr[i]
// 		arr[i] = arr[l-i-1]
// 		arr[l-i-1] = temp
// 		// fmt.Println("arr after first iteration", arr)

// 	}
// 	// fmt.Println("arr inside", arr)

// 	return arr
// }

func Rotate(arr []int, k int) []int {
	l := len(arr)
	if l == 0 {
		return arr
	}
	reverse(arr)
	reverse(arr[:k])
	reverse(arr[k:])
	return  arr
}

func reverse(arr []int) {
	left, right := 0, len(arr)-1
	for left < right {
		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}
}
