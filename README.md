# Go Map Accessing Non-Existent Key

This example demonstrates a common issue in Go involving accessing a non-existent key in a map.  Unlike some languages that would throw an exception or panic, Go maps gracefully return the zero value for the associated type.  This behavior, while convenient in some cases, can also lead to subtle bugs if not explicitly checked for.

The `bug.go` file illustrates the unexpected behavior. The solution is shown in `bugSolution.go`.

## How to Reproduce
1.  Clone this repository.
2.  Run `go run bug.go` (Prints 0).
3. Run `go run bugSolution.go` (Prints Key not found)

## Solution
The `bugSolution.go` demonstrates a robust way to handle this scenario by explicitly checking the existence of the key before accessing it using the `map[key]` operator.  Alternatively, the solution uses the comma ok idiom.