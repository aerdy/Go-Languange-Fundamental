package main

import "fmt"

func main() {
	var a int = 10
	var b int = 20

	jumlah(a, b)


	//function kurang dengn balikan
	fmt.Println(kurang(a, b))
}

func jumlah(a int, b int) {
	var c int = a + b
	fmt.Println(c)
}


//retunt
func kurang(a int, b int) int {
	var c int = a - b
	return c
}