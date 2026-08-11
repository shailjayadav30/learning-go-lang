package main

import (
	"fmt"
	"shailjayadav30/golang/slices"
	// slices "shailjayadav30/golang/slices"
	// arrays "shailjayadav30/golang/arrays"
	// lecture_one "shailjayadav30/golang/lecture1"
	// loops "shailjayadav30/golang/loops"
	// simplevariable "shailjayadav30/golang/simple_variable"
)

func main() {
	// lecture_one.Hello()
	// simplevariable.Simplevalues()
	// loops.Loop()
	// arrays.Array()
	// num := []int{1, 2, 3, 4, 5, 6}
	// index,value:=slices.Search(num,4)
	// fmt.Println("ans",index,value)
	num2:=[] int{1,2,3,4,5}
	ans:=slices.Rotate(num2,2)
fmt.Println("rotate",ans)
}

