## Uninitialized slice

Here we declare a slice without initializing it. An uninitialized slice is `nil`. Its length is `0`, and comparing it to `nil` returns `true`.

`var s []string fmt.Println("uninit:", s, s == nil, len(s) == 0)`

---

## Creating a slice with make

We use `make` to create a slice with a specific length. At this point, the slice is initialized, but all elements are set to their zero value, which for strings is an empty string.

`s = make([]string, 3) fmt.Println("Emp:", s, "Len:", len(s), "Cap:", cap(s))`

---

## Setting and getting slice elements

We can set values using index assignment and retrieve them using indexing, just like arrays. Since slices are backed by arrays, indexing works the same way.

`s[0] = "a" s[1] = "b" s[2] = "c" fmt.Println("Set:", s) fmt.Println("Get:", s[2])`

---

## Slice length

The `len` function returns the current number of elements in the slice. Unlike arrays, this value can change as elements are added or removed.

`fmt.Println("Len:", len(s))`

---

## Appending to a slice

Slices grow dynamically using `append`. If the underlying array does not have enough capacity, Go automatically allocates a new array and copies the elements.

`s = append(s, "d") s = append(s, "e", "f") fmt.Println("apd:", s)`

---

## Copying a slice

To copy a slice, we create a new slice with the same length and use the built-in `copy` function. This copies the elements, not just the reference.

`c := make([]string, len(s)) copy(c, s) fmt.Println("cpy:", c)`

---

## Slice expressions

Slice expressions allow us to create a new view into an existing slice. The resulting slice shares the same underlying array.

`l := s[2:5] fmt.Println("sl1:", l)`

Here we omit the starting index, which defaults to `0`.

`l = s[:5] fmt.Println("sl2:", l)`

---

## Declaring and initializing a slice literal

Slices can also be declared and initialized directly using a slice literal. This is often the cleanest and most common way to create slices.

`t := []string{"g", "h", "i"} fmt.Println("dcl:", t)`

---

## Comparing slices

Slices cannot be compared using `==`, except against `nil`. The `slices.Equal` function compares elements one by one and returns true if both slices contain the same values in the same order.

`t2 := []string{"g", "h", "i"} if slices.Equal(t, t2) { 	fmt.Println("t==t2") }`

---
## Two-dimensional slices

Here we create a slice of slices. Each inner slice can have a different length, which is something arrays cannot do.

`twoD := make([][]int, 3) for i := range 3 { 	innerLen := i + 1 	twoD[i] = make([]int, innerLen) 	for j := range innerLen { 		twoD[i][j] = i + j 	} } fmt.Println("2D:", twoD)`