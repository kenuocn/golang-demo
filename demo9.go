package main

import "fmt"

func updateSlice(arr []int)  {
	arr[0] = 100
}

func main() {
	// Slice 学习.

	arr := [...]int{0,1,2,3,4,5,6,7,8,9}

	fmt.Println(arr[2:6])  //2 , 3, 4, 5
	fmt.Println(arr[:6]) // 0, 1,2,3,4,5

	s1 := arr[2:]
	fmt.Println(s1) //2,3,4,5,6,7,8,9

	s2 := arr[:]
	fmt.Println(s2) // 0,1,2,3,4,5,6,7,8,9

	updateSlice(s1)
	fmt.Println(s1)  // 100,3,4,5,6,7,8,9
	fmt.Println(arr)

	updateSlice(s2)

	fmt.Println(s2) // 100,1,100,3,4,5,6,7,8,9
	fmt.Println(arr) // 100,1,100,3,4,5,6,7,8,9


	arr1 := []int{0,1,2,3,4,5,6,7}

	s := arr1[2:6]
	s[0] = 10
	fmt.Println(arr1) // 0,1,10,2,3,4,5,6,7

	/*
		Slice 本身没有数据,是对底层array的一个view
		arr 的值变为[0,1,10,2,3,4,5,6,7]
	 */
}
