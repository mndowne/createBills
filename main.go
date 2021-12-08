package main

import "fmt"

func main() {
	mybill := newBill("marty's bill")

	fmt.Println(mybill.format())

}
