package main

import (
  "fmt"
)

type sapeca int
var x sapeca

func main() {
  fmt.Println(x)
  fmt.Printf("%T\n", x)
  x = sapeca(42)
  fmt.Println(x)
}


