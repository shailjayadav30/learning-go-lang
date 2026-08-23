package functions

// func Functi(fn func(name string) string)string {
// 	fn("Hello")
// 	return  "shailja"
// }
//? a function that returns another function
func Function() func(a, b int) int {
	return func(a, b int) int {
		return a + b
	}
}
