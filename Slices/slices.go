package main

import (
	"fmt"
	"slices"
)

func main() {

	var s []string
	fmt.Println("uninit: ", s, s == nil, len(s) == 0)

	s = make([]string, 3)
	fmt.Println("Emp: ", s, "Len: ", len(s), "Cap: ", cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("Set: ", s)
	fmt.Println("Get: ", s[2])

	fmt.Println("Len: ", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd: ", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy: ", c)

	l := s[2:5]
	fmt.Println("sl1: ", l)

	l = s[:5]
	fmt.Println("sl2: ", l)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl: ", t)

	t2 := []string{"g", "h", "i"}
	if slices.Equal(t, t2) {
		fmt.Println("t==t2")
	}
	twoD := make([][]int, 3)
	for i := range 3 {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := range innerLen {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2D:", twoD)

}
