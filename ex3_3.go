package main

import (
  "fmt"
)

var x int = 42
var y string = "James Bond"
var z bool = true

func main() {
  fmt.Print("x: ")
  fmt.Println(x)
  fmt.Print("y: ")
  fmt.Println(y)
  fmt.Print("z: ")
  fmt.Println(z)
  fmt.Printf("Tipos:\nx: %T\ny: %T\nz: %T\n", x, y, z)

  s := fmt.Sprintf("%v\t %v\t %v", x, y, z)
  fmt.Println(s)
}
