package main

import "fmt"

func main() {
	var a [5] int
	a[1] = 100
	a[2] = 100
	a[3] = 100
	a[4] = 100

	fmt.Println(a)
	fmt.Println(a[1])
	fmt.Println(len(a))


	//type 2

	var b [] int8
	b =int8 {1,2,3,4}
	fmt.Println(b)
}
