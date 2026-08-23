package mapss

import (
	"fmt"
	// "maps"
	"strings"
)

//  maps are like hash,object and dictionary in any lang

func Mapsss() {
	// creating map with make keyword
	//  variable_name:=make(map[key datatype]value datatype)

	m := make(map[string]string) //!uninitialized map will return map[] nil map
	// fmt.Println(m)  
	// setting an element
	m["name"] = "eagle"
	m["type"] = "bird"
	m["color"] = "black"
	// map[color:black name:eagle type:bird] o/p
	fmt.Println(m)
	// fmt.Println(m["type"])
	//! if key doesn't exists in map then it returns default value of the datatype
	// fmt.Println(m["legs"])
	// delete(m, "color")
	// fmt.Println(len(m))

	// creating map without make keyword
	// mapp := map[string]int{"price": 40}
	// fmt.Println(mapp)
	// k, ok := mapp["prices"]
	// if ok {
	// 	fmt.Println("All ok", k)
	// } else {
	// 	fmt.Println("Not ok", k)

	// }
	// how to check if 2 maps are equal

	// m1 := map[string]string{"name": "shailja"}
	// m2 := map[string]string{"name": "shailja"}
	// fmt.Println(maps.Equal(m1, m2))

}

// 1. Word Frequency CounterWrite a function wordCount(s string) map[string]int that takes a string sentence, splits it into words, and returns a map counting how many times each word appears.

func WordFrequency(s string) map[string]int {
	splitedString := strings.Fields(s)
	m := make(map[string]int)

	for i := 0; i < len(splitedString); i++ {
		_, ok := m[splitedString[i]]
		if ok {
			m[splitedString[i]]++
		} else {
			m[splitedString[i]]++
		}
	}
	return m
}
