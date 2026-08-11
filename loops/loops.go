package loops

import (
	"fmt"
)

func Loop() {
	//? implement while loop as we don't have while loop in go
	// var i  =1
	// for i<=5{
	// 	fmt.Println(i)
	// 	i+=1
	// }

	//? syntax of for loop

	// for i:=1;i<=6;i++{
	// 	fmt.Println(i)
	// }

	//? infinite loop

	// for{
	// 	fmt.Println(1)
	// }

	//? range- used to do something no of times  starts from 0 and goes till range-1

	// for i := range 7 {
	// 	fmt.Println(i)
	// }

	//? if else
	// role := "user"
	// hasPermission := false
	// age := 10
	// if age >= 18 {
	// 	fmt.Println("Can vote")
	// } else if age < 18 && age > 12 {
	// 	fmt.Println("Cannot vote")
	// } else {
	// 	fmt.Println("Child")
	// }
	// if role == "user" || hasPermission {
	// 	fmt.Println("yes")
	// }

	//* Direct variable declaration inside if

	if isvalid := false; isvalid {
		fmt.Println("is valid", isvalid)
	} else {
		fmt.Println("is valid", isvalid)
	}

	//! go does not have ternary operator so you have to use basic  if else

	//? switch  case
	//! we dont need to write break keyword go internally handles it and default case is also optional if not given then also no error occurs
	//* simple switch case syntax
	// i := 7
	// switch i {
	// case 1:
	// 	fmt.Println("case 1")
	// case 2:
	// 	fmt.Println("case 2")
	// case 3:
	// 	fmt.Println("case 3")

	// case 4:
	// 	fmt.Println("case 4")
	// default:
	// 	fmt.Println("No case matched")
	// }

	//! multiple condition switch
	// role := "developer"
	// switch role {
	// case "developer", "admin":
	// 	fmt.Println("can access db")
	// case "user":
	// 	fmt.Println("cannot access fb")
	// }

	//! type switch
	// interface{} if we do like this then it means it is of any type means it can take any value

	whoami := func(i interface{}) {
		switch t := i.(type) {
		case int:
			fmt.Println("the value  is integer", t)
		case string:
			fmt.Println("the value is string")
		case bool:
			fmt.Println("the value is bool")
		case float32:
			fmt.Println("the value is float32")
		default:
			fmt.Println("the value is default")
		}
	}
	whoami(111.2222222222)
}
