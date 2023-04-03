package main

import (
  "fmt"
)

var x int
var y string
var z bool

func main() {
  fmt.Print("x: ")
  fmt.Println(x)
  fmt.Print("y: ")
  fmt.Println(y)
  fmt.Print("z: ")
  fmt.Println(z)
  fmt.Printf("Tipos:\nx: %T\ny: %T\nz: %T\n", x, y, z)
}
