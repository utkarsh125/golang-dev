# Constants in Go
Go supports constants of `character`, `string`, `boolean`, and numeric values.

`const` declares a constant value, something that won't change (isn't a variable)

A `const` statement can also appear inside a function body.

Constant expressions perform arithmetic with arbitrary precision.

A numeric constant has no type until it's given one, such as by an explicit conversion.

A number can be given a type by using it in a context that requires one, such as a variable assignment or function call. For example, here `math.Sin` expects a `float64`

