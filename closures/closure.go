package closures

func Closure() func(a int) int {
	count := 7
	return func(a int) int {
		return a + count
	}
}
