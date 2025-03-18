package main

import "fmt"

func main() {
	var array1 [5]string

	array1[0] = "Posição 1"

	fmt.Println(array1)

	array2 := [5]string{"Posição 1", "Posição 2", "Posição 4", "Posição 5"}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

	slice := []int{1, 2, 3, 4, 5, 6, 87, 7, 7, 7}
	fmt.Println(slice)

	slice = append(slice, 8)
	fmt.Println(slice)

	slice2 := array2[1:3]
	fmt.Println(slice2)

}
