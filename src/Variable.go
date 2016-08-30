package main

import "fmt"

var x string  = "scope"
func main() {
	var a string = "Anna arthdi putra"
	var b string = "Putra"
	fmt.Println(a)
	fmt.Println(a == b)

	//scope
	fmt.Println(x)

	//function
	data()

	//constant
	const d string  = "anne"
	fmt.Println(d)

}

func data()  {
	fmt.Println(x)
}