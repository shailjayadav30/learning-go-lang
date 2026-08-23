package ranges

import "fmt"

// ranges used for iteration over data structures
func Rangess() {
	// nums := []int{3, 4, 5, 6}
	// sum := 0
	// for _, num := range nums {
	// 	sum = sum + num
	// 	fmt.Println(sum)
	// }

	// for i, char := range "Shailja" {
	// 	fmt.Println(i, string(char))
	// }

	user := map[string]string{"name": "shailja", "age": "22"}
	for k, v := range user {
		fmt.Println(k, v)
	}
}
