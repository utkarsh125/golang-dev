# For Loop in Go
`for` is Go's only looping construct. Here are some basic types of loops.

The most basic type, with a single condition.

```go
i := 1
    for i <= 3 {
        fmt.Println(i)
        i = i + 1
    }
...


A classic initial/condition/after `for` loop.
```go
for j := 0; j < 3; j++ {
    fmt.Println(j)
}
```

Another way of accomplishing the basic "do this N times" iteration is `range` over an integer.

```go
for i := range 3 {
    fmt.Println("range", i)
}
```

`for` without a condition will loop repeatedly until you `break` out of the loop or `return` from the enclosing function.

```go
for {
	fmt.Println("loop")
	break
}
```

You can also use `continue` for the next iteration of the loop.

```go
for n := range 6 {
	if n%2 == 0 {
		continue
	}
	fmt.Println(n)
}
```

