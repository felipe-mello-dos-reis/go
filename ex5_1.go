package main

import("fmt")

func main() {
  a := 10
  minhafunc(a)
}

func minhafunc(x int){
  fmt.Printf("binário: %b\ndecimal: %d\nhexadecimal: %#x\n", x, x, x)
}
