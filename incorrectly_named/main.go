package main

import (
    "fmt"
    "github.com/chmeliik/gomod-multi/wrongly_named/some-package"
)

func main() {
    fmt.Println("Hello from wrongly_named.")
    some_package.Hi()
}
