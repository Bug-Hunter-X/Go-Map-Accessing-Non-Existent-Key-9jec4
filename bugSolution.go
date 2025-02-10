```go
package main

import "fmt"

func main() {
    var m map[string]int
    val, ok := m["a"]
    if ok {
        fmt.Println(val)
    } else {
        fmt.Println("Key not found")
    }
}
```