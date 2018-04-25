package main

import "fmt"

func main() {

	//b , kb , mb,gb ,tb,pb
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)

	fmt.Println(b , kb , mb,gb ,tb,pb)
}
