package main

import("fmt")

func main(){
  x := 123
  fmt.Printf("%v\t%d\t%b\t%#x\n", x,x,x,x)
  y := x << 1
  fmt.Printf("%v\t%d\t%b\t%#x\n", y,y,y,y)
}
