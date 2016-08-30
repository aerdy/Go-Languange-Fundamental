package main

import "fmt"

func main() {
	var a  = new (Nama)
	a = Nama{x:1,y:3}

	fmt.Println(a)
}

type Nama struct {
	x,y int
}
