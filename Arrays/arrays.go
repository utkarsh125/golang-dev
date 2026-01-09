package main

import "fmt"

func main() {

	//Here we create an array 'a' that wil hold exactly 5 ints.
	//The type of elements and lengths are both part of the array's type.
	//By default an array is zero-valued, which for ints means 0s.
	var a [5]int //array of size 5
	fmt.Println("Emp: ", a)

	//We can set a value at an index using the array[index] = value syntax.
	//and get a value with array[index]
	a[4] = 100
	fmt.Println("Set: ", a)
	fmt.Println("Get: ", a[4])

	fmt.Println("Len: ", len(a)) //length of array 'a'

	//Use this syntax to declare and initialize an array in one line.
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl: ", b)

	//you can also have the compiler count the number of elements for you with ...
	b = [...]int{1, 2, 3, 4, 5}
	fmt.Println("idx: ", b)

	//if you specify the index with :, the elements in between will be zeroed.

	b = [...]int{100, 3: 400, 500}
	fmt.Println("idx: ", b)

	var twoD [2][3]int
	for i := range 2 {
		for j := range 3 {
			twoD[i][j] = i + j
		}
	}

	fmt.Println("2d: ", twoD)

	twoD = [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println("2d:", twoD)

}
