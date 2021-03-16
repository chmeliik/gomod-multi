package main

import (
    "fmt"
    "github.com/chmeliik/gomod-multi/wrong_v2_local/v2/some-package"
    "github.com/chmeliik/some-module"
    "github.com/chmeliik/some-module/other-package"
)

func main() {
    fmt.Println("Hello from wrong_v2_local.")
    some_package.Hi()
    some_module.Hi()
    other_package.Hi()
}
