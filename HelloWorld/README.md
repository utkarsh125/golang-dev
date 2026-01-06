# Hello World
The very first program when learning any new language should be the `hello_world` file.

It makes you take some steps
- You write the very first lines of code to know about the nature of the syntax
- You learn to execute the code

To run the program, put the code in `hello-world.go` and use `go run <fileName>`.
Sometimes we'll want to build our programs into binaries. We can do this using `go build`.
We can then execute the built binary directly.

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello World")
}
```