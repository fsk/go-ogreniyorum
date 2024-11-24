// Loops
package main

import "fmt"

func main() {
	// Klasik for loop
	// parantez kullanilmaz
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// for loop with uint
	for i := uint(0); i < 10; i++ {
		fmt.Println(i)
	}

	// While benzeri loop
	// sonsuz loop asagidaki gibi yazilir

	// for {
	// 	fmt.Println("Sonsuz loop")
	// }

	// condition loop
	// while loop'a benzer
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}

	// Range loop
	arr := []int{1, 2, 3, 4, 5}
	for index, value := range arr {
		fmt.Println("Index: ", index, "Value: ", value)
	}

	arr2 := []string{"ankara", "istanbul", "izmir", "adana", "bursa"}
	for _, value := range arr2 {
		fmt.Println("Value: ", value)
	}

	for index := range arr2 {
		fmt.Println("Index: ", index)
	}

}
