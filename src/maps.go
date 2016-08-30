package main

import "fmt"

func main() {
	var a = make(map[int]int)
	a[1] = 10
	fmt.Println(a[1])


	var b = make(map[string]int)
	b["a"] = 10
	fmt.Println(b["a"])

	var c = map[string]string{
		"a": "a",
		"b": "b",
	}
	fmt.Println(c)
	fmt.Println(c["b"])

	//delete maps
	delete(a,1)
	fmt.Println(a[1])


}
