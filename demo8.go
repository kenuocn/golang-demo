package main

import "fmt"

func main() {
	arr := [...] int{0,1,2,3,4,5,6,7,8,9}

	fmt.Println(arr[2:6])  //2 , 3, 4, 5
	fmt.Println(arr[:6]) // 0, 1,2,3,4,5

	s1 := arr[2:]
	fmt.Println(s1) //2,3,4,5,6,7,8,9

	s2 := arr[:]
	fmt.Println(s2) // 0,1,2,3,4,5,6,7,8,9
}
