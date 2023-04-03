package main

import (
  "fmt"
)

type sapeca int
var x sapeca

func main() {
  fmt.Println(x)
  fmt.Printf("%T\n", x)
  x = 42
  fmt.Println(x)
  y := int(x)
  fmt.Printf("%v %T\n", y,y)
}

