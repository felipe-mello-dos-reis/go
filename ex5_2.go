package main

import("fmt")
const(
  a = iota + 1
  b
  c
  d
)
func main(){
  fmt.Println(a==b)
  fmt.Println(a<=b)
  fmt.Println(c>b)
  fmt.Println(b!=a)
}
